package backend

import (
	"context"

	"github.com/txchat/dtalk/app/services/version/versionclient"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVersionLogic {
	return &UpdateVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateVersionLogic) UpdateVersion(req *types.UpdateVersionReq) (resp *types.UpdateVersionResp, err error) {
	createResp, err := l.svcCtx.VersionRPC.Update(l.ctx, &versionclient.UpdateReq{
		VersionInfo: &versionclient.VersionInfo{
			Id:          req.Id,
			VersionName: req.VersionName,
			VersionCode: req.VersionCode,
			Url:         req.Url,
			Force:       req.Force,
			Description: req.Description,
			OpeUser:     req.OpeUser,
			Md5:         req.Md5,
			Size:        req.Size,
		},
	})
	if err != nil {
		return
	}
	queryResp, err := l.svcCtx.VersionRPC.Query(l.ctx, &versionclient.QueryReq{
		Id: createResp.GetId(),
	})
	if err != nil {
		return
	}
	resp = &types.UpdateVersionResp{
		Version: types.VersionInfo{
			Id:          queryResp.GetVersionInfo().GetId(),
			Platform:    queryResp.GetVersionInfo().GetPlatform(),
			Status:      queryResp.GetVersionInfo().GetStatus(),
			DeviceType:  queryResp.GetVersionInfo().GetDeviceType(),
			VersionName: queryResp.GetVersionInfo().GetVersionName(),
			VersionCode: queryResp.GetVersionInfo().GetVersionCode(),
			Url:         queryResp.GetVersionInfo().GetUrl(),
			Force:       queryResp.GetVersionInfo().GetForce(),
			Description: queryResp.GetVersionInfo().GetDescription(),
			OpeUser:     queryResp.GetVersionInfo().GetOpeUser(),
			Md5:         queryResp.GetVersionInfo().GetMd5(),
			Size:        queryResp.GetVersionInfo().GetSize(),
			UpdateTime:  queryResp.GetVersionInfo().GetUpdateTime(),
			CreateTime:  queryResp.GetVersionInfo().GetCreateTime(),
		},
	}
	return
}
