package logic

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/services/oss/internal/svc"
	"github.com/txchat/dtalk/app/services/oss/oss"
	xoss "github.com/txchat/dtalk/pkg/oss"
	"github.com/zeromicro/go-zero/core/logx"
)

type CompleteUploadMultiPartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompleteUploadMultiPartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteUploadMultiPartLogic {
	return &CompleteUploadMultiPartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompleteUploadMultiPartLogic) CompleteUploadMultiPart(in *oss.CompleteUploadMultiPartReq) (*oss.CompleteUploadMultiPartResp, error) {
	engine, err := l.svcCtx.AppOssEngines.GetEngine(in.GetBase().GetAppId(), in.GetBase().GetOssType())
	if err != nil {
		return nil, xerror.ErrFeaturesUnSupported
	}
	parts := make([]xoss.Part, 0, len(in.GetParts()))
	for _, p := range in.GetParts() {
		parts = append(parts, xoss.Part{
			ETag:       p.GetETag(),
			PartNumber: p.GetPartNumber(),
		})
	}
	url, uri, err := engine.CompleteMultipartUpload(in.GetKey(), in.GetTaskId(), parts)
	if err != nil {
		return nil, err
	}
	return &oss.CompleteUploadMultiPartResp{
		Url: url,
		Uri: uri,
	}, nil
}
