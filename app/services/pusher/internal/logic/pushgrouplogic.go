package logic

import (
	"context"

	"github.com/txchat/im/api/protocol"

	"github.com/txchat/im/app/logic/logicclient"

	"github.com/txchat/dtalk/app/services/pusher/internal/svc"
	"github.com/txchat/dtalk/app/services/pusher/pusher"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushGroupLogic {
	return &PushGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushGroupLogic) PushGroup(in *pusher.PushGroupReq) (*pusher.PushGroupResp, error) {
	msg := &logicclient.PushGroupReq{
		AppId: in.GetApp(),
		Group: in.GetGid(),
		Op:    protocol.Op_Message,
		Body:  in.GetBody(),
	}

	_, err := l.svcCtx.LogicRPC.PushGroup(l.ctx, msg)
	if err != nil {
		l.Error("GroupCast PushGroup Failed", "err", err, "appId", msg.GetAppId(), "to group", msg.GetGroup(), "len of msg", len(msg.GetBody()))
		return nil, err
	}
	return &pusher.PushGroupResp{}, nil
}
