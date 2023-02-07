package oss

import (
	"context"

	"github.com/txchat/dtalk/app/services/oss/ossclient"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type GetHostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewGetHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHostLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &GetHostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *GetHostLogic) GetHost(req *types.GetHostReq) (resp *types.GetHostResp, err error) {
	rpcResp, err := l.svcCtx.OssRPC.EngineHost(l.ctx, &ossclient.EngineHostReq{
		Base: &ossclient.BaseInfo{
			AppId:   req.AppId,
			OssType: req.OssType,
		},
	})
	if err != nil {
		return nil, err
	}

	resp = &types.GetHostResp{
		Host: rpcResp.GetHost(),
	}
	return
}
