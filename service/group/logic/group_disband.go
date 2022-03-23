package logic

import (
	"context"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type GroupDisbandLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewGroupDisbandLogic(ctx context.Context, svc *service.Service) *GroupDisbandLogic {
	return &GroupDisbandLogic{
		ctx: ctx,
		svc: svc,
	}
}

// GroupDisband 解散群
func (l *GroupDisbandLogic) GroupDisband(req *pb.GroupDisbandReq) (*pb.GroupDisbandResp, error) {
	groupId := req.GroupId
	personId := req.PersonId

	person, err := l.svc.GetPersonByMemberIdAndGroupId(l.ctx, personId, groupId)
	if err != nil {
		return nil, err
	}

	// 只有群主可以解散群
	if err = person.IsOwner(); err != nil {
		return nil, err
	}

	err = l.svc.GroupDisband(l.ctx, groupId, personId)
	if err != nil {
		return nil, err
	}

	return &pb.GroupDisbandResp{}, nil
}
