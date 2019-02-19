// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/chat.proto

package go_jebb_chat

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	LoggedIn             bool     `protobuf:"varint,1,opt,name=logged_in,json=loggedIn,proto3" json:"logged_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetLoggedIn() bool {
	if m != nil {
		return m.LoggedIn
	}
	return false
}

type OnlineRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OnlineRequest) Reset()         { *m = OnlineRequest{} }
func (m *OnlineRequest) String() string { return proto.CompactTextString(m) }
func (*OnlineRequest) ProtoMessage()    {}
func (*OnlineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{2}
}

func (m *OnlineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnlineRequest.Unmarshal(m, b)
}
func (m *OnlineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnlineRequest.Marshal(b, m, deterministic)
}
func (m *OnlineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnlineRequest.Merge(m, src)
}
func (m *OnlineRequest) XXX_Size() int {
	return xxx_messageInfo_OnlineRequest.Size(m)
}
func (m *OnlineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OnlineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OnlineRequest proto.InternalMessageInfo

type OnlineResponse struct {
	Users                []string `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OnlineResponse) Reset()         { *m = OnlineResponse{} }
func (m *OnlineResponse) String() string { return proto.CompactTextString(m) }
func (*OnlineResponse) ProtoMessage()    {}
func (*OnlineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{3}
}

func (m *OnlineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnlineResponse.Unmarshal(m, b)
}
func (m *OnlineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnlineResponse.Marshal(b, m, deterministic)
}
func (m *OnlineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnlineResponse.Merge(m, src)
}
func (m *OnlineResponse) XXX_Size() int {
	return xxx_messageInfo_OnlineResponse.Size(m)
}
func (m *OnlineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OnlineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OnlineResponse proto.InternalMessageInfo

func (m *OnlineResponse) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

type Message struct {
	MessageType          int32    `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	Sender               string   `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver             string   `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Msg                  string   `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{4}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessageType() int32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

func (m *Message) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Message) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *Message) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type LogoutRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{5}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type LogoutResponse struct {
	LoggedOut            bool     `protobuf:"varint,1,opt,name=logged_out,json=loggedOut,proto3" json:"logged_out,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{6}
}

func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (m *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(m, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

func (m *LogoutResponse) GetLoggedOut() bool {
	if m != nil {
		return m.LoggedOut
	}
	return false
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "go.jebb.chat.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "go.jebb.chat.LoginResponse")
	proto.RegisterType((*OnlineRequest)(nil), "go.jebb.chat.OnlineRequest")
	proto.RegisterType((*OnlineResponse)(nil), "go.jebb.chat.OnlineResponse")
	proto.RegisterType((*Message)(nil), "go.jebb.chat.Message")
	proto.RegisterType((*LogoutRequest)(nil), "go.jebb.chat.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "go.jebb.chat.LogoutResponse")
}

func init() { proto.RegisterFile("proto/chat.proto", fileDescriptor_ed7e7dde45555b7d) }

var fileDescriptor_ed7e7dde45555b7d = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0x29, 0xbc, 0xf4, 0x85, 0xe1, 0x8f, 0x64, 0xa3, 0xa6, 0x29, 0x9a, 0xe0, 0x1e, 0x0c,
	0x89, 0xa6, 0x18, 0xbd, 0x7a, 0xd2, 0x68, 0x62, 0xa2, 0x21, 0x29, 0xde, 0x49, 0x29, 0x93, 0x52,
	0x85, 0xdd, 0xba, 0xbb, 0xc5, 0xf0, 0x61, 0xfc, 0xae, 0xa6, 0xdd, 0x2d, 0x02, 0x72, 0xf0, 0xb6,
	0xf3, 0xcc, 0x74, 0xe6, 0x99, 0xdf, 0x14, 0x3a, 0x89, 0xe0, 0x8a, 0x0f, 0xc2, 0x59, 0xa0, 0xbc,
	0xfc, 0x49, 0x9a, 0x11, 0xf7, 0xde, 0x70, 0x32, 0xf1, 0x32, 0x8d, 0x3e, 0x42, 0xf3, 0x99, 0x47,
	0x31, 0xf3, 0xf1, 0x23, 0x45, 0xa9, 0x88, 0x0b, 0xb5, 0x54, 0xa2, 0x60, 0xc1, 0x02, 0x1d, 0xab,
	0x67, 0xf5, 0xeb, 0xfe, 0x3a, 0xce, 0x72, 0x49, 0x20, 0xe5, 0x27, 0x17, 0x53, 0xa7, 0xac, 0x73,
	0x45, 0x4c, 0x2f, 0xa1, 0x65, 0xfa, 0xc8, 0x84, 0x33, 0x89, 0xa4, 0x0b, 0xf5, 0x39, 0x8f, 0x22,
	0x9c, 0x8e, 0x63, 0x96, 0x77, 0xaa, 0xf9, 0x35, 0x2d, 0x3c, 0x31, 0x7a, 0x00, 0xad, 0x21, 0x9b,
	0xc7, 0x0c, 0xcd, 0x58, 0x7a, 0x0e, 0xed, 0x42, 0x30, 0xdf, 0x1f, 0x42, 0x35, 0x1b, 0x2c, 0x1d,
	0xab, 0x57, 0xe9, 0xd7, 0x7d, 0x1d, 0x50, 0x01, 0xff, 0x5f, 0x50, 0xca, 0x20, 0x42, 0x72, 0x06,
	0xcd, 0x85, 0x7e, 0x8e, 0xd5, 0x2a, 0xd1, 0x6e, 0xab, 0x7e, 0xc3, 0x68, 0xaf, 0xab, 0x04, 0xc9,
	0x31, 0xd8, 0x12, 0xd9, 0x14, 0x85, 0xb1, 0x6b, 0xa2, 0x6c, 0x11, 0x81, 0x21, 0xc6, 0x4b, 0x14,
	0x4e, 0x45, 0x2f, 0x52, 0xc4, 0xa4, 0x03, 0x95, 0x85, 0x8c, 0x9c, 0x7f, 0xb9, 0x9c, 0x3d, 0xe9,
	0x45, 0xbe, 0x1a, 0x4f, 0xd5, 0x1f, 0x18, 0xd1, 0x01, 0xb4, 0x8b, 0x62, 0xb3, 0xc8, 0x29, 0x80,
	0x01, 0xc1, 0x53, 0x65, 0x48, 0x18, 0x34, 0xc3, 0x54, 0x5d, 0x7f, 0x95, 0xa1, 0x71, 0x3f, 0x0b,
	0xd4, 0x08, 0xc5, 0x32, 0x0e, 0x91, 0xdc, 0x41, 0x35, 0x07, 0x49, 0x5c, 0x6f, 0xf3, 0x50, 0xde,
	0xe6, 0x95, 0xdc, 0xee, 0xde, 0x9c, 0x1e, 0x48, 0x4b, 0xe4, 0x01, 0x6c, 0x4d, 0x93, 0xec, 0x14,
	0x6e, 0x41, 0x77, 0x4f, 0xf6, 0x27, 0xd7, 0x6d, 0x6e, 0xc1, 0x1e, 0xf1, 0xf0, 0x1d, 0x15, 0x39,
	0xda, 0xae, 0x34, 0x27, 0x70, 0xf7, 0xcb, 0xb4, 0xd4, 0xb7, 0xae, 0xac, 0xcc, 0x84, 0x26, 0x41,
	0x7e, 0xbb, 0xfd, 0x81, 0xb9, 0x6b, 0x62, 0x1b, 0x1e, 0x2d, 0x4d, 0xec, 0xfc, 0xaf, 0xbd, 0xf9,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0xda, 0x92, 0x6a, 0xf1, 0xc9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Online(ctx context.Context, in *OnlineRequest, opts ...grpc.CallOption) (*OnlineResponse, error)
	Socket(ctx context.Context, opts ...grpc.CallOption) (ChatService_SocketClient, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/go.jebb.chat.ChatService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Online(ctx context.Context, in *OnlineRequest, opts ...grpc.CallOption) (*OnlineResponse, error) {
	out := new(OnlineResponse)
	err := c.cc.Invoke(ctx, "/go.jebb.chat.ChatService/Online", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Socket(ctx context.Context, opts ...grpc.CallOption) (ChatService_SocketClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/go.jebb.chat.ChatService/Socket", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSocketClient{stream}
	return x, nil
}

type ChatService_SocketClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatServiceSocketClient struct {
	grpc.ClientStream
}

func (x *chatServiceSocketClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSocketClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/go.jebb.chat.ChatService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Online(context.Context, *OnlineRequest) (*OnlineResponse, error)
	Socket(ChatService_SocketServer) error
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.jebb.chat.ChatService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Online_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Online(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.jebb.chat.ChatService/Online",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Online(ctx, req.(*OnlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Socket_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Socket(&chatServiceSocketServer{stream})
}

type ChatService_SocketServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatServiceSocketServer struct {
	grpc.ServerStream
}

func (x *chatServiceSocketServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSocketServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.jebb.chat.ChatService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.jebb.chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _ChatService_Login_Handler,
		},
		{
			MethodName: "Online",
			Handler:    _ChatService_Online_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _ChatService_Logout_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Socket",
			Handler:       _ChatService_Socket_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/chat.proto",
}
