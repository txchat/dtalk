package logic

import (
	"context"
	"github.com/rs/zerolog"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/group/model/types"
	"github.com/txchat/dtalk/service/group/service"
)

type ChangeOwnerLogic struct {
	ctx context.Context
	svc *service.Service
	log zerolog.Logger
}

func NewChangeOwnerLogic(ctx context.Context, svc *service.Service) *ChangeOwnerLogic {
	return &ChangeOwnerLogic{
		ctx: ctx,
		svc: svc,
		log: svc.GetLog(),
	}
}

func (l *ChangeOwnerLogic) ChangeOwner(req *types.ChangeOwnerRequest) (res *types.ChangeOwnerResponse, err error) {
	groupId := req.Id
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

	return res, nil
}
