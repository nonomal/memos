// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api/v1/resource_service.proto

package apiv1

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ResourceService_CreateResource_FullMethodName    = "/memos.api.v1.ResourceService/CreateResource"
	ResourceService_ListResources_FullMethodName     = "/memos.api.v1.ResourceService/ListResources"
	ResourceService_GetResource_FullMethodName       = "/memos.api.v1.ResourceService/GetResource"
	ResourceService_GetResourceBinary_FullMethodName = "/memos.api.v1.ResourceService/GetResourceBinary"
	ResourceService_UpdateResource_FullMethodName    = "/memos.api.v1.ResourceService/UpdateResource"
	ResourceService_DeleteResource_FullMethodName    = "/memos.api.v1.ResourceService/DeleteResource"
)

// ResourceServiceClient is the client API for ResourceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceServiceClient interface {
	// CreateResource creates a new resource.
	CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...grpc.CallOption) (*Resource, error)
	// ListResources lists all resources.
	ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error)
	// GetResource returns a resource by name.
	GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*Resource, error)
	// GetResourceBinary returns a resource binary by name.
	GetResourceBinary(ctx context.Context, in *GetResourceBinaryRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	// UpdateResource updates a resource.
	UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...grpc.CallOption) (*Resource, error)
	// DeleteResource deletes a resource by name.
	DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type resourceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceServiceClient(cc grpc.ClientConnInterface) ResourceServiceClient {
	return &resourceServiceClient{cc}
}

func (c *resourceServiceClient) CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...grpc.CallOption) (*Resource, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Resource)
	err := c.cc.Invoke(ctx, ResourceService_CreateResource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResourcesResponse)
	err := c.cc.Invoke(ctx, ResourceService_ListResources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*Resource, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Resource)
	err := c.cc.Invoke(ctx, ResourceService_GetResource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) GetResourceBinary(ctx context.Context, in *GetResourceBinaryRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, ResourceService_GetResourceBinary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...grpc.CallOption) (*Resource, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Resource)
	err := c.cc.Invoke(ctx, ResourceService_UpdateResource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ResourceService_DeleteResource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceServiceServer is the server API for ResourceService service.
// All implementations must embed UnimplementedResourceServiceServer
// for forward compatibility.
type ResourceServiceServer interface {
	// CreateResource creates a new resource.
	CreateResource(context.Context, *CreateResourceRequest) (*Resource, error)
	// ListResources lists all resources.
	ListResources(context.Context, *ListResourcesRequest) (*ListResourcesResponse, error)
	// GetResource returns a resource by name.
	GetResource(context.Context, *GetResourceRequest) (*Resource, error)
	// GetResourceBinary returns a resource binary by name.
	GetResourceBinary(context.Context, *GetResourceBinaryRequest) (*httpbody.HttpBody, error)
	// UpdateResource updates a resource.
	UpdateResource(context.Context, *UpdateResourceRequest) (*Resource, error)
	// DeleteResource deletes a resource by name.
	DeleteResource(context.Context, *DeleteResourceRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedResourceServiceServer()
}

// UnimplementedResourceServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedResourceServiceServer struct{}

func (UnimplementedResourceServiceServer) CreateResource(context.Context, *CreateResourceRequest) (*Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}
func (UnimplementedResourceServiceServer) ListResources(context.Context, *ListResourcesRequest) (*ListResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResources not implemented")
}
func (UnimplementedResourceServiceServer) GetResource(context.Context, *GetResourceRequest) (*Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResource not implemented")
}
func (UnimplementedResourceServiceServer) GetResourceBinary(context.Context, *GetResourceBinaryRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResourceBinary not implemented")
}
func (UnimplementedResourceServiceServer) UpdateResource(context.Context, *UpdateResourceRequest) (*Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResource not implemented")
}
func (UnimplementedResourceServiceServer) DeleteResource(context.Context, *DeleteResourceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}
func (UnimplementedResourceServiceServer) mustEmbedUnimplementedResourceServiceServer() {}
func (UnimplementedResourceServiceServer) testEmbeddedByValue()                         {}

// UnsafeResourceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceServiceServer will
// result in compilation errors.
type UnsafeResourceServiceServer interface {
	mustEmbedUnimplementedResourceServiceServer()
}

func RegisterResourceServiceServer(s grpc.ServiceRegistrar, srv ResourceServiceServer) {
	// If the following call pancis, it indicates UnimplementedResourceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ResourceService_ServiceDesc, srv)
}

func _ResourceService_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceService_CreateResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).CreateResource(ctx, req.(*CreateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_ListResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).ListResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceService_ListResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).ListResources(ctx, req.(*ListResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceService_GetResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).GetResource(ctx, req.(*GetResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_GetResourceBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).GetResourceBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceService_GetResourceBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).GetResourceBinary(ctx, req.(*GetResourceBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_UpdateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).UpdateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceService_UpdateResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).UpdateResource(ctx, req.(*UpdateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceService_DeleteResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).DeleteResource(ctx, req.(*DeleteResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceService_ServiceDesc is the grpc.ServiceDesc for ResourceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "memos.api.v1.ResourceService",
	HandlerType: (*ResourceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResource",
			Handler:    _ResourceService_CreateResource_Handler,
		},
		{
			MethodName: "ListResources",
			Handler:    _ResourceService_ListResources_Handler,
		},
		{
			MethodName: "GetResource",
			Handler:    _ResourceService_GetResource_Handler,
		},
		{
			MethodName: "GetResourceBinary",
			Handler:    _ResourceService_GetResourceBinary_Handler,
		},
		{
			MethodName: "UpdateResource",
			Handler:    _ResourceService_UpdateResource_Handler,
		},
		{
			MethodName: "DeleteResource",
			Handler:    _ResourceService_DeleteResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/resource_service.proto",
}
