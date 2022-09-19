package logic

import (
	"context"

	"github.com/txchat/dtalk/pkg/util"
	record "github.com/txchat/dtalk/service/record/proto"

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
	innerL := NewPusherLogic(l.ctx, l.svcCtx)
	err := innerL.UniCastDevices(&record.PushMsg{
		AppId:     l.svcCtx.Config.AppID,
		FromId:    in.GetFrom(),
		Mid:       util.MustToInt64(in.GetMid()),
		Key:       in.GetKey(),
		Target:    in.GetTarget(),
		Msg:       in.GetData(),
		Type:      in.GetType(),
		FrameType: in.GetFrameType(),
	})
	return &pusher.PushReply{}, err
}
