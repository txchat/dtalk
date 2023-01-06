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

type AbortMultiUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewAbortMultiUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AbortMultiUploadLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &AbortMultiUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *AbortMultiUploadLogic) AbortMultiUpload(req *types.AbortMultiUploadReq) (resp *types.AbortMultiUploadResp, err error) {
	// key 非空 且 key 不包含 ..
	if strings.TrimSpace(req.Key) == "" || strings.Contains(req.Key, "..") {
		return nil, xerror.ErrOssKeyIllegal
	}

	_, err = l.svcCtx.OssRPC.AbortUploadMultiPart(l.ctx, &ossclient.AbortUploadMultiPartReq{
		Base: &ossclient.BaseInfo{
			AppId:   req.AppId,
			OssType: req.OssType,
		},
		Key:    req.Key,
		TaskId: req.UploadId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.AbortMultiUploadResp{}
	return
}
