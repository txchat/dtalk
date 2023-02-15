package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/generator/generatorclient"
	"github.com/txchat/dtalk/app/services/transfer/internal/model"
	"github.com/txchat/dtalk/app/services/transfer/internal/svc"
	"github.com/txchat/dtalk/app/services/transfer/transfer"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/im/api/protocol"
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
	idGenResp, err := l.svcCtx.IDGenClient.GetID(l.ctx, &generatorclient.GetIDReq{})
	if err != nil {
		return nil, err
	}
	// 组装消息协议
	p := &protocol.Proto{
		Ver:     model.NowProtoVersion,
		Op:      int32(protocol.Op_ReceiveMsg),
		Seq:     0,
		Ack:     0,
		Mid:     idGenResp.GetId(),
		Channel: protocol.Channel(in.GetChannelType()),
		Target:  in.GetTarget(),
		Time:    util.TimeNowUnixMilli(),
		Body:    in.GetBody(),
	}
	return &transfer.TransferMessageResp{}, l.svcCtx.TransferMessage(l.ctx, in.GetFrom(), p)
}
