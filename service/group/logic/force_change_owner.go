package logic

import (
	"context"

	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type ForceChangeOwnerLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewForceChangeOwnerLogic(ctx context.Context, svc *service.Service) *ForceChangeOwnerLogic {
	return &ForceChangeOwnerLogic{
		ctx: ctx,
		svc: svc,
	}
}

// ForceChangeOwner .
func (l *ForceChangeOwnerLogic) ForceChangeOwner(req *pb.ForceChangeOwnerReq) (*pb.ForceChangeOwnerResp, error) {
	_, err := FilteredMemberId(req.Member.Id)
	if err != nil {
		return nil, err
	}

	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, req.GroupId)
	if err != nil {
		return nil, err
	}

	// 判断这个人是否在群里
	member, err := l.svc.GetMemberByMemberIdAndGroupId(l.ctx, req.Member.Id, group.GroupId)
	if err != nil {
		return nil, err
	}

	owner, err := l.svc.GetMemberByMemberIdAndGroupId(l.ctx, group.GroupOwnerId, group.GroupId)
	if err != nil {
		return nil, err
	}

	if member.GroupMemberId == owner.GroupMemberId {
		return &pb.ForceChangeOwnerResp{}, nil
	}

	err = l.svc.ChangeOwner(l.ctx, group, owner, member)
	if err != nil {
		return nil, err
	}

	return &pb.ForceChangeOwnerResp{}, nil
}
