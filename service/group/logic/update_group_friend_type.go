package logic

import (
	"context"

	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type UpdateGroupFriendTypeLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewUpdateGroupFriendTypeLogic(ctx context.Context, svc *service.Service) *UpdateGroupFriendTypeLogic {
	return &UpdateGroupFriendTypeLogic{
		ctx: ctx,
		svc: svc,
	}
}

// UpdateGroupFriendType 更新群内加好友设置
func (l *UpdateGroupFriendTypeLogic) UpdateGroupFriendType(req *pb.UpdateGroupFriendTypeReq) (*pb.UpdateGroupFriendTypeResp, error) {
	groupId := req.GroupId
	personId := req.PersonId
	friendType := int32(req.GroupFriendType)

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

	if err := l.svc.UpdateGroupFriendType(l.ctx, group, friendType); err != nil {
		return nil, err
	}
	return &pb.UpdateGroupFriendTypeResp{}, nil
}
