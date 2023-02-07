package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/answer/answer"
	"github.com/txchat/dtalk/app/services/answer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushCommonMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushCommonMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushCommonMsgLogic {
	return &PushCommonMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushCommonMsgLogic) PushCommonMsg(in *answer.PushCommonMsgReq) (*answer.PushCommonMsgReply, error) {
	innerLogic := NewInnerPushLogic(l.ctx, l.svcCtx)
	mid, createTime, err := innerLogic.PushToClient(l.svcCtx.AnswerInter4rpc, in.GetKey(), in.GetFrom(), in.GetBody())
	if err != nil {
		return nil, err
	}
	return &answer.PushCommonMsgReply{
		Mid:  mid,
		Time: createTime,
	}, nil
}
