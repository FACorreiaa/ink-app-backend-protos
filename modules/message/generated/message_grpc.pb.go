// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: message.proto

package message

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
	MessagingService_SendMessage_FullMethodName           = "/inkMe.message.MessagingService/SendMessage"
	MessagingService_StreamMessages_FullMethodName        = "/inkMe.message.MessagingService/StreamMessages"
	MessagingService_GetMessage_FullMethodName            = "/inkMe.message.MessagingService/GetMessage"
	MessagingService_ListMessages_FullMethodName          = "/inkMe.message.MessagingService/ListMessages"
	MessagingService_MarkMessageAsRead_FullMethodName     = "/inkMe.message.MessagingService/MarkMessageAsRead"
	MessagingService_CreateMessageThread_FullMethodName   = "/inkMe.message.MessagingService/CreateMessageThread"
	MessagingService_GetMessageThread_FullMethodName      = "/inkMe.message.MessagingService/GetMessageThread"
	MessagingService_ListMessageThreads_FullMethodName    = "/inkMe.message.MessagingService/ListMessageThreads"
	MessagingService_CreateMessageTemplate_FullMethodName = "/inkMe.message.MessagingService/CreateMessageTemplate"
	MessagingService_ListMessageTemplates_FullMethodName  = "/inkMe.message.MessagingService/ListMessageTemplates"
	MessagingService_SendTemplatedMessage_FullMethodName  = "/inkMe.message.MessagingService/SendTemplatedMessage"
)

// MessagingServiceClient is the client API for MessagingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagingServiceClient interface {
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	StreamMessages(ctx context.Context, in *StreamMessagesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Message], error)
	GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponse, error)
	ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (*ListMessagesResponse, error)
	MarkMessageAsRead(ctx context.Context, in *MarkMessageAsReadRequest, opts ...grpc.CallOption) (*MarkMessageAsReadResponse, error)
	// Thread management
	CreateMessageThread(ctx context.Context, in *CreateMessageThreadRequest, opts ...grpc.CallOption) (*CreateMessageThreadResponse, error)
	GetMessageThread(ctx context.Context, in *GetMessageThreadRequest, opts ...grpc.CallOption) (*GetMessageThreadResponse, error)
	ListMessageThreads(ctx context.Context, in *ListMessageThreadsRequest, opts ...grpc.CallOption) (*ListMessageThreadsResponse, error)
	// Templates for common messages
	CreateMessageTemplate(ctx context.Context, in *CreateMessageTemplateRequest, opts ...grpc.CallOption) (*CreateMessageTemplateResponse, error)
	ListMessageTemplates(ctx context.Context, in *ListMessageTemplatesRequest, opts ...grpc.CallOption) (*ListMessageTemplatesResponse, error)
	SendTemplatedMessage(ctx context.Context, in *SendTemplatedMessageRequest, opts ...grpc.CallOption) (*SendTemplatedMessageResponse, error)
}

type messagingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingServiceClient(cc grpc.ClientConnInterface) MessagingServiceClient {
	return &messagingServiceClient{cc}
}

func (c *messagingServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, MessagingService_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) StreamMessages(ctx context.Context, in *StreamMessagesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Message], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MessagingService_ServiceDesc.Streams[0], MessagingService_StreamMessages_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamMessagesRequest, Message]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MessagingService_StreamMessagesClient = grpc.ServerStreamingClient[Message]

func (c *messagingServiceClient) GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessageResponse)
	err := c.cc.Invoke(ctx, MessagingService_GetMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (*ListMessagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMessagesResponse)
	err := c.cc.Invoke(ctx, MessagingService_ListMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) MarkMessageAsRead(ctx context.Context, in *MarkMessageAsReadRequest, opts ...grpc.CallOption) (*MarkMessageAsReadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarkMessageAsReadResponse)
	err := c.cc.Invoke(ctx, MessagingService_MarkMessageAsRead_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) CreateMessageThread(ctx context.Context, in *CreateMessageThreadRequest, opts ...grpc.CallOption) (*CreateMessageThreadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMessageThreadResponse)
	err := c.cc.Invoke(ctx, MessagingService_CreateMessageThread_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) GetMessageThread(ctx context.Context, in *GetMessageThreadRequest, opts ...grpc.CallOption) (*GetMessageThreadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessageThreadResponse)
	err := c.cc.Invoke(ctx, MessagingService_GetMessageThread_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) ListMessageThreads(ctx context.Context, in *ListMessageThreadsRequest, opts ...grpc.CallOption) (*ListMessageThreadsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMessageThreadsResponse)
	err := c.cc.Invoke(ctx, MessagingService_ListMessageThreads_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) CreateMessageTemplate(ctx context.Context, in *CreateMessageTemplateRequest, opts ...grpc.CallOption) (*CreateMessageTemplateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMessageTemplateResponse)
	err := c.cc.Invoke(ctx, MessagingService_CreateMessageTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) ListMessageTemplates(ctx context.Context, in *ListMessageTemplatesRequest, opts ...grpc.CallOption) (*ListMessageTemplatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMessageTemplatesResponse)
	err := c.cc.Invoke(ctx, MessagingService_ListMessageTemplates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) SendTemplatedMessage(ctx context.Context, in *SendTemplatedMessageRequest, opts ...grpc.CallOption) (*SendTemplatedMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendTemplatedMessageResponse)
	err := c.cc.Invoke(ctx, MessagingService_SendTemplatedMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagingServiceServer is the server API for MessagingService service.
// All implementations must embed UnimplementedMessagingServiceServer
// for forward compatibility.
type MessagingServiceServer interface {
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	StreamMessages(*StreamMessagesRequest, grpc.ServerStreamingServer[Message]) error
	GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error)
	ListMessages(context.Context, *ListMessagesRequest) (*ListMessagesResponse, error)
	MarkMessageAsRead(context.Context, *MarkMessageAsReadRequest) (*MarkMessageAsReadResponse, error)
	// Thread management
	CreateMessageThread(context.Context, *CreateMessageThreadRequest) (*CreateMessageThreadResponse, error)
	GetMessageThread(context.Context, *GetMessageThreadRequest) (*GetMessageThreadResponse, error)
	ListMessageThreads(context.Context, *ListMessageThreadsRequest) (*ListMessageThreadsResponse, error)
	// Templates for common messages
	CreateMessageTemplate(context.Context, *CreateMessageTemplateRequest) (*CreateMessageTemplateResponse, error)
	ListMessageTemplates(context.Context, *ListMessageTemplatesRequest) (*ListMessageTemplatesResponse, error)
	SendTemplatedMessage(context.Context, *SendTemplatedMessageRequest) (*SendTemplatedMessageResponse, error)
	mustEmbedUnimplementedMessagingServiceServer()
}

// UnimplementedMessagingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessagingServiceServer struct{}

func (UnimplementedMessagingServiceServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessagingServiceServer) StreamMessages(*StreamMessagesRequest, grpc.ServerStreamingServer[Message]) error {
	return status.Errorf(codes.Unimplemented, "method StreamMessages not implemented")
}
func (UnimplementedMessagingServiceServer) GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedMessagingServiceServer) ListMessages(context.Context, *ListMessagesRequest) (*ListMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessages not implemented")
}
func (UnimplementedMessagingServiceServer) MarkMessageAsRead(context.Context, *MarkMessageAsReadRequest) (*MarkMessageAsReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkMessageAsRead not implemented")
}
func (UnimplementedMessagingServiceServer) CreateMessageThread(context.Context, *CreateMessageThreadRequest) (*CreateMessageThreadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessageThread not implemented")
}
func (UnimplementedMessagingServiceServer) GetMessageThread(context.Context, *GetMessageThreadRequest) (*GetMessageThreadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageThread not implemented")
}
func (UnimplementedMessagingServiceServer) ListMessageThreads(context.Context, *ListMessageThreadsRequest) (*ListMessageThreadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessageThreads not implemented")
}
func (UnimplementedMessagingServiceServer) CreateMessageTemplate(context.Context, *CreateMessageTemplateRequest) (*CreateMessageTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessageTemplate not implemented")
}
func (UnimplementedMessagingServiceServer) ListMessageTemplates(context.Context, *ListMessageTemplatesRequest) (*ListMessageTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessageTemplates not implemented")
}
func (UnimplementedMessagingServiceServer) SendTemplatedMessage(context.Context, *SendTemplatedMessageRequest) (*SendTemplatedMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTemplatedMessage not implemented")
}
func (UnimplementedMessagingServiceServer) mustEmbedUnimplementedMessagingServiceServer() {}
func (UnimplementedMessagingServiceServer) testEmbeddedByValue()                          {}

// UnsafeMessagingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingServiceServer will
// result in compilation errors.
type UnsafeMessagingServiceServer interface {
	mustEmbedUnimplementedMessagingServiceServer()
}

func RegisterMessagingServiceServer(s grpc.ServiceRegistrar, srv MessagingServiceServer) {
	// If the following call pancis, it indicates UnimplementedMessagingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MessagingService_ServiceDesc, srv)
}

func _MessagingService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_StreamMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamMessagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessagingServiceServer).StreamMessages(m, &grpc.GenericServerStream[StreamMessagesRequest, Message]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MessagingService_StreamMessagesServer = grpc.ServerStreamingServer[Message]

func _MessagingService_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_GetMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).GetMessage(ctx, req.(*GetMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_ListMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).ListMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_ListMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).ListMessages(ctx, req.(*ListMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_MarkMessageAsRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkMessageAsReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).MarkMessageAsRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_MarkMessageAsRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).MarkMessageAsRead(ctx, req.(*MarkMessageAsReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_CreateMessageThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).CreateMessageThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_CreateMessageThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).CreateMessageThread(ctx, req.(*CreateMessageThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_GetMessageThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).GetMessageThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_GetMessageThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).GetMessageThread(ctx, req.(*GetMessageThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_ListMessageThreads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessageThreadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).ListMessageThreads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_ListMessageThreads_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).ListMessageThreads(ctx, req.(*ListMessageThreadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_CreateMessageTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).CreateMessageTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_CreateMessageTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).CreateMessageTemplate(ctx, req.(*CreateMessageTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_ListMessageTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessageTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).ListMessageTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_ListMessageTemplates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).ListMessageTemplates(ctx, req.(*ListMessageTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_SendTemplatedMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTemplatedMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).SendTemplatedMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_SendTemplatedMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).SendTemplatedMessage(ctx, req.(*SendTemplatedMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessagingService_ServiceDesc is the grpc.ServiceDesc for MessagingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessagingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inkMe.message.MessagingService",
	HandlerType: (*MessagingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _MessagingService_SendMessage_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _MessagingService_GetMessage_Handler,
		},
		{
			MethodName: "ListMessages",
			Handler:    _MessagingService_ListMessages_Handler,
		},
		{
			MethodName: "MarkMessageAsRead",
			Handler:    _MessagingService_MarkMessageAsRead_Handler,
		},
		{
			MethodName: "CreateMessageThread",
			Handler:    _MessagingService_CreateMessageThread_Handler,
		},
		{
			MethodName: "GetMessageThread",
			Handler:    _MessagingService_GetMessageThread_Handler,
		},
		{
			MethodName: "ListMessageThreads",
			Handler:    _MessagingService_ListMessageThreads_Handler,
		},
		{
			MethodName: "CreateMessageTemplate",
			Handler:    _MessagingService_CreateMessageTemplate_Handler,
		},
		{
			MethodName: "ListMessageTemplates",
			Handler:    _MessagingService_ListMessageTemplates_Handler,
		},
		{
			MethodName: "SendTemplatedMessage",
			Handler:    _MessagingService_SendTemplatedMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMessages",
			Handler:       _MessagingService_StreamMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "message.proto",
}
