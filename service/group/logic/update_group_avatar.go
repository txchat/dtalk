package logic

import (
	"context"

	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type UpdateGroupAvatarLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewUpdateGroupAvatarLogic(ctx context.Context, svc *service.Service) *UpdateGroupAvatarLogic {
	return &UpdateGroupAvatarLogic{
		ctx: ctx,
		svc: svc,
	}
}

// UpdateGroupAvatar 更新群头像
func (l *UpdateGroupAvatarLogic) UpdateGroupAvatar(req *pb.UpdateGroupAvatarReq) (*pb.UpdateGroupAvatarResp, error) {
	groupId := req.GroupId
	personId := req.PersonId
	groupAvatar := req.Avatar

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

	if err := l.svc.UpdateGroupAvatar(l.ctx, group, groupAvatar); err != nil {
		return nil, err
	}

	return &pb.UpdateGroupAvatarResp{}, nil
}
