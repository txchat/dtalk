package logic

import (
	"context"

	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type GetGroupMemberListLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewGetGroupMemberListLogic(ctx context.Context, svc *service.Service) *GetGroupMemberListLogic {
	return &GetGroupMemberListLogic{
		ctx: ctx,
		svc: svc,
	}
}

// GetGroupMemberList 查询群成员列表
func (l *GetGroupMemberListLogic) GetGroupMemberList(req *pb.GetGroupMemberListReq) (*pb.GetGroupMemberListResp, error) {
	groupId := req.GroupId
	personId := req.PersonId

	_, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, err
	}

	_, err = l.svc.GetPersonByMemberIdAndGroupId(l.ctx, personId, groupId)
	if err != nil {
		return nil, err
	}

	groupMembers, err := l.svc.GetMembersByGroupId(groupId)
	if err != nil {
		return nil, err
	}

	return &pb.GetGroupMemberListResp{
		Members: NewRPCGroupMemberInfos(groupMembers),
	}, nil
}
