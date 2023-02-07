package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/txchat/dtalk/app/services/version/internal/model"

	"github.com/txchat/dtalk/app/services/version/internal/svc"
	"github.com/txchat/dtalk/app/services/version/version"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *version.UpdateReq) (*version.UpdateResp, error) {
	v := in.VersionInfo
	if v == nil {
		return &version.UpdateResp{}, fmt.Errorf("请求参数为空")
	}
	_, _, err := l.svcCtx.Repo.UpdateVersionInfo(l.ctx, &model.VersionForm{
		Id:          v.GetId(),
		Platform:    v.GetPlatform(),
		Status:      v.GetStatus(),
		DeviceType:  v.GetDeviceType(),
		VersionName: v.GetVersionName(),
		VersionCode: v.GetVersionCode(),
		URL:         v.GetUrl(),
		Force:       v.GetForce(),
		Description: v.GetDescription(),
		OpeUser:     v.GetOpeUser(),
		Md5:         v.GetMd5(),
		Size:        v.GetSize(),
		UpdateTime:  time.Now().UnixNano() / 1e6,
	})
	return &version.UpdateResp{
		Id: v.GetId(),
	}, err
}
