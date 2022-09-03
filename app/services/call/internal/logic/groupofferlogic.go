package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/call/call"
	"github.com/txchat/dtalk/app/services/call/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupOfferLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupOfferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupOfferLogic {
	return &GroupOfferLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupOfferLogic) GroupOffer(in *call.GroupOfferReq) (*call.GroupOfferResp, error) {
	// todo: add your logic here and delete this line

	return &call.GroupOfferResp{}, nil
}
