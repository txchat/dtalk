// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package version

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

// VersionClient is the client API for Version service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VersionClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error)
	Query(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*QueryResp, error)
	ReleaseSpecificVersion(ctx context.Context, in *ReleaseSpecificVersionReq, opts ...grpc.CallOption) (*ReleaseSpecificVersionResp, error)
	SpecificPlatformAndDeviceTypeVersions(ctx context.Context, in *SpecificPlatformAndDeviceTypeVersionsReq, opts ...grpc.CallOption) (*SpecificPlatformAndDeviceTypeVersionsReqResp, error)
	SpecificPlatformAndDeviceTypeCount(ctx context.Context, in *SpecificPlatformAndDeviceTypeCountReq, opts ...grpc.CallOption) (*SpecificPlatformAndDeviceTypeCountResp, error)
	LastReleaseVersion(ctx context.Context, in *LastReleaseVersionReq, opts ...grpc.CallOption) (*LastReleaseVersionResp, error)
	ForceNumberBetween(ctx context.Context, in *ForceNumberBetweenReq, opts ...grpc.CallOption) (*ForceNumberBetweenResp, error)
}

type versionClient struct {
	cc grpc.ClientConnInterface
}

func NewVersionClient(cc grpc.ClientConnInterface) VersionClient {
	return &versionClient{cc}
}

func (c *versionClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	out := new(CreateResp)
	err := c.cc.Invoke(ctx, "/version.Version/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error) {
	out := new(UpdateResp)
	err := c.cc.Invoke(ctx, "/version.Version/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionClient) Query(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*QueryResp, error) {
	out := new(QueryResp)
	err := c.cc.Invoke(ctx, "/version.Version/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionClient) ReleaseSpecificVersion(ctx context.Context, in *ReleaseSpecificVersionReq, opts ...grpc.CallOption) (*ReleaseSpecificVersionResp, error) {
	out := new(ReleaseSpecificVersionResp)
	err := c.cc.Invoke(ctx, "/version.Version/ReleaseSpecificVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionClient) SpecificPlatformAndDeviceTypeVersions(ctx context.Context, in *SpecificPlatformAndDeviceTypeVersionsReq, opts ...grpc.CallOption) (*SpecificPlatformAndDeviceTypeVersionsReqResp, error) {
	out := new(SpecificPlatformAndDeviceTypeVersionsReqResp)
	err := c.cc.Invoke(ctx, "/version.Version/SpecificPlatformAndDeviceTypeVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionClient) SpecificPlatformAndDeviceTypeCount(ctx context.Context, in *SpecificPlatformAndDeviceTypeCountReq, opts ...grpc.CallOption) (*SpecificPlatformAndDeviceTypeCountResp, error) {
	out := new(SpecificPlatformAndDeviceTypeCountResp)
	err := c.cc.Invoke(ctx, "/version.Version/SpecificPlatformAndDeviceTypeCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionClient) LastReleaseVersion(ctx context.Context, in *LastReleaseVersionReq, opts ...grpc.CallOption) (*LastReleaseVersionResp, error) {
	out := new(LastReleaseVersionResp)
	err := c.cc.Invoke(ctx, "/version.Version/LastReleaseVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionClient) ForceNumberBetween(ctx context.Context, in *ForceNumberBetweenReq, opts ...grpc.CallOption) (*ForceNumberBetweenResp, error) {
	out := new(ForceNumberBetweenResp)
	err := c.cc.Invoke(ctx, "/version.Version/ForceNumberBetween", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionServer is the server API for Version service.
// All implementations must embed UnimplementedVersionServer
// for forward compatibility
type VersionServer interface {
	Create(context.Context, *CreateReq) (*CreateResp, error)
	Update(context.Context, *UpdateReq) (*UpdateResp, error)
	Query(context.Context, *QueryReq) (*QueryResp, error)
	ReleaseSpecificVersion(context.Context, *ReleaseSpecificVersionReq) (*ReleaseSpecificVersionResp, error)
	SpecificPlatformAndDeviceTypeVersions(context.Context, *SpecificPlatformAndDeviceTypeVersionsReq) (*SpecificPlatformAndDeviceTypeVersionsReqResp, error)
	SpecificPlatformAndDeviceTypeCount(context.Context, *SpecificPlatformAndDeviceTypeCountReq) (*SpecificPlatformAndDeviceTypeCountResp, error)
	LastReleaseVersion(context.Context, *LastReleaseVersionReq) (*LastReleaseVersionResp, error)
	ForceNumberBetween(context.Context, *ForceNumberBetweenReq) (*ForceNumberBetweenResp, error)
	mustEmbedUnimplementedVersionServer()
}

// UnimplementedVersionServer must be embedded to have forward compatible implementations.
type UnimplementedVersionServer struct {
}

func (UnimplementedVersionServer) Create(context.Context, *CreateReq) (*CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedVersionServer) Update(context.Context, *UpdateReq) (*UpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedVersionServer) Query(context.Context, *QueryReq) (*QueryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedVersionServer) ReleaseSpecificVersion(context.Context, *ReleaseSpecificVersionReq) (*ReleaseSpecificVersionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseSpecificVersion not implemented")
}
func (UnimplementedVersionServer) SpecificPlatformAndDeviceTypeVersions(context.Context, *SpecificPlatformAndDeviceTypeVersionsReq) (*SpecificPlatformAndDeviceTypeVersionsReqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpecificPlatformAndDeviceTypeVersions not implemented")
}
func (UnimplementedVersionServer) SpecificPlatformAndDeviceTypeCount(context.Context, *SpecificPlatformAndDeviceTypeCountReq) (*SpecificPlatformAndDeviceTypeCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpecificPlatformAndDeviceTypeCount not implemented")
}
func (UnimplementedVersionServer) LastReleaseVersion(context.Context, *LastReleaseVersionReq) (*LastReleaseVersionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastReleaseVersion not implemented")
}
func (UnimplementedVersionServer) ForceNumberBetween(context.Context, *ForceNumberBetweenReq) (*ForceNumberBetweenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForceNumberBetween not implemented")
}
func (UnimplementedVersionServer) mustEmbedUnimplementedVersionServer() {}

// UnsafeVersionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VersionServer will
// result in compilation errors.
type UnsafeVersionServer interface {
	mustEmbedUnimplementedVersionServer()
}

func RegisterVersionServer(s grpc.ServiceRegistrar, srv VersionServer) {
	s.RegisterService(&Version_ServiceDesc, srv)
}

func _Version_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.Version/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.Version/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.Version/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServer).Query(ctx, req.(*QueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version_ReleaseSpecificVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseSpecificVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServer).ReleaseSpecificVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.Version/ReleaseSpecificVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServer).ReleaseSpecificVersion(ctx, req.(*ReleaseSpecificVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version_SpecificPlatformAndDeviceTypeVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpecificPlatformAndDeviceTypeVersionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServer).SpecificPlatformAndDeviceTypeVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.Version/SpecificPlatformAndDeviceTypeVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServer).SpecificPlatformAndDeviceTypeVersions(ctx, req.(*SpecificPlatformAndDeviceTypeVersionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version_SpecificPlatformAndDeviceTypeCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpecificPlatformAndDeviceTypeCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServer).SpecificPlatformAndDeviceTypeCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.Version/SpecificPlatformAndDeviceTypeCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServer).SpecificPlatformAndDeviceTypeCount(ctx, req.(*SpecificPlatformAndDeviceTypeCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version_LastReleaseVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastReleaseVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServer).LastReleaseVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.Version/LastReleaseVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServer).LastReleaseVersion(ctx, req.(*LastReleaseVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version_ForceNumberBetween_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForceNumberBetweenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServer).ForceNumberBetween(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.Version/ForceNumberBetween",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServer).ForceNumberBetween(ctx, req.(*ForceNumberBetweenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Version_ServiceDesc is the grpc.ServiceDesc for Version service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Version_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "version.Version",
	HandlerType: (*VersionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Version_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Version_Update_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Version_Query_Handler,
		},
		{
			MethodName: "ReleaseSpecificVersion",
			Handler:    _Version_ReleaseSpecificVersion_Handler,
		},
		{
			MethodName: "SpecificPlatformAndDeviceTypeVersions",
			Handler:    _Version_SpecificPlatformAndDeviceTypeVersions_Handler,
		},
		{
			MethodName: "SpecificPlatformAndDeviceTypeCount",
			Handler:    _Version_SpecificPlatformAndDeviceTypeCount_Handler,
		},
		{
			MethodName: "LastReleaseVersion",
			Handler:    _Version_LastReleaseVersion_Handler,
		},
		{
			MethodName: "ForceNumberBetween",
			Handler:    _Version_ForceNumberBetween_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "version.proto",
}