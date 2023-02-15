package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/pusher/internal/svc"
	"github.com/txchat/dtalk/app/services/pusher/pusher"
	"github.com/txchat/im/app/logic/logicclient"
	"github.com/zeromicro/go-zero/core/logx"
)

type PushListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushListLogic {
	return &PushListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushListLogic) PushList(in *pusher.PushListReq) (*pusher.PushListResp, error) {
	msg := &logicclient.PushByUIDReq{
		AppId: in.GetApp(),
		ToId:  append(in.GetUid(), in.GetFrom()),
		Msg:   in.GetBody(),
	}

	if l.svcCtx.Config.OffPushEnabled {
		l.svcCtx.PushOffline(l.ctx, in.GetApp(), in.GetFrom(), in.GetUid())
	}

	_, err := l.svcCtx.LogicRPC.PushByUID(l.ctx, msg)
	if err != nil {
		l.Error("PushByUID Failed", "err", err, "appId", msg.GetAppId(), "toId", msg.GetToId(), "len of msg", len(msg.GetMsg()))
		return nil, err
	}
	return &pusher.PushListResp{}, nil
}
