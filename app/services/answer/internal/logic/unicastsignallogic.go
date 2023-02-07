package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/answer/answer"
	"github.com/txchat/dtalk/app/services/answer/internal/svc"
	"github.com/txchat/imparse"

	"github.com/zeromicro/go-zero/core/logx"
)

type UniCastSignalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUniCastSignalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UniCastSignalLogic {
	return &UniCastSignalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UniCastSignalLogic) UniCastSignal(in *answer.UniCastSignalReq) (*answer.UniCastSignalReply, error) {
	data, err := signalBody(in.GetTarget(), in.GetType(), in.GetBody())
	if err != nil {
		return &answer.UniCastSignalReply{}, err
	}

	innerLogic := NewInnerPushLogic(l.ctx, l.svcCtx)
	mid, err := innerLogic.InnerPushToClient(l.svcCtx.AnswerInter4rpc, "", "", in.GetTarget(), imparse.UniCast, data)
	if err != nil {
		return &answer.UniCastSignalReply{}, err
	}
	return &answer.UniCastSignalReply{
		Mid: mid,
	}, nil
}
