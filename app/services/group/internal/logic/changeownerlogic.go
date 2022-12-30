package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/app/services/group/internal/svc"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeOwnerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeOwnerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeOwnerLogic {
	return &ChangeOwnerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeOwnerLogic) ChangeOwner(in *group.ChangeOwnerReq) (*group.ChangeOwnerResp, error) {
	// todo: add your logic here and delete this line

	nowTS := util.TimeNowUnixMilli()
	gid := in.GetGid()

	g, err := l.svcCtx.Repo.GetGroupById(gid)
	if err != nil {
		return nil, err
	}

	tx, err := l.svcCtx.Repo.NewTx()
	if err != nil {
		return nil, err
	}
	defer tx.RollBack()

	_, _, err = l.svcCtx.Repo.UpdateGroupOwner(tx, &model.GroupInfo{
		GroupId:         gid,
		GroupUpdateTime: nowTS,
		GroupOwnerId:    in.GetNew(),
	})
	if err != nil {
		return nil, err
	}
	_, _, err = l.svcCtx.Repo.UpdateGroupMemberRole(tx, &model.GroupMember{
		GroupId:               gid,
		GroupMemberId:         g.GroupOwnerId,
		GroupMemberType:       model.GroupMemberTypeNormal,
		GroupMemberUpdateTime: nowTS,
	})
	if err != nil {
		return nil, err
	}
	_, _, err = l.svcCtx.Repo.UpdateGroupMemberRole(tx, &model.GroupMember{
		GroupId:               gid,
		GroupMemberId:         in.GetNew(),
		GroupMemberType:       model.GroupMemberTypeOwner,
		GroupMemberUpdateTime: nowTS,
	})
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return &group.ChangeOwnerResp{}, nil
}
