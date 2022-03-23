package logic

import (
	"context"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type UpdateGroupNameLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewUpdateGroupNameLogic(ctx context.Context, svc *service.Service) *UpdateGroupNameLogic {
	return &UpdateGroupNameLogic{
		ctx: ctx,
		svc: svc,
	}
}

// UpdateGroupName 更新群名称
func (l *UpdateGroupNameLogic) UpdateGroupName(req *pb.UpdateGroupNameReq) (*pb.UpdateGroupNameResp, error) {
	groupId := req.GroupId
	personId := req.PersonId
	groupName := req.Name
	groupPubName := req.PublicName

	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, err
	}

	person, err := l.svc.GetPersonByMemberIdAndGroupId(l.ctx, personId, groupId)
	if err != nil {
		return nil, err
	}
	if err = person.IsAdmin(); err != nil {
		return nil, err
	}

	if err := l.svc.UpdateGroupName(l.ctx, group, groupName, groupPubName); err != nil {
		return nil, err
	}

	return &pb.UpdateGroupNameResp{}, nil
}
