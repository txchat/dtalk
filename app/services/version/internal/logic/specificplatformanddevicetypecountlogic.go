package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/version/internal/svc"
	"github.com/txchat/dtalk/app/services/version/version"

	"github.com/zeromicro/go-zero/core/logx"
)

type SpecificPlatformAndDeviceTypeCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSpecificPlatformAndDeviceTypeCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SpecificPlatformAndDeviceTypeCountLogic {
	return &SpecificPlatformAndDeviceTypeCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SpecificPlatformAndDeviceTypeCountLogic) SpecificPlatformAndDeviceTypeCount(in *version.SpecificPlatformAndDeviceTypeCountReq) (*version.SpecificPlatformAndDeviceTypeCountResp, error) {
	totalCount, err := l.svcCtx.Repo.SpecificPlatformAndDeviceTypeCount(l.ctx, in.GetPlatform(), in.GetDeviceType())
	return &version.SpecificPlatformAndDeviceTypeCountResp{
		TotalCount: totalCount,
	}, err
}
