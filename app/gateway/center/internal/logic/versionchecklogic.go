package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/version/versionclient"
	xhttp "github.com/txchat/dtalk/pkg/net/http"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VersionCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewVersionCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VersionCheckLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &VersionCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *VersionCheckLogic) VersionCheck(req *types.VersionCheckReq) (resp *types.VersionCheckResp, err error) {
	req.DeviceType = l.custom.Device

	lastReleaseVersionRPCResp, err := l.svcCtx.VersionRPC.LastReleaseVersion(l.ctx, &versionclient.LastReleaseVersionReq{
		Platform:   l.svcCtx.Config.Backend.Platform,
		DeviceType: req.DeviceType,
	})

	isForce := lastReleaseVersionRPCResp.GetVersionInfo().GetForce()
	if !isForce {
		var forceNumberRPCResp *versionclient.ForceNumberBetweenResp
		forceNumberRPCResp, err = l.svcCtx.VersionRPC.ForceNumberBetween(l.ctx, &versionclient.ForceNumberBetweenReq{
			Platform:   l.svcCtx.Config.Backend.Platform,
			DeviceType: req.DeviceType,
			Begin:      req.VersionCode,
			End:        lastReleaseVersionRPCResp.GetVersionInfo().GetVersionCode(),
		})
		if err != nil {
			return nil, err
		}
		isForce = forceNumberRPCResp.Num > 0
	}

	resp = &types.VersionCheckResp{
		VersionInfo: types.VersionInfo{
			Id:          lastReleaseVersionRPCResp.GetVersionInfo().GetId(),
			Platform:    lastReleaseVersionRPCResp.GetVersionInfo().GetPlatform(),
			Status:      lastReleaseVersionRPCResp.GetVersionInfo().GetStatus(),
			DeviceType:  lastReleaseVersionRPCResp.GetVersionInfo().GetDeviceType(),
			VersionName: lastReleaseVersionRPCResp.GetVersionInfo().GetVersionName(),
			VersionCode: lastReleaseVersionRPCResp.GetVersionInfo().GetVersionCode(),
			Url:         lastReleaseVersionRPCResp.GetVersionInfo().GetUrl(),
			Force:       isForce,
			Description: lastReleaseVersionRPCResp.GetVersionInfo().GetDescription(),
			OpeUser:     lastReleaseVersionRPCResp.GetVersionInfo().GetOpeUser(),
			Md5:         lastReleaseVersionRPCResp.GetVersionInfo().GetMd5(),
			Size:        lastReleaseVersionRPCResp.GetVersionInfo().GetSize(),
			UpdateTime:  lastReleaseVersionRPCResp.GetVersionInfo().GetUpdateTime(),
			CreateTime:  lastReleaseVersionRPCResp.GetVersionInfo().GetCreateTime(),
		},
	}
	return
}
