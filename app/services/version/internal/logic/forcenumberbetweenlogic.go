package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/version/internal/svc"
	"github.com/txchat/dtalk/app/services/version/version"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForceNumberBetweenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewForceNumberBetweenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForceNumberBetweenLogic {
	return &ForceNumberBetweenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ForceNumberBetweenLogic) ForceNumberBetween(in *version.ForceNumberBetweenReq) (*version.ForceNumberBetweenResp, error) {
	num, err := l.svcCtx.Repo.ForceNumberBetween(l.ctx, in.GetPlatform(), in.GetDeviceType(), in.GetBegin(), in.GetEnd())
	return &version.ForceNumberBetweenResp{
		Num: num,
	}, err
}
