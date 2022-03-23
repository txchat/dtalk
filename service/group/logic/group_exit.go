package logic

import (
	"context"
	xerror "github.com/txchat/dtalk/pkg/error"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type GroupExitLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewGroupExitLogic(ctx context.Context, svc *service.Service) *GroupExitLogic {
	return &GroupExitLogic{
		ctx: ctx,
		svc: svc,
	}
}

// GroupExit 退出群
func (l *GroupExitLogic) GroupExit(req *pb.GroupExitReq) (*pb.GroupExitResp, error) {
	groupId := req.GroupId
	personId := req.PersonId

	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, err
	}

	person, err := l.svc.GetPersonByMemberIdAndGroupId(l.ctx, personId, groupId)
	if err != nil {
		return nil, err
	}

	if err = person.IsOwner(); err == nil {
		return nil, xerror.NewError(xerror.GroupOwnerExit)
	}

	err = l.svc.ExitGroup(l.ctx, group, person)
	if err != nil {
		return nil, err
	}

	return &pb.GroupExitResp{}, nil
}
