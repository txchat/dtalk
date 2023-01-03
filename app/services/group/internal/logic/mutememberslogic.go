package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MuteMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMuteMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MuteMembersLogic {
	return &MuteMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MuteMembersLogic) MuteMembers(in *group.MuteMembersReq) (*group.MuteMembersResp, error) {
	// todo: add your logic here and delete this line

	return &group.MuteMembersResp{}, nil
}
