package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/dtalk/app/services/storage/storage"
	"github.com/txchat/imparse/proto/common"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecordsAfterMidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecordsAfterMidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecordsAfterMidLogic {
	return &GetRecordsAfterMidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRecordsAfterMidLogic) GetRecordsAfterMid(in *storage.GetRecordsAfterMidReq) (*storage.GetRecordsAfterMidReply, error) {
	var items []*model.MsgContent
	var err error
	switch in.GetTp() {
	case common.Channel_ToUser:
		items, err = l.svcCtx.Repo.GetPriRecord(in.GetFrom(), in.GetTarget(), in.GetMid(), in.GetCount())
	case common.Channel_ToGroup:
		err = model.ErrCustomNotSupport
	default:
		err = model.ErrRecordNotFind
	}
	if err != nil {
		return nil, err
	}
	records := make([]*storage.GetRecordReply, len(items))
	for i, item := range items {
		records[i] = &storage.GetRecordReply{
			Mid:        item.Mid,
			Seq:        item.Seq,
			SenderId:   item.SenderId,
			ReceiverId: item.ReceiverId,
			MsgType:    item.MsgType,
			Content:    item.Content,
			CreateTime: item.CreateTime,
			Source:     item.Source,
		}
	}
	return &storage.GetRecordsAfterMidReply{
		Records: records,
	}, nil
}
