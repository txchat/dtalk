package logic

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/services/oss/internal/svc"
	"github.com/txchat/dtalk/app/services/oss/oss"

	"github.com/zeromicro/go-zero/core/logx"
)

type EngineHostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEngineHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EngineHostLogic {
	return &EngineHostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EngineHostLogic) EngineHost(in *oss.EngineHostReq) (*oss.EngineHostResp, error) {
	engine, err := l.svcCtx.AppOssEngines.GetEngine(in.GetBase().GetAppId(), in.GetBase().GetOssType())
	if err != nil {
		return nil, xerror.ErrFeaturesUnSupported
	}
	return &oss.EngineHostResp{
		Host: engine.GetHost(),
	}, nil
}
