package logic

import (
	"bytes"
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/services/oss/internal/svc"
	"github.com/txchat/dtalk/app/services/oss/oss"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadMultiPartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadMultiPartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadMultiPartLogic {
	return &UploadMultiPartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadMultiPartLogic) UploadMultiPart(in *oss.UploadMultiPartReq) (*oss.UploadMultiPartResp, error) {
	engine, err := l.svcCtx.AppOssEngines.GetEngine(in.GetBase().GetAppId(), in.GetBase().GetOssType())
	if err != nil {
		return nil, xerror.ErrFeaturesUnSupported
	}
	eTag, err := engine.UploadPart(in.GetKey(), in.GetTaskId(), bytes.NewReader(in.GetBody()), in.GetPartNumber(), 0, 0)
	if err != nil {
		return nil, err
	}
	return &oss.UploadMultiPartResp{
		Part: &oss.PartInfo{
			ETag:       eTag,
			PartNumber: 0,
		},
	}, nil
}
