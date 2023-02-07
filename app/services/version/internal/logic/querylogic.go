package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/version/internal/svc"
	"github.com/txchat/dtalk/app/services/version/version"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLogic {
	return &QueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryLogic) Query(in *version.QueryReq) (*version.QueryResp, error) {
	vInfo, err := l.svcCtx.Repo.GetVersionInfo(l.ctx, in.GetId())
	if err != nil {
		return &version.QueryResp{}, err
	}
	return &version.QueryResp{
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
