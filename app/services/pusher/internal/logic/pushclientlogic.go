package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/pusher/internal/svc"
	"github.com/txchat/dtalk/app/services/pusher/pusher"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushClientLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushClientLogic {
	return &PushClientLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushClientLogic) PushClient(in *pusher.PushReq) (*pusher.PushReply, error) {
	// todo: add your logic here and delete this line

	return &pusher.PushReply{}, nil
}
