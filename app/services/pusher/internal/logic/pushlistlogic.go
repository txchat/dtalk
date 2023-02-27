package logic

import (
	"context"

	"github.com/txchat/im/api/protocol"

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
		ToId:  in.GetUid(),
		Op:    protocol.Op_Message,
		Body:  in.GetBody(),
	}

	_, err := l.svcCtx.LogicRPC.PushByUID(l.ctx, msg)
	if err != nil {
		if l.svcCtx.Config.OffPushEnabled {
			err = l.svcCtx.PublishThirdPartyPushMQ(l.ctx, in.GetFrom(), in.GetUid(), in.GetBody())
			if err != nil {
				l.Error(err)
			}
		}
		return nil, err
	}
	return &pusher.PushListResp{}, nil
}
