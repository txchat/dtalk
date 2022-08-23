package logic

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetModulesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetModulesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetModulesLogic {
	return &GetModulesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetModulesLogic) GetModules(req *types.GetModulesReq) (resp []types.GetModulesResp, err error) {
	modules := l.svcCtx.Config.Modules
	resp = make([]types.GetModulesResp, len(modules))
	for i, module := range modules {
		resp[i] = types.GetModulesResp{
			Name:      module.Name,
			IsEnabled: module.IsEnabled,
			EndPoints: module.EndPoints,
		}
	}
	return
}