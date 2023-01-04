// Code generated by goctl. DO NOT EDIT!
// Source: group.proto

package groupclient

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/group"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ChangeMemberRoleReq          = group.ChangeMemberRoleReq
	ChangeMemberRoleResp         = group.ChangeMemberRoleResp
	ChangeOwnerReq               = group.ChangeOwnerReq
	ChangeOwnerResp              = group.ChangeOwnerResp
	CreateGroupReq               = group.CreateGroupReq
	CreateGroupReq_MemberMinData = group.CreateGroupReq_MemberMinData
	CreateGroupResp              = group.CreateGroupResp
	DisbandGroupReq              = group.DisbandGroupReq
	DisbandGroupResp             = group.DisbandGroupResp
	GetMuteListReq               = group.GetMuteListReq
	GetMuteListResp              = group.GetMuteListResp
	GroupInfo                    = group.GroupInfo
	GroupInfoReq                 = group.GroupInfoReq
	GroupInfoResp                = group.GroupInfoResp
	GroupLimitedMembersReq       = group.GroupLimitedMembersReq
	GroupLimitedMembersResp      = group.GroupLimitedMembersResp
	InviteMembersReq             = group.InviteMembersReq
	InviteMembersResp            = group.InviteMembersResp
	JoinedGroupsReq              = group.JoinedGroupsReq
	JoinedGroupsResp             = group.JoinedGroupsResp
	KickOutMembersReq            = group.KickOutMembersReq
	KickOutMembersResp           = group.KickOutMembersResp
	MemberExitReq                = group.MemberExitReq
	MemberExitResp               = group.MemberExitResp
	MemberInfo                   = group.MemberInfo
	MemberInfoReq                = group.MemberInfoReq
	MemberInfoResp               = group.MemberInfoResp
	MembersInfoReq               = group.MembersInfoReq
	MembersInfoResp              = group.MembersInfoResp
	MuteMembersReq               = group.MuteMembersReq
	MuteMembersResp              = group.MuteMembersResp
	UnMuteMembersReq             = group.UnMuteMembersReq
	UnMuteMembersResp            = group.UnMuteMembersResp
	UpdateGroupAvatarReq         = group.UpdateGroupAvatarReq
	UpdateGroupAvatarResp        = group.UpdateGroupAvatarResp
	UpdateGroupFriendlyTypeReq   = group.UpdateGroupFriendlyTypeReq
	UpdateGroupFriendlyTypeResp  = group.UpdateGroupFriendlyTypeResp
	UpdateGroupJoinTypeReq       = group.UpdateGroupJoinTypeReq
	UpdateGroupJoinTypeResp      = group.UpdateGroupJoinTypeResp
	UpdateGroupMemberNameReq     = group.UpdateGroupMemberNameReq
	UpdateGroupMemberNameResp    = group.UpdateGroupMemberNameResp
	UpdateGroupMuteTypeReq       = group.UpdateGroupMuteTypeReq
	UpdateGroupMuteTypeResp      = group.UpdateGroupMuteTypeResp
	UpdateGroupNameReq           = group.UpdateGroupNameReq
	UpdateGroupNameResp          = group.UpdateGroupNameResp

	Group interface {
		GroupInfo(ctx context.Context, in *GroupInfoReq, opts ...grpc.CallOption) (*GroupInfoResp, error)
		MemberInfo(ctx context.Context, in *MemberInfoReq, opts ...grpc.CallOption) (*MemberInfoResp, error)
		MembersInfo(ctx context.Context, in *MembersInfoReq, opts ...grpc.CallOption) (*MembersInfoResp, error)
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
		GetMuteList(ctx context.Context, in *GetMuteListReq, opts ...grpc.CallOption) (*GetMuteListResp, error)
		InviteMembers(ctx context.Context, in *InviteMembersReq, opts ...grpc.CallOption) (*InviteMembersResp, error)
		KickOutMembers(ctx context.Context, in *KickOutMembersReq, opts ...grpc.CallOption) (*KickOutMembersResp, error)
		MemberExit(ctx context.Context, in *MemberExitReq, opts ...grpc.CallOption) (*MemberExitResp, error)
		ChangeMemberRole(ctx context.Context, in *ChangeMemberRoleReq, opts ...grpc.CallOption) (*ChangeMemberRoleResp, error)
		MuteMembers(ctx context.Context, in *MuteMembersReq, opts ...grpc.CallOption) (*MuteMembersResp, error)
		UnMuteMembers(ctx context.Context, in *UnMuteMembersReq, opts ...grpc.CallOption) (*UnMuteMembersResp, error)
		UpdateGroupMemberName(ctx context.Context, in *UpdateGroupMemberNameReq, opts ...grpc.CallOption) (*UpdateGroupMemberNameResp, error)
	}

	defaultGroup struct {
		cli zrpc.Client
	}
)

func NewGroup(cli zrpc.Client) Group {
	return &defaultGroup{
		cli: cli,
	}
}

func (m *defaultGroup) GroupInfo(ctx context.Context, in *GroupInfoReq, opts ...grpc.CallOption) (*GroupInfoResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.GroupInfo(ctx, in, opts...)
}

func (m *defaultGroup) MemberInfo(ctx context.Context, in *MemberInfoReq, opts ...grpc.CallOption) (*MemberInfoResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.MemberInfo(ctx, in, opts...)
}

func (m *defaultGroup) MembersInfo(ctx context.Context, in *MembersInfoReq, opts ...grpc.CallOption) (*MembersInfoResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.MembersInfo(ctx, in, opts...)
}

func (m *defaultGroup) GroupLimitedMembers(ctx context.Context, in *GroupLimitedMembersReq, opts ...grpc.CallOption) (*GroupLimitedMembersResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.GroupLimitedMembers(ctx, in, opts...)
}

func (m *defaultGroup) CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*CreateGroupResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.CreateGroup(ctx, in, opts...)
}

func (m *defaultGroup) JoinedGroups(ctx context.Context, in *JoinedGroupsReq, opts ...grpc.CallOption) (*JoinedGroupsResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.JoinedGroups(ctx, in, opts...)
}

func (m *defaultGroup) ChangeOwner(ctx context.Context, in *ChangeOwnerReq, opts ...grpc.CallOption) (*ChangeOwnerResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.ChangeOwner(ctx, in, opts...)
}

func (m *defaultGroup) DisbandGroup(ctx context.Context, in *DisbandGroupReq, opts ...grpc.CallOption) (*DisbandGroupResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.DisbandGroup(ctx, in, opts...)
}

func (m *defaultGroup) UpdateGroupName(ctx context.Context, in *UpdateGroupNameReq, opts ...grpc.CallOption) (*UpdateGroupNameResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.UpdateGroupName(ctx, in, opts...)
}

func (m *defaultGroup) UpdateGroupAvatar(ctx context.Context, in *UpdateGroupAvatarReq, opts ...grpc.CallOption) (*UpdateGroupAvatarResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.UpdateGroupAvatar(ctx, in, opts...)
}

func (m *defaultGroup) UpdateGroupJoinType(ctx context.Context, in *UpdateGroupJoinTypeReq, opts ...grpc.CallOption) (*UpdateGroupJoinTypeResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.UpdateGroupJoinType(ctx, in, opts...)
}

func (m *defaultGroup) UpdateGroupFriendlyType(ctx context.Context, in *UpdateGroupFriendlyTypeReq, opts ...grpc.CallOption) (*UpdateGroupFriendlyTypeResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.UpdateGroupFriendlyType(ctx, in, opts...)
}

func (m *defaultGroup) UpdateGroupMuteType(ctx context.Context, in *UpdateGroupMuteTypeReq, opts ...grpc.CallOption) (*UpdateGroupMuteTypeResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.UpdateGroupMuteType(ctx, in, opts...)
}

func (m *defaultGroup) GetMuteList(ctx context.Context, in *GetMuteListReq, opts ...grpc.CallOption) (*GetMuteListResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.GetMuteList(ctx, in, opts...)
}

func (m *defaultGroup) InviteMembers(ctx context.Context, in *InviteMembersReq, opts ...grpc.CallOption) (*InviteMembersResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.InviteMembers(ctx, in, opts...)
}

func (m *defaultGroup) KickOutMembers(ctx context.Context, in *KickOutMembersReq, opts ...grpc.CallOption) (*KickOutMembersResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.KickOutMembers(ctx, in, opts...)
}

func (m *defaultGroup) MemberExit(ctx context.Context, in *MemberExitReq, opts ...grpc.CallOption) (*MemberExitResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.MemberExit(ctx, in, opts...)
}

func (m *defaultGroup) ChangeMemberRole(ctx context.Context, in *ChangeMemberRoleReq, opts ...grpc.CallOption) (*ChangeMemberRoleResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.ChangeMemberRole(ctx, in, opts...)
}

func (m *defaultGroup) MuteMembers(ctx context.Context, in *MuteMembersReq, opts ...grpc.CallOption) (*MuteMembersResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.MuteMembers(ctx, in, opts...)
}

func (m *defaultGroup) UnMuteMembers(ctx context.Context, in *UnMuteMembersReq, opts ...grpc.CallOption) (*UnMuteMembersResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.UnMuteMembers(ctx, in, opts...)
}

func (m *defaultGroup) UpdateGroupMemberName(ctx context.Context, in *UpdateGroupMemberNameReq, opts ...grpc.CallOption) (*UpdateGroupMemberNameResp, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.UpdateGroupMemberName(ctx, in, opts...)
}
