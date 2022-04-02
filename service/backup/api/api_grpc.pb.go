// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package backup

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

// BackupClient is the client API for Backup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackupClient interface {
	Retrieve(ctx context.Context, in *RetrieveReq, opts ...grpc.CallOption) (*RetrieveReply, error)
}

type backupClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupClient(cc grpc.ClientConnInterface) BackupClient {
	return &backupClient{cc}
}

func (c *backupClient) Retrieve(ctx context.Context, in *RetrieveReq, opts ...grpc.CallOption) (*RetrieveReply, error) {
	out := new(RetrieveReply)
	err := c.cc.Invoke(ctx, "/dtalk.backup.Backup/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupServer is the server API for Backup service.
// All implementations must embed UnimplementedBackupServer
// for forward compatibility
type BackupServer interface {
	Retrieve(context.Context, *RetrieveReq) (*RetrieveReply, error)
	mustEmbedUnimplementedBackupServer()
}

// UnimplementedBackupServer must be embedded to have forward compatible implementations.
type UnimplementedBackupServer struct {
}

func (UnimplementedBackupServer) Retrieve(context.Context, *RetrieveReq) (*RetrieveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedBackupServer) mustEmbedUnimplementedBackupServer() {}

// UnsafeBackupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackupServer will
// result in compilation errors.
type UnsafeBackupServer interface {
	mustEmbedUnimplementedBackupServer()
}

func RegisterBackupServer(s grpc.ServiceRegistrar, srv BackupServer) {
	s.RegisterService(&Backup_ServiceDesc, srv)
}

func _Backup_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtalk.backup.Backup/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).Retrieve(ctx, req.(*RetrieveReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Backup_ServiceDesc is the grpc.ServiceDesc for Backup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Backup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dtalk.backup.Backup",
	HandlerType: (*BackupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Retrieve",
			Handler:    _Backup_Retrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}