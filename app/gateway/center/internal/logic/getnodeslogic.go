package logic

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetNodesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodesLogic {
	return &GetNodesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNodesLogic) GetNodes(req *types.GetNodesReq) (resp *types.GetNodesResp, err error) {
	resp = &types.GetNodesResp{
		Servers: l.svcCtx.Config.ChatNodes,
		Nodes:   l.svcCtx.Config.ContractNodes,
	}
	return
}
