// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bspchange.proto

package bsp_transaction

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type RequestViewChangeMessage struct {
	Timestamp            uint64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Endpoint             string   `protobuf:"bytes,2,opt,name=Endpoint,proto3" json:"Endpoint,omitempty"`
	PeerId               uint64   `protobuf:"varint,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	ViewNumber           uint64   `protobuf:"varint,4,opt,name=view_number,json=viewNumber,proto3" json:"view_number,omitempty"`
	CurrChannelId        string   `protobuf:"bytes,5,opt,name=CurrChannelId,proto3" json:"CurrChannelId,omitempty"`
	NextChannelId        string   `protobuf:"bytes,6,opt,name=NextChannelId,proto3" json:"NextChannelId,omitempty"`
	OldEndpoint          string   `protobuf:"bytes,7,opt,name=OldEndpoint,proto3" json:"OldEndpoint,omitempty"`
	PreparedBlockHeight  uint64   `protobuf:"varint,9,opt,name=PreparedBlockHeight,proto3" json:"PreparedBlockHeight,omitempty"`
	PreparedBlockHash    []byte   `protobuf:"bytes,11,opt,name=PreparedBlockHash,proto3" json:"PreparedBlockHash,omitempty"`
	ProofBlocks          [][]byte `protobuf:"bytes,12,rep,name=proof_blocks,json=proofBlocks,proto3" json:"proof_blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestViewChangeMessage) Reset()         { *m = RequestViewChangeMessage{} }
func (m *RequestViewChangeMessage) String() string { return proto.CompactTextString(m) }
func (*RequestViewChangeMessage) ProtoMessage()    {}
func (*RequestViewChangeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f00fded0f4841af, []int{0}
}

func (m *RequestViewChangeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestViewChangeMessage.Unmarshal(m, b)
}
func (m *RequestViewChangeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestViewChangeMessage.Marshal(b, m, deterministic)
}
func (m *RequestViewChangeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestViewChangeMessage.Merge(m, src)
}
func (m *RequestViewChangeMessage) XXX_Size() int {
	return xxx_messageInfo_RequestViewChangeMessage.Size(m)
}
func (m *RequestViewChangeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestViewChangeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RequestViewChangeMessage proto.InternalMessageInfo

func (m *RequestViewChangeMessage) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *RequestViewChangeMessage) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *RequestViewChangeMessage) GetPeerId() uint64 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *RequestViewChangeMessage) GetViewNumber() uint64 {
	if m != nil {
		return m.ViewNumber
	}
	return 0
}

func (m *RequestViewChangeMessage) GetCurrChannelId() string {
	if m != nil {
		return m.CurrChannelId
	}
	return ""
}

func (m *RequestViewChangeMessage) GetNextChannelId() string {
	if m != nil {
		return m.NextChannelId
	}
	return ""
}

func (m *RequestViewChangeMessage) GetOldEndpoint() string {
	if m != nil {
		return m.OldEndpoint
	}
	return ""
}

func (m *RequestViewChangeMessage) GetPreparedBlockHeight() uint64 {
	if m != nil {
		return m.PreparedBlockHeight
	}
	return 0
}

func (m *RequestViewChangeMessage) GetPreparedBlockHash() []byte {
	if m != nil {
		return m.PreparedBlockHash
	}
	return nil
}

func (m *RequestViewChangeMessage) GetProofBlocks() [][]byte {
	if m != nil {
		return m.ProofBlocks
	}
	return nil
}

type NewViewTx struct {
	Timestamp            uint64                      `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ViewNumber           uint64                      `protobuf:"varint,2,opt,name=view_number,json=viewNumber,proto3" json:"view_number,omitempty"`
	ChannelId            string                      `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Endpoint             string                      `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	OldBspEndpoint       string                      `protobuf:"bytes,5,opt,name=oldBspEndpoint,proto3" json:"oldBspEndpoint,omitempty"`
	StatedbHash          []byte                      `protobuf:"bytes,6,opt,name=statedb_hash,json=statedbHash,proto3" json:"statedb_hash,omitempty"`
	Height               uint64                      `protobuf:"varint,7,opt,name=height,proto3" json:"height,omitempty"`
	CurrentBlockHash     []byte                      `protobuf:"bytes,9,opt,name=current_block_hash,json=currentBlockHash,proto3" json:"current_block_hash,omitempty"`
	Vcm                  []*RequestViewChangeMessage `protobuf:"bytes,11,rep,name=vcm,proto3" json:"vcm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *NewViewTx) Reset()         { *m = NewViewTx{} }
func (m *NewViewTx) String() string { return proto.CompactTextString(m) }
func (*NewViewTx) ProtoMessage()    {}
func (*NewViewTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f00fded0f4841af, []int{1}
}

func (m *NewViewTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewViewTx.Unmarshal(m, b)
}
func (m *NewViewTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewViewTx.Marshal(b, m, deterministic)
}
func (m *NewViewTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewViewTx.Merge(m, src)
}
func (m *NewViewTx) XXX_Size() int {
	return xxx_messageInfo_NewViewTx.Size(m)
}
func (m *NewViewTx) XXX_DiscardUnknown() {
	xxx_messageInfo_NewViewTx.DiscardUnknown(m)
}

var xxx_messageInfo_NewViewTx proto.InternalMessageInfo

func (m *NewViewTx) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *NewViewTx) GetViewNumber() uint64 {
	if m != nil {
		return m.ViewNumber
	}
	return 0
}

func (m *NewViewTx) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *NewViewTx) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *NewViewTx) GetOldBspEndpoint() string {
	if m != nil {
		return m.OldBspEndpoint
	}
	return ""
}

func (m *NewViewTx) GetStatedbHash() []byte {
	if m != nil {
		return m.StatedbHash
	}
	return nil
}

func (m *NewViewTx) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *NewViewTx) GetCurrentBlockHash() []byte {
	if m != nil {
		return m.CurrentBlockHash
	}
	return nil
}

func (m *NewViewTx) GetVcm() []*RequestViewChangeMessage {
	if m != nil {
		return m.Vcm
	}
	return nil
}

type NewViewTransaction struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
	SignatureHeader      []byte   `protobuf:"bytes,3,opt,name=SignatureHeader,proto3" json:"SignatureHeader,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewViewTransaction) Reset()         { *m = NewViewTransaction{} }
func (m *NewViewTransaction) String() string { return proto.CompactTextString(m) }
func (*NewViewTransaction) ProtoMessage()    {}
func (*NewViewTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f00fded0f4841af, []int{2}
}

func (m *NewViewTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewViewTransaction.Unmarshal(m, b)
}
func (m *NewViewTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewViewTransaction.Marshal(b, m, deterministic)
}
func (m *NewViewTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewViewTransaction.Merge(m, src)
}
func (m *NewViewTransaction) XXX_Size() int {
	return xxx_messageInfo_NewViewTransaction.Size(m)
}
func (m *NewViewTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_NewViewTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_NewViewTransaction proto.InternalMessageInfo

func (m *NewViewTransaction) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *NewViewTransaction) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *NewViewTransaction) GetSignatureHeader() []byte {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

type PingMessage struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f00fded0f4841af, []int{3}
}

func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (m *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(m, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type Resp struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f00fded0f4841af, []int{4}
}

func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (m *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(m, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func (m *Resp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resp) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ViewChangeCommand struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Destination          string   `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewChangeCommand) Reset()         { *m = ViewChangeCommand{} }
func (m *ViewChangeCommand) String() string { return proto.CompactTextString(m) }
func (*ViewChangeCommand) ProtoMessage()    {}
func (*ViewChangeCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f00fded0f4841af, []int{5}
}

func (m *ViewChangeCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewChangeCommand.Unmarshal(m, b)
}
func (m *ViewChangeCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewChangeCommand.Marshal(b, m, deterministic)
}
func (m *ViewChangeCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewChangeCommand.Merge(m, src)
}
func (m *ViewChangeCommand) XXX_Size() int {
	return xxx_messageInfo_ViewChangeCommand.Size(m)
}
func (m *ViewChangeCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewChangeCommand.DiscardUnknown(m)
}

var xxx_messageInfo_ViewChangeCommand proto.InternalMessageInfo

func (m *ViewChangeCommand) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *ViewChangeCommand) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ViewChangeCommand) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

type ChangeResultEvent struct {
	PrevViewNumber  uint64 `protobuf:"varint,1,opt,name=PrevViewNumber,proto3" json:"PrevViewNumber,omitempty"`
	NextViewNumber  uint64 `protobuf:"varint,2,opt,name=NextViewNumber,proto3" json:"NextViewNumber,omitempty"`
	OrdererEndpoint string `protobuf:"bytes,3,opt,name=OrdererEndpoint,proto3" json:"OrdererEndpoint,omitempty"`
	// Expected to be unmarshaled to repeated KVWrites(i.e., an array of key/value pairs)
	NextInitialState     [][]byte `protobuf:"bytes,4,rep,name=NextInitialState,proto3" json:"NextInitialState,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeResultEvent) Reset()         { *m = ChangeResultEvent{} }
func (m *ChangeResultEvent) String() string { return proto.CompactTextString(m) }
func (*ChangeResultEvent) ProtoMessage()    {}
func (*ChangeResultEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f00fded0f4841af, []int{6}
}

func (m *ChangeResultEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeResultEvent.Unmarshal(m, b)
}
func (m *ChangeResultEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeResultEvent.Marshal(b, m, deterministic)
}
func (m *ChangeResultEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeResultEvent.Merge(m, src)
}
func (m *ChangeResultEvent) XXX_Size() int {
	return xxx_messageInfo_ChangeResultEvent.Size(m)
}
func (m *ChangeResultEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeResultEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeResultEvent proto.InternalMessageInfo

func (m *ChangeResultEvent) GetPrevViewNumber() uint64 {
	if m != nil {
		return m.PrevViewNumber
	}
	return 0
}

func (m *ChangeResultEvent) GetNextViewNumber() uint64 {
	if m != nil {
		return m.NextViewNumber
	}
	return 0
}

func (m *ChangeResultEvent) GetOrdererEndpoint() string {
	if m != nil {
		return m.OrdererEndpoint
	}
	return ""
}

func (m *ChangeResultEvent) GetNextInitialState() [][]byte {
	if m != nil {
		return m.NextInitialState
	}
	return nil
}

type ServeRequest struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServeRequest) Reset()         { *m = ServeRequest{} }
func (m *ServeRequest) String() string { return proto.CompactTextString(m) }
func (*ServeRequest) ProtoMessage()    {}
func (*ServeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f00fded0f4841af, []int{7}
}

func (m *ServeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServeRequest.Unmarshal(m, b)
}
func (m *ServeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServeRequest.Marshal(b, m, deterministic)
}
func (m *ServeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServeRequest.Merge(m, src)
}
func (m *ServeRequest) XXX_Size() int {
	return xxx_messageInfo_ServeRequest.Size(m)
}
func (m *ServeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServeRequest proto.InternalMessageInfo

func (m *ServeRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ServeResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServeResponse) Reset()         { *m = ServeResponse{} }
func (m *ServeResponse) String() string { return proto.CompactTextString(m) }
func (*ServeResponse) ProtoMessage()    {}
func (*ServeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f00fded0f4841af, []int{8}
}

func (m *ServeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServeResponse.Unmarshal(m, b)
}
func (m *ServeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServeResponse.Marshal(b, m, deterministic)
}
func (m *ServeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServeResponse.Merge(m, src)
}
func (m *ServeResponse) XXX_Size() int {
	return xxx_messageInfo_ServeResponse.Size(m)
}
func (m *ServeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServeResponse proto.InternalMessageInfo

func (m *ServeResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestViewChangeMessage)(nil), "bsp_transaction.RequestViewChangeMessage")
	proto.RegisterType((*NewViewTx)(nil), "bsp_transaction.NewViewTx")
	proto.RegisterType((*NewViewTransaction)(nil), "bsp_transaction.NewViewTransaction")
	proto.RegisterType((*PingMessage)(nil), "bsp_transaction.PingMessage")
	proto.RegisterType((*Resp)(nil), "bsp_transaction.Resp")
	proto.RegisterType((*ViewChangeCommand)(nil), "bsp_transaction.ViewChangeCommand")
	proto.RegisterType((*ChangeResultEvent)(nil), "bsp_transaction.ChangeResultEvent")
	proto.RegisterType((*ServeRequest)(nil), "bsp_transaction.ServeRequest")
	proto.RegisterType((*ServeResponse)(nil), "bsp_transaction.ServeResponse")
}

func init() { proto.RegisterFile("bspchange.proto", fileDescriptor_0f00fded0f4841af) }

var fileDescriptor_0f00fded0f4841af = []byte{
	// 718 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdd, 0x4e, 0xdb, 0x48,
	0x14, 0xc7, 0xc9, 0x07, 0x09, 0x3e, 0x36, 0x0b, 0x99, 0xd5, 0xee, 0x5a, 0x11, 0xec, 0x06, 0x6b,
	0xb5, 0x0a, 0xab, 0x15, 0x5a, 0xd1, 0xcb, 0x5e, 0x15, 0x8a, 0x44, 0x2e, 0x0a, 0x91, 0xa9, 0xe8,
	0x65, 0x34, 0xb1, 0x4f, 0x1d, 0x8b, 0x78, 0xec, 0xce, 0x8c, 0x03, 0x7d, 0xb4, 0x3e, 0x47, 0xfb,
	0x3e, 0xad, 0x66, 0xc6, 0x1f, 0x21, 0x06, 0x95, 0x3b, 0x9f, 0xff, 0x9c, 0x39, 0x33, 0xe7, 0x37,
	0xff, 0x23, 0xc3, 0xde, 0x5c, 0x64, 0xc1, 0x82, 0xb2, 0x08, 0x4f, 0x32, 0x9e, 0xca, 0x94, 0x28,
	0x61, 0x26, 0x39, 0x65, 0x82, 0x06, 0x32, 0x4e, 0x99, 0xf7, 0xbd, 0x0d, 0xae, 0x8f, 0x9f, 0x72,
	0x14, 0xf2, 0x36, 0xc6, 0xfb, 0x73, 0x9d, 0xfc, 0x0e, 0x85, 0xa0, 0x11, 0x92, 0x03, 0xb0, 0x64,
	0x9c, 0xa0, 0x90, 0x34, 0xc9, 0xdc, 0xd6, 0xa8, 0x35, 0xee, 0xfa, 0xb5, 0x40, 0x86, 0xb0, 0x73,
	0xc1, 0xc2, 0x2c, 0x8d, 0x99, 0x74, 0xdb, 0xa3, 0xd6, 0xd8, 0xf2, 0xab, 0x98, 0xfc, 0x01, 0xfd,
	0x0c, 0x91, 0xcf, 0xe2, 0xd0, 0xed, 0xe8, 0x7d, 0x3d, 0x15, 0x4e, 0x42, 0xf2, 0x17, 0xd8, 0xab,
	0x18, 0xef, 0x67, 0x2c, 0x4f, 0xe6, 0xc8, 0xdd, 0xae, 0x5e, 0x04, 0x25, 0x5d, 0x69, 0x85, 0xfc,
	0x0d, 0xbb, 0xe7, 0x39, 0xe7, 0xea, 0x22, 0x0c, 0x97, 0x93, 0xd0, 0xdd, 0xd6, 0xa5, 0x1f, 0x8b,
	0x2a, 0xeb, 0x0a, 0x1f, 0x64, 0x9d, 0xd5, 0x33, 0x59, 0x8f, 0x44, 0x32, 0x02, 0xfb, 0x7a, 0x19,
	0x56, 0x97, 0xec, 0xeb, 0x9c, 0x75, 0x89, 0xfc, 0x0f, 0xbf, 0x4e, 0x39, 0x66, 0x94, 0x63, 0x78,
	0xb6, 0x4c, 0x83, 0xbb, 0x4b, 0x8c, 0xa3, 0x85, 0x74, 0x2d, 0x7d, 0xad, 0xa7, 0x96, 0xc8, 0x7f,
	0x30, 0x78, 0x2c, 0x53, 0xb1, 0x70, 0xed, 0x51, 0x6b, 0xec, 0xf8, 0xcd, 0x05, 0x72, 0x04, 0x4e,
	0xc6, 0xd3, 0xf4, 0xe3, 0x6c, 0xae, 0x24, 0xe1, 0x3a, 0xa3, 0xce, 0xd8, 0xf1, 0x6d, 0xad, 0xe9,
	0x2c, 0xe1, 0x7d, 0x6d, 0x83, 0x75, 0x85, 0xf7, 0x8a, 0xfe, 0xfb, 0x87, 0x9f, 0x20, 0xdf, 0xa0,
	0xd7, 0x6e, 0xd0, 0x3b, 0x04, 0x08, 0x4c, 0xfb, 0x25, 0x7a, 0xcb, 0xb7, 0x82, 0x0a, 0xc8, 0x10,
	0x76, 0xb0, 0xa4, 0xd1, 0x35, 0x4f, 0x56, 0xc6, 0xe4, 0x1f, 0xf8, 0x25, 0x5d, 0x86, 0x67, 0x22,
	0xab, 0x78, 0x19, 0xf2, 0x1b, 0xaa, 0x6a, 0x49, 0x48, 0x2a, 0x31, 0x9c, 0xcf, 0x16, 0xaa, 0xf7,
	0x9e, 0xee, 0xdd, 0x2e, 0x34, 0xdd, 0xf5, 0xef, 0xd0, 0x5b, 0x18, 0x90, 0x7d, 0xf3, 0xf8, 0x8b,
	0x92, 0x1d, 0x09, 0x72, 0xce, 0x91, 0x49, 0xc3, 0xc3, 0x14, 0xb0, 0x74, 0x81, 0xfd, 0x62, 0xa5,
	0x66, 0xf7, 0x1a, 0x3a, 0xab, 0x20, 0x71, 0xed, 0x51, 0x67, 0x6c, 0x9f, 0x1e, 0x9f, 0x6c, 0x38,
	0xf7, 0xe4, 0x39, 0xd7, 0xfa, 0x6a, 0x97, 0xb7, 0x02, 0x52, 0x42, 0xad, 0xf7, 0x10, 0x17, 0xfa,
	0x53, 0xfa, 0x79, 0x99, 0xd2, 0x50, 0xb3, 0x75, 0xfc, 0x32, 0x54, 0xdc, 0x6f, 0xe2, 0x88, 0x51,
	0x99, 0x73, 0xd4, 0x5c, 0x1d, 0xbf, 0x16, 0xc8, 0x18, 0xf6, 0xaa, 0xe0, 0x12, 0x69, 0x88, 0x5c,
	0xb3, 0x75, 0xfc, 0x4d, 0xd9, 0x3b, 0x06, 0x7b, 0x1a, 0xb3, 0xa8, 0x9c, 0xa0, 0x21, 0xec, 0x44,
	0x1c, 0x51, 0xc6, 0x2c, 0xd2, 0x27, 0x5a, 0x7e, 0x15, 0x7b, 0xa7, 0xd0, 0xf5, 0x51, 0x64, 0x84,
	0x40, 0x97, 0xd1, 0x04, 0x8b, 0x75, 0xfd, 0xad, 0x08, 0x2a, 0xa0, 0xb9, 0x28, 0x26, 0xab, 0x88,
	0xbc, 0x08, 0x06, 0x75, 0xc3, 0xe7, 0x69, 0x92, 0x50, 0x16, 0xaa, 0xae, 0x02, 0xf3, 0x59, 0xd4,
	0x28, 0x43, 0x5d, 0x26, 0xcd, 0x79, 0x80, 0x55, 0x19, 0x1d, 0xa9, 0xc1, 0x08, 0x51, 0xc8, 0x98,
	0x51, 0x85, 0xa5, 0xf0, 0xc9, 0xba, 0xe4, 0x7d, 0x69, 0xc1, 0xc0, 0x9c, 0xe2, 0xa3, 0xc8, 0x97,
	0xf2, 0x62, 0x85, 0xc6, 0x23, 0x53, 0x8e, 0xab, 0xdb, 0xca, 0x70, 0x85, 0x45, 0x37, 0x54, 0x95,
	0xa7, 0x26, 0x71, 0x2d, 0xcf, 0x58, 0x75, 0x43, 0x55, 0x5c, 0xaf, 0x79, 0x88, 0x1c, 0x79, 0x65,
	0x3a, 0x73, 0x97, 0x4d, 0x99, 0xfc, 0x0b, 0xfb, 0x6a, 0xef, 0x84, 0xc5, 0x32, 0xa6, 0xcb, 0x1b,
	0x65, 0x36, 0xb7, 0xab, 0x87, 0xa9, 0xa1, 0x7b, 0x23, 0x70, 0x6e, 0x90, 0xaf, 0xb0, 0x70, 0x08,
	0xd9, 0x87, 0x4e, 0x22, 0x4a, 0xfe, 0xea, 0xd3, 0x3b, 0x82, 0xdd, 0x22, 0x43, 0x64, 0x29, 0x13,
	0xd8, 0x4c, 0x39, 0xfd, 0xd6, 0x06, 0xbb, 0x46, 0xcd, 0xc9, 0x25, 0x6c, 0xeb, 0x2d, 0xe4, 0xb0,
	0xe1, 0xc4, 0xf5, 0xc3, 0x86, 0x7f, 0x3e, 0xb7, 0x6c, 0x4e, 0xf2, 0xb6, 0xc8, 0x04, 0xa0, 0x2e,
	0x4c, 0xbc, 0x46, 0x7e, 0xe3, 0x81, 0x87, 0xbf, 0x3d, 0x61, 0x7e, 0x91, 0x79, 0x5b, 0xe4, 0x03,
	0x0c, 0xde, 0x88, 0xbb, 0x7a, 0xc3, 0xdb, 0x94, 0xbd, 0xac, 0xe2, 0x41, 0x23, 0x67, 0xcd, 0xb5,
	0xa6, 0x70, 0x63, 0xbe, 0xc8, 0xcb, 0x67, 0xf0, 0xd9, 0x1b, 0xcf, 0x7b, 0xfa, 0x3f, 0xf4, 0xea,
	0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x6d, 0x8b, 0x3a, 0x9a, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ViewChangerClient is the client API for ViewChanger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ViewChangerClient interface {
	Serve(ctx context.Context, in *ServeRequest, opts ...grpc.CallOption) (*ServeResponse, error)
	ViewChange(ctx context.Context, in *ViewChangeCommand, opts ...grpc.CallOption) (*Resp, error)
	AskViewChangeDone(ctx context.Context, in *ViewChangeCommand, opts ...grpc.CallOption) (*PingMessage, error)
	RequestViewChange(ctx context.Context, in *RequestViewChangeMessage, opts ...grpc.CallOption) (*Resp, error)
}

type viewChangerClient struct {
	cc *grpc.ClientConn
}

func NewViewChangerClient(cc *grpc.ClientConn) ViewChangerClient {
	return &viewChangerClient{cc}
}

func (c *viewChangerClient) Serve(ctx context.Context, in *ServeRequest, opts ...grpc.CallOption) (*ServeResponse, error) {
	out := new(ServeResponse)
	err := c.cc.Invoke(ctx, "/bsp_transaction.ViewChanger/Serve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewChangerClient) ViewChange(ctx context.Context, in *ViewChangeCommand, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/bsp_transaction.ViewChanger/ViewChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewChangerClient) AskViewChangeDone(ctx context.Context, in *ViewChangeCommand, opts ...grpc.CallOption) (*PingMessage, error) {
	out := new(PingMessage)
	err := c.cc.Invoke(ctx, "/bsp_transaction.ViewChanger/AskViewChangeDone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewChangerClient) RequestViewChange(ctx context.Context, in *RequestViewChangeMessage, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/bsp_transaction.ViewChanger/RequestViewChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViewChangerServer is the server API for ViewChanger service.
type ViewChangerServer interface {
	Serve(context.Context, *ServeRequest) (*ServeResponse, error)
	ViewChange(context.Context, *ViewChangeCommand) (*Resp, error)
	AskViewChangeDone(context.Context, *ViewChangeCommand) (*PingMessage, error)
	RequestViewChange(context.Context, *RequestViewChangeMessage) (*Resp, error)
}

// UnimplementedViewChangerServer can be embedded to have forward compatible implementations.
type UnimplementedViewChangerServer struct {
}

func (*UnimplementedViewChangerServer) Serve(ctx context.Context, req *ServeRequest) (*ServeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Serve not implemented")
}
func (*UnimplementedViewChangerServer) ViewChange(ctx context.Context, req *ViewChangeCommand) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewChange not implemented")
}
func (*UnimplementedViewChangerServer) AskViewChangeDone(ctx context.Context, req *ViewChangeCommand) (*PingMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskViewChangeDone not implemented")
}
func (*UnimplementedViewChangerServer) RequestViewChange(ctx context.Context, req *RequestViewChangeMessage) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestViewChange not implemented")
}

func RegisterViewChangerServer(s *grpc.Server, srv ViewChangerServer) {
	s.RegisterService(&_ViewChanger_serviceDesc, srv)
}

func _ViewChanger_Serve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewChangerServer).Serve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bsp_transaction.ViewChanger/Serve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewChangerServer).Serve(ctx, req.(*ServeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewChanger_ViewChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewChangeCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewChangerServer).ViewChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bsp_transaction.ViewChanger/ViewChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewChangerServer).ViewChange(ctx, req.(*ViewChangeCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewChanger_AskViewChangeDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewChangeCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewChangerServer).AskViewChangeDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bsp_transaction.ViewChanger/AskViewChangeDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewChangerServer).AskViewChangeDone(ctx, req.(*ViewChangeCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewChanger_RequestViewChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestViewChangeMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewChangerServer).RequestViewChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bsp_transaction.ViewChanger/RequestViewChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewChangerServer).RequestViewChange(ctx, req.(*RequestViewChangeMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _ViewChanger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bsp_transaction.ViewChanger",
	HandlerType: (*ViewChangerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Serve",
			Handler:    _ViewChanger_Serve_Handler,
		},
		{
			MethodName: "ViewChange",
			Handler:    _ViewChanger_ViewChange_Handler,
		},
		{
			MethodName: "AskViewChangeDone",
			Handler:    _ViewChanger_AskViewChangeDone_Handler,
		},
		{
			MethodName: "RequestViewChange",
			Handler:    _ViewChanger_RequestViewChange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bspchange.proto",
}
