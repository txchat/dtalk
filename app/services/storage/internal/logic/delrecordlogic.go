package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/dtalk/app/services/storage/storage"
	"github.com/txchat/imparse/proto/common"
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
	// todo: add your logic here and delete this line
	var err error
	switch in.GetTp() {
	case common.Channel_ToUser:
		_, _, err = l.svcCtx.Repo.DelMsgContent(in.GetMid())
	case common.Channel_ToGroup:
		_, _, err = l.svcCtx.Repo.DelGroupMsgContent(in.GetMid())
	default:
		err = model.ErrRecordNotFind
	}
	if err != nil {
		return nil, err
	}
	return &storage.DelRecordReply{}, nil
}
