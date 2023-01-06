package logic

import (
	"bytes"
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/services/oss/internal/svc"
	"github.com/txchat/dtalk/app/services/oss/oss"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadLogic) Upload(in *oss.UploadReq) (*oss.UploadResp, error) {
	engine, err := l.svcCtx.AppOssEngines.GetEngine(in.GetBase().GetAppId(), in.GetBase().GetOssType())
	if err != nil {
		return nil, xerror.ErrFeaturesUnSupported
	}
	url, uri, err := engine.Upload(in.GetKey(), bytes.NewReader(in.GetBody()), in.GetSize())
	if err != nil {
		return nil, err
	}
	return &oss.UploadResp{
		Url: url,
		Uri: uri,
	}, nil
}
