// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: pkg/proto/observiability.proto

package observer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerClient interface {
	GetPluginNames(ctx context.Context, in *GetPluginNameRequest, opts ...grpc.CallOption) (*GetPluginNameResponse, error)
	PluginCapabilities(ctx context.Context, in *PluginCapabilitiesRequest, opts ...grpc.CallOption) (*PluginCapabilitiesResponse, error)
	GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
}

type serverClient struct {
	cc grpc.ClientConnInterface
}

func NewServerClient(cc grpc.ClientConnInterface) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) GetPluginNames(ctx context.Context, in *GetPluginNameRequest, opts ...grpc.CallOption) (*GetPluginNameResponse, error) {
	out := new(GetPluginNameResponse)
	err := c.cc.Invoke(ctx, "/obi.v1.Server/GetPluginNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) PluginCapabilities(ctx context.Context, in *PluginCapabilitiesRequest, opts ...grpc.CallOption) (*PluginCapabilitiesResponse, error) {
	out := new(PluginCapabilitiesResponse)
	err := c.cc.Invoke(ctx, "/obi.v1.Server/PluginCapabilities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, "/obi.v1.Server/GetMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
// All implementations must embed UnimplementedServerServer
// for forward compatibility
type ServerServer interface {
	GetPluginNames(context.Context, *GetPluginNameRequest) (*GetPluginNameResponse, error)
	PluginCapabilities(context.Context, *PluginCapabilitiesRequest) (*PluginCapabilitiesResponse, error)
	GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
	mustEmbedUnimplementedServerServer()
}

// UnimplementedServerServer must be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (UnimplementedServerServer) GetPluginNames(context.Context, *GetPluginNameRequest) (*GetPluginNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginNames not implemented")
}
func (UnimplementedServerServer) PluginCapabilities(context.Context, *PluginCapabilitiesRequest) (*PluginCapabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginCapabilities not implemented")
}
func (UnimplementedServerServer) GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}
func (UnimplementedServerServer) mustEmbedUnimplementedServerServer() {}

// UnsafeServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerServer will
// result in compilation errors.
type UnsafeServerServer interface {
	mustEmbedUnimplementedServerServer()
}

func RegisterServerServer(s grpc.ServiceRegistrar, srv ServerServer) {
	s.RegisterService(&Server_ServiceDesc, srv)
}

func _Server_GetPluginNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPluginNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).GetPluginNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/obi.v1.Server/GetPluginNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).GetPluginNames(ctx, req.(*GetPluginNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_PluginCapabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginCapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).PluginCapabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/obi.v1.Server/PluginCapabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).PluginCapabilities(ctx, req.(*PluginCapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/obi.v1.Server/GetMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).GetMetrics(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Server_ServiceDesc is the grpc.ServiceDesc for Server service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Server_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "obi.v1.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPluginNames",
			Handler:    _Server_GetPluginNames_Handler,
		},
		{
			MethodName: "PluginCapabilities",
			Handler:    _Server_PluginCapabilities_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _Server_GetMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/observiability.proto",
}
