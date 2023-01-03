package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type KickOutMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKickOutMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KickOutMembersLogic {
	return &KickOutMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KickOutMembersLogic) KickOutMembers(in *group.KickOutMembersReq) (*group.KickOutMembersResp, error) {
	// todo: add your logic here and delete this line

	return &group.KickOutMembersResp{}, nil
}
