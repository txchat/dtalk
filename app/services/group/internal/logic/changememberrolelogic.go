package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/internal/model"
	xgroup "github.com/txchat/dtalk/internal/group"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeMemberRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeMemberRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeMemberRoleLogic {
	return &ChangeMemberRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeMemberRoleLogic) ChangeMemberRole(in *group.ChangeMemberRoleReq) (*group.ChangeMemberRoleResp, error) {
	tx, err := l.svcCtx.Repo.NewTx()
	if err != nil {
		return nil, err
	}
	defer tx.RollBack()

	_, _, err = l.svcCtx.Repo.UpdateGroupMemberRole(tx, &model.GroupMember{
		GroupId:               in.GetGid(),
		GroupMemberId:         in.GetMid(),
		GroupMemberType:       int32(in.GetRole()),
		GroupMemberUpdateTime: util.TimeNowUnixMilli(),
	})
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	err = l.svcCtx.SignalHub.UpdateGroupMemberRole(l.ctx, in.GetGid(), in.GetMid(), xgroup.RoleType(in.GetRole()))
	return &group.ChangeMemberRoleResp{}, nil
}
