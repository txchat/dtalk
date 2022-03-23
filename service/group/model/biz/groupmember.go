package biz

import (
	"context"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/group/model/types"
)

type groupMemberContextKey string

const (
	// GroupMemberContextKey .
	GroupMemberContextKey = groupMemberContextKey("FZM-Group-Member-Token")
)

const (
	GroupMemberTypeOwner  = 2  // 群主
	GroupMemberTypeAdmin  = 1  // 管理员
	GroupMemberTypeNormal = 0  // 群员
	GroupMemberTypeOther  = 10 // 退群
)

type GroupMember struct {
	GroupId         int64
	GroupMemberId   string
	GroupMemberName string
	// 用户角色，0=群员,1=管理员, 2=群主，10=退群
	GroupMemberType int32
	// 该用户被禁言结束的时间 9223372036854775807=永久禁言
	GroupMemberMuteTime int64
	GroupMemberJoinTime int64
}

func (m *GroupMember) ToTypes() *types.GroupMember {
	return &types.GroupMember{
		MemberId:       m.GroupMemberId,
		MemberName:     m.GroupMemberName,
		MemberType:     m.GroupMemberType,
		MemberMuteTime: m.GroupMemberMuteTime,
	}
}

func (m *GroupMember) IsAdmin() error {
	if m.GroupMemberType < GroupMemberTypeAdmin {
		return xerror.NewError(xerror.GroupAdminDeny)
	}
	return nil
}

func (m *GroupMember) IsOwner() error {
	if m.GroupMemberType < GroupMemberTypeOwner {
		return xerror.NewError(xerror.GroupOwnerDeny)
	}
	return nil
}

func (m *GroupMember) RemoveOneMember(member *GroupMember) error {
	if m.GroupId != member.GroupId || m.GroupMemberType == GroupMemberTypeOther || member.GroupMemberType == GroupMemberTypeOwner {
		return xerror.NewError(xerror.CodeInnerError)
	}

	if m.GroupMemberType <= member.GroupMemberType {
		return xerror.NewError(xerror.GroupHigherPermission)
	}

	return nil
}

type InviteFlow int

const (
	InviteOk    InviteFlow = 1
	InviteApply InviteFlow = 2
	InviteFail  InviteFlow = 3
)

func (m *GroupMember) TryInvite(group *GroupInfo) InviteFlow {
	if m.GroupMemberType == GroupMemberTypeNormal {
		switch group.GroupJoinType {
		case GroupJoinTypeApply:
			return InviteApply
		case GroupJoinTypeAny:
			return InviteOk
		case GroupJoinTypeAdmin:
			return InviteFail
		default:
			return InviteFail

		}
	}

	return InviteOk
}

func (m *GroupMember) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, GroupMemberContextKey, m)
}

func NewGroupMemberFromContext(ctx context.Context) (*GroupMember, error) {
	if groupMember, ok := ctx.Value(GroupMemberContextKey).(*GroupMember); ok {
		return groupMember, nil
	}

	return nil, xerror.NewError(xerror.GroupPersonNotExist)
}
