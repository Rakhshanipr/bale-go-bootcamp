// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: messaging.proto

package grpc2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Messaging_AddUser_FullMethodName         = "/Messaging/AddUser"
	Messaging_SendMessage_FullMethodName     = "/Messaging/SendMessage"
	Messaging_FetchMessage_FullMethodName    = "/Messaging/FetchMessage"
	Messaging_GetUserMessages_FullMethodName = "/Messaging/GetUserMessages"
)

// MessagingClient is the client API for Messaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagingClient interface {
	AddUser(ctx context.Context, in *RequestAddUser, opts ...grpc.CallOption) (*ResponseAddUser, error)
	SendMessage(ctx context.Context, in *RequestSendMessage, opts ...grpc.CallOption) (*ResponseSendMessage, error)
	FetchMessage(ctx context.Context, in *RequestFetchMessage, opts ...grpc.CallOption) (*ResponseFetchMessage, error)
	GetUserMessages(ctx context.Context, in *RequestGetUserMessages, opts ...grpc.CallOption) (*ResponseGetUserMessages, error)
}

type messagingClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingClient(cc grpc.ClientConnInterface) MessagingClient {
	return &messagingClient{cc}
}

func (c *messagingClient) AddUser(ctx context.Context, in *RequestAddUser, opts ...grpc.CallOption) (*ResponseAddUser, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseAddUser)
	err := c.cc.Invoke(ctx, Messaging_AddUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) SendMessage(ctx context.Context, in *RequestSendMessage, opts ...grpc.CallOption) (*ResponseSendMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseSendMessage)
	err := c.cc.Invoke(ctx, Messaging_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) FetchMessage(ctx context.Context, in *RequestFetchMessage, opts ...grpc.CallOption) (*ResponseFetchMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseFetchMessage)
	err := c.cc.Invoke(ctx, Messaging_FetchMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) GetUserMessages(ctx context.Context, in *RequestGetUserMessages, opts ...grpc.CallOption) (*ResponseGetUserMessages, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseGetUserMessages)
	err := c.cc.Invoke(ctx, Messaging_GetUserMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagingServer is the server API for Messaging service.
// All implementations must embed UnimplementedMessagingServer
// for forward compatibility.
type MessagingServer interface {
	AddUser(context.Context, *RequestAddUser) (*ResponseAddUser, error)
	SendMessage(context.Context, *RequestSendMessage) (*ResponseSendMessage, error)
	FetchMessage(context.Context, *RequestFetchMessage) (*ResponseFetchMessage, error)
	GetUserMessages(context.Context, *RequestGetUserMessages) (*ResponseGetUserMessages, error)
	mustEmbedUnimplementedMessagingServer()
}

// UnimplementedMessagingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessagingServer struct{}

func (UnimplementedMessagingServer) AddUser(context.Context, *RequestAddUser) (*ResponseAddUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedMessagingServer) SendMessage(context.Context, *RequestSendMessage) (*ResponseSendMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessagingServer) FetchMessage(context.Context, *RequestFetchMessage) (*ResponseFetchMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchMessage not implemented")
}
func (UnimplementedMessagingServer) GetUserMessages(context.Context, *RequestGetUserMessages) (*ResponseGetUserMessages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMessages not implemented")
}
func (UnimplementedMessagingServer) mustEmbedUnimplementedMessagingServer() {}
func (UnimplementedMessagingServer) testEmbeddedByValue()                   {}

// UnsafeMessagingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingServer will
// result in compilation errors.
type UnsafeMessagingServer interface {
	mustEmbedUnimplementedMessagingServer()
}

func RegisterMessagingServer(s grpc.ServiceRegistrar, srv MessagingServer) {
	// If the following call pancis, it indicates UnimplementedMessagingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Messaging_ServiceDesc, srv)
}

func _Messaging_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAddUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).AddUser(ctx, req.(*RequestAddUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSendMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).SendMessage(ctx, req.(*RequestSendMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_FetchMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestFetchMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).FetchMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_FetchMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).FetchMessage(ctx, req.(*RequestFetchMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_GetUserMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetUserMessages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).GetUserMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_GetUserMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).GetUserMessages(ctx, req.(*RequestGetUserMessages))
	}
	return interceptor(ctx, in, info, handler)
}

// Messaging_ServiceDesc is the grpc.ServiceDesc for Messaging service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messaging_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Messaging",
	HandlerType: (*MessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _Messaging_AddUser_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Messaging_SendMessage_Handler,
		},
		{
			MethodName: "FetchMessage",
			Handler:    _Messaging_FetchMessage_Handler,
		},
		{
			MethodName: "GetUserMessages",
			Handler:    _Messaging_GetUserMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messaging.proto",
}
