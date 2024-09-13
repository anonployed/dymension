// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/rollapp/block_descriptor.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

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
	// timestamp is the time from the block header
	Timestamp time.Time `protobuf:"bytes,3,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
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

func (m *BlockDescriptor) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
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
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4e, 0xc2, 0x30,
	0x18, 0xc7, 0x57, 0x20, 0x44, 0x8a, 0x89, 0x66, 0x31, 0x66, 0x21, 0xa6, 0x2c, 0x9c, 0x76, 0x6a,
	0x13, 0xd0, 0x17, 0x68, 0xf0, 0xea, 0x61, 0xf1, 0xa2, 0x17, 0xc3, 0xa0, 0x96, 0xc5, 0x8d, 0xaf,
	0x59, 0x8b, 0x61, 0xbe, 0x82, 0x17, 0x1e, 0x8b, 0x23, 0x47, 0x4f, 0x6a, 0xb6, 0x17, 0x31, 0xdb,
	0x60, 0x18, 0x12, 0xbd, 0xf5, 0xdf, 0xfc, 0x7f, 0xfd, 0xb5, 0xfd, 0xf0, 0xcd, 0x2c, 0x8d, 0xc5,
	0x42, 0x87, 0xb0, 0x58, 0xa5, 0x6f, 0xac, 0x0e, 0x2c, 0x81, 0x28, 0x9a, 0x28, 0xc5, 0x82, 0x08,
	0xa6, 0x2f, 0x4f, 0x33, 0xa1, 0xa7, 0x49, 0xa8, 0x0c, 0x24, 0x54, 0x25, 0x60, 0xc0, 0x26, 0xbf,
	0x31, 0x5a, 0x07, 0xba, 0xc3, 0x7a, 0x17, 0x12, 0x24, 0x94, 0x55, 0x56, 0xac, 0x2a, 0xaa, 0xd7,
	0x97, 0x00, 0x32, 0x12, 0xac, 0x4c, 0xc1, 0xf2, 0x99, 0x99, 0x30, 0x16, 0xda, 0x4c, 0x62, 0x55,
	0x15, 0x06, 0xef, 0x08, 0x9f, 0xf1, 0xc2, 0x38, 0xae, 0x85, 0xf6, 0x25, 0x6e, 0xcf, 0x45, 0x28,
	0xe7, 0xc6, 0x41, 0x2e, 0xf2, 0x5a, 0xfe, 0x2e, 0xd9, 0x57, 0xb8, 0xa3, 0xcd, 0xc4, 0x08, 0x1f,
	0xc0, 0x38, 0x0d, 0x17, 0x79, 0xa7, 0xfe, 0x61, 0xc3, 0xe6, 0xb8, 0x53, 0x1f, 0xee, 0x34, 0x5d,
	0xe4, 0x75, 0x87, 0x3d, 0x5a, 0xe9, 0xe9, 0x5e, 0x4f, 0xef, 0xf7, 0x0d, 0x7e, 0xb2, 0xf9, 0xec,
	0x5b, 0xeb, 0xaf, 0x3e, 0xf2, 0x0f, 0xd8, 0xe0, 0x01, 0x9f, 0x1f, 0x5d, 0x46, 0xdb, 0xb7, 0xb8,
	0xc1, 0xc7, 0x0e, 0x72, 0x9b, 0x5e, 0x77, 0xc8, 0xe8, 0xff, 0xbf, 0x40, 0x8f, 0x68, 0xde, 0x2a,
	0x2c, 0x7e, 0x83, 0x8f, 0xf9, 0xdd, 0x26, 0x23, 0x68, 0x9b, 0x11, 0xf4, 0x9d, 0x11, 0xb4, 0xce,
	0x89, 0xb5, 0xcd, 0x89, 0xf5, 0x91, 0x13, 0xeb, 0xf1, 0x5a, 0x86, 0x66, 0xbe, 0x0c, 0xe8, 0x14,
	0x62, 0xf6, 0xc7, 0x6c, 0x5e, 0x47, 0x6c, 0x55, 0x0f, 0xc8, 0xa4, 0x4a, 0xe8, 0xa0, 0x5d, 0xbe,
	0x69, 0xf4, 0x13, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x62, 0xb7, 0x2c, 0xcf, 0x01, 0x00, 0x00,
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
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintBlockDescriptor(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
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
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovBlockDescriptor(uint64(l))
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
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
