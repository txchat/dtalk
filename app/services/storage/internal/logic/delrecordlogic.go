package logic

import (
	"context"

	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/dtalk/app/services/storage/storage"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRecordLogic {
	return &DelRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelRecordLogic) DelRecord(in *storage.DelRecordReq) (*storage.DelRecordReply, error) {
	var err error
	switch in.GetTp() {
	case message.Channel_Private:
		_, _, err = l.svcCtx.Repo.DelPrivateMsg(in.GetMid())
	case message.Channel_Group:
		_, _, err = l.svcCtx.Repo.DelGroupMsg(in.GetMid())
	default:
		err = model.ErrRecordNotFind
	}
	if err != nil {
		return nil, err
	}
	return &storage.DelRecordReply{}, nil
}
