// COPIED FROM main

package types

import (
	"fmt"
	"io"
	"math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	"github.com/cosmos/gogoproto/proto"
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

// BlockDescriptor defines a single rollapp chain block description.
type BlockDescriptor struct {
	// height is the height of the block
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// stateRoot is a 32 byte array of the hash of the block (state root of the block)
	StateRoot []byte `protobuf:"bytes,2,opt,name=stateRoot,proto3" json:"stateRoot,omitempty"`
}

func (m *BlockDescriptor) Reset()         { *m = BlockDescriptor{} }
func (m *BlockDescriptor) String() string { return proto.CompactTextString(m) }
func (*BlockDescriptor) ProtoMessage()    {}
func (*BlockDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb4c1d0c21c2e68, []int{0}
}
func (m *BlockDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockDescriptor.Merge(m, src)
}
func (m *BlockDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *BlockDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_BlockDescriptor proto.InternalMessageInfo

func (m *BlockDescriptor) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockDescriptor) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

// BlockDescriptors defines list of BlockDescriptor.
type BlockDescriptors struct {
	BD []BlockDescriptor `protobuf:"bytes,1,rep,name=BD,proto3" json:"BD"`
}

func (m *BlockDescriptors) Reset()         { *m = BlockDescriptors{} }
func (m *BlockDescriptors) String() string { return proto.CompactTextString(m) }
func (*BlockDescriptors) ProtoMessage()    {}
func (*BlockDescriptors) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb4c1d0c21c2e68, []int{1}
}
func (m *BlockDescriptors) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockDescriptors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockDescriptors.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockDescriptors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockDescriptors.Merge(m, src)
}
func (m *BlockDescriptors) XXX_Size() int {
	return m.Size()
}
func (m *BlockDescriptors) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockDescriptors.DiscardUnknown(m)
}

var xxx_messageInfo_BlockDescriptors proto.InternalMessageInfo

func (m *BlockDescriptors) GetBD() []BlockDescriptor {
	if m != nil {
		return m.BD
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockDescriptor)(nil), "dymensionxyz.dymension.rollapp.BlockDescriptor")
	proto.RegisterType((*BlockDescriptors)(nil), "dymensionxyz.dymension.rollapp.BlockDescriptors")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/rollapp/block_descriptor.proto", fileDescriptor_6eb4c1d0c21c2e68)
}

var fileDescriptor_6eb4c1d0c21c2e68 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xa9, 0xcc, 0x4d,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0xab, 0xa8, 0xac, 0xd2, 0x87, 0x73, 0xf4, 0x8b, 0xf2, 0x73, 0x72,
	0x12, 0x0b, 0x0a, 0xf4, 0x93, 0x72, 0xf2, 0x93, 0xb3, 0xe3, 0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32,
	0x0b, 0x4a, 0xf2, 0x8b, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xe4, 0x90, 0xb5, 0xe9, 0xc1,
	0x39, 0x7a, 0x50, 0x6d, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xa5, 0xfa, 0x20, 0x16, 0x44,
	0x97, 0x92, 0x3b, 0x17, 0xbf, 0x13, 0xc8, 0x3c, 0x17, 0xb8, 0x71, 0x42, 0x62, 0x5c, 0x6c, 0x19,
	0xa9, 0x99, 0xe9, 0x19, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x50, 0x9e, 0x90, 0x0c,
	0x17, 0x67, 0x71, 0x49, 0x62, 0x49, 0x6a, 0x50, 0x7e, 0x7e, 0x89, 0x04, 0x93, 0x02, 0xa3, 0x06,
	0x4f, 0x10, 0x42, 0x40, 0x29, 0x92, 0x4b, 0x00, 0xcd, 0xa0, 0x62, 0x21, 0x57, 0x2e, 0x26, 0x27,
	0x17, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x7d, 0x3d, 0xfc, 0xee, 0xd3, 0x43, 0xd3, 0xed,
	0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10, 0x93, 0x93, 0x8b, 0x93, 0xdf, 0x89, 0x47, 0x72, 0x8c,
	0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72,
	0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x99, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0xe3, 0x08, 0xb5, 0x32, 0x63, 0xfd, 0x0a, 0x78, 0xd0, 0x95, 0x54, 0x16, 0xa4, 0x16,
	0x27, 0xb1, 0x81, 0xbd, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x40, 0x09, 0x0c, 0xda, 0x69,
	0x01, 0x00, 0x00,
}

func (m *BlockDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StateRoot) > 0 {
		i -= len(m.StateRoot)
		copy(dAtA[i:], m.StateRoot)
		i = encodeVarintBlockDescriptor(dAtA, i, uint64(len(m.StateRoot)))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintBlockDescriptor(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockDescriptors) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockDescriptors) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockDescriptors) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BD) > 0 {
		for iNdEx := len(m.BD) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BD[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlockDescriptor(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlockDescriptor(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlockDescriptor(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlockDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovBlockDescriptor(uint64(m.Height))
	}
	l = len(m.StateRoot)
	if l > 0 {
		n += 1 + l + sovBlockDescriptor(uint64(l))
	}
	return n
}

func (m *BlockDescriptors) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BD) > 0 {
		for _, e := range m.BD {
			l = e.Size()
			n += 1 + l + sovBlockDescriptor(uint64(l))
		}
	}
	return n
}

func sovBlockDescriptor(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlockDescriptor(x uint64) (n int) {
	return sovBlockDescriptor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockDescriptor
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
			return fmt.Errorf("proto: BlockDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockDescriptor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockDescriptor
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
				return ErrInvalidLengthBlockDescriptor
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockDescriptor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateRoot = append(m.StateRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.StateRoot == nil {
				m.StateRoot = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockDescriptor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockDescriptor
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockDescriptors) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockDescriptor
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
			return fmt.Errorf("proto: BlockDescriptors: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockDescriptors: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BD", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockDescriptor
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
				return ErrInvalidLengthBlockDescriptor
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlockDescriptor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BD = append(m.BD, BlockDescriptor{})
			if err := m.BD[len(m.BD)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockDescriptor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockDescriptor
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBlockDescriptor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlockDescriptor
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
					return 0, ErrIntOverflowBlockDescriptor
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
					return 0, ErrIntOverflowBlockDescriptor
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
				return 0, ErrInvalidLengthBlockDescriptor
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlockDescriptor
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlockDescriptor
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlockDescriptor        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlockDescriptor          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlockDescriptor = fmt.Errorf("proto: unexpected end of group")
)
