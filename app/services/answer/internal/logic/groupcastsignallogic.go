package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/answer/answer"
	"github.com/txchat/dtalk/app/services/answer/internal/svc"
	"github.com/txchat/imparse"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupCastSignalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupCastSignalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCastSignalLogic {
	return &GroupCastSignalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupCastSignalLogic) GroupCastSignal(in *answer.GroupCastSignalReq) (*answer.GroupCastSignalReply, error) {
	data, err := signalBody(in.GetTarget(), in.GetType(), in.GetBody())
	if err != nil {
		return &answer.GroupCastSignalReply{}, err
	}
	innerLogic := NewInnerPushLogic(l.ctx, l.svcCtx)
	mid, err := innerLogic.InnerPushToClient(l.svcCtx.AnswerInter4rpc, "", "", in.GetTarget(), imparse.GroupCast, data)
	if err != nil {
		return &answer.GroupCastSignalReply{}, err
	}
	return &answer.GroupCastSignalReply{
		Mid: mid,
	}, nil
}
