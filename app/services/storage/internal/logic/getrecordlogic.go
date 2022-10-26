package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/dtalk/app/services/storage/storage"
	"github.com/txchat/imparse/proto/common"
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
	var item *model.MsgContent
	var err error
	switch in.GetTp() {
	case common.Channel_ToUser:
		item, err = l.svcCtx.Repo.GetSpecifyRecord(in.GetMid())
	case common.Channel_ToGroup:
		item, err = l.svcCtx.Repo.GetSpecifyGroupRecord(in.GetMid())
	default:
		err = model.ErrRecordNotFind
	}
	if err != nil {
		return nil, err
	}
	return &storage.GetRecordReply{
		Mid:        item.Mid,
		Seq:        item.Seq,
		SenderId:   item.SenderId,
		ReceiverId: item.ReceiverId,
		MsgType:    item.MsgType,
		Content:    item.Content,
		CreateTime: item.CreateTime,
		Source:     item.Source,
	}, nil
}
