/**
 * Copyright (c) F5, Inc.
 *
 * This source code is licensed under the Apache License, Version 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package plugins

import (
	"context"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"go.uber.org/atomic"

	"github.com/nginx/agent/sdk/v2/proto"
	"github.com/nginx/agent/v2/src/core"
	"github.com/nginx/agent/v2/src/core/config"
	"github.com/nginx/agent/v2/src/core/metrics"
)

const (
	reportStaggeringStartTime = 5555 * time.Millisecond // want to start the report cycle staggering of the collection time
)

type MetricsThrottle struct {
	messagePipeline    core.MessagePipeInterface
	BulkSize           int
	metricBuffer       []core.Payload
	ticker             *time.Ticker
	reportsReady       *atomic.Bool
	collectorsUpdate   *atomic.Bool
	metricsAggregation bool
	metricsCollections metrics.Collections
	ctx                context.Context
	wg                 sync.WaitGroup
	mu                 sync.Mutex
	env                core.Environment
	conf               *config.Config
	errors             chan error
}

func NewMetricsThrottle(conf *config.Config, env core.Environment) *MetricsThrottle {
	metricsCollections := metrics.Collections{
		Count: 0,
		Data:  make(map[string]metrics.PerDimension),
	}

	return &MetricsThrottle{
		metricBuffer:       make([]core.Payload, 0),
		BulkSize:           conf.AgentMetrics.BulkSize,
		ticker:             time.NewTicker(conf.AgentMetrics.ReportInterval + reportStaggeringStartTime),
		reportsReady:       atomic.NewBool(false),
		collectorsUpdate:   atomic.NewBool(false),
		metricsAggregation: conf.AgentMetrics.Mode == "aggregated",
		metricsCollections: metricsCollections,
		wg:                 sync.WaitGroup{},
		env:                env,
		conf:               conf,
		errors:             make(chan error),
	}
}

func (r *MetricsThrottle) Init(pipeline core.MessagePipeInterface) {
	r.messagePipeline = pipeline
	r.ctx = pipeline.Context()
	if r.metricsAggregation {
		r.wg.Add(1)
		go r.metricsReportGoroutine(r.ctx, &r.wg)
	}
	log.Info("MetricsThrottle initializing")
}

func (r *MetricsThrottle) Close() {
	log.Info("MetricsThrottle is wrapping up")
	r.reportsReady.Store(false) // allow metricsReportGoroutine to shutdown gracefully
	r.ticker.Stop()
}

func (r *MetricsThrottle) Info() *core.Info {
	return core.NewInfo("MetricsThrottle", "v0.0.1")
}

func (r *MetricsThrottle) Process(msg *core.Message) {
	switch {
	case msg.Exact(core.AgentConfigChanged):
		// If the agent config on disk changed update MetricsThrottle with relevant config info
		r.syncAgentConfigChange()
		r.collectorsUpdate.Store(true)
		return
	case msg.Exact(core.MetricReport):
		if r.metricsAggregation {
			switch report := msg.Data().(type) {
			case *proto.MetricsReport:
				r.mu.Lock()
				r.metricsCollections = metrics.SaveCollections(r.metricsCollections, report)
				r.mu.Unlock()
				log.Debug("MetricsThrottle: Metrics collection saved")
				r.reportsReady.Store(true)
			}
		} else {
			r.metricBuffer = append(r.metricBuffer,
				generateMetricsReports(getAllStatsEntities(msg.Data()), false)...)
			log.Tracef("MetricsThrottle buffer size: %d of %d", len(r.metricBuffer), r.BulkSize)
			if len(r.metricBuffer) >= r.BulkSize {
				log.Info("MetricsThrottle buffer flush")
				r.messagePipeline.Process(core.NewMessage(core.CommMetrics, r.metricBuffer))
				r.metricBuffer = make([]core.Payload, 0)
			}
		}
	}
}

func (r *MetricsThrottle) Subscriptions() []string {
	return []string{core.MetricReport, core.AgentConfigChanged, core.LoggerLevel}
}

func (r *MetricsThrottle) metricsReportGoroutine(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer r.ticker.Stop()
	log.Info("MetricsThrottle waiting for report ready")
	for {
		if !r.reportsReady.Load() {
			continue
		}

		select {
		case <-ctx.Done():
			err := r.ctx.Err()
			if err != nil {
				log.Errorf("error in done context metricsReportGoroutine %v", err)
			}
			return
		case <-r.ticker.C:
			reports := r.getAggregatedReports()
			if len(reports) > 0 {
				r.messagePipeline.Process(core.NewMessage(core.CommMetrics, reports))
			}
			if r.collectorsUpdate.Load() {
				r.BulkSize = r.conf.AgentMetrics.BulkSize
				r.metricsAggregation = r.conf.AgentMetrics.Mode == "aggregated"
				r.ticker.Stop()
				r.ticker = time.NewTicker(r.conf.AgentMetrics.ReportInterval + reportStaggeringStartTime)
				r.messagePipeline.Process(core.NewMessage(core.AgentCollectorsUpdate, ""))
				r.collectorsUpdate.Store(false)
			}
		case err := <-r.errors:
			log.Errorf("Error in metricsReportGoroutine %v", err)
		}
	}
}

func (r *MetricsThrottle) syncAgentConfigChange() {
	conf, err := config.GetConfig(r.env.GetSystemUUID())
	if err != nil {
		log.Errorf("Failed to load config for updating: %v", err)
		return
	}
	if conf.DisplayName == "" {
		conf.DisplayName = r.env.GetHostname()
		log.Infof("setting displayName to %s", conf.DisplayName)
	}
	log.Debugf("MetricsThrottle is updating to a new config - %v", conf)

	r.conf = conf
}

func (r *MetricsThrottle) getAggregatedReports() []core.Payload {
	r.mu.Lock()
	defer r.mu.Unlock()
	reports := generateMetricsReports(metrics.GenerateMetrics(r.metricsCollections), false)
	r.metricsCollections = metrics.Collections{
		Count: 0,
		Data:  make(map[string]metrics.PerDimension),
	}

	return reports
}
