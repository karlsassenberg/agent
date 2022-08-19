// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dpenv.proto

package proto

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type EnvReport_Type int32

const (
	EnvReport_SYSTEM   EnvReport_Type = 0
	EnvReport_INSTANCE EnvReport_Type = 1
	EnvReport_AGENT    EnvReport_Type = 2
)

var EnvReport_Type_name = map[int32]string{
	0: "SYSTEM",
	1: "INSTANCE",
	2: "AGENT",
}

var EnvReport_Type_value = map[string]int32{
	"SYSTEM":   0,
	"INSTANCE": 1,
	"AGENT":    2,
}

func (x EnvReport_Type) String() string {
	return proto.EnumName(EnvReport_Type_name, int32(x))
}

func (EnvReport_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a2ec482b9ece7397, []int{0, 0}
}

// MetasReport a report containing status entities for a specific metric type
type EnvReport struct {
	Meta                 *Metadata         `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta"`
	Type                 EnvReport_Type    `protobuf:"varint,2,opt,name=type,proto3,enum=f5.nginx.agent.sdk.EnvReport_Type" json:"type"`
	PropertySets         []*EnvPropertySet `protobuf:"bytes,3,rep,name=property_sets,json=propertySets,proto3" json:"property_sets"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EnvReport) Reset()         { *m = EnvReport{} }
func (m *EnvReport) String() string { return proto.CompactTextString(m) }
func (*EnvReport) ProtoMessage()    {}
func (*EnvReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2ec482b9ece7397, []int{0}
}
func (m *EnvReport) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EnvReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EnvReport.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EnvReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvReport.Merge(m, src)
}
func (m *EnvReport) XXX_Size() int {
	return m.Size()
}
func (m *EnvReport) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvReport.DiscardUnknown(m)
}

var xxx_messageInfo_EnvReport proto.InternalMessageInfo

func (m *EnvReport) GetMeta() *Metadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *EnvReport) GetType() EnvReport_Type {
	if m != nil {
		return m.Type
	}
	return EnvReport_SYSTEM
}

func (m *EnvReport) GetPropertySets() []*EnvPropertySet {
	if m != nil {
		return m.PropertySets
	}
	return nil
}

// EnvPropety - a container for a Dataplane Environment property.
type EnvProperty struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Types that are valid to be assigned to Value:
	//	*EnvProperty_Ival
	//	*EnvProperty_Dval
	//	*EnvProperty_Sval
	Value                isEnvProperty_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *EnvProperty) Reset()         { *m = EnvProperty{} }
func (m *EnvProperty) String() string { return proto.CompactTextString(m) }
func (*EnvProperty) ProtoMessage()    {}
func (*EnvProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2ec482b9ece7397, []int{1}
}
func (m *EnvProperty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EnvProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EnvProperty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EnvProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvProperty.Merge(m, src)
}
func (m *EnvProperty) XXX_Size() int {
	return m.Size()
}
func (m *EnvProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvProperty.DiscardUnknown(m)
}

var xxx_messageInfo_EnvProperty proto.InternalMessageInfo

type isEnvProperty_Value interface {
	isEnvProperty_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type EnvProperty_Ival struct {
	Ival int64 `protobuf:"varint,10,opt,name=ival,proto3,oneof" json:"ival"`
}
type EnvProperty_Dval struct {
	Dval float64 `protobuf:"fixed64,11,opt,name=dval,proto3,oneof" json:"dval"`
}
type EnvProperty_Sval struct {
	Sval string `protobuf:"bytes,12,opt,name=sval,proto3,oneof" json:"sval"`
}

func (*EnvProperty_Ival) isEnvProperty_Value() {}
func (*EnvProperty_Dval) isEnvProperty_Value() {}
func (*EnvProperty_Sval) isEnvProperty_Value() {}

func (m *EnvProperty) GetValue() isEnvProperty_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *EnvProperty) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EnvProperty) GetIval() int64 {
	if x, ok := m.GetValue().(*EnvProperty_Ival); ok {
		return x.Ival
	}
	return 0
}

func (m *EnvProperty) GetDval() float64 {
	if x, ok := m.GetValue().(*EnvProperty_Dval); ok {
		return x.Dval
	}
	return 0
}

func (m *EnvProperty) GetSval() string {
	if x, ok := m.GetValue().(*EnvProperty_Sval); ok {
		return x.Sval
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EnvProperty) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EnvProperty_Ival)(nil),
		(*EnvProperty_Dval)(nil),
		(*EnvProperty_Sval)(nil),
	}
}

type EnvPropertySet struct {
	Dimensions           []*Dimension   `protobuf:"bytes,1,rep,name=dimensions,proto3" json:"dimensions"`
	Properties           []*EnvProperty `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *EnvPropertySet) Reset()         { *m = EnvPropertySet{} }
func (m *EnvPropertySet) String() string { return proto.CompactTextString(m) }
func (*EnvPropertySet) ProtoMessage()    {}
func (*EnvPropertySet) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2ec482b9ece7397, []int{2}
}
func (m *EnvPropertySet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EnvPropertySet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EnvPropertySet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EnvPropertySet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvPropertySet.Merge(m, src)
}
func (m *EnvPropertySet) XXX_Size() int {
	return m.Size()
}
func (m *EnvPropertySet) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvPropertySet.DiscardUnknown(m)
}

var xxx_messageInfo_EnvPropertySet proto.InternalMessageInfo

func (m *EnvPropertySet) GetDimensions() []*Dimension {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *EnvPropertySet) GetProperties() []*EnvProperty {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterEnum("f5.nginx.agent.sdk.EnvReport_Type", EnvReport_Type_name, EnvReport_Type_value)
	proto.RegisterType((*EnvReport)(nil), "f5.nginx.agent.sdk.EnvReport")
	proto.RegisterType((*EnvProperty)(nil), "f5.nginx.agent.sdk.EnvProperty")
	proto.RegisterType((*EnvPropertySet)(nil), "f5.nginx.agent.sdk.EnvPropertySet")
}

func init() { proto.RegisterFile("dpenv.proto", fileDescriptor_a2ec482b9ece7397) }

var fileDescriptor_a2ec482b9ece7397 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xbb, 0x89, 0xdb, 0xaf, 0x19, 0xa7, 0x51, 0xbe, 0x3d, 0x59, 0x55, 0x49, 0x2c, 0x9f,
	0x2c, 0x21, 0x6c, 0x29, 0xa8, 0x97, 0x72, 0x69, 0x0d, 0x11, 0x70, 0x48, 0x40, 0x4e, 0x2e, 0xf4,
	0x82, 0xb6, 0xd9, 0xad, 0x65, 0x35, 0xde, 0x5d, 0x79, 0xb7, 0x16, 0x79, 0x0d, 0xc4, 0x83, 0xf0,
	0x18, 0x1c, 0x79, 0x82, 0x08, 0xe5, 0x98, 0xa7, 0x40, 0xde, 0x35, 0x24, 0x15, 0x45, 0x5c, 0x66,
	0x76, 0xf6, 0x37, 0xf3, 0x1f, 0xfb, 0xaf, 0x05, 0x97, 0x4a, 0xc6, 0xab, 0x48, 0x96, 0x42, 0x0b,
	0x8c, 0x6f, 0xcf, 0x23, 0x9e, 0xe5, 0xfc, 0x53, 0x44, 0x32, 0xc6, 0x75, 0xa4, 0xe8, 0xdd, 0x29,
	0x64, 0x22, 0x13, 0x96, 0x9f, 0x76, 0x17, 0xa2, 0x28, 0x04, 0x6f, 0xaa, 0x93, 0x82, 0xe9, 0x32,
	0x5f, 0x28, 0x5b, 0x06, 0x9f, 0x5b, 0xd0, 0x19, 0xf3, 0x2a, 0x65, 0x52, 0x94, 0x1a, 0x5f, 0x80,
	0x53, 0x30, 0x4d, 0x3c, 0xe4, 0xa3, 0xd0, 0x1d, 0x9d, 0x45, 0x7f, 0x2a, 0x47, 0x13, 0xa6, 0x09,
	0x25, 0x9a, 0x24, 0xc7, 0xdb, 0xf5, 0xd0, 0x74, 0xa7, 0x26, 0xe2, 0x4b, 0x70, 0xf4, 0x4a, 0x32,
	0xaf, 0xe5, 0xa3, 0xb0, 0x37, 0x0a, 0x1e, 0x9b, 0xfd, 0xbd, 0x28, 0x9a, 0xaf, 0x24, 0xb3, 0x0a,
	0xf5, 0x4c, 0x6a, 0x22, 0xbe, 0x86, 0x13, 0x59, 0x0a, 0xc9, 0x4a, 0xbd, 0xfa, 0xa8, 0x98, 0x56,
	0x5e, 0xdb, 0x6f, 0x87, 0xee, 0x5f, 0xa5, 0xde, 0x37, 0xbd, 0x33, 0xa6, 0x93, 0xff, 0xb7, 0xeb,
	0xe1, 0xc3, 0xe1, 0xb4, 0x2b, 0x77, 0x5c, 0x05, 0x4f, 0xc1, 0xa9, 0x77, 0x62, 0x80, 0xa3, 0xd9,
	0x87, 0xd9, 0x7c, 0x3c, 0xe9, 0x1f, 0xe0, 0x2e, 0x1c, 0xbf, 0x9d, 0xce, 0xe6, 0x57, 0xd3, 0x97,
	0xe3, 0x3e, 0xc2, 0x1d, 0x38, 0xbc, 0x7a, 0x3d, 0x9e, 0xce, 0xfb, 0xad, 0xe0, 0x0b, 0x02, 0x77,
	0x6f, 0x01, 0x3e, 0x03, 0x87, 0x93, 0x82, 0x19, 0x5b, 0x3a, 0xf6, 0xb3, 0xeb, 0x3a, 0x35, 0x11,
	0x0f, 0xc0, 0xc9, 0x2b, 0xb2, 0xf4, 0xc0, 0x47, 0x61, 0xdb, 0xd2, 0xba, 0x7e, 0x73, 0x90, 0x9a,
	0x5c, 0x73, 0x5a, 0x73, 0xd7, 0x47, 0x21, 0xb2, 0x9c, 0x36, 0x9c, 0x36, 0x5c, 0xd5, 0xbc, 0xbb,
	0x53, 0x57, 0x0d, 0xaf, 0x73, 0xf2, 0x1f, 0x1c, 0x56, 0x64, 0x79, 0xcf, 0x82, 0xaf, 0x08, 0x7a,
	0x0f, 0xff, 0x1b, 0x4f, 0x00, 0x68, 0x5e, 0x30, 0xae, 0x72, 0xc1, 0x95, 0x87, 0x8c, 0x5f, 0x4f,
	0x1e, 0xf3, 0xeb, 0xd5, 0xaf, 0xae, 0xa4, 0xb7, 0x5d, 0x0f, 0xf7, 0x86, 0xd2, 0xbd, 0x33, 0x7e,
	0x07, 0xd0, 0xb8, 0x96, 0x33, 0xe5, 0xb5, 0x8c, 0xdc, 0xf0, 0x1f, 0xf6, 0x5b, 0xc1, 0xdd, 0x58,
	0xba, 0x77, 0x4e, 0xa6, 0xdf, 0x36, 0x03, 0xf4, 0x7d, 0x33, 0x40, 0x3f, 0x36, 0x03, 0x74, 0x7d,
	0x99, 0xe5, 0x7a, 0x49, 0x6e, 0xa2, 0x85, 0x28, 0x2e, 0x6e, 0xcf, 0x63, 0x23, 0x1a, 0x1b, 0xd1,
	0x58, 0x96, 0x82, 0xde, 0x2f, 0xb4, 0xbd, 0x7b, 0x66, 0xef, 0x14, 0xbd, 0x8b, 0xab, 0x51, 0x6c,
	0x1e, 0xea, 0x0b, 0x13, 0x6f, 0x8e, 0x4c, 0x7a, 0xfe, 0x33, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x4c,
	0x43, 0xe2, 0x01, 0x03, 0x00, 0x00,
}

func (m *EnvReport) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnvReport) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EnvReport) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PropertySets) > 0 {
		for iNdEx := len(m.PropertySets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PropertySets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDpenv(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Type != 0 {
		i = encodeVarintDpenv(dAtA, i, uint64(m.Type))
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
			i = encodeVarintDpenv(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EnvProperty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnvProperty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EnvProperty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Value != nil {
		{
			size := m.Value.Size()
			i -= size
			if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDpenv(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EnvProperty_Ival) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EnvProperty_Ival) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintDpenv(dAtA, i, uint64(m.Ival))
	i--
	dAtA[i] = 0x50
	return len(dAtA) - i, nil
}
func (m *EnvProperty_Dval) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EnvProperty_Dval) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= 8
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Dval))))
	i--
	dAtA[i] = 0x59
	return len(dAtA) - i, nil
}
func (m *EnvProperty_Sval) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EnvProperty_Sval) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Sval)
	copy(dAtA[i:], m.Sval)
	i = encodeVarintDpenv(dAtA, i, uint64(len(m.Sval)))
	i--
	dAtA[i] = 0x62
	return len(dAtA) - i, nil
}
func (m *EnvPropertySet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnvPropertySet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EnvPropertySet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Properties) > 0 {
		for iNdEx := len(m.Properties) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Properties[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDpenv(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
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
				i = encodeVarintDpenv(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDpenv(dAtA []byte, offset int, v uint64) int {
	offset -= sovDpenv(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EnvReport) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovDpenv(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovDpenv(uint64(m.Type))
	}
	if len(m.PropertySets) > 0 {
		for _, e := range m.PropertySets {
			l = e.Size()
			n += 1 + l + sovDpenv(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EnvProperty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDpenv(uint64(l))
	}
	if m.Value != nil {
		n += m.Value.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EnvProperty_Ival) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovDpenv(uint64(m.Ival))
	return n
}
func (m *EnvProperty_Dval) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 9
	return n
}
func (m *EnvProperty_Sval) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sval)
	n += 1 + l + sovDpenv(uint64(l))
	return n
}
func (m *EnvPropertySet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Dimensions) > 0 {
		for _, e := range m.Dimensions {
			l = e.Size()
			n += 1 + l + sovDpenv(uint64(l))
		}
	}
	if len(m.Properties) > 0 {
		for _, e := range m.Properties {
			l = e.Size()
			n += 1 + l + sovDpenv(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDpenv(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDpenv(x uint64) (n int) {
	return sovDpenv(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EnvReport) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDpenv
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
			return fmt.Errorf("proto: EnvReport: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnvReport: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpenv
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
				return ErrInvalidLengthDpenv
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDpenv
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
					return ErrIntOverflowDpenv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= EnvReport_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PropertySets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpenv
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
				return ErrInvalidLengthDpenv
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDpenv
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PropertySets = append(m.PropertySets, &EnvPropertySet{})
			if err := m.PropertySets[len(m.PropertySets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDpenv(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDpenv
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
func (m *EnvProperty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDpenv
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
			return fmt.Errorf("proto: EnvProperty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnvProperty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpenv
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
				return ErrInvalidLengthDpenv
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDpenv
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ival", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpenv
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &EnvProperty_Ival{v}
		case 11:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dval", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = &EnvProperty_Dval{float64(math.Float64frombits(v))}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sval", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpenv
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
				return ErrInvalidLengthDpenv
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDpenv
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = &EnvProperty_Sval{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDpenv(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDpenv
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
func (m *EnvPropertySet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDpenv
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
			return fmt.Errorf("proto: EnvPropertySet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnvPropertySet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dimensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpenv
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
				return ErrInvalidLengthDpenv
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDpenv
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dimensions = append(m.Dimensions, &Dimension{})
			if err := m.Dimensions[len(m.Dimensions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpenv
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
				return ErrInvalidLengthDpenv
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDpenv
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Properties = append(m.Properties, &EnvProperty{})
			if err := m.Properties[len(m.Properties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDpenv(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDpenv
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
func skipDpenv(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDpenv
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
					return 0, ErrIntOverflowDpenv
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
					return 0, ErrIntOverflowDpenv
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
				return 0, ErrInvalidLengthDpenv
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDpenv
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDpenv
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDpenv        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDpenv          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDpenv = fmt.Errorf("proto: unexpected end of group")
)