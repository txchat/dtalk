package logic

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/services/oss/internal/svc"
	"github.com/txchat/dtalk/app/services/oss/oss"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitUploadMultiPartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitUploadMultiPartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitUploadMultiPartLogic {
	return &InitUploadMultiPartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitUploadMultiPartLogic) InitUploadMultiPart(in *oss.InitUploadMultiPartReq) (*oss.InitUploadMultiPartResp, error) {
	engine, err := l.svcCtx.AppOssEngines.GetEngine(in.GetBase().GetAppId(), in.GetBase().GetOssType())
	if err != nil {
		return nil, xerror.ErrFeaturesUnSupported
	}
	taskId, err := engine.InitiateMultipartUpload(in.GetKey())
	if err != nil {
		return nil, err
	}
	return &oss.InitUploadMultiPartResp{
		TaskId: taskId,
	}, nil
}
