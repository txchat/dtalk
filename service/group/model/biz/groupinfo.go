package biz

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model/types"
)

type groupContextKey string

const (
	// GroupContextKey .
	GroupContextKey   = groupContextKey("FZM-Group-Token")
	GroupIdContextKey = groupContextKey("FZM-Group-Id")
)

const (
	GroupStatusNormal  = 0 // 正常
	GroupStatusBlock   = 1 // 封禁
	GroupStatusDisBand = 2 // 解散

	GroupJoinTypeAny   = 0 // 无需审批（默认）
	GroupJoinTypeAdmin = 1 // 禁止加群，群主和管理员邀请加群
	GroupJoinTypeApply = 2 // 普通人邀请需要审批,群主和管理员直接加群

	GroupMuteTypeAny   = 0 // 全员可发言
	GroupMuteTypeAdmin = 1 // 全员禁言(除群主和管理员)

	GroupFriendTypeAllow = 0 // 群内可加好友
	GroupFriendTypeDeny  = 1 // 群内禁止加好友

	// 群类型
	GroupTypeNormal = 0 // 普通群
	GroupTypeEnt    = 1 // 全员群
	GroupTypeDep    = 2 // 部门群
)

type GroupInfo struct {
	GroupId     int64
	GroupMarkId string
	GroupName   string
	GroupAvatar string
	// 群人数
	GroupMemberNum int32
	// 群人数上限
	GroupMaximum   int32
	GroupIntroduce string
	// 群状态，0=正常 1=封禁 2=解散
	GroupStatus     int32
	GroupOwnerId    string
	GroupCreateTime int64
	GroupUpdateTime int64
	// 加群方式，0=无需审批（默认），1=禁止加群，群主和管理员邀请加群, 2=普通人邀请需要审批,群主和管理员直接加群
	GroupJoinType int32
	// 禁言， 0=全员可发言， 1=全员禁言(除群主和管理员)
	GroupMuteType int32
	// 加好友限制， 0=群内可加好友，1=群内禁止加好友
	GroupFriendType int32

	// 群内当前被禁言的人数
	MuteNum int32 `json:"muteNum"`
	// 群内管理员数量
	AdminNum int32 `json:"adminNum"`

	//
	AESKey string
	//
	GroupPubName string
	// 群类型 (0: 普通群, 1: 全员群, 2: 部门群)
	GroupType int32
}

func (g *GroupInfo) IsNormal() error {
	switch g.GroupStatus {
	case GroupStatusNormal:
		return nil
	case GroupStatusBlock:
		return xerror.NewError(xerror.GroupStatusBlock)
	case GroupStatusDisBand:
		return xerror.NewError(xerror.GroupStatusDisBand)
	default:
		return xerror.NewError(xerror.CodeInnerError)
	}
}

func (g *GroupInfo) ToTypes(owner *types.GroupMember, person *types.GroupMember) *types.GroupInfo {
	return &types.GroupInfo{
		Id:         g.GroupId,
		IdStr:      util.ToString(g.GroupId),
		MarkId:     g.GroupMarkId,
		Name:       g.GroupName,
		Avatar:     g.GroupAvatar,
		Introduce:  g.GroupIntroduce,
		Owner:      owner,
		Person:     person,
		MemberNum:  g.GroupMemberNum,
		Maximum:    g.GroupMaximum,
		Status:     g.GroupStatus,
		CreateTime: g.GroupCreateTime,
		JoinType:   g.GroupJoinType,
		MuteType:   g.GroupMuteType,
		FriendType: g.GroupFriendType,
		MuteNum:    g.MuteNum,
		AdminNum:   g.AdminNum,
		AESKey:     g.AESKey,
		PublicName: g.GroupPubName,
		GroupType:  g.GroupType,
	}
}

func (g *GroupInfo) TryJoin(add int32) error {
	if add+g.GroupMemberNum > g.GroupMaximum {
		return xerror.NewError(xerror.GroupMemberLimit)
	}
	return nil
}

func (g *GroupInfo) TrySetAdmin() error {
	if g.AdminNum+1 > AdminNum {
		return xerror.NewError(xerror.GroupAdminNumLimit)
	}

	return nil
}

func (g *GroupInfo) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, GroupContextKey, g)
}

func NewGroupInfoFromContext(ctx context.Context) (*GroupInfo, error) {
	if group, ok := ctx.Value(GroupContextKey).(*GroupInfo); ok {
		return group, nil
	}

	return nil, xerror.NewError(xerror.GroupNotExist)
}

func NewGroupIdFromContext(ctx context.Context) (int64, error) {
	if groupId, ok := ctx.Value(GroupIdContextKey).(int64); ok {
		return groupId, nil
	}

	return 0, xerror.NewError(xerror.GroupNotExist)
}
