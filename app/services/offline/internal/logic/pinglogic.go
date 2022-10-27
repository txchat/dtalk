package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/offline/internal/svc"
	"github.com/txchat/dtalk/app/services/offline/offline"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *offline.Request) (*offline.Response, error) {
	// todo: add your logic here and delete this line

	return &offline.Response{}, nil
}
