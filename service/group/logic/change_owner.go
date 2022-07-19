package logic

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type ChangeOwnerLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewChangeOwnerLogic(ctx context.Context, svc *service.Service) *ChangeOwnerLogic {
	return &ChangeOwnerLogic{
		ctx: ctx,
		svc: svc,
	}
}

// ChangeOwner 退出群
func (l *ChangeOwnerLogic) ChangeOwner(req *pb.ChangeOwnerReq) (*pb.ChangeOwnerResp, error) {
	groupId := req.GroupId
	personId := req.PersonId
	memberId := req.MemberId

	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, err
	}

	person, err := l.svc.GetPersonByMemberIdAndGroupId(l.ctx, personId, groupId)
	if err != nil {
		return nil, err
	}

	member, err := l.svc.GetMemberByMemberIdAndGroupId(l.ctx, memberId, groupId)
	if err != nil {
		return nil, err
	}

	// 只有群主可以转让群
	if err = person.IsOwner(); err != nil {
		return nil, err
	}

	if personId == memberId {
		return nil, xerror.NewError(xerror.GroupChangeOwnerSelf)
	}

	err = l.svc.ChangeOwner(l.ctx, group, person, member)
	if err != nil {
		return nil, err
	}

	return &pb.ChangeOwnerResp{}, nil
}
