package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberExitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberExitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberExitLogic {
	return &MemberExitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberExitLogic) MemberExit(in *group.MemberExitReq) (*group.MemberExitResp, error) {
	// todo: add your logic here and delete this line

	return &group.MemberExitResp{}, nil
}
