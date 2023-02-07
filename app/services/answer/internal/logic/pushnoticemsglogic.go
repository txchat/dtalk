package logic

import (
	"context"

	"github.com/txchat/imparse"

	"github.com/txchat/dtalk/app/services/answer/answer"
	"github.com/txchat/dtalk/app/services/answer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushNoticeMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushNoticeMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushNoticeMsgLogic {
	return &PushNoticeMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushNoticeMsgLogic) PushNoticeMsg(in *answer.PushNoticeMsgReq) (*answer.PushNoticeMsgReply, error) {
	data, err := noticeMsgData(in.GetChannelType(), in.GetFrom(), in.GetTarget(), in.GetSeq(), in.GetData())
	if err != nil {
		return &answer.PushNoticeMsgReply{}, err
	}
	innerLogic := NewInnerPushLogic(l.ctx, l.svcCtx)
	mid, err := innerLogic.InnerPushToClient(l.svcCtx.AnswerInter4rpc, "", in.GetFrom(), in.GetTarget(), imparse.Undefined, data)
	if err != nil {
		return &answer.PushNoticeMsgReply{}, err
	}
	return &answer.PushNoticeMsgReply{
		Mid: mid,
	}, nil
}
