// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metrics.proto

package proto

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Metric type enum
type MetricsReport_Type int32

const (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> adds updated generated files from protobuf
=======
>>>>>>> Create dedicated cache and upstream metrics reports
=======
=======
>>>>>>> adds updated generated files from protobuf
>>>>>>> adds updated generated files from protobuf
=======
>>>>>>> feat: restoring metrics.proto comments
	// System metric type
	MetricsReport_SYSTEM MetricsReport_Type = 0
	// NGINX instance metric type
	MetricsReport_INSTANCE MetricsReport_Type = 1
	// Agent metric type
	MetricsReport_AGENT MetricsReport_Type = 2
	// Cache zone metric type
<<<<<<< HEAD
<<<<<<< HEAD
	MetricsReport_CACHE_ZONE MetricsReport_Type = 3
	// Upstreams metric type
	MetricsReport_UPSTREAMS MetricsReport_Type = 4
=======
=======
>>>>>>> feat: fixed tests and rebased with nginx/agent main
	MetricsReport_SYSTEM     MetricsReport_Type = 0
	MetricsReport_INSTANCE   MetricsReport_Type = 1
	MetricsReport_AGENT      MetricsReport_Type = 2
	MetricsReport_CACHE_ZONE MetricsReport_Type = 3
	MetricsReport_UPSTREAMS  MetricsReport_Type = 4
<<<<<<< HEAD
>>>>>>> Create dedicated cache and upstream metrics reports
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> adds updated generated files from protobuf
=======
	MetricsReport_CACHE_ZONE MetricsReport_Type = 3
	// Upstreams metric type
	MetricsReport_UPSTREAMS MetricsReport_Type = 4
>>>>>>> adds updated generated files from protobuf
<<<<<<< HEAD
=======
>>>>>>> Create dedicated cache and upstream metrics reports
=======
>>>>>>> adds updated generated files from protobuf
=======
>>>>>>> feat: fixed tests and rebased with nginx/agent main
=======
	MetricsReport_CACHE_ZONE MetricsReport_Type = 3
	// Upstreams metric type
	MetricsReport_UPSTREAMS MetricsReport_Type = 4
>>>>>>> feat: restoring metrics.proto comments
)

var MetricsReport_Type_name = map[int32]string{
	0: "SYSTEM",
	1: "INSTANCE",
	2: "AGENT",
	3: "CACHE_ZONE",
	4: "UPSTREAMS",
}

var MetricsReport_Type_value = map[string]int32{
	"SYSTEM":     0,
	"INSTANCE":   1,
	"AGENT":      2,
	"CACHE_ZONE": 3,
	"UPSTREAMS":  4,
}

func (x MetricsReport_Type) String() string {
	return proto.EnumName(MetricsReport_Type_name, int32(x))
}

func (MetricsReport_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0, 0}
}

// Represents a metric report
type MetricsReport struct {
	// Provides meta information about the metrics
	Meta *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta"`
	// Type of metrics
	Type MetricsReport_Type `protobuf:"varint,2,opt,name=type,proto3,enum=f5.nginx.agent.sdk.MetricsReport_Type" json:"type"`
	// List of stats entities
	Data                 []*StatsEntity `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MetricsReport) Reset()         { *m = MetricsReport{} }
func (m *MetricsReport) String() string { return proto.CompactTextString(m) }
func (*MetricsReport) ProtoMessage()    {}
func (*MetricsReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0}
}
func (m *MetricsReport) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetricsReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetricsReport.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetricsReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsReport.Merge(m, src)
}
func (m *MetricsReport) XXX_Size() int {
	return m.Size()
}
func (m *MetricsReport) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsReport.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsReport proto.InternalMessageInfo

func (m *MetricsReport) GetMeta() *Metadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *MetricsReport) GetType() MetricsReport_Type {
	if m != nil {
		return m.Type
	}
	return MetricsReport_SYSTEM
}

func (m *MetricsReport) GetData() []*StatsEntity {
	if m != nil {
		return m.Data
	}
	return nil
}

// Represents a simple metric
type SimpleMetric struct {
	// Metric name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Metric value
	Value                float64  `protobuf:"fixed64,2,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleMetric) Reset()         { *m = SimpleMetric{} }
func (m *SimpleMetric) String() string { return proto.CompactTextString(m) }
func (*SimpleMetric) ProtoMessage()    {}
func (*SimpleMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{1}
}
func (m *SimpleMetric) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimpleMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SimpleMetric.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SimpleMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleMetric.Merge(m, src)
}
func (m *SimpleMetric) XXX_Size() int {
	return m.Size()
}
func (m *SimpleMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleMetric.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleMetric proto.InternalMessageInfo

func (m *SimpleMetric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SimpleMetric) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Represents a dimension which is a dimensional attribute used when classifying and categorizing data
type Dimension struct {
	// Dimension name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Dimension value
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dimension) Reset()         { *m = Dimension{} }
func (m *Dimension) String() string { return proto.CompactTextString(m) }
func (*Dimension) ProtoMessage()    {}
func (*Dimension) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{2}
}
func (m *Dimension) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Dimension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Dimension.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Dimension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dimension.Merge(m, src)
}
func (m *Dimension) XXX_Size() int {
	return m.Size()
}
func (m *Dimension) XXX_DiscardUnknown() {
	xxx_messageInfo_Dimension.DiscardUnknown(m)
}

var xxx_messageInfo_Dimension proto.InternalMessageInfo

func (m *Dimension) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Dimension) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Represents a stats entity which is a timestamped entry for dimensions and metrics
type StatsEntity struct {
	// Timestamp defines the time of stats entity creation
	Timestamp *types.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// List of dimensions
	Dimensions []*Dimension `protobuf:"bytes,3,rep,name=dimensions,proto3" json:"dimensions"`
	// List of metrics
	Simplemetrics        []*SimpleMetric `protobuf:"bytes,4,rep,name=simplemetrics,proto3" json:"simplemetrics"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StatsEntity) Reset()         { *m = StatsEntity{} }
func (m *StatsEntity) String() string { return proto.CompactTextString(m) }
func (*StatsEntity) ProtoMessage()    {}
func (*StatsEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{3}
}
func (m *StatsEntity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatsEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatsEntity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StatsEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsEntity.Merge(m, src)
}
func (m *StatsEntity) XXX_Size() int {
	return m.Size()
}
func (m *StatsEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsEntity.DiscardUnknown(m)
}

var xxx_messageInfo_StatsEntity proto.InternalMessageInfo

func (m *StatsEntity) GetTimestamp() *types.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *StatsEntity) GetDimensions() []*Dimension {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *StatsEntity) GetSimplemetrics() []*SimpleMetric {
	if m != nil {
		return m.Simplemetrics
	}
	return nil
}

func init() {
	proto.RegisterEnum("f5.nginx.agent.sdk.MetricsReport_Type", MetricsReport_Type_name, MetricsReport_Type_value)
	proto.RegisterType((*MetricsReport)(nil), "f5.nginx.agent.sdk.MetricsReport")
	proto.RegisterType((*SimpleMetric)(nil), "f5.nginx.agent.sdk.SimpleMetric")
	proto.RegisterType((*Dimension)(nil), "f5.nginx.agent.sdk.Dimension")
	proto.RegisterType((*StatsEntity)(nil), "f5.nginx.agent.sdk.StatsEntity")
}

func init() { proto.RegisterFile("metrics.proto", fileDescriptor_6039342a2ba47b72) }

var fileDescriptor_6039342a2ba47b72 = []byte{
<<<<<<< HEAD
<<<<<<< HEAD
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x8e, 0x5b, 0xe2, 0xc9, 0x87, 0xcc, 0x9e, 0xa2, 0xa8, 0xc4, 0x51, 0x0e, 0x28, 0x5c,
	0xd6, 0x52, 0x10, 0x02, 0x81, 0x38, 0xc4, 0xa9, 0x05, 0x44, 0x72, 0x40, 0x6b, 0x73, 0x68, 0x2f,
	0xc8, 0x4d, 0xb6, 0xc6, 0x6a, 0xd6, 0x6b, 0xc5, 0x9b, 0x8a, 0xdc, 0xf8, 0x21, 0xfc, 0x20, 0x8e,
	0xfc, 0x82, 0x08, 0xe5, 0x98, 0x5f, 0x81, 0xbc, 0xeb, 0xb4, 0x89, 0x08, 0x87, 0x5e, 0x66, 0x77,
	0xac, 0x79, 0xef, 0x8d, 0xdf, 0xb3, 0xa1, 0xc1, 0xa8, 0x58, 0x24, 0xd3, 0x1c, 0x67, 0x0b, 0x2e,
	0x38, 0x42, 0xd7, 0x2f, 0x71, 0x1a, 0x27, 0xe9, 0x77, 0x1c, 0xc5, 0x34, 0x15, 0x38, 0x9f, 0xdd,
	0xb4, 0xeb, 0x53, 0xce, 0x18, 0x4f, 0xd5, 0x44, 0x1b, 0x62, 0x1e, 0xf3, 0xf2, 0x6e, 0xc7, 0x9c,
	0xc7, 0x73, 0xea, 0xc8, 0xee, 0x6a, 0x79, 0xed, 0x88, 0x84, 0xd1, 0x5c, 0x44, 0x2c, 0x53, 0x03,
	0xbd, 0x9f, 0x3a, 0x34, 0x7c, 0x25, 0x40, 0x68, 0xc6, 0x17, 0x02, 0xbd, 0x01, 0x83, 0x51, 0x11,
	0xb5, 0xb4, 0xae, 0xd6, 0xaf, 0x0d, 0xce, 0xf0, 0xbf, 0x7a, 0xd8, 0xa7, 0x22, 0x9a, 0x45, 0x22,
	0x72, 0xab, 0xdb, 0xb5, 0x2d, 0xa7, 0x89, 0xac, 0xe8, 0x1c, 0x0c, 0xb1, 0xca, 0x68, 0x4b, 0xef,
	0x6a, 0xfd, 0xe6, 0xe0, 0xd9, 0x7f, 0xb0, 0xf7, 0x62, 0x38, 0x5c, 0x65, 0x54, 0xb1, 0x14, 0x38,
	0x22, 0x2b, 0x7a, 0x07, 0x46, 0xc1, 0xde, 0xaa, 0x74, 0x2b, 0xfd, 0xda, 0xc0, 0x3e, 0xc6, 0x12,
	0x88, 0x48, 0xe4, 0x5e, 0x2a, 0x12, 0xb1, 0x52, 0xf0, 0x02, 0x40, 0x64, 0xed, 0x8d, 0xc1, 0x28,
	0x68, 0x11, 0xc0, 0x69, 0x70, 0x11, 0x84, 0x9e, 0x6f, 0x3d, 0x42, 0x75, 0xa8, 0x7e, 0x9c, 0x04,
	0xe1, 0x70, 0x32, 0xf2, 0x2c, 0x0d, 0x99, 0x70, 0x32, 0x7c, 0xef, 0x4d, 0x42, 0x4b, 0x47, 0x4d,
	0x80, 0xd1, 0x70, 0xf4, 0xc1, 0xfb, 0x7a, 0xf9, 0x69, 0xe2, 0x59, 0x15, 0xd4, 0x00, 0xf3, 0xcb,
	0xe7, 0x20, 0x24, 0xde, 0xd0, 0x0f, 0x2c, 0xa3, 0xe7, 0x43, 0x3d, 0x48, 0x58, 0x36, 0xa7, 0x6a,
	0x6d, 0x74, 0x06, 0x46, 0x1a, 0x31, 0x2a, 0xcd, 0x31, 0x95, 0x72, 0xd1, 0x13, 0x59, 0x91, 0x0d,
	0x27, 0xb7, 0xd1, 0x7c, 0xa9, 0xde, 0x5f, 0x73, 0xcd, 0xed, 0xda, 0x56, 0x0f, 0x88, 0x3a, 0x7a,
	0x63, 0x30, 0xcf, 0x13, 0x46, 0xd3, 0x3c, 0xe1, 0xe9, 0x43, 0xb8, 0xcc, 0x23, 0x5c, 0x3f, 0x74,
	0xa8, 0xed, 0xd9, 0x80, 0x5e, 0x83, 0x79, 0x17, 0x6e, 0x19, 0x5e, 0x1b, 0xab, 0xf8, 0xf1, 0x2e,
	0x7e, 0x1c, 0xee, 0x26, 0xc8, 0xfd, 0x30, 0xf2, 0x01, 0x66, 0xbb, 0xad, 0xf2, 0xd2, 0xf5, 0xa7,
	0xc7, 0x5c, 0xbf, 0xdb, 0xdd, 0x6d, 0x6e, 0xd7, 0xf6, 0x1e, 0x88, 0xec, 0xdd, 0xd1, 0x05, 0x34,
	0x72, 0xe9, 0x59, 0xf9, 0xe1, 0xb6, 0x0c, 0xc9, 0xd8, 0x3d, 0x9a, 0xe3, 0x9e, 0xb9, 0xee, 0x93,
	0xed, 0xda, 0x3e, 0x84, 0x92, 0xc3, 0x76, 0x6c, 0x54, 0x75, 0xab, 0x42, 0x1e, 0x97, 0xad, 0xfb,
	0xea, 0xd7, 0xa6, 0xa3, 0xfd, 0xde, 0x74, 0xb4, 0x3f, 0x9b, 0x8e, 0x76, 0xf9, 0x3c, 0x4e, 0xc4,
	0xb7, 0xe5, 0x15, 0x9e, 0x72, 0xe6, 0x48, 0x1d, 0x47, 0xea, 0x38, 0xf9, 0xec, 0xc6, 0xb9, 0x1d,
	0xa8, 0x5f, 0xe0, 0xad, 0x72, 0xe2, 0x54, 0x1e, 0x2f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x87,
	0xa1, 0xae, 0x17, 0x5c, 0x03, 0x00, 0x00,
=======
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x8e, 0x5b, 0xe2, 0xc9, 0x8f, 0xcc, 0x9e, 0xa2, 0xa8, 0xc4, 0x51, 0x0e, 0x28, 0x5c,
=======
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x8e, 0x5b, 0xe2, 0xc9, 0x87, 0xcc, 0x9e, 0xa2, 0xa8, 0xc4, 0x51, 0x0e, 0x28, 0x5c,
>>>>>>> feat: restoring metrics.proto comments
	0xd6, 0x52, 0x10, 0x02, 0x81, 0x38, 0xc4, 0xa9, 0x05, 0x44, 0x72, 0x40, 0x6b, 0x73, 0x68, 0x2f,
	0xc8, 0x4d, 0xb6, 0xc6, 0x6a, 0xd6, 0x6b, 0xc5, 0x9b, 0x8a, 0xdc, 0xf8, 0x21, 0xfc, 0x20, 0x8e,
	0xfc, 0x82, 0x08, 0xe5, 0x98, 0x5f, 0x81, 0xbc, 0xeb, 0xb4, 0x89, 0x08, 0x87, 0x5e, 0x66, 0x77,
	0xac, 0x79, 0xef, 0x8d, 0xdf, 0xb3, 0xa1, 0xc1, 0xa8, 0x58, 0x24, 0xd3, 0x1c, 0x67, 0x0b, 0x2e,
	0x38, 0x42, 0xd7, 0x2f, 0x71, 0x1a, 0x27, 0xe9, 0x77, 0x1c, 0xc5, 0x34, 0x15, 0x38, 0x9f, 0xdd,
	0xb4, 0xeb, 0x53, 0xce, 0x18, 0x4f, 0xd5, 0x44, 0x1b, 0x62, 0x1e, 0xf3, 0xf2, 0x6e, 0xc7, 0x9c,
	0xc7, 0x73, 0xea, 0xc8, 0xee, 0x6a, 0x79, 0xed, 0x88, 0x84, 0xd1, 0x5c, 0x44, 0x2c, 0x53, 0x03,
	0xbd, 0x9f, 0x3a, 0x34, 0x7c, 0x25, 0x40, 0x68, 0xc6, 0x17, 0x02, 0xbd, 0x01, 0x83, 0x51, 0x11,
	0xb5, 0xb4, 0xae, 0xd6, 0xaf, 0x0d, 0xce, 0xf0, 0xbf, 0x7a, 0xd8, 0xa7, 0x22, 0x9a, 0x45, 0x22,
	0x72, 0xab, 0xdb, 0xb5, 0x2d, 0xa7, 0x89, 0xac, 0xe8, 0x1c, 0x0c, 0xb1, 0xca, 0x68, 0x4b, 0xef,
	0x6a, 0xfd, 0xe6, 0xe0, 0xd9, 0x7f, 0xb0, 0xf7, 0x62, 0x38, 0x5c, 0x65, 0x54, 0xb1, 0x14, 0x38,
	0x22, 0x2b, 0x7a, 0x07, 0x46, 0xc1, 0xde, 0xaa, 0x74, 0x2b, 0xfd, 0xda, 0xc0, 0x3e, 0xc6, 0x12,
	0x88, 0x48, 0xe4, 0x5e, 0x2a, 0x12, 0xb1, 0x52, 0xf0, 0x02, 0x40, 0x64, 0xed, 0x8d, 0xc1, 0x28,
	0x68, 0x11, 0xc0, 0x69, 0x70, 0x11, 0x84, 0x9e, 0x6f, 0x3d, 0x42, 0x75, 0xa8, 0x7e, 0x9c, 0x04,
	0xe1, 0x70, 0x32, 0xf2, 0x2c, 0x0d, 0x99, 0x70, 0x32, 0x7c, 0xef, 0x4d, 0x42, 0x4b, 0x47, 0x4d,
	0x80, 0xd1, 0x70, 0xf4, 0xc1, 0xfb, 0x7a, 0xf9, 0x69, 0xe2, 0x59, 0x15, 0xd4, 0x00, 0xf3, 0xcb,
	0xe7, 0x20, 0x24, 0xde, 0xd0, 0x0f, 0x2c, 0xa3, 0xe7, 0x43, 0x3d, 0x48, 0x58, 0x36, 0xa7, 0x6a,
	0x6d, 0x74, 0x06, 0x46, 0x1a, 0x31, 0x2a, 0xcd, 0x31, 0x95, 0x72, 0xd1, 0x13, 0x59, 0x91, 0x0d,
	0x27, 0xb7, 0xd1, 0x7c, 0xa9, 0xde, 0x5f, 0x73, 0xcd, 0xed, 0xda, 0x56, 0x0f, 0x88, 0x3a, 0x7a,
	0x63, 0x30, 0xcf, 0x13, 0x46, 0xd3, 0x3c, 0xe1, 0xe9, 0x43, 0xb8, 0xcc, 0x23, 0x5c, 0x3f, 0x74,
<<<<<<< HEAD
	0xa8, 0xed, 0xd9, 0x80, 0x5e, 0x83, 0x79, 0x37, 0xf6, 0x72, 0x78, 0x6d, 0xac, 0x16, 0x03, 0xef,
	0x16, 0x03, 0x87, 0xbb, 0x0a, 0x72, 0x5f, 0x8c, 0x7c, 0x80, 0xd9, 0xae, 0xab, 0xbc, 0x74, 0xfd,
	0xe9, 0x31, 0xd7, 0xef, 0x7a, 0x77, 0x9b, 0xdb, 0xb5, 0xbd, 0x07, 0x22, 0x7b, 0x77, 0x74, 0x01,
	0x8d, 0x5c, 0x7a, 0x56, 0x2e, 0x6e, 0xcb, 0x90, 0x8c, 0xdd, 0xa3, 0x73, 0xdc, 0x33, 0xd7, 0x7d,
	0xb2, 0x5d, 0xdb, 0x87, 0x50, 0x72, 0x98, 0x8e, 0x8d, 0xaa, 0x6e, 0x55, 0xc8, 0xe3, 0x32, 0x75,
	0x5f, 0xfd, 0xda, 0x74, 0xb4, 0xdf, 0x9b, 0x8e, 0xf6, 0x67, 0xd3, 0xd1, 0x2e, 0x9f, 0xc7, 0x89,
	0xf8, 0xb6, 0xbc, 0xc2, 0x53, 0xce, 0x1c, 0xa9, 0xe3, 0x48, 0x1d, 0x27, 0x9f, 0xdd, 0x38, 0xb7,
	0x03, 0xf5, 0x38, 0xde, 0x2a, 0x27, 0x4e, 0xe5, 0xf1, 0xe2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xee, 0x4a, 0xc6, 0xb7, 0x5c, 0x03, 0x00, 0x00,
>>>>>>> Create dedicated cache and upstream metrics reports
=======
	0xa8, 0xed, 0xd9, 0x80, 0x5e, 0x83, 0x79, 0x17, 0x6e, 0x19, 0x5e, 0x1b, 0xab, 0xf8, 0xf1, 0x2e,
	0x7e, 0x1c, 0xee, 0x26, 0xc8, 0xfd, 0x30, 0xf2, 0x01, 0x66, 0xbb, 0xad, 0xf2, 0xd2, 0xf5, 0xa7,
	0xc7, 0x5c, 0xbf, 0xdb, 0xdd, 0x6d, 0x6e, 0xd7, 0xf6, 0x1e, 0x88, 0xec, 0xdd, 0xd1, 0x05, 0x34,
	0x72, 0xe9, 0x59, 0xf9, 0xe1, 0xb6, 0x0c, 0xc9, 0xd8, 0x3d, 0x9a, 0xe3, 0x9e, 0xb9, 0xee, 0x93,
	0xed, 0xda, 0x3e, 0x84, 0x92, 0xc3, 0x76, 0x6c, 0x54, 0x75, 0xab, 0x42, 0x1e, 0x97, 0xad, 0xfb,
	0xea, 0xd7, 0xa6, 0xa3, 0xfd, 0xde, 0x74, 0xb4, 0x3f, 0x9b, 0x8e, 0x76, 0xf9, 0x3c, 0x4e, 0xc4,
	0xb7, 0xe5, 0x15, 0x9e, 0x72, 0xe6, 0x48, 0x1d, 0x47, 0xea, 0x38, 0xf9, 0xec, 0xc6, 0xb9, 0x1d,
	0xa8, 0x5f, 0xe0, 0xad, 0x72, 0xe2, 0x54, 0x1e, 0x2f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x87,
	0xa1, 0xae, 0x17, 0x5c, 0x03, 0x00, 0x00,
>>>>>>> feat: restoring metrics.proto comments
}

func (m *MetricsReport) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricsReport) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetricsReport) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMetrics(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Type != 0 {
		i = encodeVarintMetrics(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if m.Meta != nil {
		{
			size, err := m.Meta.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMetrics(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SimpleMetric) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimpleMetric) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SimpleMetric) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Value != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x11
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMetrics(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Dimension) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Dimension) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Dimension) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintMetrics(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMetrics(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StatsEntity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatsEntity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StatsEntity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Simplemetrics) > 0 {
		for iNdEx := len(m.Simplemetrics) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Simplemetrics[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMetrics(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Dimensions) > 0 {
		for iNdEx := len(m.Dimensions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Dimensions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMetrics(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMetrics(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMetrics(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetrics(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MetricsReport) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovMetrics(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovMetrics(uint64(m.Type))
	}
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovMetrics(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SimpleMetric) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMetrics(uint64(l))
	}
	if m.Value != 0 {
		n += 9
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Dimension) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMetrics(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovMetrics(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StatsEntity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovMetrics(uint64(l))
	}
	if len(m.Dimensions) > 0 {
		for _, e := range m.Dimensions {
			l = e.Size()
			n += 1 + l + sovMetrics(uint64(l))
		}
	}
	if len(m.Simplemetrics) > 0 {
		for _, e := range m.Simplemetrics {
			l = e.Size()
			n += 1 + l + sovMetrics(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMetrics(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetrics(x uint64) (n int) {
	return sovMetrics(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MetricsReport) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MetricsReport: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricsReport: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = &Metadata{}
			}
			if err := m.Meta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= MetricsReport_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &StatsEntity{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SimpleMetric) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SimpleMetric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimpleMetric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Dimension) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Dimension: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Dimension: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatsEntity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatsEntity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatsEntity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &types.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dimensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dimensions = append(m.Dimensions, &Dimension{})
			if err := m.Dimensions[len(m.Dimensions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Simplemetrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Simplemetrics = append(m.Simplemetrics, &SimpleMetric{})
			if err := m.Simplemetrics[len(m.Simplemetrics)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMetrics(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMetrics
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetrics
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetrics
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetrics        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetrics          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetrics = fmt.Errorf("proto: unexpected end of group")
)
