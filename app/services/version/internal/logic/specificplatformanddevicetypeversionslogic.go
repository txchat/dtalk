package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/version/internal/svc"
	"github.com/txchat/dtalk/app/services/version/version"

	"github.com/zeromicro/go-zero/core/logx"
)

type SpecificPlatformAndDeviceTypeVersionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSpecificPlatformAndDeviceTypeVersionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SpecificPlatformAndDeviceTypeVersionsLogic {
	return &SpecificPlatformAndDeviceTypeVersionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SpecificPlatformAndDeviceTypeVersionsLogic) SpecificPlatformAndDeviceTypeVersions(in *version.SpecificPlatformAndDeviceTypeVersionsReq) (*version.SpecificPlatformAndDeviceTypeVersionsReqResp, error) {
	versions, err := l.svcCtx.Repo.SpecificPlatformAndDeviceTypeVersions(l.ctx, in.GetPlatform(), in.GetDeviceType(), in.GetPage(), in.GetSize())
	if err != nil {
		return &version.SpecificPlatformAndDeviceTypeVersionsReqResp{}, err
	}

	versionInfo := make([]*version.VersionInfo, len(versions))
	for i, v := range versions {
		versionInfo[i] = &version.VersionInfo{
			Id:          v.Id,
			Platform:    v.Platform,
			Status:      v.Status,
			DeviceType:  v.DeviceType,
			VersionName: v.VersionName,
			VersionCode: v.VersionCode,
			Url:         v.URL,
			Force:       v.Force,
			Description: v.Description,
			OpeUser:     v.OpeUser,
			Md5:         v.Md5,
			Size:        v.Size,
			UpdateTime:  v.UpdateTime,
			CreateTime:  v.CreateTime,
		}
	}
	return &version.SpecificPlatformAndDeviceTypeVersionsReqResp{
		VersionInfo: versionInfo,
	}, nil
}
