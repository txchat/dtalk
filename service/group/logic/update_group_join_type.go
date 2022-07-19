package logic

import (
	"context"

	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type UpdateGroupJoinTypeLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewUpdateGroupJoinTypeLogic(ctx context.Context, svc *service.Service) *UpdateGroupJoinTypeLogic {
	return &UpdateGroupJoinTypeLogic{
		ctx: ctx,
		svc: svc,
	}
}

// UpdateGroupJoinType 更新加群设置
func (l *UpdateGroupJoinTypeLogic) UpdateGroupJoinType(req *pb.UpdateGroupJoinTypeReq) (*pb.UpdateGroupJoinTypeResp, error) {
	groupId := req.GroupId
	personId := req.PersonId
	joinType := int32(req.GroupJoinType)

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

	if err := l.svc.UpdateGroupJoinType(l.ctx, group, joinType); err != nil {
		return nil, err
	}

	return &pb.UpdateGroupJoinTypeResp{}, nil
}
