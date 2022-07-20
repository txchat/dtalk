// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pusher

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

// PusherClient is the client API for Pusher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PusherClient interface {
	PushClient(ctx context.Context, in *PushReq, opts ...grpc.CallOption) (*PushReply, error)
}

type pusherClient struct {
	cc grpc.ClientConnInterface
}

func NewPusherClient(cc grpc.ClientConnInterface) PusherClient {
	return &pusherClient{cc}
}

func (c *pusherClient) PushClient(ctx context.Context, in *PushReq, opts ...grpc.CallOption) (*PushReply, error) {
	out := new(PushReply)
	err := c.cc.Invoke(ctx, "/dtalk.pusher.Pusher/PushClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PusherServer is the server API for Pusher service.
// All implementations must embed UnimplementedPusherServer
// for forward compatibility
type PusherServer interface {
	PushClient(context.Context, *PushReq) (*PushReply, error)
	mustEmbedUnimplementedPusherServer()
}

// UnimplementedPusherServer must be embedded to have forward compatible implementations.
type UnimplementedPusherServer struct {
}

func (UnimplementedPusherServer) PushClient(context.Context, *PushReq) (*PushReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushClient not implemented")
}
func (UnimplementedPusherServer) mustEmbedUnimplementedPusherServer() {}

// UnsafePusherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PusherServer will
// result in compilation errors.
type UnsafePusherServer interface {
	mustEmbedUnimplementedPusherServer()
}

func RegisterPusherServer(s grpc.ServiceRegistrar, srv PusherServer) {
	s.RegisterService(&Pusher_ServiceDesc, srv)
}

func _Pusher_PushClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PusherServer).PushClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtalk.pusher.Pusher/PushClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PusherServer).PushClient(ctx, req.(*PushReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Pusher_ServiceDesc is the grpc.ServiceDesc for Pusher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pusher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dtalk.pusher.Pusher",
	HandlerType: (*PusherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushClient",
			Handler:    _Pusher_PushClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
