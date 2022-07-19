package logic

import (
	"context"

	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type GetGroupListLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewGetGroupListLogic(ctx context.Context, svc *service.Service) *GetGroupListLogic {
	return &GetGroupListLogic{
		ctx: ctx,
		svc: svc,
	}
}

// GetGroupList 查询加入的群列表
func (l *GetGroupListLogic) GetGroupList(req *pb.GetGroupListReq) (*pb.GetGroupListResp, error) {
	groupIds, err := l.svc.GetGroupIdsByMemberId(req.PersonId)
	if err != nil {
		return nil, err
	}

	groupInfos := make([]*pb.GroupBizInfo, 0, len(groupIds))
	for _, groupId := range groupIds {
		group, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
		if err != nil {
			return nil, err
		}

		owner, err := l.svc.GetMemberByMemberIdAndGroupId(l.ctx, group.GroupOwnerId, group.GroupId)
		if err != nil {
			return nil, err
		}
		person, err := l.svc.GetPersonByMemberIdAndGroupId(l.ctx, req.PersonId, group.GroupId)
		if err != nil {
			return nil, err
		}
		groupInfo := NewRPCGroupInfo(group)
		groupInfo.Owner = NewRPCGroupMemberInfo(owner)
		groupInfo.Person = NewRPCGroupMemberInfo(person)
		groupInfos = append(groupInfos, groupInfo)
	}

	return &pb.GetGroupListResp{
		Groups: groupInfos,
	}, nil
}
