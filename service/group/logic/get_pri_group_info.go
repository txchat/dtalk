package logic

import (
	"context"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type GetPriGroupInfoLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewGetPriGroupInfoLogic(ctx context.Context, svc *service.Service) *GetPriGroupInfoLogic {
	return &GetPriGroupInfoLogic{
		ctx: ctx,
		svc: svc,
	}
}

// GetPriGroupInfo 查询群全部信息
func (l *GetPriGroupInfoLogic) GetPriGroupInfo(req *pb.GetPriGroupInfoReq) (*pb.GetPriGroupInfoResp, error) {
	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, req.GroupId)
	if err != nil {
		return nil, err
	}

	owner, err := l.svc.GetMemberByMemberIdAndGroupId(l.ctx, group.GroupOwnerId, group.GroupId)
	if err != nil {
		return nil, err
	}

	person, err := l.svc.GetPersonByMemberIdAndGroupId(l.ctx, req.PersonId, req.GroupId)
	if err != nil {
		return nil, err
	}

	members, err := l.svc.GetGroupMembersByGroupIdWithLimit(group.GroupId, 0, int64(req.DisplayNum))
	if err != nil {
		return nil, err
	}

	groupInfo := NewRPCGroupInfo(group)
	groupInfo.Owner = NewRPCGroupMemberInfo(owner)
	groupInfo.Person = NewRPCGroupMemberInfo(person)
	groupInfo.Members = NewRPCGroupMemberInfos(members)

	return &pb.GetPriGroupInfoResp{
		Group: groupInfo,
	}, nil
}
