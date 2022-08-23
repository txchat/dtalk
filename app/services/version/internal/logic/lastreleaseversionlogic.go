package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/version/internal/svc"
	"github.com/txchat/dtalk/app/services/version/version"

	"github.com/zeromicro/go-zero/core/logx"
)

type LastReleaseVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLastReleaseVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LastReleaseVersionLogic {
	return &LastReleaseVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LastReleaseVersionLogic) LastReleaseVersion(in *version.LastReleaseVersionReq) (*version.LastReleaseVersionResp, error) {
	vInfo, err := l.svcCtx.Repo.LastReleaseVersion(l.ctx, in.GetPlatform(), in.GetDeviceType())
	if err != nil {
		return &version.LastReleaseVersionResp{}, err
	}
	return &version.LastReleaseVersionResp{
		VersionInfo: &version.VersionInfo{
			Id:          vInfo.Id,
			Platform:    vInfo.Platform,
			Status:      vInfo.Status,
			DeviceType:  vInfo.DeviceType,
			VersionName: vInfo.VersionName,
			VersionCode: vInfo.VersionCode,
			Url:         vInfo.URL,
			Force:       vInfo.Force,
			Description: vInfo.Description,
			OpeUser:     vInfo.OpeUser,
			Md5:         vInfo.Md5,
			Size:        vInfo.Size,
			UpdateTime:  vInfo.UpdateTime,
			CreateTime:  vInfo.CreateTime,
		},
	}, nil
}
