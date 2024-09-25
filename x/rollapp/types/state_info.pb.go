// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/rollapp/state_info.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	types "github.com/dymensionxyz/dymension/v3/x/common/types"
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

// StateInfoIndex is the data used for indexing and retrieving a StateInfo
// it updated and saved with every UpdateState in StateInfo.
// We use the this structure also for:
// 1. LatestStateInfoIndex which defines the rollapps' current (latest) index of the last UpdateState
// 2. LatestFinalizedStateIndex which defines the rollapps' current (latest) index of the latest StateInfo that was finalized
type StateInfoIndex struct {
	// rollappId is the rollapp that the sequencer belongs to and asking to update
	// it used to identify the what rollapp a StateInfo belongs
	// The rollappId follows the same standard as cosmos chain_id
	RollappId string `protobuf:"bytes,1,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// index is a sequential increasing number, updating on each
	// state update used for indexing to a specific state info, the first index is 1
	Index uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *StateInfoIndex) Reset()         { *m = StateInfoIndex{} }
func (m *StateInfoIndex) String() string { return proto.CompactTextString(m) }
func (*StateInfoIndex) ProtoMessage()    {}
func (*StateInfoIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_750f3a9f16533ec4, []int{0}
}
func (m *StateInfoIndex) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateInfoIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StateInfoIndex.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StateInfoIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateInfoIndex.Merge(m, src)
}
func (m *StateInfoIndex) XXX_Size() int {
	return m.Size()
}
func (m *StateInfoIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_StateInfoIndex.DiscardUnknown(m)
}

var xxx_messageInfo_StateInfoIndex proto.InternalMessageInfo

func (m *StateInfoIndex) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *StateInfoIndex) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

// StateInfo defines a rollapps' state.
type StateInfo struct {
	// stateInfoIndex defines what rollapp the state belongs to
	// and in which index it can be referenced
	StateInfoIndex StateInfoIndex `protobuf:"bytes,1,opt,name=stateInfoIndex,proto3" json:"stateInfoIndex"`
	// sequencer is the bech32-encoded address of the sequencer sent the update
	Sequencer string `protobuf:"bytes,2,opt,name=sequencer,proto3" json:"sequencer,omitempty"`
	// startHeight is the block height of the first block in the batch
	StartHeight uint64 `protobuf:"varint,3,opt,name=startHeight,proto3" json:"startHeight,omitempty"`
	// DAPath is the description of the location on the DA layer
	DAPath string `protobuf:"bytes,5,opt,name=DAPath,proto3" json:"DAPath,omitempty"`
	// creationHeight is the height at which the UpdateState took place
	CreationHeight uint64 `protobuf:"varint,7,opt,name=creationHeight,proto3" json:"creationHeight,omitempty"`
	// status is the status of the state update
	Status types.Status `protobuf:"varint,8,opt,name=status,proto3,enum=dymensionxyz.dymension.common.Status" json:"status,omitempty"`
	// BDs is a list of block description objects (one per block)
	// the list must be ordered by height, starting from startHeight to startHeight+numBlocks-1
	BDs BlockDescriptors `protobuf:"bytes,9,opt,name=BDs,proto3" json:"BDs"`
	// created_at is the timestamp at which the StateInfo was created
	CreatedAt time.Time `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at" yaml:"created_at"`
	// DrsVersion is a DRS version used by the rollapp.
	DrsVersion string `protobuf:"bytes,11,opt,name=drs_version,json=drsVersion,proto3" json:"drs_version,omitempty"`
}

func (m *StateInfo) Reset()         { *m = StateInfo{} }
func (m *StateInfo) String() string { return proto.CompactTextString(m) }
func (*StateInfo) ProtoMessage()    {}
func (*StateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_750f3a9f16533ec4, []int{1}
}
func (m *StateInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StateInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateInfo.Merge(m, src)
}
func (m *StateInfo) XXX_Size() int {
	return m.Size()
}
func (m *StateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StateInfo proto.InternalMessageInfo

func (m *StateInfo) GetStateInfoIndex() StateInfoIndex {
	if m != nil {
		return m.StateInfoIndex
	}
	return StateInfoIndex{}
}

func (m *StateInfo) GetSequencer() string {
	if m != nil {
		return m.Sequencer
	}
	return ""
}

func (m *StateInfo) GetStartHeight() uint64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *StateInfo) GetDAPath() string {
	if m != nil {
		return m.DAPath
	}
	return ""
}

func (m *StateInfo) GetCreationHeight() uint64 {
	if m != nil {
		return m.CreationHeight
	}
	return 0
}

func (m *StateInfo) GetStatus() types.Status {
	if m != nil {
		return m.Status
	}
	return types.Status_PENDING
}

func (m *StateInfo) GetBDs() BlockDescriptors {
	if m != nil {
		return m.BDs
	}
	return BlockDescriptors{}
}

func (m *StateInfo) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *StateInfo) GetDrsVersion() string {
	if m != nil {
		return m.DrsVersion
	}
	return ""
}

// StateInfoSummary is a compact representation of StateInfo
type StateInfoSummary struct {
	// stateInfoIndex defines what rollapp the state belongs to
	// and in which index it can be referenced
	StateInfoIndex StateInfoIndex `protobuf:"bytes,1,opt,name=stateInfoIndex,proto3" json:"stateInfoIndex"`
	// status is the status of the state update
	Status types.Status `protobuf:"varint,2,opt,name=status,proto3,enum=dymensionxyz.dymension.common.Status" json:"status,omitempty"`
	// creationHeight is the height at which the UpdateState took place
	CreationHeight uint64 `protobuf:"varint,3,opt,name=creationHeight,proto3" json:"creationHeight,omitempty"`
}

func (m *StateInfoSummary) Reset()         { *m = StateInfoSummary{} }
func (m *StateInfoSummary) String() string { return proto.CompactTextString(m) }
func (*StateInfoSummary) ProtoMessage()    {}
func (*StateInfoSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_750f3a9f16533ec4, []int{2}
}
func (m *StateInfoSummary) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateInfoSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StateInfoSummary.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StateInfoSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateInfoSummary.Merge(m, src)
}
func (m *StateInfoSummary) XXX_Size() int {
	return m.Size()
}
func (m *StateInfoSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_StateInfoSummary.DiscardUnknown(m)
}

var xxx_messageInfo_StateInfoSummary proto.InternalMessageInfo

func (m *StateInfoSummary) GetStateInfoIndex() StateInfoIndex {
	if m != nil {
		return m.StateInfoIndex
	}
	return StateInfoIndex{}
}

func (m *StateInfoSummary) GetStatus() types.Status {
	if m != nil {
		return m.Status
	}
	return types.Status_PENDING
}

func (m *StateInfoSummary) GetCreationHeight() uint64 {
	if m != nil {
		return m.CreationHeight
	}
	return 0
}

// BlockHeightToFinalizationQueue defines a map from block height to list of states to finalized
type BlockHeightToFinalizationQueue struct {
	// creationHeight is the block height that the state should be finalized
	CreationHeight uint64 `protobuf:"varint,1,opt,name=creationHeight,proto3" json:"creationHeight,omitempty"`
	// finalizationQueue is a list of states that are waiting to be finalized
	// when the block height becomes creationHeight
	FinalizationQueue []StateInfoIndex `protobuf:"bytes,2,rep,name=finalizationQueue,proto3" json:"finalizationQueue"`
}

func (m *BlockHeightToFinalizationQueue) Reset()         { *m = BlockHeightToFinalizationQueue{} }
func (m *BlockHeightToFinalizationQueue) String() string { return proto.CompactTextString(m) }
func (*BlockHeightToFinalizationQueue) ProtoMessage()    {}
func (*BlockHeightToFinalizationQueue) Descriptor() ([]byte, []int) {
	return fileDescriptor_750f3a9f16533ec4, []int{3}
}
func (m *BlockHeightToFinalizationQueue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeightToFinalizationQueue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeightToFinalizationQueue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeightToFinalizationQueue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeightToFinalizationQueue.Merge(m, src)
}
func (m *BlockHeightToFinalizationQueue) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeightToFinalizationQueue) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeightToFinalizationQueue.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeightToFinalizationQueue proto.InternalMessageInfo

func (m *BlockHeightToFinalizationQueue) GetCreationHeight() uint64 {
	if m != nil {
		return m.CreationHeight
	}
	return 0
}

func (m *BlockHeightToFinalizationQueue) GetFinalizationQueue() []StateInfoIndex {
	if m != nil {
		return m.FinalizationQueue
	}
	return nil
}

func init() {
	proto.RegisterType((*StateInfoIndex)(nil), "dymensionxyz.dymension.rollapp.StateInfoIndex")
	proto.RegisterType((*StateInfo)(nil), "dymensionxyz.dymension.rollapp.StateInfo")
	proto.RegisterType((*StateInfoSummary)(nil), "dymensionxyz.dymension.rollapp.StateInfoSummary")
	proto.RegisterType((*BlockHeightToFinalizationQueue)(nil), "dymensionxyz.dymension.rollapp.BlockHeightToFinalizationQueue")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/rollapp/state_info.proto", fileDescriptor_750f3a9f16533ec4)
}

var fileDescriptor_750f3a9f16533ec4 = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6a, 0xdb, 0x40,
	0x18, 0xf4, 0xfa, 0x2f, 0xd6, 0x1a, 0x8c, 0xb3, 0x84, 0x22, 0x4c, 0x2b, 0x1b, 0x41, 0x8b, 0xe9,
	0x41, 0x2a, 0x49, 0x7b, 0x29, 0xf4, 0x10, 0x63, 0x4a, 0x9c, 0x43, 0x69, 0x95, 0x50, 0x4a, 0x29,
	0x98, 0xb5, 0xb5, 0x96, 0x97, 0x4a, 0x5a, 0x75, 0x77, 0x15, 0xec, 0x3c, 0x45, 0x1e, 0xa4, 0x0f,
	0x92, 0x63, 0x6e, 0xed, 0x29, 0x2d, 0xf6, 0x1b, 0xf4, 0xd6, 0x5b, 0xd1, 0x4a, 0xb1, 0x13, 0xc7,
	0x6e, 0x20, 0x90, 0x9b, 0xbe, 0x8f, 0x6f, 0x86, 0xd9, 0x99, 0x41, 0xd0, 0x76, 0xa7, 0x01, 0x09,
	0x05, 0x65, 0xe1, 0x64, 0x7a, 0xba, 0x1c, 0x6c, 0xce, 0x7c, 0x1f, 0x47, 0x91, 0x2d, 0x24, 0x96,
	0xa4, 0x4f, 0xc3, 0x11, 0xb3, 0x22, 0xce, 0x24, 0x43, 0xc6, 0x75, 0x80, 0xb5, 0x18, 0xac, 0x0c,
	0xd0, 0xd8, 0xf1, 0x98, 0xc7, 0xd4, 0xa9, 0x9d, 0x7c, 0xa5, 0xa8, 0x46, 0xd3, 0x63, 0xcc, 0xf3,
	0x89, 0xad, 0xa6, 0x41, 0x3c, 0xb2, 0x25, 0x0d, 0x88, 0x90, 0x38, 0x88, 0xb2, 0x83, 0x57, 0x77,
	0xe8, 0x18, 0xf8, 0x6c, 0xf8, 0xb5, 0xef, 0x12, 0x31, 0xe4, 0x34, 0x92, 0x8c, 0x67, 0xb0, 0xe7,
	0x1b, 0x60, 0x43, 0x16, 0x04, 0x2c, 0x54, 0xea, 0x63, 0x91, 0xde, 0x9a, 0x5d, 0x58, 0x3b, 0x4a,
	0x5e, 0xd3, 0x0b, 0x47, 0xac, 0x17, 0xba, 0x64, 0x82, 0x1e, 0x43, 0x2d, 0xe3, 0xef, 0xb9, 0x3a,
	0x68, 0x81, 0xb6, 0xe6, 0x2c, 0x17, 0x68, 0x07, 0x96, 0x68, 0x72, 0xa6, 0xe7, 0x5b, 0xa0, 0x5d,
	0x74, 0xd2, 0xc1, 0xfc, 0x5b, 0x80, 0xda, 0x82, 0x06, 0x7d, 0x81, 0x35, 0x71, 0x83, 0x53, 0xd1,
	0x54, 0x77, 0x2d, 0xeb, 0xff, 0x36, 0x59, 0x37, 0x95, 0x74, 0x8a, 0xe7, 0x97, 0xcd, 0x9c, 0xb3,
	0xc2, 0x95, 0xe8, 0x13, 0xe4, 0x5b, 0x4c, 0xc2, 0x21, 0xe1, 0x4a, 0x85, 0xe6, 0x2c, 0x17, 0xa8,
	0x05, 0xab, 0x42, 0x62, 0x2e, 0x0f, 0x08, 0xf5, 0xc6, 0x52, 0x2f, 0x28, 0x95, 0xd7, 0x57, 0xe8,
	0x11, 0x2c, 0x77, 0xf7, 0xdf, 0x63, 0x39, 0xd6, 0x4b, 0x0a, 0x9c, 0x4d, 0xe8, 0x19, 0xac, 0x0d,
	0x39, 0xc1, 0x92, 0xb2, 0x30, 0x03, 0x6f, 0x29, 0xf0, 0xca, 0x16, 0xbd, 0x81, 0xe5, 0xd4, 0x41,
	0xbd, 0xd2, 0x02, 0xed, 0xda, 0xee, 0xd3, 0x4d, 0xaf, 0x4a, 0xed, 0x56, 0x8f, 0x8a, 0x85, 0x93,
	0x81, 0xd0, 0x01, 0x2c, 0x74, 0xba, 0x42, 0xd7, 0x94, 0x23, 0x2f, 0xee, 0x72, 0xa4, 0x93, 0x24,
	0xdc, 0x5d, 0x04, 0x2c, 0x32, 0x4f, 0x12, 0x0a, 0xf4, 0x09, 0x42, 0x25, 0x8d, 0xb8, 0x7d, 0x2c,
	0x75, 0xa8, 0x08, 0x1b, 0x56, 0xda, 0x29, 0xeb, 0xaa, 0x53, 0xd6, 0xf1, 0x55, 0xa7, 0x3a, 0x4f,
	0x12, 0xe8, 0x9f, 0xcb, 0xe6, 0xf6, 0x14, 0x07, 0xfe, 0x6b, 0x73, 0x89, 0x35, 0xcf, 0x7e, 0x35,
	0x81, 0xa3, 0x65, 0x8b, 0x7d, 0x89, 0x9a, 0xb0, 0xea, 0x72, 0xd1, 0x3f, 0x21, 0x3c, 0x11, 0xa3,
	0x57, 0x95, 0x4f, 0xd0, 0xe5, 0xe2, 0x63, 0xba, 0x39, 0x2c, 0x56, 0x8a, 0xf5, 0xd2, 0x61, 0xb1,
	0x52, 0xae, 0x6f, 0x99, 0x3f, 0x00, 0xac, 0x2f, 0x82, 0x3b, 0x8a, 0x83, 0x00, 0xf3, 0xe9, 0x03,
	0x57, 0x60, 0x19, 0x41, 0xfe, 0x3e, 0x11, 0xdc, 0x4e, 0xba, 0xb0, 0x2e, 0x69, 0xf3, 0x3b, 0x80,
	0x86, 0x0a, 0x20, 0x9d, 0x8f, 0xd9, 0x5b, 0x1a, 0x62, 0x9f, 0x9e, 0xaa, 0x9b, 0x0f, 0x31, 0x89,
	0xc9, 0x1a, 0x2a, 0xb0, 0xb6, 0x34, 0x03, 0xb8, 0x3d, 0x5a, 0x05, 0xeb, 0xf9, 0x56, 0xe1, 0xde,
	0x96, 0xdc, 0xa6, 0xeb, 0xbc, 0x3b, 0x9f, 0x19, 0xe0, 0x62, 0x66, 0x80, 0xdf, 0x33, 0x03, 0x9c,
	0xcd, 0x8d, 0xdc, 0xc5, 0xdc, 0xc8, 0xfd, 0x9c, 0x1b, 0xb9, 0xcf, 0x2f, 0x3d, 0x2a, 0xc7, 0xf1,
	0x20, 0xb1, 0x63, 0xd3, 0xaf, 0xed, 0x64, 0xcf, 0x9e, 0x2c, 0xfe, 0x2b, 0x72, 0x1a, 0x11, 0x31,
	0x28, 0xab, 0x0e, 0xed, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x71, 0x15, 0xb3, 0x0e, 0x05,
	0x00, 0x00,
}

func (m *StateInfoIndex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateInfoIndex) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateInfoIndex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintStateInfo(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StateInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DrsVersion) > 0 {
		i -= len(m.DrsVersion)
		copy(dAtA[i:], m.DrsVersion)
		i = encodeVarintStateInfo(dAtA, i, uint64(len(m.DrsVersion)))
		i--
		dAtA[i] = 0x5a
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CreatedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintStateInfo(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x52
	{
		size, err := m.BDs.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStateInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.Status != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if m.CreationHeight != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.CreationHeight))
		i--
		dAtA[i] = 0x38
	}
	if len(m.DAPath) > 0 {
		i -= len(m.DAPath)
		copy(dAtA[i:], m.DAPath)
		i = encodeVarintStateInfo(dAtA, i, uint64(len(m.DAPath)))
		i--
		dAtA[i] = 0x2a
	}
	if m.StartHeight != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Sequencer) > 0 {
		i -= len(m.Sequencer)
		copy(dAtA[i:], m.Sequencer)
		i = encodeVarintStateInfo(dAtA, i, uint64(len(m.Sequencer)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.StateInfoIndex.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStateInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *StateInfoSummary) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateInfoSummary) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateInfoSummary) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreationHeight != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.CreationHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.Status != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.StateInfoIndex.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStateInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BlockHeightToFinalizationQueue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeightToFinalizationQueue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeightToFinalizationQueue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FinalizationQueue) > 0 {
		for iNdEx := len(m.FinalizationQueue) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FinalizationQueue[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStateInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CreationHeight != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.CreationHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStateInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovStateInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StateInfoIndex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovStateInfo(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovStateInfo(uint64(m.Index))
	}
	return n
}

func (m *StateInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.StateInfoIndex.Size()
	n += 1 + l + sovStateInfo(uint64(l))
	l = len(m.Sequencer)
	if l > 0 {
		n += 1 + l + sovStateInfo(uint64(l))
	}
	if m.StartHeight != 0 {
		n += 1 + sovStateInfo(uint64(m.StartHeight))
	}
	l = len(m.DAPath)
	if l > 0 {
		n += 1 + l + sovStateInfo(uint64(l))
	}
	if m.CreationHeight != 0 {
		n += 1 + sovStateInfo(uint64(m.CreationHeight))
	}
	if m.Status != 0 {
		n += 1 + sovStateInfo(uint64(m.Status))
	}
	l = m.BDs.Size()
	n += 1 + l + sovStateInfo(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovStateInfo(uint64(l))
	l = len(m.DrsVersion)
	if l > 0 {
		n += 1 + l + sovStateInfo(uint64(l))
	}
	return n
}

func (m *StateInfoSummary) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.StateInfoIndex.Size()
	n += 1 + l + sovStateInfo(uint64(l))
	if m.Status != 0 {
		n += 1 + sovStateInfo(uint64(m.Status))
	}
	if m.CreationHeight != 0 {
		n += 1 + sovStateInfo(uint64(m.CreationHeight))
	}
	return n
}

func (m *BlockHeightToFinalizationQueue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreationHeight != 0 {
		n += 1 + sovStateInfo(uint64(m.CreationHeight))
	}
	if len(m.FinalizationQueue) > 0 {
		for _, e := range m.FinalizationQueue {
			l = e.Size()
			n += 1 + l + sovStateInfo(uint64(l))
		}
	}
	return n
}

func sovStateInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStateInfo(x uint64) (n int) {
	return sovStateInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StateInfoIndex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateInfo
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
			return fmt.Errorf("proto: StateInfoIndex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateInfoIndex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
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
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStateInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStateInfo
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
func (m *StateInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateInfo
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
			return fmt.Errorf("proto: StateInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateInfoIndex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
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
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StateInfoIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequencer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
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
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sequencer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DAPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
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
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DAPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
			}
			m.CreationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BDs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
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
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BDs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
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
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DrsVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
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
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DrsVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStateInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStateInfo
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
func (m *StateInfoSummary) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateInfo
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
			return fmt.Errorf("proto: StateInfoSummary: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateInfoSummary: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateInfoIndex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
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
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StateInfoIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
			}
			m.CreationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStateInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStateInfo
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
func (m *BlockHeightToFinalizationQueue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateInfo
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
			return fmt.Errorf("proto: BlockHeightToFinalizationQueue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeightToFinalizationQueue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
			}
			m.CreationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizationQueue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
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
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FinalizationQueue = append(m.FinalizationQueue, StateInfoIndex{})
			if err := m.FinalizationQueue[len(m.FinalizationQueue)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStateInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStateInfo
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
func skipStateInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStateInfo
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
					return 0, ErrIntOverflowStateInfo
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
					return 0, ErrIntOverflowStateInfo
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
				return 0, ErrInvalidLengthStateInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStateInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStateInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStateInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStateInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStateInfo = fmt.Errorf("proto: unexpected end of group")
)
