// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/rollapp/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgCreateRollapp creates a new rollapp chain on the hub.
type MsgCreateRollapp struct {
	// creator is the bech32-encoded address of the rollapp creator
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// rollappId is the unique identifier of the rollapp chain.
	// The rollapp_id follows the same standard as cosmos chain_id
	RollappId string `protobuf:"bytes,2,opt,name=rollapp_id,json=rollappId,proto3" json:"rollapp_id,omitempty"`
	// initial_sequencer_address is a bech32-encoded address of the
	// sequencer that is allowed to initially serve this rollappId.
	InitialSequencerAddress string `protobuf:"bytes,3,opt,name=initial_sequencer_address,json=initialSequencerAddress,proto3" json:"initial_sequencer_address,omitempty"`
	// the unique rollapp address bech32 prefix.
	Bech32Prefix string `protobuf:"bytes,4,opt,name=bech32_prefix,json=bech32Prefix,proto3" json:"bech32_prefix,omitempty"`
	// genesis doc identifier
	GenesisInfo GenesisInfo `protobuf:"bytes,5,opt,name=genesis_info,json=genesisInfo,proto3" json:"genesis_info"`
}

func (m *MsgCreateRollapp) Reset()         { *m = MsgCreateRollapp{} }
func (m *MsgCreateRollapp) String() string { return proto.CompactTextString(m) }
func (*MsgCreateRollapp) ProtoMessage()    {}
func (*MsgCreateRollapp) Descriptor() ([]byte, []int) {
	return fileDescriptor_935cc363af28220c, []int{0}
}
func (m *MsgCreateRollapp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateRollapp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateRollapp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateRollapp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateRollapp.Merge(m, src)
}
func (m *MsgCreateRollapp) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateRollapp) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateRollapp.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateRollapp proto.InternalMessageInfo

func (m *MsgCreateRollapp) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateRollapp) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *MsgCreateRollapp) GetInitialSequencerAddress() string {
	if m != nil {
		return m.InitialSequencerAddress
	}
	return ""
}

func (m *MsgCreateRollapp) GetBech32Prefix() string {
	if m != nil {
		return m.Bech32Prefix
	}
	return ""
}

func (m *MsgCreateRollapp) GetGenesisInfo() GenesisInfo {
	if m != nil {
		return m.GenesisInfo
	}
	return GenesisInfo{}
}

type MsgCreateRollappResponse struct {
}

func (m *MsgCreateRollappResponse) Reset()         { *m = MsgCreateRollappResponse{} }
func (m *MsgCreateRollappResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateRollappResponse) ProtoMessage()    {}
func (*MsgCreateRollappResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_935cc363af28220c, []int{1}
}
func (m *MsgCreateRollappResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateRollappResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateRollappResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateRollappResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateRollappResponse.Merge(m, src)
}
func (m *MsgCreateRollappResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateRollappResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateRollappResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateRollappResponse proto.InternalMessageInfo

// MsgUpdateState updates a rollapp state with a block batch.
// a block batch is a list of ordered blocks (by height)
type MsgUpdateState struct {
	// creator is the bech32-encoded address of the sequencer sending the update
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// rollappId is the rollapp that the sequencer belongs to and asking to update
	// The rollappId follows the same standard as cosmos chain_id
	RollappId string `protobuf:"bytes,2,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// startHeight is the block height of the first block in the batch
	StartHeight uint64 `protobuf:"varint,3,opt,name=startHeight,proto3" json:"startHeight,omitempty"`
	// numBlocks is the number of blocks included in this batch update
	NumBlocks uint64 `protobuf:"varint,4,opt,name=numBlocks,proto3" json:"numBlocks,omitempty"`
	// DAPath is the description of the location on the DA layer
	DAPath string `protobuf:"bytes,5,opt,name=DAPath,proto3" json:"DAPath,omitempty"`
	// version is the version of the rollapp
	Version uint64 `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	// BDs is a list of block description objects (one per block)
	// the list must be ordered by height, starting from startHeight to startHeight+numBlocks-1
	BDs BlockDescriptors `protobuf:"bytes,7,opt,name=BDs,proto3" json:"BDs"`
}

func (m *MsgUpdateState) Reset()         { *m = MsgUpdateState{} }
func (m *MsgUpdateState) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateState) ProtoMessage()    {}
func (*MsgUpdateState) Descriptor() ([]byte, []int) {
	return fileDescriptor_935cc363af28220c, []int{2}
}
func (m *MsgUpdateState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateState.Merge(m, src)
}
func (m *MsgUpdateState) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateState) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateState.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateState proto.InternalMessageInfo

func (m *MsgUpdateState) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateState) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *MsgUpdateState) GetStartHeight() uint64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *MsgUpdateState) GetNumBlocks() uint64 {
	if m != nil {
		return m.NumBlocks
	}
	return 0
}

func (m *MsgUpdateState) GetDAPath() string {
	if m != nil {
		return m.DAPath
	}
	return ""
}

func (m *MsgUpdateState) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *MsgUpdateState) GetBDs() BlockDescriptors {
	if m != nil {
		return m.BDs
	}
	return BlockDescriptors{}
}

type MsgUpdateStateResponse struct {
}

func (m *MsgUpdateStateResponse) Reset()         { *m = MsgUpdateStateResponse{} }
func (m *MsgUpdateStateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateStateResponse) ProtoMessage()    {}
func (*MsgUpdateStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_935cc363af28220c, []int{3}
}
func (m *MsgUpdateStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateStateResponse.Merge(m, src)
}
func (m *MsgUpdateStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateStateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateRollapp)(nil), "dymensionxyz.dymension.rollapp.MsgCreateRollapp")
	proto.RegisterType((*MsgCreateRollappResponse)(nil), "dymensionxyz.dymension.rollapp.MsgCreateRollappResponse")
	proto.RegisterType((*MsgUpdateState)(nil), "dymensionxyz.dymension.rollapp.MsgUpdateState")
	proto.RegisterType((*MsgUpdateStateResponse)(nil), "dymensionxyz.dymension.rollapp.MsgUpdateStateResponse")
}

func init() { proto.RegisterFile("dymension/rollapp/tx.proto", fileDescriptor_935cc363af28220c) }

var fileDescriptor_935cc363af28220c = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0xb6, 0x74, 0xaa, 0xbb, 0x21, 0x64, 0xa1, 0x61, 0x22, 0xc8, 0xaa, 0x72, 0xa9,
	0x84, 0x94, 0xa0, 0x16, 0x21, 0xc4, 0x6d, 0xa5, 0x12, 0xdb, 0xa1, 0x68, 0xca, 0xe0, 0xc2, 0x25,
	0x4a, 0x93, 0xd7, 0xd4, 0xa2, 0xb5, 0x83, 0xed, 0x4e, 0x2d, 0xdc, 0x10, 0x1f, 0x80, 0x8f, 0xb5,
	0xe3, 0x8e, 0x9c, 0x10, 0x6a, 0x3f, 0x05, 0x37, 0x14, 0x27, 0x69, 0xba, 0x55, 0xda, 0xe8, 0xa9,
	0x7e, 0xff, 0xf7, 0x7f, 0xcf, 0x7e, 0x3f, 0xbb, 0x41, 0x66, 0xb8, 0x98, 0x02, 0x93, 0x94, 0x33,
	0x47, 0xf0, 0xc9, 0xc4, 0x8f, 0x63, 0x47, 0xcd, 0xed, 0x58, 0x70, 0xc5, 0xb1, 0xb5, 0xce, 0xcd,
	0x17, 0x5f, 0xed, 0x75, 0x60, 0x67, 0x46, 0xb3, 0xbd, 0x5d, 0x3b, 0x9c, 0xf0, 0xe0, 0xb3, 0x17,
	0x82, 0x0c, 0x04, 0x8d, 0x15, 0x17, 0x69, 0x27, 0xf3, 0x68, 0xdb, 0x99, 0xfd, 0x66, 0x86, 0x87,
	0x11, 0x8f, 0xb8, 0x5e, 0x3a, 0xc9, 0x2a, 0x55, 0x5b, 0xdf, 0xcb, 0xe8, 0xc1, 0x40, 0x46, 0x6f,
	0x05, 0xf8, 0x0a, 0xdc, 0xb4, 0x00, 0x13, 0xb4, 0x17, 0x24, 0x02, 0x17, 0xc4, 0x68, 0x1a, 0xed,
	0xba, 0x9b, 0x87, 0xf8, 0x29, 0x42, 0x59, 0x57, 0x8f, 0x86, 0xa4, 0xac, 0x93, 0xf5, 0x4c, 0x39,
	0x0d, 0xf1, 0x1b, 0xf4, 0x98, 0x32, 0xaa, 0xa8, 0x3f, 0xf1, 0x24, 0x7c, 0x99, 0x01, 0x0b, 0x40,
	0x78, 0x7e, 0x18, 0x0a, 0x90, 0x92, 0x54, 0xb4, 0xfb, 0x51, 0x66, 0x38, 0xcf, 0xf3, 0xc7, 0x69,
	0x1a, 0x3f, 0x43, 0x07, 0x43, 0x08, 0xc6, 0xdd, 0x8e, 0x17, 0x0b, 0x18, 0xd1, 0x39, 0xa9, 0x6a,
	0xff, 0x7e, 0x2a, 0x9e, 0x69, 0x0d, 0x7f, 0x40, 0xfb, 0x11, 0x30, 0x90, 0x54, 0x7a, 0x94, 0x8d,
	0x38, 0xb9, 0xd7, 0x34, 0xda, 0x8d, 0xce, 0x73, 0xfb, 0x76, 0x8c, 0xf6, 0xbb, 0xb4, 0xe6, 0x94,
	0x8d, 0x78, 0xaf, 0x7a, 0xf9, 0xfb, 0xa8, 0xe4, 0x36, 0xa2, 0x42, 0x6a, 0x99, 0x88, 0xdc, 0x64,
	0xe0, 0x82, 0x8c, 0x39, 0x93, 0xd0, 0xfa, 0x51, 0x46, 0xf7, 0x07, 0x32, 0xfa, 0x18, 0x87, 0xbe,
	0x82, 0x73, 0xe5, 0x2b, 0xb8, 0x05, 0xcf, 0x13, 0x54, 0xc0, 0xd8, 0xa6, 0xd3, 0x44, 0x0d, 0xa9,
	0x7c, 0xa1, 0x4e, 0x80, 0x46, 0x63, 0xa5, 0x79, 0x54, 0xdd, 0x4d, 0x29, 0xa9, 0x67, 0xb3, 0x69,
	0x2f, 0xb9, 0x61, 0xa9, 0xe7, 0xaf, 0xba, 0x85, 0x80, 0x0f, 0x51, 0xad, 0x7f, 0x7c, 0xe6, 0xab,
	0xb1, 0x1e, 0xbb, 0xee, 0x66, 0x51, 0x72, 0x9e, 0x0b, 0x10, 0xc9, 0xc0, 0xa4, 0xa6, 0x6b, 0xf2,
	0x10, 0x9f, 0xa0, 0x4a, 0xaf, 0x2f, 0xc9, 0x9e, 0xa6, 0xf4, 0xe2, 0x2e, 0x4a, 0x7a, 0x9b, 0xfe,
	0xfa, 0x61, 0xc9, 0x0c, 0x55, 0xd2, 0xa2, 0x45, 0xd0, 0xe1, 0x75, 0x0a, 0x39, 0xa0, 0xce, 0x5f,
	0x03, 0x55, 0x06, 0x32, 0xc2, 0xdf, 0xd0, 0xc1, 0xf5, 0x57, 0x74, 0xe7, 0x7e, 0x37, 0x99, 0x9b,
	0xaf, 0x77, 0xad, 0xc8, 0x0f, 0x81, 0x67, 0xa8, 0xb1, 0x79, 0x43, 0xf6, 0x7f, 0x34, 0xda, 0xf0,
	0x9b, 0xaf, 0x76, 0xf3, 0xe7, 0xdb, 0xf6, 0xde, 0x5f, 0x2e, 0x2d, 0xe3, 0x6a, 0x69, 0x19, 0x7f,
	0x96, 0x96, 0xf1, 0x73, 0x65, 0x95, 0xae, 0x56, 0x56, 0xe9, 0xd7, 0xca, 0x2a, 0x7d, 0x7a, 0x19,
	0x51, 0x35, 0x9e, 0x0d, 0xed, 0x80, 0x4f, 0x9d, 0xcd, 0xde, 0x45, 0xe0, 0x5c, 0x74, 0x9d, 0x79,
	0xf1, 0x45, 0x58, 0xc4, 0x20, 0x87, 0x35, 0xfd, 0xa7, 0xec, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0x54, 0x6d, 0xcc, 0x4d, 0x33, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateRollapp(ctx context.Context, in *MsgCreateRollapp, opts ...grpc.CallOption) (*MsgCreateRollappResponse, error)
	UpdateState(ctx context.Context, in *MsgUpdateState, opts ...grpc.CallOption) (*MsgUpdateStateResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateRollapp(ctx context.Context, in *MsgCreateRollapp, opts ...grpc.CallOption) (*MsgCreateRollappResponse, error) {
	out := new(MsgCreateRollappResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.rollapp.Msg/CreateRollapp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateState(ctx context.Context, in *MsgUpdateState, opts ...grpc.CallOption) (*MsgUpdateStateResponse, error) {
	out := new(MsgUpdateStateResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.rollapp.Msg/UpdateState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateRollapp(context.Context, *MsgCreateRollapp) (*MsgCreateRollappResponse, error)
	UpdateState(context.Context, *MsgUpdateState) (*MsgUpdateStateResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateRollapp(ctx context.Context, req *MsgCreateRollapp) (*MsgCreateRollappResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRollapp not implemented")
}
func (*UnimplementedMsgServer) UpdateState(ctx context.Context, req *MsgUpdateState) (*MsgUpdateStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateState not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateRollapp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateRollapp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateRollapp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.rollapp.Msg/CreateRollapp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateRollapp(ctx, req.(*MsgCreateRollapp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.rollapp.Msg/UpdateState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateState(ctx, req.(*MsgUpdateState))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dymensionxyz.dymension.rollapp.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRollapp",
			Handler:    _Msg_CreateRollapp_Handler,
		},
		{
			MethodName: "UpdateState",
			Handler:    _Msg_UpdateState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dymension/rollapp/tx.proto",
}

func (m *MsgCreateRollapp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateRollapp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateRollapp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.GenesisInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Bech32Prefix) > 0 {
		i -= len(m.Bech32Prefix)
		copy(dAtA[i:], m.Bech32Prefix)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Bech32Prefix)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.InitialSequencerAddress) > 0 {
		i -= len(m.InitialSequencerAddress)
		copy(dAtA[i:], m.InitialSequencerAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.InitialSequencerAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateRollappResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateRollappResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateRollappResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BDs.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.Version != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x30
	}
	if len(m.DAPath) > 0 {
		i -= len(m.DAPath)
		copy(dAtA[i:], m.DAPath)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DAPath)))
		i--
		dAtA[i] = 0x2a
	}
	if m.NumBlocks != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.NumBlocks))
		i--
		dAtA[i] = 0x20
	}
	if m.StartHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateRollapp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.InitialSequencerAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Bech32Prefix)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.GenesisInfo.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgCreateRollappResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.StartHeight != 0 {
		n += 1 + sovTx(uint64(m.StartHeight))
	}
	if m.NumBlocks != 0 {
		n += 1 + sovTx(uint64(m.NumBlocks))
	}
	l = len(m.DAPath)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovTx(uint64(m.Version))
	}
	l = m.BDs.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateRollapp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateRollapp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateRollapp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialSequencerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialSequencerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bech32Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bech32Prefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateRollappResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateRollappResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateRollappResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumBlocks", wireType)
			}
			m.NumBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumBlocks |= uint64(b&0x7F) << shift
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
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DAPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BDs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BDs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
