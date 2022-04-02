// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package backend

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

// GeneratorClient is the client API for Generator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeneratorClient interface {
	CreateVersion(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error)
}

type generatorClient struct {
	cc grpc.ClientConnInterface
}

func NewGeneratorClient(cc grpc.ClientConnInterface) GeneratorClient {
	return &generatorClient{cc}
}

func (c *generatorClient) CreateVersion(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := c.cc.Invoke(ctx, "/dtalk.backend.Generator/CreateVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeneratorServer is the server API for Generator service.
// All implementations must embed UnimplementedGeneratorServer
// for forward compatibility
type GeneratorServer interface {
	CreateVersion(context.Context, *CreateRequest) (*CreateReply, error)
	mustEmbedUnimplementedGeneratorServer()
}

// UnimplementedGeneratorServer must be embedded to have forward compatible implementations.
type UnimplementedGeneratorServer struct {
}

func (UnimplementedGeneratorServer) CreateVersion(context.Context, *CreateRequest) (*CreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVersion not implemented")
}
func (UnimplementedGeneratorServer) mustEmbedUnimplementedGeneratorServer() {}

// UnsafeGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeneratorServer will
// result in compilation errors.
type UnsafeGeneratorServer interface {
	mustEmbedUnimplementedGeneratorServer()
}

func RegisterGeneratorServer(s grpc.ServiceRegistrar, srv GeneratorServer) {
	s.RegisterService(&Generator_ServiceDesc, srv)
}

func _Generator_CreateVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).CreateVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtalk.backend.Generator/CreateVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).CreateVersion(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Generator_ServiceDesc is the grpc.ServiceDesc for Generator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Generator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dtalk.backend.Generator",
	HandlerType: (*GeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVersion",
			Handler:    _Generator_CreateVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}