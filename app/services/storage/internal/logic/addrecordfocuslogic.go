package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/dtalk/app/services/storage/storage"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRecordFocusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRecordFocusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRecordFocusLogic {
	return &AddRecordFocusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRecordFocusLogic) AddRecordFocus(in *storage.AddRecordFocusReq) (*storage.AddRecordFocusReply, error) {
	err := l.svcCtx.Repo.AddRecordFocus(in.GetUid(), in.GetMid(), in.GetTime())
	if err != nil {
		return nil, err
	}
	currentNum, err := l.svcCtx.Repo.GetRecordFocusNumber(in.GetMid())
	if err != nil {
		return nil, err
	}
	return &storage.AddRecordFocusReply{
		CurrentNum: currentNum,
	}, nil
}
