package logic

import (
	"context"

	xgroup "github.com/txchat/dtalk/internal/group"

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
	nowTS := util.TimeNowUnixMilli()
	gid := in.GetGid()

	gInfo, err := l.svcCtx.Repo.GetGroupById(gid)
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
		GroupMemberId:         gInfo.GroupOwnerId,
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

	err = l.svcCtx.SignalHub.UpdateGroupMemberRole(l.ctx, gid, in.GetNew(), xgroup.RoleOwner)
	err = l.svcCtx.SignalHub.UpdateGroupMemberRole(l.ctx, gid, in.GetOperator(), xgroup.RoleNormal)
	err = l.svcCtx.NoticeHub.UpdateGroupMemberRole(l.ctx, gid, in.GetOperator(), in.GetNew(), xgroup.RoleOwner)
	return &group.ChangeOwnerResp{}, nil
}
