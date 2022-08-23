package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/version/versionclient"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVersionLogic {
	return &CreateVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateVersionLogic) CreateVersion(req *types.CreateVersionReq) (resp *types.CreateVersionResp, err error) {
	createResp, err := l.svcCtx.VersionRPC.Create(l.ctx, &versionclient.CreateReq{
		VersionInfo: &versionclient.VersionInfo{
			Platform:    req.Platform,
			DeviceType:  req.DeviceType,
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
	resp = &types.CreateVersionResp{
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
