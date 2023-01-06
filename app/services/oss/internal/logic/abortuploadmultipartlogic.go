package logic

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/services/oss/internal/svc"
	"github.com/txchat/dtalk/app/services/oss/oss"

	"github.com/zeromicro/go-zero/core/logx"
)

type AbortUploadMultiPartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAbortUploadMultiPartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AbortUploadMultiPartLogic {
	return &AbortUploadMultiPartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AbortUploadMultiPartLogic) AbortUploadMultiPart(in *oss.AbortUploadMultiPartReq) (*oss.AbortUploadMultiPartResp, error) {
	engine, err := l.svcCtx.AppOssEngines.GetEngine(in.GetBase().GetAppId(), in.GetBase().GetOssType())
	if err != nil {
		return nil, xerror.ErrFeaturesUnSupported
	}
	err = engine.AbortMultipartUpload(in.GetKey(), in.GetTaskId())
	if err != nil {
		return nil, err
	}
	return &oss.AbortUploadMultiPartResp{}, nil
}
