// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dp_software_details.proto

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Represents dataplane software details which contains details for additional software running on the dataplane that pertains to NGINX Agent
type DataplaneSoftwareDetails struct {
	// Types that are valid to be assigned to Data:
	//	*DataplaneSoftwareDetails_AppProtectWafDetails
	//	*DataplaneSoftwareDetails_NginxDetails
	Data                 isDataplaneSoftwareDetails_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *DataplaneSoftwareDetails) Reset()         { *m = DataplaneSoftwareDetails{} }
func (m *DataplaneSoftwareDetails) String() string { return proto.CompactTextString(m) }
func (*DataplaneSoftwareDetails) ProtoMessage()    {}
func (*DataplaneSoftwareDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a59b96dc90da7, []int{0}
}
func (m *DataplaneSoftwareDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataplaneSoftwareDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataplaneSoftwareDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataplaneSoftwareDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneSoftwareDetails.Merge(m, src)
}
func (m *DataplaneSoftwareDetails) XXX_Size() int {
	return m.Size()
}
func (m *DataplaneSoftwareDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneSoftwareDetails.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneSoftwareDetails proto.InternalMessageInfo

type isDataplaneSoftwareDetails_Data interface {
	isDataplaneSoftwareDetails_Data()
	MarshalTo([]byte) (int, error)
	Size() int
}

type DataplaneSoftwareDetails_AppProtectWafDetails struct {
	AppProtectWafDetails *AppProtectWAFDetails `protobuf:"bytes,1,opt,name=app_protect_waf_details,json=appProtectWafDetails,proto3,oneof"`
}
type DataplaneSoftwareDetails_NginxDetails struct {
	NginxDetails *NginxDetails `protobuf:"bytes,2,opt,name=nginx_details,json=nginxDetails,proto3,oneof"`
}

func (*DataplaneSoftwareDetails_AppProtectWafDetails) isDataplaneSoftwareDetails_Data() {}
func (*DataplaneSoftwareDetails_NginxDetails) isDataplaneSoftwareDetails_Data()         {}

func (m *DataplaneSoftwareDetails) GetData() isDataplaneSoftwareDetails_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DataplaneSoftwareDetails) GetAppProtectWafDetails() *AppProtectWAFDetails {
	if x, ok := m.GetData().(*DataplaneSoftwareDetails_AppProtectWafDetails); ok {
		return x.AppProtectWafDetails
	}
	return nil
}

func (m *DataplaneSoftwareDetails) GetNginxDetails() *NginxDetails {
	if x, ok := m.GetData().(*DataplaneSoftwareDetails_NginxDetails); ok {
		return x.NginxDetails
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DataplaneSoftwareDetails) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DataplaneSoftwareDetails_OneofMarshaler, _DataplaneSoftwareDetails_OneofUnmarshaler, _DataplaneSoftwareDetails_OneofSizer, []interface{}{
		(*DataplaneSoftwareDetails_AppProtectWafDetails)(nil),
		(*DataplaneSoftwareDetails_NginxDetails)(nil),
	}
}

func _DataplaneSoftwareDetails_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DataplaneSoftwareDetails)
	// data
	switch x := m.Data.(type) {
	case *DataplaneSoftwareDetails_AppProtectWafDetails:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AppProtectWafDetails); err != nil {
			return err
		}
	case *DataplaneSoftwareDetails_NginxDetails:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NginxDetails); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DataplaneSoftwareDetails.Data has unexpected type %T", x)
	}
	return nil
}

func _DataplaneSoftwareDetails_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DataplaneSoftwareDetails)
	switch tag {
	case 1: // data.app_protect_waf_details
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AppProtectWAFDetails)
		err := b.DecodeMessage(msg)
		m.Data = &DataplaneSoftwareDetails_AppProtectWafDetails{msg}
		return true, err
	case 2: // data.nginx_details
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NginxDetails)
		err := b.DecodeMessage(msg)
		m.Data = &DataplaneSoftwareDetails_NginxDetails{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DataplaneSoftwareDetails_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DataplaneSoftwareDetails)
	// data
	switch x := m.Data.(type) {
	case *DataplaneSoftwareDetails_AppProtectWafDetails:
		s := proto.Size(x.AppProtectWafDetails)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DataplaneSoftwareDetails_NginxDetails:
		s := proto.Size(x.NginxDetails)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*DataplaneSoftwareDetails)(nil), "f5.nginx.agent.sdk.DataplaneSoftwareDetails")
}

func init() { proto.RegisterFile("dp_software_details.proto", fileDescriptor_c38a59b96dc90da7) }

var fileDescriptor_c38a59b96dc90da7 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x29, 0x88, 0x2f,
	0xce, 0x4f, 0x2b, 0x29, 0x4f, 0x2c, 0x4a, 0x8d, 0x4f, 0x49, 0x2d, 0x49, 0xcc, 0xcc, 0x29, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0x33, 0xd5, 0xcb, 0x4b, 0xcf, 0xcc, 0xab, 0xd0,
	0x4b, 0x4c, 0x4f, 0xcd, 0x2b, 0xd1, 0x2b, 0x4e, 0xc9, 0x96, 0xe2, 0x4a, 0xcf, 0x4f, 0xcf, 0x87,
	0xc8, 0x4b, 0x71, 0xe6, 0x25, 0x16, 0x40, 0x99, 0xdc, 0x10, 0x75, 0x60, 0x8e, 0xd2, 0x3f, 0x46,
	0x2e, 0x09, 0x97, 0xc4, 0x92, 0xc4, 0x82, 0x9c, 0xc4, 0xbc, 0xd4, 0x60, 0xa8, 0xd9, 0x2e, 0x10,
	0xa3, 0x85, 0xea, 0xb8, 0xc4, 0x13, 0x0b, 0x0a, 0xe2, 0x41, 0x2a, 0x53, 0x93, 0x4b, 0xe2, 0xcb,
	0x13, 0xd3, 0x60, 0xb6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x69, 0xe8, 0x61, 0x5a, 0xab,
	0xe7, 0x58, 0x50, 0x10, 0x00, 0xd1, 0x11, 0xee, 0xe8, 0x06, 0x35, 0xca, 0x49, 0xfa, 0xd5, 0x3d,
	0x79, 0x5c, 0x86, 0x79, 0x30, 0x04, 0x89, 0x24, 0x22, 0x34, 0x25, 0xa6, 0xc1, 0xec, 0x8f, 0xe2,
	0xe2, 0x05, 0x1b, 0x0e, 0xb7, 0x95, 0x09, 0x6c, 0xab, 0x02, 0x36, 0x5b, 0xfd, 0x40, 0x7c, 0x98,
	0x6d, 0x82, 0xaf, 0xee, 0xc9, 0xa3, 0x6a, 0xf5, 0x60, 0x08, 0xe2, 0xc9, 0x43, 0x56, 0xc2, 0xc6,
	0xc5, 0x92, 0x92, 0x58, 0x92, 0xe8, 0x64, 0x7e, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x46, 0x69, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7,
	0xea, 0x83, 0x95, 0xeb, 0x83, 0x6d, 0xd0, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x33, 0xd2, 0x07, 0x87,
	0x99, 0x35, 0x98, 0x4c, 0x62, 0x03, 0x53, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xcc,
	0xb1, 0x3e, 0x95, 0x01, 0x00, 0x00,
}

func (m *DataplaneSoftwareDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataplaneSoftwareDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		nn1, err1 := m.Data.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += nn1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DataplaneSoftwareDetails_AppProtectWafDetails) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.AppProtectWafDetails != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDpSoftwareDetails(dAtA, i, uint64(m.AppProtectWafDetails.Size()))
		n2, err2 := m.AppProtectWafDetails.MarshalTo(dAtA[i:])
		if err2 != nil {
			return 0, err2
		}
		i += n2
	}
	return i, nil
}
func (m *DataplaneSoftwareDetails_NginxDetails) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.NginxDetails != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDpSoftwareDetails(dAtA, i, uint64(m.NginxDetails.Size()))
		n3, err3 := m.NginxDetails.MarshalTo(dAtA[i:])
		if err3 != nil {
			return 0, err3
		}
		i += n3
	}
	return i, nil
}
func encodeVarintDpSoftwareDetails(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DataplaneSoftwareDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		n += m.Data.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DataplaneSoftwareDetails_AppProtectWafDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AppProtectWafDetails != nil {
		l = m.AppProtectWafDetails.Size()
		n += 1 + l + sovDpSoftwareDetails(uint64(l))
	}
	return n
}
func (m *DataplaneSoftwareDetails_NginxDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NginxDetails != nil {
		l = m.NginxDetails.Size()
		n += 1 + l + sovDpSoftwareDetails(uint64(l))
	}
	return n
}

func sovDpSoftwareDetails(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDpSoftwareDetails(x uint64) (n int) {
	return sovDpSoftwareDetails(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataplaneSoftwareDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDpSoftwareDetails
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
			return fmt.Errorf("proto: DataplaneSoftwareDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataplaneSoftwareDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppProtectWafDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpSoftwareDetails
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
				return ErrInvalidLengthDpSoftwareDetails
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDpSoftwareDetails
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AppProtectWAFDetails{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Data = &DataplaneSoftwareDetails_AppProtectWafDetails{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NginxDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpSoftwareDetails
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
				return ErrInvalidLengthDpSoftwareDetails
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDpSoftwareDetails
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NginxDetails{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Data = &DataplaneSoftwareDetails_NginxDetails{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDpSoftwareDetails(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDpSoftwareDetails
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDpSoftwareDetails
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
func skipDpSoftwareDetails(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDpSoftwareDetails
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
					return 0, ErrIntOverflowDpSoftwareDetails
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDpSoftwareDetails
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
				return 0, ErrInvalidLengthDpSoftwareDetails
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthDpSoftwareDetails
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDpSoftwareDetails
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDpSoftwareDetails(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthDpSoftwareDetails
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDpSoftwareDetails = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDpSoftwareDetails   = fmt.Errorf("proto: integer overflow")
)
