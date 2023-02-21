package logic

import (
	"context"

	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/dtalk/app/services/storage/storage"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecordLogic {
	return &GetRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRecordLogic) GetRecord(in *storage.GetRecordReq) (*storage.GetRecordReply, error) {
	var record *model.MsgContent
	var err error
	switch in.GetTp() {
	case message.Channel_Private:
		record, err = l.svcCtx.Repo.GetPrivateMsgByMid(in.GetMid())
	case message.Channel_Group:
		record, err = l.svcCtx.Repo.GetGroupMsgByMid(in.GetMid())
	default:
		err = model.ErrRecordNotFind
	}
	if err != nil {
		return nil, err
	}
	return &storage.GetRecordReply{
		Record: model.ChatMsgRepoToRPC(record),
	}, nil
}
