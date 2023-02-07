package backend

import (
	"context"
	"math"

	"github.com/txchat/dtalk/app/services/version/versionclient"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVersionLogic {
	return &ListVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListVersionLogic) ListVersion(req *types.ListVersionReq) (resp *types.ListVersionResp, err error) {
	size := int64(20)
	req.Platform = l.svcCtx.Config.Backend.Platform
	versionInfoRPCResp, err := l.svcCtx.VersionRPC.SpecificPlatformAndDeviceTypeVersions(l.ctx, &versionclient.SpecificPlatformAndDeviceTypeVersionsReq{
		Platform:   req.Platform,
		DeviceType: req.DeviceType,
		Page:       req.Page,
		Size:       size,
	})
	if err != nil {
		return
	}
	versionCountRPCResp, err := l.svcCtx.VersionRPC.SpecificPlatformAndDeviceTypeCount(l.ctx, &versionclient.SpecificPlatformAndDeviceTypeCountReq{
		Platform:   req.Platform,
		DeviceType: req.DeviceType,
	})
	if err != nil {
		return
	}

	versionList := make([]types.VersionInfo, len(versionInfoRPCResp.GetVersionInfo()))
	for i, info := range versionInfoRPCResp.GetVersionInfo() {
		versionList[i] = types.VersionInfo{
			Id:          info.Id,
			Platform:    info.Platform,
			Status:      info.Status,
			DeviceType:  info.DeviceType,
			VersionName: info.VersionName,
			VersionCode: info.VersionCode,
			Url:         info.Url,
			Force:       info.Force,
			Description: info.Description,
			OpeUser:     info.OpeUser,
			Md5:         info.Md5,
			Size:        info.Size,
			UpdateTime:  info.UpdateTime,
			CreateTime:  info.CreateTime,
		}
	}
	resp = &types.ListVersionResp{
		TotalElements: versionCountRPCResp.GetTotalCount(),
		TotalPages:    int64(math.Ceil(float64(versionCountRPCResp.GetTotalCount()) / float64(size))),
		VersionList:   versionList,
	}
	return
}
