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
	QueryBind(ctx context.Context, in *QueryBindReq, opts ...grpc.CallOption) (*QueryBindResp, error)
	QueryRelated(ctx context.Context, in *QueryRelatedReq, opts ...grpc.CallOption) (*QueryRelatedResp, error)
	UpdateAddressBackup(ctx context.Context, in *UpdateAddressBackupReq, opts ...grpc.CallOption) (*UpdateAddressBackupResp, error)
	UpdateAddressRelated(ctx context.Context, in *UpdateAddressRelatedReq, opts ...grpc.CallOption) (*UpdateAddressRelatedResp, error)
	UpdateMnemonic(ctx context.Context, in *UpdateMnemonicReq, opts ...grpc.CallOption) (*UpdateMnemonicResp, error)
}

type backupClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupClient(cc grpc.ClientConnInterface) BackupClient {
	return &backupClient{cc}
}

func (c *backupClient) QueryBind(ctx context.Context, in *QueryBindReq, opts ...grpc.CallOption) (*QueryBindResp, error) {
	out := new(QueryBindResp)
	err := c.cc.Invoke(ctx, "/backup.Backup/QueryBind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupClient) QueryRelated(ctx context.Context, in *QueryRelatedReq, opts ...grpc.CallOption) (*QueryRelatedResp, error) {
	out := new(QueryRelatedResp)
	err := c.cc.Invoke(ctx, "/backup.Backup/QueryRelated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupClient) UpdateAddressBackup(ctx context.Context, in *UpdateAddressBackupReq, opts ...grpc.CallOption) (*UpdateAddressBackupResp, error) {
	out := new(UpdateAddressBackupResp)
	err := c.cc.Invoke(ctx, "/backup.Backup/UpdateAddressBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupClient) UpdateAddressRelated(ctx context.Context, in *UpdateAddressRelatedReq, opts ...grpc.CallOption) (*UpdateAddressRelatedResp, error) {
	out := new(UpdateAddressRelatedResp)
	err := c.cc.Invoke(ctx, "/backup.Backup/UpdateAddressRelated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupClient) UpdateMnemonic(ctx context.Context, in *UpdateMnemonicReq, opts ...grpc.CallOption) (*UpdateMnemonicResp, error) {
	out := new(UpdateMnemonicResp)
	err := c.cc.Invoke(ctx, "/backup.Backup/UpdateMnemonic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupServer is the server API for Backup service.
// All implementations must embed UnimplementedBackupServer
// for forward compatibility
type BackupServer interface {
	QueryBind(context.Context, *QueryBindReq) (*QueryBindResp, error)
	QueryRelated(context.Context, *QueryRelatedReq) (*QueryRelatedResp, error)
	UpdateAddressBackup(context.Context, *UpdateAddressBackupReq) (*UpdateAddressBackupResp, error)
	UpdateAddressRelated(context.Context, *UpdateAddressRelatedReq) (*UpdateAddressRelatedResp, error)
	UpdateMnemonic(context.Context, *UpdateMnemonicReq) (*UpdateMnemonicResp, error)
	mustEmbedUnimplementedBackupServer()
}

// UnimplementedBackupServer must be embedded to have forward compatible implementations.
type UnimplementedBackupServer struct {
}

func (UnimplementedBackupServer) QueryBind(context.Context, *QueryBindReq) (*QueryBindResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBind not implemented")
}
func (UnimplementedBackupServer) QueryRelated(context.Context, *QueryRelatedReq) (*QueryRelatedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRelated not implemented")
}
func (UnimplementedBackupServer) UpdateAddressBackup(context.Context, *UpdateAddressBackupReq) (*UpdateAddressBackupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddressBackup not implemented")
}
func (UnimplementedBackupServer) UpdateAddressRelated(context.Context, *UpdateAddressRelatedReq) (*UpdateAddressRelatedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddressRelated not implemented")
}
func (UnimplementedBackupServer) UpdateMnemonic(context.Context, *UpdateMnemonicReq) (*UpdateMnemonicResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMnemonic not implemented")
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

func _Backup_QueryBind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).QueryBind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.Backup/QueryBind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).QueryBind(ctx, req.(*QueryBindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backup_QueryRelated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRelatedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).QueryRelated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.Backup/QueryRelated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).QueryRelated(ctx, req.(*QueryRelatedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backup_UpdateAddressBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAddressBackupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).UpdateAddressBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.Backup/UpdateAddressBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).UpdateAddressBackup(ctx, req.(*UpdateAddressBackupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backup_UpdateAddressRelated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAddressRelatedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).UpdateAddressRelated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.Backup/UpdateAddressRelated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).UpdateAddressRelated(ctx, req.(*UpdateAddressRelatedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backup_UpdateMnemonic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMnemonicReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).UpdateMnemonic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.Backup/UpdateMnemonic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).UpdateMnemonic(ctx, req.(*UpdateMnemonicReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Backup_ServiceDesc is the grpc.ServiceDesc for Backup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Backup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backup.Backup",
	HandlerType: (*BackupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryBind",
			Handler:    _Backup_QueryBind_Handler,
		},
		{
			MethodName: "QueryRelated",
			Handler:    _Backup_QueryRelated_Handler,
		},
		{
			MethodName: "UpdateAddressBackup",
			Handler:    _Backup_UpdateAddressBackup_Handler,
		},
		{
			MethodName: "UpdateAddressRelated",
			Handler:    _Backup_UpdateAddressRelated_Handler,
		},
		{
			MethodName: "UpdateMnemonic",
			Handler:    _Backup_UpdateMnemonic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backup.proto",
}