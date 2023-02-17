package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/transfer/internal/svc"
	"github.com/txchat/dtalk/app/services/transfer/transfer"
	"github.com/zeromicro/go-zero/core/logx"
)

type TransferMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransferMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferMessageLogic {
	return &TransferMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TransferMessageLogic) TransferMessage(in *transfer.TransferMessageReq) (*transfer.TransferMessageResp, error) {
	return &transfer.TransferMessageResp{}, l.svcCtx.TransferMessage(l.ctx, in.GetChannelType(), in.GetFrom(), in.GetTarget(), in.GetBody())
}
