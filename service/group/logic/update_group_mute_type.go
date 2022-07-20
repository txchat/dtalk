package logic

import (
	"context"

	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type UpdateGroupMuteTypeLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewUpdateGroupMuteTypeLogic(ctx context.Context, svc *service.Service) *UpdateGroupMuteTypeLogic {
	return &UpdateGroupMuteTypeLogic{
		ctx: ctx,
		svc: svc,
	}
}

// UpdateGroupMuteType 更新群内禁言设置
func (l *UpdateGroupMuteTypeLogic) UpdateGroupMuteType(req *pb.UpdateGroupMuteTypeReq) (*pb.UpdateGroupMuteTypeResp, error) {
	groupId := req.GroupId
	personId := req.PersonId
	muteType := int32(req.GroupMuteType)

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

	if err := l.svc.UpdateGroupMuteType(l.ctx, group, muteType); err != nil {
		return nil, err
	}

	return &pb.UpdateGroupMuteTypeResp{}, nil
}
