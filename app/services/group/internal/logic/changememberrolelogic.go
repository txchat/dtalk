package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &group.ChangeMemberRoleResp{}, nil
}
