// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: portfolio.proto

package portfolio

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
	PortfolioService_UploadItem_FullMethodName = "/inkMe.portfolio.PortfolioService/UploadItem"
	PortfolioService_ListItems_FullMethodName  = "/inkMe.portfolio.PortfolioService/ListItems"
	PortfolioService_UpdateItem_FullMethodName = "/inkMe.portfolio.PortfolioService/UpdateItem"
)

// PortfolioServiceClient is the client API for PortfolioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortfolioServiceClient interface {
	UploadItem(ctx context.Context, in *UploadItemRequest, opts ...grpc.CallOption) (*UploadItemResponse, error)
	ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (*ListItemsResponse, error)
	UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error)
}

type portfolioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortfolioServiceClient(cc grpc.ClientConnInterface) PortfolioServiceClient {
	return &portfolioServiceClient{cc}
}

func (c *portfolioServiceClient) UploadItem(ctx context.Context, in *UploadItemRequest, opts ...grpc.CallOption) (*UploadItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadItemResponse)
	err := c.cc.Invoke(ctx, PortfolioService_UploadItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (*ListItemsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListItemsResponse)
	err := c.cc.Invoke(ctx, PortfolioService_ListItems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioServiceClient) UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateItemResponse)
	err := c.cc.Invoke(ctx, PortfolioService_UpdateItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortfolioServiceServer is the server API for PortfolioService service.
// All implementations must embed UnimplementedPortfolioServiceServer
// for forward compatibility.
type PortfolioServiceServer interface {
	UploadItem(context.Context, *UploadItemRequest) (*UploadItemResponse, error)
	ListItems(context.Context, *ListItemsRequest) (*ListItemsResponse, error)
	UpdateItem(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error)
	mustEmbedUnimplementedPortfolioServiceServer()
}

// UnimplementedPortfolioServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPortfolioServiceServer struct{}

func (UnimplementedPortfolioServiceServer) UploadItem(context.Context, *UploadItemRequest) (*UploadItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadItem not implemented")
}
func (UnimplementedPortfolioServiceServer) ListItems(context.Context, *ListItemsRequest) (*ListItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListItems not implemented")
}
func (UnimplementedPortfolioServiceServer) UpdateItem(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItem not implemented")
}
func (UnimplementedPortfolioServiceServer) mustEmbedUnimplementedPortfolioServiceServer() {}
func (UnimplementedPortfolioServiceServer) testEmbeddedByValue()                          {}

// UnsafePortfolioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortfolioServiceServer will
// result in compilation errors.
type UnsafePortfolioServiceServer interface {
	mustEmbedUnimplementedPortfolioServiceServer()
}

func RegisterPortfolioServiceServer(s grpc.ServiceRegistrar, srv PortfolioServiceServer) {
	// If the following call pancis, it indicates UnimplementedPortfolioServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PortfolioService_ServiceDesc, srv)
}

func _PortfolioService_UploadItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).UploadItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_UploadItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).UploadItem(ctx, req.(*UploadItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_ListItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).ListItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_ListItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).ListItems(ctx, req.(*ListItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortfolioService_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortfolioService_UpdateItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).UpdateItem(ctx, req.(*UpdateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PortfolioService_ServiceDesc is the grpc.ServiceDesc for PortfolioService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortfolioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inkMe.portfolio.PortfolioService",
	HandlerType: (*PortfolioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadItem",
			Handler:    _PortfolioService_UploadItem_Handler,
		},
		{
			MethodName: "ListItems",
			Handler:    _PortfolioService_ListItems_Handler,
		},
		{
			MethodName: "UpdateItem",
			Handler:    _PortfolioService_UpdateItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portfolio.proto",
}
