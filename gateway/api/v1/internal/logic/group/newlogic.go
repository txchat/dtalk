package logic

import (
	"context"

	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	"github.com/txchat/dtalk/pkg/api"
	"github.com/txchat/dtalk/pkg/util"
	pb "github.com/txchat/dtalk/service/group/api"
)

type GroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) GroupLogic {
	return GroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupLogic) getOpe() string {
	return api.NewAddrWithContext(l.ctx)
}

func NewTypesGroupInfo(do *pb.GroupBizInfo) *types.GroupInfo {
	if do == nil {
		return &types.GroupInfo{}
	}

	res := &types.GroupInfo{
		Id:         do.Id,
		IdStr:      util.MustToString(do.Id),
		MarkId:     do.MarkId,
		Name:       do.Name,
		PublicName: do.PubName,
		Avatar:     do.Avatar,
		Introduce:  do.Introduce,
		Owner:      nil,
		Person:     nil,
		MemberNum:  do.MemberNum,
		Maximum:    do.MemberMaximum,
		Status:     int32(do.Status),
		CreateTime: do.CreateTime,
		JoinType:   int32(do.JoinType),
		MuteType:   int32(do.MuteType),
		FriendType: int32(do.FriendType),
		MuteNum:    do.MuteNum,
		AdminNum:   do.AdminNum,
		AESKey:     do.AESKey,
		GroupType:  int32(do.GetType()),
	}

	if do.Owner != nil {
		res.Owner = NewTypesGroupMemberInfo(do.Owner)
	}

	if do.Person != nil {
		res.Person = NewTypesGroupMemberInfo(do.Person)
	}

	return res
}

func NewTypesGroupInfos(dos []*pb.GroupBizInfo) []*types.GroupInfo {
	dtos := make([]*types.GroupInfo, 0, len(dos))
	for _, do := range dos {
		dtos = append(dtos, NewTypesGroupInfo(do))
	}

	return dtos
}

func NewTypesGroupMemberInfo(do *pb.GroupMemberBizInfo) *types.GroupMember {
	if do == nil {
		return &types.GroupMember{}
	}

	return &types.GroupMember{
		MemberId:       do.Id,
		MemberName:     do.Name,
		MemberType:     int32(do.Type),
		MemberMuteTime: do.MuteTime,
	}
}

func NewTypesGroupMemberInfos(dos []*pb.GroupMemberBizInfo) []*types.GroupMember {
	dtos := make([]*types.GroupMember, 0, len(dos))
	for _, do := range dos {
		dtos = append(dtos, NewTypesGroupMemberInfo(do))
	}

	return dtos
}
