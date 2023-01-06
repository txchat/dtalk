package oss

import (
	"context"
	"strings"

	"github.com/txchat/dtalk/app/services/oss/ossclient"
	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type CompleteMultiUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewCompleteMultiUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteMultiUploadLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &CompleteMultiUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *CompleteMultiUploadLogic) CompleteMultiUpload(req *types.CompleteMultiUploadReq) (resp *types.CompleteMultiUploadResp, err error) {
	// key 非空 且 key 不包含 ..
	if strings.TrimSpace(req.Key) == "" || strings.Contains(req.Key, "..") {
		return nil, xerror.ErrOssKeyIllegal
	}

	parts := make([]*ossclient.PartInfo, 0, len(req.Parts))
	for _, part := range req.Parts {
		parts = append(parts, &ossclient.PartInfo{
			ETag:       part.ETag,
			PartNumber: part.PartNumber,
		})
	}
	rpcResp, err := l.svcCtx.OssRPC.CompleteUploadMultiPart(l.ctx, &ossclient.CompleteUploadMultiPartReq{
		Base: &ossclient.BaseInfo{
			AppId:   req.AppId,
			OssType: req.OssType,
		},
		Key:    req.Key,
		TaskId: req.UploadId,
		Parts:  parts,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.CompleteMultiUploadResp{
		Url: rpcResp.GetUrl(),
		Uri: rpcResp.GetUri(),
	}
	return
}
