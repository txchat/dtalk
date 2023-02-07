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

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *version.CreateReq) (*version.CreateResp, error) {
	v := in.VersionInfo
	if v == nil {
		return &version.CreateResp{}, fmt.Errorf("请求参数为空")
	}
	_, lastId, err := l.svcCtx.Repo.AddVersionInfo(l.ctx, &model.VersionForm{
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
		CreateTime:  time.Now().UnixNano() / 1e6,
	})
	return &version.CreateResp{
		Id: lastId,
	}, err
}
