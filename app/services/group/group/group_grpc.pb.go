// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package group

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

// GroupClient is the client API for Group service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupClient interface {
	GroupInfo(ctx context.Context, in *GroupInfoReq, opts ...grpc.CallOption) (*GroupInfoResp, error)
	MemberInfo(ctx context.Context, in *MemberInfoReq, opts ...grpc.CallOption) (*MemberInfoResp, error)
	GroupLimitedMembers(ctx context.Context, in *GroupLimitedMembersReq, opts ...grpc.CallOption) (*GroupLimitedMembersResp, error)
	CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*CreateGroupResp, error)
	JoinedGroups(ctx context.Context, in *JoinedGroupsReq, opts ...grpc.CallOption) (*JoinedGroupsResp, error)
	ChangeOwner(ctx context.Context, in *ChangeOwnerReq, opts ...grpc.CallOption) (*ChangeOwnerResp, error)
	DisbandGroup(ctx context.Context, in *DisbandGroupReq, opts ...grpc.CallOption) (*DisbandGroupResp, error)
	UpdateGroupName(ctx context.Context, in *UpdateGroupNameReq, opts ...grpc.CallOption) (*UpdateGroupNameResp, error)
	UpdateGroupAvatar(ctx context.Context, in *UpdateGroupAvatarReq, opts ...grpc.CallOption) (*UpdateGroupAvatarResp, error)
	UpdateGroupJoinType(ctx context.Context, in *UpdateGroupJoinTypeReq, opts ...grpc.CallOption) (*UpdateGroupJoinTypeResp, error)
	UpdateGroupFriendlyType(ctx context.Context, in *UpdateGroupFriendlyTypeReq, opts ...grpc.CallOption) (*UpdateGroupFriendlyTypeResp, error)
	UpdateGroupMuteType(ctx context.Context, in *UpdateGroupMuteTypeReq, opts ...grpc.CallOption) (*UpdateGroupMuteTypeResp, error)
	InviteMembers(ctx context.Context, in *InviteMembersReq, opts ...grpc.CallOption) (*InviteMembersResp, error)
	KickOutMembers(ctx context.Context, in *KickOutMembersReq, opts ...grpc.CallOption) (*KickOutMembersResp, error)
	MemberExit(ctx context.Context, in *MemberExitReq, opts ...grpc.CallOption) (*MemberExitResp, error)
	ChangeMemberRole(ctx context.Context, in *ChangeMemberRoleReq, opts ...grpc.CallOption) (*ChangeMemberRoleResp, error)
	MuteMembers(ctx context.Context, in *MuteMembersReq, opts ...grpc.CallOption) (*MuteMembersResp, error)
	UnMuteMembers(ctx context.Context, in *UnMuteMembersReq, opts ...grpc.CallOption) (*UnMuteMembersResp, error)
}

type groupClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupClient(cc grpc.ClientConnInterface) GroupClient {
	return &groupClient{cc}
}

func (c *groupClient) GroupInfo(ctx context.Context, in *GroupInfoReq, opts ...grpc.CallOption) (*GroupInfoResp, error) {
	out := new(GroupInfoResp)
	err := c.cc.Invoke(ctx, "/group.Group/GroupInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) MemberInfo(ctx context.Context, in *MemberInfoReq, opts ...grpc.CallOption) (*MemberInfoResp, error) {
	out := new(MemberInfoResp)
	err := c.cc.Invoke(ctx, "/group.Group/MemberInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GroupLimitedMembers(ctx context.Context, in *GroupLimitedMembersReq, opts ...grpc.CallOption) (*GroupLimitedMembersResp, error) {
	out := new(GroupLimitedMembersResp)
	err := c.cc.Invoke(ctx, "/group.Group/GroupLimitedMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*CreateGroupResp, error) {
	out := new(CreateGroupResp)
	err := c.cc.Invoke(ctx, "/group.Group/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) JoinedGroups(ctx context.Context, in *JoinedGroupsReq, opts ...grpc.CallOption) (*JoinedGroupsResp, error) {
	out := new(JoinedGroupsResp)
	err := c.cc.Invoke(ctx, "/group.Group/JoinedGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) ChangeOwner(ctx context.Context, in *ChangeOwnerReq, opts ...grpc.CallOption) (*ChangeOwnerResp, error) {
	out := new(ChangeOwnerResp)
	err := c.cc.Invoke(ctx, "/group.Group/ChangeOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) DisbandGroup(ctx context.Context, in *DisbandGroupReq, opts ...grpc.CallOption) (*DisbandGroupResp, error) {
	out := new(DisbandGroupResp)
	err := c.cc.Invoke(ctx, "/group.Group/DisbandGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) UpdateGroupName(ctx context.Context, in *UpdateGroupNameReq, opts ...grpc.CallOption) (*UpdateGroupNameResp, error) {
	out := new(UpdateGroupNameResp)
	err := c.cc.Invoke(ctx, "/group.Group/UpdateGroupName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) UpdateGroupAvatar(ctx context.Context, in *UpdateGroupAvatarReq, opts ...grpc.CallOption) (*UpdateGroupAvatarResp, error) {
	out := new(UpdateGroupAvatarResp)
	err := c.cc.Invoke(ctx, "/group.Group/UpdateGroupAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) UpdateGroupJoinType(ctx context.Context, in *UpdateGroupJoinTypeReq, opts ...grpc.CallOption) (*UpdateGroupJoinTypeResp, error) {
	out := new(UpdateGroupJoinTypeResp)
	err := c.cc.Invoke(ctx, "/group.Group/UpdateGroupJoinType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) UpdateGroupFriendlyType(ctx context.Context, in *UpdateGroupFriendlyTypeReq, opts ...grpc.CallOption) (*UpdateGroupFriendlyTypeResp, error) {
	out := new(UpdateGroupFriendlyTypeResp)
	err := c.cc.Invoke(ctx, "/group.Group/UpdateGroupFriendlyType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) UpdateGroupMuteType(ctx context.Context, in *UpdateGroupMuteTypeReq, opts ...grpc.CallOption) (*UpdateGroupMuteTypeResp, error) {
	out := new(UpdateGroupMuteTypeResp)
	err := c.cc.Invoke(ctx, "/group.Group/UpdateGroupMuteType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) InviteMembers(ctx context.Context, in *InviteMembersReq, opts ...grpc.CallOption) (*InviteMembersResp, error) {
	out := new(InviteMembersResp)
	err := c.cc.Invoke(ctx, "/group.Group/InviteMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) KickOutMembers(ctx context.Context, in *KickOutMembersReq, opts ...grpc.CallOption) (*KickOutMembersResp, error) {
	out := new(KickOutMembersResp)
	err := c.cc.Invoke(ctx, "/group.Group/KickOutMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) MemberExit(ctx context.Context, in *MemberExitReq, opts ...grpc.CallOption) (*MemberExitResp, error) {
	out := new(MemberExitResp)
	err := c.cc.Invoke(ctx, "/group.Group/MemberExit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) ChangeMemberRole(ctx context.Context, in *ChangeMemberRoleReq, opts ...grpc.CallOption) (*ChangeMemberRoleResp, error) {
	out := new(ChangeMemberRoleResp)
	err := c.cc.Invoke(ctx, "/group.Group/ChangeMemberRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) MuteMembers(ctx context.Context, in *MuteMembersReq, opts ...grpc.CallOption) (*MuteMembersResp, error) {
	out := new(MuteMembersResp)
	err := c.cc.Invoke(ctx, "/group.Group/MuteMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) UnMuteMembers(ctx context.Context, in *UnMuteMembersReq, opts ...grpc.CallOption) (*UnMuteMembersResp, error) {
	out := new(UnMuteMembersResp)
	err := c.cc.Invoke(ctx, "/group.Group/UnMuteMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServer is the server API for Group service.
// All implementations must embed UnimplementedGroupServer
// for forward compatibility
type GroupServer interface {
	GroupInfo(context.Context, *GroupInfoReq) (*GroupInfoResp, error)
	MemberInfo(context.Context, *MemberInfoReq) (*MemberInfoResp, error)
	GroupLimitedMembers(context.Context, *GroupLimitedMembersReq) (*GroupLimitedMembersResp, error)
	CreateGroup(context.Context, *CreateGroupReq) (*CreateGroupResp, error)
	JoinedGroups(context.Context, *JoinedGroupsReq) (*JoinedGroupsResp, error)
	ChangeOwner(context.Context, *ChangeOwnerReq) (*ChangeOwnerResp, error)
	DisbandGroup(context.Context, *DisbandGroupReq) (*DisbandGroupResp, error)
	UpdateGroupName(context.Context, *UpdateGroupNameReq) (*UpdateGroupNameResp, error)
	UpdateGroupAvatar(context.Context, *UpdateGroupAvatarReq) (*UpdateGroupAvatarResp, error)
	UpdateGroupJoinType(context.Context, *UpdateGroupJoinTypeReq) (*UpdateGroupJoinTypeResp, error)
	UpdateGroupFriendlyType(context.Context, *UpdateGroupFriendlyTypeReq) (*UpdateGroupFriendlyTypeResp, error)
	UpdateGroupMuteType(context.Context, *UpdateGroupMuteTypeReq) (*UpdateGroupMuteTypeResp, error)
	InviteMembers(context.Context, *InviteMembersReq) (*InviteMembersResp, error)
	KickOutMembers(context.Context, *KickOutMembersReq) (*KickOutMembersResp, error)
	MemberExit(context.Context, *MemberExitReq) (*MemberExitResp, error)
	ChangeMemberRole(context.Context, *ChangeMemberRoleReq) (*ChangeMemberRoleResp, error)
	MuteMembers(context.Context, *MuteMembersReq) (*MuteMembersResp, error)
	UnMuteMembers(context.Context, *UnMuteMembersReq) (*UnMuteMembersResp, error)
	mustEmbedUnimplementedGroupServer()
}

// UnimplementedGroupServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServer struct {
}

func (UnimplementedGroupServer) GroupInfo(context.Context, *GroupInfoReq) (*GroupInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupInfo not implemented")
}
func (UnimplementedGroupServer) MemberInfo(context.Context, *MemberInfoReq) (*MemberInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberInfo not implemented")
}
func (UnimplementedGroupServer) GroupLimitedMembers(context.Context, *GroupLimitedMembersReq) (*GroupLimitedMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupLimitedMembers not implemented")
}
func (UnimplementedGroupServer) CreateGroup(context.Context, *CreateGroupReq) (*CreateGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedGroupServer) JoinedGroups(context.Context, *JoinedGroupsReq) (*JoinedGroupsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinedGroups not implemented")
}
func (UnimplementedGroupServer) ChangeOwner(context.Context, *ChangeOwnerReq) (*ChangeOwnerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeOwner not implemented")
}
func (UnimplementedGroupServer) DisbandGroup(context.Context, *DisbandGroupReq) (*DisbandGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisbandGroup not implemented")
}
func (UnimplementedGroupServer) UpdateGroupName(context.Context, *UpdateGroupNameReq) (*UpdateGroupNameResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupName not implemented")
}
func (UnimplementedGroupServer) UpdateGroupAvatar(context.Context, *UpdateGroupAvatarReq) (*UpdateGroupAvatarResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupAvatar not implemented")
}
func (UnimplementedGroupServer) UpdateGroupJoinType(context.Context, *UpdateGroupJoinTypeReq) (*UpdateGroupJoinTypeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupJoinType not implemented")
}
func (UnimplementedGroupServer) UpdateGroupFriendlyType(context.Context, *UpdateGroupFriendlyTypeReq) (*UpdateGroupFriendlyTypeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupFriendlyType not implemented")
}
func (UnimplementedGroupServer) UpdateGroupMuteType(context.Context, *UpdateGroupMuteTypeReq) (*UpdateGroupMuteTypeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupMuteType not implemented")
}
func (UnimplementedGroupServer) InviteMembers(context.Context, *InviteMembersReq) (*InviteMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteMembers not implemented")
}
func (UnimplementedGroupServer) KickOutMembers(context.Context, *KickOutMembersReq) (*KickOutMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickOutMembers not implemented")
}
func (UnimplementedGroupServer) MemberExit(context.Context, *MemberExitReq) (*MemberExitResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberExit not implemented")
}
func (UnimplementedGroupServer) ChangeMemberRole(context.Context, *ChangeMemberRoleReq) (*ChangeMemberRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeMemberRole not implemented")
}
func (UnimplementedGroupServer) MuteMembers(context.Context, *MuteMembersReq) (*MuteMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MuteMembers not implemented")
}
func (UnimplementedGroupServer) UnMuteMembers(context.Context, *UnMuteMembersReq) (*UnMuteMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnMuteMembers not implemented")
}
func (UnimplementedGroupServer) mustEmbedUnimplementedGroupServer() {}

// UnsafeGroupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServer will
// result in compilation errors.
type UnsafeGroupServer interface {
	mustEmbedUnimplementedGroupServer()
}

func RegisterGroupServer(s grpc.ServiceRegistrar, srv GroupServer) {
	s.RegisterService(&Group_ServiceDesc, srv)
}

func _Group_GroupInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GroupInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/GroupInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GroupInfo(ctx, req.(*GroupInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_MemberInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).MemberInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/MemberInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).MemberInfo(ctx, req.(*MemberInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GroupLimitedMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupLimitedMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GroupLimitedMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/GroupLimitedMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GroupLimitedMembers(ctx, req.(*GroupLimitedMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).CreateGroup(ctx, req.(*CreateGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_JoinedGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinedGroupsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).JoinedGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/JoinedGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).JoinedGroups(ctx, req.(*JoinedGroupsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_ChangeOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeOwnerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).ChangeOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/ChangeOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).ChangeOwner(ctx, req.(*ChangeOwnerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_DisbandGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisbandGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).DisbandGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/DisbandGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).DisbandGroup(ctx, req.(*DisbandGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_UpdateGroupName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).UpdateGroupName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/UpdateGroupName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).UpdateGroupName(ctx, req.(*UpdateGroupNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_UpdateGroupAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupAvatarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).UpdateGroupAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/UpdateGroupAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).UpdateGroupAvatar(ctx, req.(*UpdateGroupAvatarReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_UpdateGroupJoinType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupJoinTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).UpdateGroupJoinType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/UpdateGroupJoinType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).UpdateGroupJoinType(ctx, req.(*UpdateGroupJoinTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_UpdateGroupFriendlyType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupFriendlyTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).UpdateGroupFriendlyType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/UpdateGroupFriendlyType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).UpdateGroupFriendlyType(ctx, req.(*UpdateGroupFriendlyTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_UpdateGroupMuteType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupMuteTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).UpdateGroupMuteType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/UpdateGroupMuteType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).UpdateGroupMuteType(ctx, req.(*UpdateGroupMuteTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_InviteMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).InviteMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/InviteMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).InviteMembers(ctx, req.(*InviteMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_KickOutMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickOutMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).KickOutMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/KickOutMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).KickOutMembers(ctx, req.(*KickOutMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_MemberExit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberExitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).MemberExit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/MemberExit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).MemberExit(ctx, req.(*MemberExitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_ChangeMemberRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeMemberRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).ChangeMemberRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/ChangeMemberRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).ChangeMemberRole(ctx, req.(*ChangeMemberRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_MuteMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MuteMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).MuteMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/MuteMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).MuteMembers(ctx, req.(*MuteMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_UnMuteMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnMuteMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).UnMuteMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.Group/UnMuteMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).UnMuteMembers(ctx, req.(*UnMuteMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Group_ServiceDesc is the grpc.ServiceDesc for Group service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Group_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "group.Group",
	HandlerType: (*GroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GroupInfo",
			Handler:    _Group_GroupInfo_Handler,
		},
		{
			MethodName: "MemberInfo",
			Handler:    _Group_MemberInfo_Handler,
		},
		{
			MethodName: "GroupLimitedMembers",
			Handler:    _Group_GroupLimitedMembers_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _Group_CreateGroup_Handler,
		},
		{
			MethodName: "JoinedGroups",
			Handler:    _Group_JoinedGroups_Handler,
		},
		{
			MethodName: "ChangeOwner",
			Handler:    _Group_ChangeOwner_Handler,
		},
		{
			MethodName: "DisbandGroup",
			Handler:    _Group_DisbandGroup_Handler,
		},
		{
			MethodName: "UpdateGroupName",
			Handler:    _Group_UpdateGroupName_Handler,
		},
		{
			MethodName: "UpdateGroupAvatar",
			Handler:    _Group_UpdateGroupAvatar_Handler,
		},
		{
			MethodName: "UpdateGroupJoinType",
			Handler:    _Group_UpdateGroupJoinType_Handler,
		},
		{
			MethodName: "UpdateGroupFriendlyType",
			Handler:    _Group_UpdateGroupFriendlyType_Handler,
		},
		{
			MethodName: "UpdateGroupMuteType",
			Handler:    _Group_UpdateGroupMuteType_Handler,
		},
		{
			MethodName: "InviteMembers",
			Handler:    _Group_InviteMembers_Handler,
		},
		{
			MethodName: "KickOutMembers",
			Handler:    _Group_KickOutMembers_Handler,
		},
		{
			MethodName: "MemberExit",
			Handler:    _Group_MemberExit_Handler,
		},
		{
			MethodName: "ChangeMemberRole",
			Handler:    _Group_ChangeMemberRole_Handler,
		},
		{
			MethodName: "MuteMembers",
			Handler:    _Group_MuteMembers_Handler,
		},
		{
			MethodName: "UnMuteMembers",
			Handler:    _Group_UnMuteMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group.proto",
}
