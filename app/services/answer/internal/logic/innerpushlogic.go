package logic

import (
	"bytes"
	"context"

	"github.com/txchat/dtalk/internal/bizproto"
	"github.com/txchat/imparse"

	"github.com/txchat/dtalk/app/services/answer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type InnerPushLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInnerPushLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InnerPushLogic {
	return &InnerPushLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InnerPushLogic) InnerPushToClient(exec imparse.Answer, key, from, target string, pushType imparse.Channel, body []byte) (int64, error) {
	var mid int64
	frame, err := l.svcCtx.Parser.NewFrame(key, from, bytes.NewReader(body), bizproto.WithTarget(target), bizproto.WithTransmissionMethod(pushType))
	if err != nil {
		return 0, err
	}

	_, err = exec.Filter(l.ctx, frame)
	if err != nil {
		return 0, err
	}
	if err = exec.Transport(l.ctx, frame); err != nil {
		return 0, err
	}
	mid, err = exec.Ack(l.ctx, frame)
	if err != nil {
		return 0, err
	}

	return mid, nil
}

func (l *InnerPushLogic) PushToClient(exec imparse.Answer, key, from string, body []byte) (int64, uint64, error) {
	frame, err := l.svcCtx.Parser.NewFrame(key, from, bytes.NewReader(body))
	if err != nil {
		return 0, 0, err
	}

	if err = exec.Check(l.ctx, l.svcCtx.MsgChecker, frame); err != nil {
		return 0, 0, err
	}
	createTime, err := exec.Filter(l.ctx, frame)
	if err != nil {
		return 0, 0, err
	}
	if err = exec.Transport(l.ctx, frame); err != nil {
		return 0, 0, err
	}
	mid, err := exec.Ack(l.ctx, frame)
	if err != nil {
		return 0, 0, err
	}

	return mid, createTime, nil
}
