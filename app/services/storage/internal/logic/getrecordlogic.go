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
	var item *model.PrivateMsgContent
	var err error
	switch in.GetTp() {
	case message.Channel_Private:
		item, err = l.svcCtx.Repo.GetPrivateRecordByMid(in.GetMid())
	case message.Channel_Group:
		item, err = l.svcCtx.Repo.GetGroupRecordByMid(in.GetMid())
	default:
		err = model.ErrRecordNotFind
	}
	if err != nil {
		return nil, err
	}
	return &storage.GetRecordReply{
		Mid:        item.Mid,
		Seq:        item.Cid,
		SenderId:   item.SenderId,
		ReceiverId: item.ReceiverId,
		MsgType:    item.MsgType,
		Content:    item.Content,
		CreateTime: item.CreateTime,
		Source:     item.Source,
	}, nil
}
