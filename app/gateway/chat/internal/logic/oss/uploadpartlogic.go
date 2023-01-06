package oss

import (
	"context"
	"io/ioutil"
	"mime/multipart"
	"strings"

	"github.com/txchat/dtalk/app/gateway/chat/internal/model"
	"github.com/txchat/dtalk/app/services/oss/ossclient"
	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type UploadPartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewUploadPartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPartLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &UploadPartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *UploadPartLogic) UploadPart(req *types.UploadPartReq, fh *multipart.FileHeader) (resp *types.UploadPartResp, err error) {
	if fh.Size > model.MaxPartSize {
		return nil, xerror.ErrOssFileTooBig
	}

	// key 非空 且 key 不包含 ..
	if strings.TrimSpace(req.Key) == "" || strings.Contains(req.Key, "..") {
		return nil, xerror.ErrOssKeyIllegal
	}

	f, err := fh.Open()
	if err != nil {
		l.Errorf("UploadFile fh.Open err, err: %v", err)
		return nil, err
	}
	defer f.Close()

	body, err := ioutil.ReadAll(f)
	if err != nil {
		err = xerror.ErrExec
		return
	}
	if len(body) == 0 {
		err = xerror.ErrExec
		return
	}

	rpcResp, err := l.svcCtx.OssRPC.UploadMultiPart(l.ctx, &ossclient.UploadMultiPartReq{
		Base: &ossclient.BaseInfo{
			AppId:   req.AppId,
			OssType: req.OssType,
		},
		Key:        req.Key,
		Body:       body,
		TaskId:     req.UploadId,
		PartNumber: req.PartNumber,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.UploadPartResp{
		Part: types.Part{
			ETag:       rpcResp.GetPart().GetETag(),
			PartNumber: rpcResp.GetPart().GetPartNumber(),
		},
		UploadId: req.UploadId,
		Key:      req.Key,
	}
	return
}
