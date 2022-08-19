package sources

import (
	"context"
	"sync"

	"github.com/nginx/agent/sdk/v2/proto"
	"github.com/nginx/agent/v2/src/core/metrics"
	"github.com/shirou/gopsutil/v3/load"
	log "github.com/sirupsen/logrus"
)

type Load struct {
	*namedMetric
	avgStatsFunc func() (*load.AvgStat, error)
}

func NewLoadSource(namespace string) *Load {
	return &Load{namedMetric: &namedMetric{namespace, "load"}, avgStatsFunc: load.Avg}
}

func (c *Load) Collect(ctx context.Context, wg *sync.WaitGroup, m chan<- *proto.StatsEntity) {
	defer wg.Done()
	loadStats, err := c.avgStatsFunc()
	if err != nil {
		log.Errorf("Failed to collect Load metrics: %v", err)
		return
	}

	simpleMetrics := c.convertSamplesToSimpleMetrics(map[string]float64{
		"1":  loadStats.Load1,
		"5":  loadStats.Load5,
		"15": loadStats.Load15,
	})

	select {
	case <-ctx.Done():
	case m <- metrics.NewStatsEntity([]*proto.Dimension{}, simpleMetrics):
	}
}