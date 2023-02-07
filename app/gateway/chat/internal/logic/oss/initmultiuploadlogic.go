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

type InitMultiUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewInitMultiUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitMultiUploadLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &InitMultiUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *InitMultiUploadLogic) InitMultiUpload(req *types.InitMultiUploadReq) (resp *types.InitMultiUploadResp, err error) {
	// key 非空 且 key 不包含 ..
	if strings.TrimSpace(req.Key) == "" || strings.Contains(req.Key, "..") {
		return nil, xerror.ErrOssKeyIllegal
	}

	rpcResp, err := l.svcCtx.OssRPC.InitUploadMultiPart(l.ctx, &ossclient.InitUploadMultiPartReq{
		Base: &ossclient.BaseInfo{
			AppId:   req.AppId,
			OssType: req.OssType,
		},
		Key: req.Key,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.InitMultiUploadResp{
		UploadId: rpcResp.GetTaskId(),
		Key:      req.Key,
	}
	return
}
