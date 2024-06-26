// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proxy.proto

package proxy

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ExtraInfo struct {
	Salt                 []byte            `protobuf:"bytes,1,opt,name=Salt,proto3" json:"Salt,omitempty"`
	InternalConn         bool              `protobuf:"varint,2,opt,name=InternalConn,proto3" json:"InternalConn,omitempty"`
	ConnectionID         uint32            `protobuf:"varint,3,opt,name=ConnectionID,proto3" json:"ConnectionID,omitempty"`
	Label                map[string]string `protobuf:"bytes,4,rep,name=Label,proto3" json:"Label,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ClientAddr           string            `protobuf:"bytes,5,opt,name=ClientAddr,proto3" json:"ClientAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExtraInfo) Reset()         { *m = ExtraInfo{} }
func (m *ExtraInfo) String() string { return proto.CompactTextString(m) }
func (*ExtraInfo) ProtoMessage()    {}
func (*ExtraInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{0}
}
func (m *ExtraInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtraInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtraInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtraInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraInfo.Merge(m, src)
}
func (m *ExtraInfo) XXX_Size() int {
	return m.ProtoSize()
}
func (m *ExtraInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraInfo proto.InternalMessageInfo

func (m *ExtraInfo) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *ExtraInfo) GetInternalConn() bool {
	if m != nil {
		return m.InternalConn
	}
	return false
}

func (m *ExtraInfo) GetConnectionID() uint32 {
	if m != nil {
		return m.ConnectionID
	}
	return 0
}

func (m *ExtraInfo) GetLabel() map[string]string {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *ExtraInfo) GetClientAddr() string {
	if m != nil {
		return m.ClientAddr
	}
	return ""
}

func init() {
	proto.RegisterType((*ExtraInfo)(nil), "proxy.ExtraInfo")
	proto.RegisterMapType((map[string]string)(nil), "proxy.ExtraInfo.LabelEntry")
}

func init() { proto.RegisterFile("proxy.proto", fileDescriptor_700b50b08ed8dbaf) }

var fileDescriptor_700b50b08ed8dbaf = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x9d, 0xa6, 0x11, 0xf3, 0x5a, 0x41, 0x06, 0x17, 0xa1, 0x42, 0x08, 0x5d, 0x65, 0x63,
	0x82, 0xba, 0x29, 0xae, 0xd4, 0xda, 0x45, 0xc0, 0xd5, 0xb8, 0x73, 0x37, 0x69, 0xa7, 0x71, 0x68,
	0x3a, 0x13, 0x86, 0x89, 0x24, 0xb7, 0xf1, 0x38, 0x2e, 0x3d, 0x82, 0xc4, 0x2b, 0x78, 0x00, 0xc9,
	0x0b, 0xda, 0x76, 0x35, 0xff, 0xf7, 0xbf, 0xf7, 0x98, 0x9f, 0x1f, 0x46, 0xa5, 0xd1, 0x75, 0x13,
	0x97, 0x46, 0x5b, 0x4d, 0x5d, 0x84, 0xc9, 0x65, 0x2e, 0xed, 0x6b, 0x95, 0xc5, 0x4b, 0xbd, 0x4d,
	0x72, 0x9d, 0xeb, 0x04, 0xa7, 0x59, 0xb5, 0x46, 0x42, 0x40, 0xd5, 0x5f, 0x4d, 0x7f, 0x08, 0x78,
	0x8b, 0xda, 0x1a, 0x9e, 0xaa, 0xb5, 0xa6, 0x14, 0x86, 0xcf, 0xbc, 0xb0, 0x3e, 0x09, 0x49, 0x34,
	0x66, 0xa8, 0xe9, 0x14, 0xc6, 0xa9, 0xb2, 0xc2, 0x28, 0x5e, 0xcc, 0xb5, 0x52, 0xfe, 0x20, 0x24,
	0xd1, 0x09, 0x3b, 0xf0, 0xba, 0x9d, 0xee, 0x15, 0x4b, 0x2b, 0xb5, 0x4a, 0x1f, 0x7d, 0x27, 0x24,
	0xd1, 0x29, 0x3b, 0xf0, 0xe8, 0x15, 0xb8, 0x4f, 0x3c, 0x13, 0x85, 0x3f, 0x0c, 0x9d, 0x68, 0x74,
	0x7d, 0x11, 0xf7, 0xe1, 0xff, 0x3f, 0x8f, 0x71, 0xba, 0x50, 0xd6, 0x34, 0xac, 0xdf, 0xa4, 0x01,
	0xc0, 0xbc, 0x90, 0x42, 0xd9, 0xfb, 0xd5, 0xca, 0xf8, 0x6e, 0x48, 0x22, 0x8f, 0xed, 0x39, 0x93,
	0x19, 0xc0, 0xee, 0x88, 0x9e, 0x81, 0xb3, 0x11, 0x0d, 0x66, 0xf7, 0x58, 0x27, 0xe9, 0x39, 0xb8,
	0x6f, 0xbc, 0xa8, 0x04, 0x66, 0xf6, 0x58, 0x0f, 0xb7, 0x83, 0x19, 0x79, 0xb8, 0xfb, 0x68, 0x03,
	0xf2, 0xd9, 0x06, 0xe4, 0xab, 0x0d, 0x8e, 0xde, 0xbf, 0x03, 0xf2, 0x12, 0xef, 0xf5, 0xb6, 0xe5,
	0xd6, 0xc8, 0x5a, 0x1b, 0x99, 0x4b, 0xf5, 0x07, 0x4a, 0x24, 0xe5, 0x26, 0x4f, 0xca, 0x2c, 0xc1,
	0xdc, 0xd9, 0x31, 0xf6, 0x77, 0xf3, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x00, 0x37, 0x22, 0x84, 0x84,
	0x01, 0x00, 0x00,
}

func (m *ExtraInfo) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtraInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtraInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ClientAddr) > 0 {
		i -= len(m.ClientAddr)
		copy(dAtA[i:], m.ClientAddr)
		i = encodeVarintProxy(dAtA, i, uint64(len(m.ClientAddr)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Label) > 0 {
		for k := range m.Label {
			v := m.Label[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintProxy(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintProxy(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintProxy(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.ConnectionID != 0 {
		i = encodeVarintProxy(dAtA, i, uint64(m.ConnectionID))
		i--
		dAtA[i] = 0x18
	}
	if m.InternalConn {
		i--
		if m.InternalConn {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Salt) > 0 {
		i -= len(m.Salt)
		copy(dAtA[i:], m.Salt)
		i = encodeVarintProxy(dAtA, i, uint64(len(m.Salt)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProxy(dAtA []byte, offset int, v uint64) int {
	offset -= sovProxy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtraInfo) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Salt)
	if l > 0 {
		n += 1 + l + sovProxy(uint64(l))
	}
	if m.InternalConn {
		n += 2
	}
	if m.ConnectionID != 0 {
		n += 1 + sovProxy(uint64(m.ConnectionID))
	}
	if len(m.Label) > 0 {
		for k, v := range m.Label {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovProxy(uint64(len(k))) + 1 + len(v) + sovProxy(uint64(len(v)))
			n += mapEntrySize + 1 + sovProxy(uint64(mapEntrySize))
		}
	}
	l = len(m.ClientAddr)
	if l > 0 {
		n += 1 + l + sovProxy(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProxy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProxy(x uint64) (n int) {
	return sovProxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtraInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProxy
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
			return fmt.Errorf("proto: ExtraInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtraInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProxy
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salt = append(m.Salt[:0], dAtA[iNdEx:postIndex]...)
			if m.Salt == nil {
				m.Salt = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalConn", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InternalConn = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionID", wireType)
			}
			m.ConnectionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConnectionID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProxy
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
				return ErrInvalidLengthProxy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Label == nil {
				m.Label = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProxy
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProxy
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthProxy
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthProxy
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProxy
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthProxy
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthProxy
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipProxy(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthProxy
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Label[mapkey] = mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProxy
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
				return ErrInvalidLengthProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProxy
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
func skipProxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProxy
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
					return 0, ErrIntOverflowProxy
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
					return 0, ErrIntOverflowProxy
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
				return 0, ErrInvalidLengthProxy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProxy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProxy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProxy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProxy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProxy = fmt.Errorf("proto: unexpected end of group")
)
