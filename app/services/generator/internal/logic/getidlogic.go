package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/generator/generator"
	"github.com/txchat/dtalk/app/services/generator/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIDLogic {
	return &GetIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetIDLogic) GetID(in *generator.GetIDReq) (*generator.GetIDReply, error) {
	return &generator.GetIDReply{
		Id: l.svcCtx.IDGenerator.NextId(),
	}, nil
}
