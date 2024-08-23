// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/rollapp/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the rollapp module's genesis state.
type GenesisState struct {
	Params                             Params                           `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	RollappList                        []Rollapp                        `protobuf:"bytes,2,rep,name=rollappList,proto3" json:"rollappList"`
	StateInfoList                      []StateInfo                      `protobuf:"bytes,3,rep,name=stateInfoList,proto3" json:"stateInfoList"`
	LatestStateInfoIndexList           []StateInfoIndex                 `protobuf:"bytes,4,rep,name=latestStateInfoIndexList,proto3" json:"latestStateInfoIndexList"`
	LatestFinalizedStateIndexList      []StateInfoIndex                 `protobuf:"bytes,5,rep,name=latestFinalizedStateIndexList,proto3" json:"latestFinalizedStateIndexList"`
	BlockHeightToFinalizationQueueList []BlockHeightToFinalizationQueue `protobuf:"bytes,6,rep,name=blockHeightToFinalizationQueueList,proto3" json:"blockHeightToFinalizationQueueList"`
	// LivenessEvents are scheduled upcoming liveness events
	LivenessEvents []LivenessEvent `protobuf:"bytes,7,rep,name=livenessEvents,proto3" json:"livenessEvents"`
	AppList        []App           `protobuf:"bytes,8,rep,name=appList,proto3" json:"appList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76890aebc09aa04, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetRollappList() []Rollapp {
	if m != nil {
		return m.RollappList
	}
	return nil
}

func (m *GenesisState) GetStateInfoList() []StateInfo {
	if m != nil {
		return m.StateInfoList
	}
	return nil
}

func (m *GenesisState) GetLatestStateInfoIndexList() []StateInfoIndex {
	if m != nil {
		return m.LatestStateInfoIndexList
	}
	return nil
}

func (m *GenesisState) GetLatestFinalizedStateIndexList() []StateInfoIndex {
	if m != nil {
		return m.LatestFinalizedStateIndexList
	}
	return nil
}

func (m *GenesisState) GetBlockHeightToFinalizationQueueList() []BlockHeightToFinalizationQueue {
	if m != nil {
		return m.BlockHeightToFinalizationQueueList
	}
	return nil
}

func (m *GenesisState) GetLivenessEvents() []LivenessEvent {
	if m != nil {
		return m.LivenessEvents
	}
	return nil
}

func (m *GenesisState) GetAppList() []App {
	if m != nil {
		return m.AppList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "dymensionxyz.dymension.rollapp.GenesisState")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/rollapp/genesis.proto", fileDescriptor_b76890aebc09aa04)
}

var fileDescriptor_b76890aebc09aa04 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6e, 0xda, 0x30,
	0x00, 0xc6, 0x93, 0xc1, 0xc2, 0x64, 0xb6, 0x1d, 0xac, 0x1d, 0x22, 0xa4, 0x65, 0x88, 0x49, 0x1b,
	0xd3, 0x46, 0x22, 0xc1, 0x76, 0x9d, 0x34, 0xf6, 0x17, 0x09, 0x6d, 0x2b, 0xb4, 0x97, 0xf6, 0x80,
	0x0c, 0x98, 0x60, 0x35, 0xd8, 0x51, 0x6c, 0x10, 0xf0, 0x14, 0x3d, 0xf4, 0x0d, 0xfa, 0x32, 0x1c,
	0x39, 0xf6, 0x54, 0x55, 0xf0, 0x22, 0x15, 0x8e, 0x13, 0xd1, 0x4a, 0x60, 0xa4, 0x9e, 0x1c, 0xcb,
	0xdf, 0xf7, 0xfb, 0x3e, 0x5b, 0x76, 0xc0, 0xa7, 0xfe, 0x6c, 0x84, 0x29, 0x27, 0x8c, 0x4e, 0x67,
	0x73, 0x2f, 0x9d, 0x78, 0x11, 0x0b, 0x02, 0x14, 0x86, 0x9e, 0x8f, 0x29, 0xe6, 0x84, 0xbb, 0x61,
	0xc4, 0x04, 0x83, 0xce, 0xb6, 0xda, 0x4d, 0x27, 0xae, 0x52, 0x17, 0x5e, 0xf9, 0xcc, 0x67, 0x52,
	0xea, 0x6d, 0xbe, 0x62, 0x57, 0xe1, 0xa3, 0x26, 0x23, 0x44, 0x11, 0x1a, 0xa9, 0x88, 0x82, 0xae,
	0x90, 0x1a, 0x95, 0xda, 0xd3, 0xa8, 0xb9, 0x40, 0x02, 0x77, 0x08, 0x1d, 0x24, 0x5d, 0xbe, 0x1c,
	0xb6, 0xdf, 0x8e, 0x88, 0x10, 0xe5, 0x03, 0x1c, 0x29, 0x5b, 0x45, 0x63, 0x0b, 0xc8, 0x64, 0x63,
	0x4c, 0x36, 0x51, 0xd6, 0xc8, 0xd3, 0x0d, 0x94, 0xae, 0x2c, 0xf0, 0xfc, 0x77, 0x9c, 0xd9, 0xde,
	0x74, 0x85, 0x3f, 0x80, 0x15, 0x9f, 0x87, 0x6d, 0x16, 0xcd, 0x72, 0xbe, 0xfa, 0xce, 0xdd, 0x7f,
	0xe6, 0xee, 0x7f, 0xa9, 0xae, 0x67, 0x17, 0x37, 0x6f, 0x8c, 0x96, 0xf2, 0xc2, 0x7f, 0x20, 0xaf,
	0xd6, 0x9b, 0x84, 0x0b, 0xfb, 0x49, 0x31, 0x53, 0xce, 0x57, 0xdf, 0xeb, 0x50, 0xad, 0x78, 0x54,
	0xac, 0x6d, 0x02, 0x3c, 0x01, 0x2f, 0xe4, 0x59, 0x36, 0xe8, 0x80, 0x49, 0x64, 0x46, 0x22, 0x3f,
	0xe8, 0x90, 0xed, 0xc4, 0xa4, 0xa0, 0xf7, 0x29, 0x30, 0x04, 0x76, 0x80, 0x04, 0xe6, 0x22, 0xd5,
	0x35, 0x68, 0x1f, 0x4f, 0x65, 0x42, 0x56, 0x26, 0xb8, 0x07, 0x27, 0x48, 0xa7, 0x8a, 0xd9, 0x49,
	0x85, 0x73, 0xf0, 0x3a, 0x5e, 0xfb, 0x45, 0x28, 0x0a, 0xc8, 0x1c, 0xf7, 0x95, 0x28, 0x89, 0x7d,
	0xfa, 0x88, 0xd8, 0xfd, 0x68, 0x78, 0x69, 0x82, 0x52, 0x37, 0x60, 0xbd, 0xf3, 0x3f, 0x98, 0xf8,
	0x43, 0x71, 0xcc, 0x94, 0x10, 0x09, 0xc2, 0xe8, 0xd1, 0x18, 0x8f, 0xb1, 0x6c, 0x60, 0xc9, 0x06,
	0x5f, 0x75, 0x0d, 0xea, 0x7b, 0x49, 0xaa, 0xd1, 0x01, 0x79, 0xf0, 0x0c, 0xbc, 0x4c, 0xee, 0xef,
	0xcf, 0x09, 0xa6, 0x82, 0xdb, 0x39, 0xd9, 0xa0, 0xa2, 0x6b, 0xd0, 0xdc, 0x76, 0xa9, 0xc0, 0x07,
	0x28, 0xf8, 0x1d, 0xe4, 0x92, 0x5b, 0xf8, 0x4c, 0x52, 0xdf, 0xea, 0xa8, 0xdf, 0xd2, 0x1b, 0x98,
	0x38, 0xeb, 0x7f, 0x17, 0x2b, 0xc7, 0x5c, 0xae, 0x1c, 0xf3, 0x76, 0xe5, 0x98, 0x17, 0x6b, 0xc7,
	0x58, 0xae, 0x1d, 0xe3, 0x7a, 0xed, 0x18, 0xa7, 0x9f, 0x7d, 0x22, 0x86, 0xe3, 0xae, 0xdb, 0x63,
	0xa3, 0x5d, 0xff, 0x82, 0x49, 0xcd, 0x9b, 0xa6, 0x2f, 0x4f, 0xcc, 0x42, 0xcc, 0xbb, 0x96, 0x7c,
	0x7c, 0xb5, 0xbb, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x26, 0x79, 0x86, 0xfe, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AppList) > 0 {
		for iNdEx := len(m.AppList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AppList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.LivenessEvents) > 0 {
		for iNdEx := len(m.LivenessEvents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LivenessEvents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.BlockHeightToFinalizationQueueList) > 0 {
		for iNdEx := len(m.BlockHeightToFinalizationQueueList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BlockHeightToFinalizationQueueList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.LatestFinalizedStateIndexList) > 0 {
		for iNdEx := len(m.LatestFinalizedStateIndexList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LatestFinalizedStateIndexList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.LatestStateInfoIndexList) > 0 {
		for iNdEx := len(m.LatestStateInfoIndexList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LatestStateInfoIndexList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.StateInfoList) > 0 {
		for iNdEx := len(m.StateInfoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StateInfoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RollappList) > 0 {
		for iNdEx := len(m.RollappList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RollappList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.RollappList) > 0 {
		for _, e := range m.RollappList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StateInfoList) > 0 {
		for _, e := range m.StateInfoList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LatestStateInfoIndexList) > 0 {
		for _, e := range m.LatestStateInfoIndexList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LatestFinalizedStateIndexList) > 0 {
		for _, e := range m.LatestFinalizedStateIndexList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BlockHeightToFinalizationQueueList) > 0 {
		for _, e := range m.BlockHeightToFinalizationQueueList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LivenessEvents) > 0 {
		for _, e := range m.LivenessEvents {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AppList) > 0 {
		for _, e := range m.AppList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappList = append(m.RollappList, Rollapp{})
			if err := m.RollappList[len(m.RollappList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateInfoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateInfoList = append(m.StateInfoList, StateInfo{})
			if err := m.StateInfoList[len(m.StateInfoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestStateInfoIndexList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LatestStateInfoIndexList = append(m.LatestStateInfoIndexList, StateInfoIndex{})
			if err := m.LatestStateInfoIndexList[len(m.LatestStateInfoIndexList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestFinalizedStateIndexList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LatestFinalizedStateIndexList = append(m.LatestFinalizedStateIndexList, StateInfoIndex{})
			if err := m.LatestFinalizedStateIndexList[len(m.LatestFinalizedStateIndexList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeightToFinalizationQueueList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHeightToFinalizationQueueList = append(m.BlockHeightToFinalizationQueueList, BlockHeightToFinalizationQueue{})
			if err := m.BlockHeightToFinalizationQueueList[len(m.BlockHeightToFinalizationQueueList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LivenessEvents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LivenessEvents = append(m.LivenessEvents, LivenessEvent{})
			if err := m.LivenessEvents[len(m.LivenessEvents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppList = append(m.AppList, App{})
			if err := m.AppList[len(m.AppList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
