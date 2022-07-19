package logic

import (
	"context"

	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type UpdateGroupMemberNameLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewUpdateGroupMemberNameLogic(ctx context.Context, svc *service.Service) *UpdateGroupMemberNameLogic {
	return &UpdateGroupMemberNameLogic{
		ctx: ctx,
		svc: svc,
	}
}

// UpdateGroupMemberName 更新群成员群昵称
func (l *UpdateGroupMemberNameLogic) UpdateGroupMemberName(req *pb.UpdateGroupMemberNameReq) (*pb.UpdateGroupMemberNameResp, error) {
	groupId := req.GroupId
	personId := req.PersonId
	memberName := req.MemberName

	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, err
	}

	person, err := l.svc.GetPersonByMemberIdAndGroupId(l.ctx, personId, groupId)
	if err != nil {
		return nil, err
	}

	if err := l.svc.UpdateMemberName(l.ctx, group, person, memberName); err != nil {
		return nil, err
	}

	return &pb.UpdateGroupMemberNameResp{}, nil
}
