// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package device

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

// DeviceSrvClient is the client API for DeviceSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceSrvClient interface {
	AddDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Empty, error)
	EnableThreadPushDevice(ctx context.Context, in *EnableThreadPushDeviceRequest, opts ...grpc.CallOption) (*Empty, error)
	GetUserAllDevices(ctx context.Context, in *GetUserAllDevicesRequest, opts ...grpc.CallOption) (*GetUserAllDevicesReply, error)
	GetDeviceByConnId(ctx context.Context, in *GetDeviceByConnIdRequest, opts ...grpc.CallOption) (*Device, error)
}

type deviceSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceSrvClient(cc grpc.ClientConnInterface) DeviceSrvClient {
	return &deviceSrvClient{cc}
}

func (c *deviceSrvClient) AddDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dtalk.device.DeviceSrv/AddDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceSrvClient) EnableThreadPushDevice(ctx context.Context, in *EnableThreadPushDeviceRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dtalk.device.DeviceSrv/EnableThreadPushDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceSrvClient) GetUserAllDevices(ctx context.Context, in *GetUserAllDevicesRequest, opts ...grpc.CallOption) (*GetUserAllDevicesReply, error) {
	out := new(GetUserAllDevicesReply)
	err := c.cc.Invoke(ctx, "/dtalk.device.DeviceSrv/GetUserAllDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceSrvClient) GetDeviceByConnId(ctx context.Context, in *GetDeviceByConnIdRequest, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/dtalk.device.DeviceSrv/GetDeviceByConnId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceSrvServer is the server API for DeviceSrv service.
// All implementations must embed UnimplementedDeviceSrvServer
// for forward compatibility
type DeviceSrvServer interface {
	AddDevice(context.Context, *Device) (*Empty, error)
	EnableThreadPushDevice(context.Context, *EnableThreadPushDeviceRequest) (*Empty, error)
	GetUserAllDevices(context.Context, *GetUserAllDevicesRequest) (*GetUserAllDevicesReply, error)
	GetDeviceByConnId(context.Context, *GetDeviceByConnIdRequest) (*Device, error)
	mustEmbedUnimplementedDeviceSrvServer()
}

// UnimplementedDeviceSrvServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceSrvServer struct {
}

func (UnimplementedDeviceSrvServer) AddDevice(context.Context, *Device) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDevice not implemented")
}
func (UnimplementedDeviceSrvServer) EnableThreadPushDevice(context.Context, *EnableThreadPushDeviceRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableThreadPushDevice not implemented")
}
func (UnimplementedDeviceSrvServer) GetUserAllDevices(context.Context, *GetUserAllDevicesRequest) (*GetUserAllDevicesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAllDevices not implemented")
}
func (UnimplementedDeviceSrvServer) GetDeviceByConnId(context.Context, *GetDeviceByConnIdRequest) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceByConnId not implemented")
}
func (UnimplementedDeviceSrvServer) mustEmbedUnimplementedDeviceSrvServer() {}

// UnsafeDeviceSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceSrvServer will
// result in compilation errors.
type UnsafeDeviceSrvServer interface {
	mustEmbedUnimplementedDeviceSrvServer()
}

func RegisterDeviceSrvServer(s grpc.ServiceRegistrar, srv DeviceSrvServer) {
	s.RegisterService(&DeviceSrv_ServiceDesc, srv)
}

func _DeviceSrv_AddDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceSrvServer).AddDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtalk.device.DeviceSrv/AddDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceSrvServer).AddDevice(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceSrv_EnableThreadPushDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableThreadPushDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceSrvServer).EnableThreadPushDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtalk.device.DeviceSrv/EnableThreadPushDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceSrvServer).EnableThreadPushDevice(ctx, req.(*EnableThreadPushDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceSrv_GetUserAllDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAllDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceSrvServer).GetUserAllDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtalk.device.DeviceSrv/GetUserAllDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceSrvServer).GetUserAllDevices(ctx, req.(*GetUserAllDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceSrv_GetDeviceByConnId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceByConnIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceSrvServer).GetDeviceByConnId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtalk.device.DeviceSrv/GetDeviceByConnId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceSrvServer).GetDeviceByConnId(ctx, req.(*GetDeviceByConnIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceSrv_ServiceDesc is the grpc.ServiceDesc for DeviceSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dtalk.device.DeviceSrv",
	HandlerType: (*DeviceSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDevice",
			Handler:    _DeviceSrv_AddDevice_Handler,
		},
		{
			MethodName: "EnableThreadPushDevice",
			Handler:    _DeviceSrv_EnableThreadPushDevice_Handler,
		},
		{
			MethodName: "GetUserAllDevices",
			Handler:    _DeviceSrv_GetUserAllDevices_Handler,
		},
		{
			MethodName: "GetDeviceByConnId",
			Handler:    _DeviceSrv_GetDeviceByConnId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
