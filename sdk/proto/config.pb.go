// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

package proto

import (
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

type ConfigReport struct {
	Meta                 *Metadata           `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta"`
	Configs              []*ConfigDescriptor `protobuf:"bytes,2,rep,name=configs,proto3" json:"configs"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ConfigReport) Reset()         { *m = ConfigReport{} }
func (m *ConfigReport) String() string { return proto.CompactTextString(m) }
func (*ConfigReport) ProtoMessage()    {}
func (*ConfigReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}
func (m *ConfigReport) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigReport.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigReport.Merge(m, src)
}
func (m *ConfigReport) XXX_Size() int {
	return m.Size()
}
func (m *ConfigReport) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigReport.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigReport proto.InternalMessageInfo

func (m *ConfigReport) GetMeta() *Metadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ConfigReport) GetConfigs() []*ConfigDescriptor {
	if m != nil {
		return m.Configs
	}
	return nil
}

type ConfigDescriptor struct {
	SystemId             string   `protobuf:"bytes,1,opt,name=system_id,json=systemId,proto3" json:"system_id"`
	NginxId              string   `protobuf:"bytes,2,opt,name=nginx_id,json=nginxId,proto3" json:"nginx_id"`
	Checksum             string   `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigDescriptor) Reset()         { *m = ConfigDescriptor{} }
func (m *ConfigDescriptor) String() string { return proto.CompactTextString(m) }
func (*ConfigDescriptor) ProtoMessage()    {}
func (*ConfigDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}
func (m *ConfigDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigDescriptor.Merge(m, src)
}
func (m *ConfigDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *ConfigDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigDescriptor proto.InternalMessageInfo

func (m *ConfigDescriptor) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *ConfigDescriptor) GetNginxId() string {
	if m != nil {
		return m.NginxId
	}
	return ""
}

func (m *ConfigDescriptor) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func init() {
	proto.RegisterType((*ConfigReport)(nil), "f5.nginx.agent.sdk.ConfigReport")
	proto.RegisterType((*ConfigDescriptor)(nil), "f5.nginx.agent.sdk.ConfigDescriptor")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xc9, 0x26, 0xae, 0xcb, 0x2a, 0x48, 0x4e, 0x63, 0xc8, 0x3a, 0x86, 0xe0, 0x10, 0x4c,
	0x61, 0xb2, 0xcb, 0xbc, 0x48, 0xf5, 0x32, 0x44, 0x0f, 0x3d, 0x7a, 0x91, 0x2c, 0xc9, 0x62, 0x99,
	0x69, 0x4a, 0x93, 0x89, 0x7e, 0x0c, 0x4f, 0x7e, 0x25, 0x8f, 0x7e, 0x82, 0x22, 0x3b, 0xf6, 0x53,
	0x88, 0x2f, 0x6c, 0x03, 0xf5, 0xf2, 0xfa, 0xef, 0x3f, 0xbf, 0xff, 0x7b, 0x2f, 0xc1, 0x21, 0x37,
	0xf9, 0x22, 0x53, 0xb4, 0x28, 0x8d, 0x33, 0x84, 0x2c, 0x26, 0x34, 0x57, 0x59, 0xfe, 0x42, 0x99,
	0x92, 0xb9, 0xa3, 0x56, 0x2c, 0x7b, 0x58, 0x19, 0x65, 0xfc, 0x79, 0x2f, 0xe4, 0x46, 0x6b, 0x93,
	0xfb, 0xbf, 0xe1, 0x3b, 0xc2, 0xe1, 0x15, 0xc4, 0x53, 0x59, 0x98, 0xd2, 0x91, 0x29, 0xde, 0xd3,
	0xd2, 0xb1, 0x2e, 0x1a, 0xa0, 0x51, 0x67, 0x7c, 0x44, 0xff, 0x76, 0xa3, 0xb7, 0xd2, 0x31, 0xc1,
	0x1c, 0x4b, 0x82, 0xba, 0x8a, 0x80, 0x4e, 0xa1, 0x92, 0x1b, 0xdc, 0xf2, 0xab, 0xd8, 0x6e, 0x63,
	0xd0, 0x1c, 0x75, 0xc6, 0xc7, 0xff, 0xc5, 0xfd, 0xb8, 0x6b, 0x69, 0x79, 0x99, 0x15, 0xce, 0x94,
	0x49, 0xa7, 0xae, 0xa2, 0x4d, 0x30, 0xdd, 0x88, 0xe1, 0x1b, 0xc2, 0x87, 0xbf, 0x51, 0x72, 0x8a,
	0xdb, 0xf6, 0xd5, 0x3a, 0xa9, 0x1f, 0x32, 0x01, 0x2b, 0xb6, 0x93, 0x83, 0xba, 0x8a, 0x76, 0x66,
	0x1a, 0x78, 0x39, 0x13, 0xe4, 0x04, 0x07, 0x30, 0xfa, 0x07, 0x6d, 0x00, 0x1a, 0xd6, 0x55, 0xb4,
	0xf5, 0xd2, 0x16, 0xa8, 0x99, 0x20, 0x23, 0x1c, 0xf0, 0x47, 0xc9, 0x97, 0x76, 0xa5, 0xbb, 0xcd,
	0x1d, 0xb8, 0xf1, 0xd2, 0xad, 0x4a, 0xee, 0x3e, 0xd6, 0x7d, 0xf4, 0xb9, 0xee, 0xa3, 0xaf, 0x75,
	0x1f, 0xdd, 0x5f, 0xaa, 0xcc, 0x3d, 0xb1, 0x39, 0xe5, 0x46, 0x4f, 0x17, 0x93, 0x18, 0xda, 0xc5,
	0x70, 0xcf, 0xb8, 0x28, 0x8d, 0x58, 0x71, 0xe7, 0xbd, 0x33, 0xef, 0x59, 0xb1, 0x8c, 0x9f, 0xc7,
	0x31, 0xbc, 0xfb, 0x05, 0xd4, 0xf9, 0x3e, 0x7c, 0xce, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa6,
	0xff, 0x38, 0xf2, 0xc2, 0x01, 0x00, 0x00,
}

func (m *ConfigReport) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigReport) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigReport) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Configs) > 0 {
		for iNdEx := len(m.Configs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Configs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Meta != nil {
		{
			size, err := m.Meta.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConfigDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Checksum) > 0 {
		i -= len(m.Checksum)
		copy(dAtA[i:], m.Checksum)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Checksum)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NginxId) > 0 {
		i -= len(m.NginxId)
		copy(dAtA[i:], m.NginxId)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.NginxId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SystemId) > 0 {
		i -= len(m.SystemId)
		copy(dAtA[i:], m.SystemId)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.SystemId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConfigReport) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if len(m.Configs) > 0 {
		for _, e := range m.Configs {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConfigDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SystemId)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.NginxId)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConfigReport) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: ConfigReport: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigReport: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Configs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Configs = append(m.Configs, &ConfigDescriptor{})
			if err := m.Configs[len(m.Configs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
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
func (m *ConfigDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: ConfigDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SystemId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NginxId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NginxId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)