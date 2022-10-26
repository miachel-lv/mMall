// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: userModels.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UserModel struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	CreateAt             int64    `protobuf:"varint,3,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	UpdateAt             int64    `protobuf:"varint,4,opt,name=UpdateAt,proto3" json:"UpdateAt,omitempty"`
	DeleteAt             int64    `protobuf:"varint,5,opt,name=DeleteAt,proto3" json:"DeleteAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserModel) Reset()         { *m = UserModel{} }
func (m *UserModel) String() string { return proto.CompactTextString(m) }
func (*UserModel) ProtoMessage()    {}
func (*UserModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab4681ca824f7c7a, []int{0}
}
func (m *UserModel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserModel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserModel.Merge(m, src)
}
func (m *UserModel) XXX_Size() int {
	return m.Size()
}
func (m *UserModel) XXX_DiscardUnknown() {
	xxx_messageInfo_UserModel.DiscardUnknown(m)
}

var xxx_messageInfo_UserModel proto.InternalMessageInfo

func (m *UserModel) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *UserModel) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserModel) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (m *UserModel) GetUpdateAt() int64 {
	if m != nil {
		return m.UpdateAt
	}
	return 0
}

func (m *UserModel) GetDeleteAt() int64 {
	if m != nil {
		return m.DeleteAt
	}
	return 0
}

func init() {
	proto.RegisterType((*UserModel)(nil), "services.UserModel")
}

func init() { proto.RegisterFile("userModels.proto", fileDescriptor_ab4681ca824f7c7a) }

var fileDescriptor_ab4681ca824f7c7a = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2d, 0x4e, 0x2d,
	0xf2, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x4e,
	0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0x56, 0xea, 0x66, 0xe4, 0xe2, 0x0c, 0x85, 0x49, 0x0b, 0xf1,
	0x71, 0x31, 0x79, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x06, 0x31, 0x79, 0xba, 0x08, 0x49,
	0x71, 0x71, 0x80, 0x24, 0xfd, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xe0,
	0x7c, 0x90, 0x9c, 0x73, 0x51, 0x6a, 0x62, 0x49, 0xaa, 0x63, 0x89, 0x04, 0xb3, 0x02, 0xa3, 0x06,
	0x73, 0x10, 0x9c, 0x0f, 0xd6, 0x57, 0x90, 0x02, 0x91, 0x63, 0x81, 0xc8, 0xc1, 0xf8, 0x20, 0x39,
	0x97, 0xd4, 0x9c, 0x54, 0xb0, 0x1c, 0x2b, 0x44, 0x0e, 0xc6, 0x77, 0x92, 0x3e, 0xf1, 0x48, 0x8e,
	0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x88, 0xe2, 0xd4,
	0xd3, 0xb7, 0x06, 0xbb, 0xb9, 0x38, 0x89, 0x0d, 0x4c, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xb4, 0xc1, 0x3e, 0xe1, 0xcf, 0x00, 0x00, 0x00,
}

func (m *UserModel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserModel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserModel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.DeleteAt != 0 {
		i = encodeVarintUserModels(dAtA, i, uint64(m.DeleteAt))
		i--
		dAtA[i] = 0x28
	}
	if m.UpdateAt != 0 {
		i = encodeVarintUserModels(dAtA, i, uint64(m.UpdateAt))
		i--
		dAtA[i] = 0x20
	}
	if m.CreateAt != 0 {
		i = encodeVarintUserModels(dAtA, i, uint64(m.CreateAt))
		i--
		dAtA[i] = 0x18
	}
	if len(m.UserName) > 0 {
		i -= len(m.UserName)
		copy(dAtA[i:], m.UserName)
		i = encodeVarintUserModels(dAtA, i, uint64(len(m.UserName)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintUserModels(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintUserModels(dAtA []byte, offset int, v uint64) int {
	offset -= sovUserModels(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserModel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovUserModels(uint64(m.ID))
	}
	l = len(m.UserName)
	if l > 0 {
		n += 1 + l + sovUserModels(uint64(l))
	}
	if m.CreateAt != 0 {
		n += 1 + sovUserModels(uint64(m.CreateAt))
	}
	if m.UpdateAt != 0 {
		n += 1 + sovUserModels(uint64(m.UpdateAt))
	}
	if m.DeleteAt != 0 {
		n += 1 + sovUserModels(uint64(m.DeleteAt))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovUserModels(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUserModels(x uint64) (n int) {
	return sovUserModels(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserModel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserModels
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
			return fmt.Errorf("proto: UserModel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserModel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserModels
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
				return ErrInvalidLengthUserModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateAt", wireType)
			}
			m.CreateAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateAt", wireType)
			}
			m.UpdateAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeleteAt", wireType)
			}
			m.DeleteAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeleteAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUserModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserModels
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
func skipUserModels(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUserModels
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
					return 0, ErrIntOverflowUserModels
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
					return 0, ErrIntOverflowUserModels
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
				return 0, ErrInvalidLengthUserModels
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUserModels
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUserModels
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUserModels        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUserModels          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUserModels = fmt.Errorf("proto: unexpected end of group")
)
