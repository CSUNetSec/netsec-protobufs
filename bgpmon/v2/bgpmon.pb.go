// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/CSUNetSec/netsec-protobufs/bgpmon/v2/bgpmon.proto

/*
Package bgpmon is a generated protocol buffer package.

It is generated from these files:
	github.com/CSUNetSec/netsec-protobufs/bgpmon/v2/bgpmon.proto

It has these top-level messages:
	Empty
	SessionType
	ListAvailableSessionsReply
	CloseSessionRequest
	ListOpenSessionsReply
	OpenSessionRequest
	OpenSessionReply
	WriteRequest
	GetRequest
	GetReply
	BGPCapture
*/
package bgpmon

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/CSUNetSec/netsec-protobufs/common"
import bgp "github.com/CSUNetSec/netsec-protobufs/protocol/bgp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WriteRequest_Type int32

const (
	WriteRequest_BGP_CAPTURE WriteRequest_Type = 0
)

var WriteRequest_Type_name = map[int32]string{
	0: "BGP_CAPTURE",
}
var WriteRequest_Type_value = map[string]int32{
	"BGP_CAPTURE": 0,
}

func (x WriteRequest_Type) String() string {
	return proto.EnumName(WriteRequest_Type_name, int32(x))
}
func (WriteRequest_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

type GetRequest_Type int32

const (
	GetRequest_BGP_CAPTURE GetRequest_Type = 0
)

var GetRequest_Type_name = map[int32]string{
	0: "BGP_CAPTURE",
}
var GetRequest_Type_value = map[string]int32{
	"BGP_CAPTURE": 0,
}

func (x GetRequest_Type) String() string {
	return proto.EnumName(GetRequest_Type_name, int32(x))
}
func (GetRequest_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 0} }

type GetReply_Type int32

const (
	GetReply_BGP_CAPTURE GetReply_Type = 0
)

var GetReply_Type_name = map[int32]string{
	0: "BGP_CAPTURE",
}
var GetReply_Type_value = map[string]int32{
	"BGP_CAPTURE": 0,
}

func (x GetReply_Type) String() string {
	return proto.EnumName(GetReply_Type_name, int32(x))
}
func (GetReply_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{9, 0} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SessionType struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Desc string `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
}

func (m *SessionType) Reset()                    { *m = SessionType{} }
func (m *SessionType) String() string            { return proto.CompactTextString(m) }
func (*SessionType) ProtoMessage()               {}
func (*SessionType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SessionType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SessionType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SessionType) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type ListAvailableSessionsReply struct {
	AvailableSessions []*SessionType `protobuf:"bytes,1,rep,name=available_sessions,json=availableSessions" json:"available_sessions,omitempty"`
}

func (m *ListAvailableSessionsReply) Reset()                    { *m = ListAvailableSessionsReply{} }
func (m *ListAvailableSessionsReply) String() string            { return proto.CompactTextString(m) }
func (*ListAvailableSessionsReply) ProtoMessage()               {}
func (*ListAvailableSessionsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListAvailableSessionsReply) GetAvailableSessions() []*SessionType {
	if m != nil {
		return m.AvailableSessions
	}
	return nil
}

//
// Session Command Messages
type CloseSessionRequest struct {
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
}

func (m *CloseSessionRequest) Reset()                    { *m = CloseSessionRequest{} }
func (m *CloseSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*CloseSessionRequest) ProtoMessage()               {}
func (*CloseSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CloseSessionRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type ListOpenSessionsReply struct {
	SessionId []string `protobuf:"bytes,1,rep,name=session_id,json=sessionId" json:"session_id,omitempty"`
}

func (m *ListOpenSessionsReply) Reset()                    { *m = ListOpenSessionsReply{} }
func (m *ListOpenSessionsReply) String() string            { return proto.CompactTextString(m) }
func (*ListOpenSessionsReply) ProtoMessage()               {}
func (*ListOpenSessionsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListOpenSessionsReply) GetSessionId() []string {
	if m != nil {
		return m.SessionId
	}
	return nil
}

type OpenSessionRequest struct {
	SessionName string `protobuf:"bytes,1,opt,name=session_name,json=sessionName" json:"session_name,omitempty"`
	SessionId   string `protobuf:"bytes,2,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	Workers     uint32 `protobuf:"varint,3,opt,name=workers" json:"workers,omitempty"`
}

func (m *OpenSessionRequest) Reset()                    { *m = OpenSessionRequest{} }
func (m *OpenSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*OpenSessionRequest) ProtoMessage()               {}
func (*OpenSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *OpenSessionRequest) GetSessionName() string {
	if m != nil {
		return m.SessionName
	}
	return ""
}

func (m *OpenSessionRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *OpenSessionRequest) GetWorkers() uint32 {
	if m != nil {
		return m.Workers
	}
	return 0
}

type OpenSessionReply struct {
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
}

func (m *OpenSessionReply) Reset()                    { *m = OpenSessionReply{} }
func (m *OpenSessionReply) String() string            { return proto.CompactTextString(m) }
func (*OpenSessionReply) ProtoMessage()               {}
func (*OpenSessionReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *OpenSessionReply) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

//
// Write Messages
type WriteRequest struct {
	Type       WriteRequest_Type `protobuf:"varint,1,opt,name=type,enum=bgpmon.WriteRequest_Type" json:"type,omitempty"`
	SessionId  string            `protobuf:"bytes,2,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	BgpCapture *BGPCapture       `protobuf:"bytes,3,opt,name=bgp_capture,json=bgpCapture" json:"bgp_capture,omitempty"`
}

func (m *WriteRequest) Reset()                    { *m = WriteRequest{} }
func (m *WriteRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()               {}
func (*WriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *WriteRequest) GetType() WriteRequest_Type {
	if m != nil {
		return m.Type
	}
	return WriteRequest_BGP_CAPTURE
}

func (m *WriteRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *WriteRequest) GetBgpCapture() *BGPCapture {
	if m != nil {
		return m.BgpCapture
	}
	return nil
}

//
// Get Messages
type GetRequest struct {
	Type           GetRequest_Type `protobuf:"varint,1,opt,name=type,enum=bgpmon.GetRequest_Type" json:"type,omitempty"`
	SessionId      string          `protobuf:"bytes,2,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	StartTimestamp string          `protobuf:"bytes,3,opt,name=start_timestamp,json=startTimestamp" json:"start_timestamp,omitempty"`
	EndTimestamp   string          `protobuf:"bytes,4,opt,name=end_timestamp,json=endTimestamp" json:"end_timestamp,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetRequest) GetType() GetRequest_Type {
	if m != nil {
		return m.Type
	}
	return GetRequest_BGP_CAPTURE
}

func (m *GetRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *GetRequest) GetStartTimestamp() string {
	if m != nil {
		return m.StartTimestamp
	}
	return ""
}

func (m *GetRequest) GetEndTimestamp() string {
	if m != nil {
		return m.EndTimestamp
	}
	return ""
}

type GetReply struct {
	Type     GetReply_Type `protobuf:"varint,1,opt,name=type,enum=bgpmon.GetReply_Type" json:"type,omitempty"`
	Captures []*BGPCapture `protobuf:"bytes,2,rep,name=captures" json:"captures,omitempty"`
}

func (m *GetReply) Reset()                    { *m = GetReply{} }
func (m *GetReply) String() string            { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()               {}
func (*GetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetReply) GetType() GetReply_Type {
	if m != nil {
		return m.Type
	}
	return GetReply_BGP_CAPTURE
}

func (m *GetReply) GetCaptures() []*BGPCapture {
	if m != nil {
		return m.Captures
	}
	return nil
}

// a message of Capture type to facilitate the represtation of a
// BGP Update from multuple sources (live/mrt etc) to bgpmon
// and other systems. the inner actual update is defined in protocol/bgp
type BGPCapture struct {
	Timestamp      uint32                   `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	PeerAs         uint32                   `protobuf:"varint,2,opt,name=peer_as,json=peerAs" json:"peer_as,omitempty"`
	LocalAs        uint32                   `protobuf:"varint,3,opt,name=local_as,json=localAs" json:"local_as,omitempty"`
	InterfaceIndex uint32                   `protobuf:"varint,4,opt,name=interface_index,json=interfaceIndex" json:"interface_index,omitempty"`
	AddressFamily  uint32                   `protobuf:"varint,5,opt,name=address_family,json=addressFamily" json:"address_family,omitempty"`
	PeerIp         *common.IPAddressWrapper `protobuf:"bytes,6,opt,name=peer_ip,json=peerIp" json:"peer_ip,omitempty"`
	LocalIp        *common.IPAddressWrapper `protobuf:"bytes,7,opt,name=local_ip,json=localIp" json:"local_ip,omitempty"`
	Update         *bgp.BGPUpdate           `protobuf:"bytes,8,opt,name=update" json:"update,omitempty"`
}

func (m *BGPCapture) Reset()                    { *m = BGPCapture{} }
func (m *BGPCapture) String() string            { return proto.CompactTextString(m) }
func (*BGPCapture) ProtoMessage()               {}
func (*BGPCapture) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *BGPCapture) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BGPCapture) GetPeerAs() uint32 {
	if m != nil {
		return m.PeerAs
	}
	return 0
}

func (m *BGPCapture) GetLocalAs() uint32 {
	if m != nil {
		return m.LocalAs
	}
	return 0
}

func (m *BGPCapture) GetInterfaceIndex() uint32 {
	if m != nil {
		return m.InterfaceIndex
	}
	return 0
}

func (m *BGPCapture) GetAddressFamily() uint32 {
	if m != nil {
		return m.AddressFamily
	}
	return 0
}

func (m *BGPCapture) GetPeerIp() *common.IPAddressWrapper {
	if m != nil {
		return m.PeerIp
	}
	return nil
}

func (m *BGPCapture) GetLocalIp() *common.IPAddressWrapper {
	if m != nil {
		return m.LocalIp
	}
	return nil
}

func (m *BGPCapture) GetUpdate() *bgp.BGPUpdate {
	if m != nil {
		return m.Update
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "bgpmon.Empty")
	proto.RegisterType((*SessionType)(nil), "bgpmon.SessionType")
	proto.RegisterType((*ListAvailableSessionsReply)(nil), "bgpmon.ListAvailableSessionsReply")
	proto.RegisterType((*CloseSessionRequest)(nil), "bgpmon.CloseSessionRequest")
	proto.RegisterType((*ListOpenSessionsReply)(nil), "bgpmon.ListOpenSessionsReply")
	proto.RegisterType((*OpenSessionRequest)(nil), "bgpmon.OpenSessionRequest")
	proto.RegisterType((*OpenSessionReply)(nil), "bgpmon.OpenSessionReply")
	proto.RegisterType((*WriteRequest)(nil), "bgpmon.WriteRequest")
	proto.RegisterType((*GetRequest)(nil), "bgpmon.GetRequest")
	proto.RegisterType((*GetReply)(nil), "bgpmon.GetReply")
	proto.RegisterType((*BGPCapture)(nil), "bgpmon.BGPCapture")
	proto.RegisterEnum("bgpmon.WriteRequest_Type", WriteRequest_Type_name, WriteRequest_Type_value)
	proto.RegisterEnum("bgpmon.GetRequest_Type", GetRequest_Type_name, GetRequest_Type_value)
	proto.RegisterEnum("bgpmon.GetReply_Type", GetReply_Type_name, GetReply_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Bgpmond service

type BgpmondClient interface {
	CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...grpc.CallOption) (*Empty, error)
	ListOpenSessions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListOpenSessionsReply, error)
	ListAvailableSessions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListAvailableSessionsReply, error)
	OpenSession(ctx context.Context, in *OpenSessionRequest, opts ...grpc.CallOption) (*OpenSessionReply, error)
	Write(ctx context.Context, opts ...grpc.CallOption) (Bgpmond_WriteClient, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (Bgpmond_GetClient, error)
}

type bgpmondClient struct {
	cc *grpc.ClientConn
}

func NewBgpmondClient(cc *grpc.ClientConn) BgpmondClient {
	return &bgpmondClient{cc}
}

func (c *bgpmondClient) CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/bgpmon.Bgpmond/CloseSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bgpmondClient) ListOpenSessions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListOpenSessionsReply, error) {
	out := new(ListOpenSessionsReply)
	err := grpc.Invoke(ctx, "/bgpmon.Bgpmond/ListOpenSessions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bgpmondClient) ListAvailableSessions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListAvailableSessionsReply, error) {
	out := new(ListAvailableSessionsReply)
	err := grpc.Invoke(ctx, "/bgpmon.Bgpmond/ListAvailableSessions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bgpmondClient) OpenSession(ctx context.Context, in *OpenSessionRequest, opts ...grpc.CallOption) (*OpenSessionReply, error) {
	out := new(OpenSessionReply)
	err := grpc.Invoke(ctx, "/bgpmon.Bgpmond/OpenSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bgpmondClient) Write(ctx context.Context, opts ...grpc.CallOption) (Bgpmond_WriteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Bgpmond_serviceDesc.Streams[0], c.cc, "/bgpmon.Bgpmond/Write", opts...)
	if err != nil {
		return nil, err
	}
	x := &bgpmondWriteClient{stream}
	return x, nil
}

type Bgpmond_WriteClient interface {
	Send(*WriteRequest) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type bgpmondWriteClient struct {
	grpc.ClientStream
}

func (x *bgpmondWriteClient) Send(m *WriteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bgpmondWriteClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bgpmondClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (Bgpmond_GetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Bgpmond_serviceDesc.Streams[1], c.cc, "/bgpmon.Bgpmond/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &bgpmondGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Bgpmond_GetClient interface {
	Recv() (*GetReply, error)
	grpc.ClientStream
}

type bgpmondGetClient struct {
	grpc.ClientStream
}

func (x *bgpmondGetClient) Recv() (*GetReply, error) {
	m := new(GetReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Bgpmond service

type BgpmondServer interface {
	CloseSession(context.Context, *CloseSessionRequest) (*Empty, error)
	ListOpenSessions(context.Context, *Empty) (*ListOpenSessionsReply, error)
	ListAvailableSessions(context.Context, *Empty) (*ListAvailableSessionsReply, error)
	OpenSession(context.Context, *OpenSessionRequest) (*OpenSessionReply, error)
	Write(Bgpmond_WriteServer) error
	Get(*GetRequest, Bgpmond_GetServer) error
}

func RegisterBgpmondServer(s *grpc.Server, srv BgpmondServer) {
	s.RegisterService(&_Bgpmond_serviceDesc, srv)
}

func _Bgpmond_CloseSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BgpmondServer).CloseSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bgpmon.Bgpmond/CloseSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BgpmondServer).CloseSession(ctx, req.(*CloseSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bgpmond_ListOpenSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BgpmondServer).ListOpenSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bgpmon.Bgpmond/ListOpenSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BgpmondServer).ListOpenSessions(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bgpmond_ListAvailableSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BgpmondServer).ListAvailableSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bgpmon.Bgpmond/ListAvailableSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BgpmondServer).ListAvailableSessions(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bgpmond_OpenSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BgpmondServer).OpenSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bgpmon.Bgpmond/OpenSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BgpmondServer).OpenSession(ctx, req.(*OpenSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bgpmond_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BgpmondServer).Write(&bgpmondWriteServer{stream})
}

type Bgpmond_WriteServer interface {
	SendAndClose(*Empty) error
	Recv() (*WriteRequest, error)
	grpc.ServerStream
}

type bgpmondWriteServer struct {
	grpc.ServerStream
}

func (x *bgpmondWriteServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bgpmondWriteServer) Recv() (*WriteRequest, error) {
	m := new(WriteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Bgpmond_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BgpmondServer).Get(m, &bgpmondGetServer{stream})
}

type Bgpmond_GetServer interface {
	Send(*GetReply) error
	grpc.ServerStream
}

type bgpmondGetServer struct {
	grpc.ServerStream
}

func (x *bgpmondGetServer) Send(m *GetReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Bgpmond_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bgpmon.Bgpmond",
	HandlerType: (*BgpmondServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CloseSession",
			Handler:    _Bgpmond_CloseSession_Handler,
		},
		{
			MethodName: "ListOpenSessions",
			Handler:    _Bgpmond_ListOpenSessions_Handler,
		},
		{
			MethodName: "ListAvailableSessions",
			Handler:    _Bgpmond_ListAvailableSessions_Handler,
		},
		{
			MethodName: "OpenSession",
			Handler:    _Bgpmond_OpenSession_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Write",
			Handler:       _Bgpmond_Write_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Get",
			Handler:       _Bgpmond_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/CSUNetSec/netsec-protobufs/bgpmon/v2/bgpmon.proto",
}

func init() {
	proto.RegisterFile("github.com/CSUNetSec/netsec-protobufs/bgpmon/v2/bgpmon.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 762 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x4e, 0xdb, 0x48,
	0x18, 0x8d, 0x13, 0x48, 0xc2, 0x97, 0x1f, 0xb2, 0xc3, 0x22, 0x8c, 0x77, 0x91, 0x58, 0xaf, 0x76,
	0x97, 0xd5, 0x8a, 0x04, 0xc2, 0x6a, 0xa5, 0x95, 0xb8, 0x49, 0x22, 0x1a, 0x45, 0x42, 0x34, 0x32,
	0x20, 0x2e, 0x5d, 0xc7, 0x1e, 0x52, 0xab, 0xfe, 0x99, 0x7a, 0x26, 0xb4, 0xb9, 0xe9, 0x4d, 0xdf,
	0xa6, 0x0f, 0xd1, 0xbe, 0x5a, 0x35, 0xe3, 0x19, 0x62, 0x27, 0x41, 0xd0, 0x0b, 0xc4, 0xcc, 0xf9,
	0x7e, 0xe6, 0xcc, 0x99, 0xcf, 0x27, 0x70, 0x3e, 0xf5, 0xd9, 0xdb, 0xd9, 0xa4, 0xed, 0xc6, 0x61,
	0x67, 0x70, 0x7d, 0x7b, 0x85, 0xd9, 0x35, 0x76, 0x3b, 0x11, 0x66, 0x14, 0xbb, 0xc7, 0x24, 0x89,
	0x59, 0x3c, 0x99, 0xdd, 0xd3, 0xce, 0x64, 0x4a, 0xc2, 0x38, 0xea, 0x3c, 0x74, 0xe5, 0xaa, 0x2d,
	0x42, 0xa8, 0x9c, 0xee, 0x8c, 0xff, 0x5f, 0xd6, 0xc5, 0x8d, 0x43, 0xde, 0x25, 0xfd, 0x97, 0xb6,
	0x30, 0x5e, 0x48, 0x40, 0xac, 0xdc, 0x38, 0xe0, 0xe7, 0xf3, 0xbf, 0xb4, 0xda, 0xac, 0xc0, 0xe6,
	0x45, 0x48, 0xd8, 0xdc, 0x1c, 0x41, 0xed, 0x1a, 0x53, 0xea, 0xc7, 0xd1, 0xcd, 0x9c, 0x60, 0x84,
	0x60, 0x23, 0x72, 0x42, 0xac, 0x6b, 0x87, 0xda, 0xd1, 0x96, 0x25, 0xd6, 0x1c, 0x63, 0x73, 0x82,
	0xf5, 0x62, 0x8a, 0x31, 0x99, 0xe7, 0x61, 0xea, 0xea, 0xa5, 0x14, 0xe3, 0x6b, 0xf3, 0x0d, 0x18,
	0x97, 0x3e, 0x65, 0xbd, 0x07, 0xc7, 0x0f, 0x9c, 0x49, 0x80, 0x65, 0x5f, 0x6a, 0x61, 0x12, 0xcc,
	0x51, 0x1f, 0x90, 0xa3, 0x22, 0x36, 0x95, 0x21, 0x5d, 0x3b, 0x2c, 0x1d, 0xd5, 0xba, 0x3b, 0x6d,
	0xa9, 0x4e, 0x86, 0x8a, 0xf5, 0x93, 0xb3, 0xdc, 0xc8, 0xfc, 0x17, 0x76, 0x06, 0x41, 0x4c, 0x15,
	0x60, 0xe1, 0xf7, 0x33, 0x4c, 0x19, 0x3a, 0x00, 0x90, 0x0d, 0x6d, 0xdf, 0x93, 0xd4, 0xb7, 0x24,
	0x32, 0xf2, 0xcc, 0xff, 0x60, 0x97, 0xf3, 0x7a, 0x4d, 0x70, 0x94, 0xa7, 0xb4, 0x5c, 0x57, 0xca,
	0xd7, 0x11, 0x40, 0x99, 0x1a, 0x75, 0xd8, 0x6f, 0x50, 0x57, 0x45, 0x19, 0xa5, 0x6a, 0x12, 0xbb,
	0xe2, 0x82, 0xe5, 0xfb, 0x16, 0x97, 0xf8, 0x20, 0x1d, 0x2a, 0x1f, 0xe2, 0xe4, 0x1d, 0x4e, 0xa8,
	0x90, 0xaf, 0x61, 0xa9, 0xad, 0x79, 0x0a, 0xad, 0xdc, 0x89, 0xeb, 0x48, 0x2e, 0x5d, 0xee, 0x8b,
	0x06, 0xf5, 0xbb, 0xc4, 0x67, 0x58, 0xf1, 0x3b, 0x96, 0xaf, 0xc5, 0x33, 0x9b, 0xdd, 0x7d, 0xa5,
	0x6c, 0x36, 0xa7, 0x2d, 0xf4, 0x4d, 0x1f, 0xf2, 0x19, 0xae, 0x67, 0x50, 0x9b, 0x4c, 0x89, 0xed,
	0x3a, 0x84, 0xcd, 0x12, 0x2c, 0xf8, 0xd6, 0xba, 0x48, 0x35, 0xed, 0x0f, 0xc7, 0x83, 0x34, 0x62,
	0xc1, 0x64, 0x4a, 0xe4, 0xda, 0xdc, 0x83, 0x0d, 0x31, 0x4c, 0xdb, 0x50, 0xeb, 0x0f, 0xc7, 0xf6,
	0xa0, 0x37, 0xbe, 0xb9, 0xb5, 0x2e, 0x5a, 0x05, 0xf3, 0x9b, 0x06, 0x30, 0xc4, 0x4c, 0x51, 0xfd,
	0x27, 0x47, 0x75, 0x4f, 0x75, 0x5d, 0x64, 0xfc, 0x00, 0xd1, 0xbf, 0x60, 0x9b, 0x32, 0x27, 0x61,
	0x36, 0xf3, 0x43, 0x4c, 0x99, 0x13, 0x12, 0x39, 0x9b, 0x4d, 0x01, 0xdf, 0x28, 0x14, 0xfd, 0x0e,
	0x0d, 0x1c, 0x79, 0x99, 0xb4, 0x0d, 0x91, 0x56, 0xc7, 0x91, 0xf7, 0x98, 0xf4, 0xf4, 0x0d, 0x3e,
	0x41, 0x55, 0xd0, 0xe3, 0x2f, 0xf3, 0x77, 0x8e, 0xfe, 0x6e, 0x8e, 0x3e, 0x09, 0xe6, 0x59, 0xf2,
	0x6d, 0xa8, 0x4a, 0x09, 0xa9, 0x5e, 0x14, 0x23, 0xbf, 0x4e, 0xc3, 0xc7, 0x9c, 0xa7, 0xcf, 0xff,
	0x5a, 0x04, 0x58, 0x54, 0xa0, 0x5f, 0x61, 0x6b, 0x71, 0x11, 0x4d, 0x0c, 0xd3, 0x02, 0x40, 0x7b,
	0x50, 0x21, 0x18, 0x27, 0xb6, 0x43, 0x85, 0x5e, 0x0d, 0xab, 0xcc, 0xb7, 0x3d, 0x8a, 0xf6, 0xa1,
	0x1a, 0xc4, 0xae, 0x13, 0xf0, 0x88, 0x1c, 0x41, 0xb1, 0xef, 0x51, 0xae, 0xa3, 0x1f, 0x31, 0x9c,
	0xdc, 0x3b, 0x2e, 0xb6, 0xfd, 0xc8, 0xc3, 0x1f, 0x85, 0x40, 0x0d, 0xab, 0xf9, 0x08, 0x8f, 0x38,
	0x8a, 0xfe, 0x80, 0xa6, 0xe3, 0x79, 0x09, 0xa6, 0xd4, 0xbe, 0x77, 0x42, 0x3f, 0x98, 0xeb, 0x9b,
	0x22, 0xaf, 0x21, 0xd1, 0x57, 0x02, 0x44, 0xa7, 0x92, 0x83, 0x4f, 0xf4, 0xb2, 0x18, 0x1e, 0xbd,
	0x2d, 0x6d, 0x6c, 0x34, 0xee, 0xa5, 0x99, 0x77, 0x89, 0x43, 0x08, 0x4e, 0x52, 0x76, 0x23, 0x82,
	0xce, 0x14, 0x3b, 0x9f, 0xe8, 0x95, 0x67, 0x6a, 0x52, 0xde, 0x23, 0x82, 0xfe, 0x84, 0xf2, 0x8c,
	0x78, 0x0e, 0xc3, 0x7a, 0x55, 0x94, 0x34, 0xb9, 0xbe, 0x5c, 0xdc, 0x5b, 0x81, 0x5a, 0x32, 0xda,
	0xfd, 0x5c, 0x82, 0x4a, 0x5f, 0x28, 0xef, 0xa1, 0x73, 0xa8, 0x67, 0xed, 0x04, 0xfd, 0xa2, 0xde,
	0x64, 0x8d, 0xc9, 0x18, 0x0d, 0x15, 0x4c, 0x7d, 0xb3, 0x80, 0xfa, 0xd0, 0x5a, 0xb6, 0x15, 0x94,
	0x4f, 0x32, 0x0e, 0xd4, 0x76, 0xad, 0xff, 0x98, 0x05, 0x74, 0x99, 0x5a, 0xd3, 0x8a, 0x65, 0x2e,
	0x37, 0x32, 0xb3, 0x8d, 0xd6, 0x1b, 0xac, 0x59, 0x40, 0x17, 0x50, 0xcb, 0x1c, 0x82, 0x0c, 0x55,
	0xb4, 0xea, 0x62, 0x86, 0xbe, 0x36, 0x96, 0xb6, 0x39, 0x81, 0x4d, 0xe1, 0x16, 0xe8, 0xe7, 0x75,
	0xe6, 0xb1, 0x22, 0xc4, 0x91, 0x86, 0x3a, 0x50, 0x1a, 0x62, 0x86, 0xd0, 0xea, 0x17, 0x6c, 0xb4,
	0x96, 0x3f, 0x0b, 0xb3, 0x70, 0xa2, 0x4d, 0xca, 0xe2, 0x57, 0xe8, 0xec, 0x7b, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x7c, 0xa6, 0x23, 0x41, 0x46, 0x07, 0x00, 0x00,
}
