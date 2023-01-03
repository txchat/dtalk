package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnMuteMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnMuteMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnMuteMembersLogic {
	return &UnMuteMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnMuteMembersLogic) UnMuteMembers(in *group.UnMuteMembersReq) (*group.UnMuteMembersResp, error) {
	// todo: add your logic here and delete this line

	return &group.UnMuteMembersResp{}, nil
}
