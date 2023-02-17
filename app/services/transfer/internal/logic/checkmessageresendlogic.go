package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/transfer/internal/model"

	"github.com/txchat/dtalk/app/services/transfer/internal/svc"
	"github.com/txchat/dtalk/app/services/transfer/transfer"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckMessageResendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckMessageResendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckMessageResendLogic {
	return &CheckMessageResendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckMessageResendLogic) CheckMessageResend(in *transfer.CheckMessageResendReq) (*transfer.CheckMessageResendResp, error) {
	// todo: add your logic here and delete this line
	l.svcCtx.Repo.GetMidByClientSeq(l.ctx, in.GetUuid(), in.GetSeq())

	l.svcCtx.Repo.MappingClientSeq(l.ctx, &model.MessageIndex{
		Mid:  0,
		Seq:  in.GetSeq(),
		UID:  in.GetFrom(),
		UUID: in.GetUuid(),
	})
	return &transfer.CheckMessageResendResp{}, nil
}
