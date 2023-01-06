package oss

import (
	"context"

	"github.com/txchat/dtalk/app/services/oss/ossclient"

	"github.com/txchat/dtalk/app/gateway/chat/internal/model"
	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type GetTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &GetTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *GetTokenLogic) GetToken(req *types.GetTokenReq) (resp *types.GetTokenResp, err error) {
	rpcResp, err := l.svcCtx.OssRPC.AssumeRole(l.ctx, &ossclient.AssumeRoleReq{
		Base: &ossclient.BaseInfo{
			AppId:   "dtalk",
			OssType: model.Oss_Aliyun,
		},
	})
	if err != nil {
		return nil, err
	}

	resp = &types.GetTokenResp{
		RequestId: rpcResp.GetRequestId(),
		Credentials: types.Credentials{
			AccessKeySecret: rpcResp.GetCredentials().GetAccessKeySecret(),
			Expiration:      rpcResp.GetCredentials().GetExpiration(),
			AccessKeyId:     rpcResp.GetCredentials().GetAccessKeyId(),
			SecurityToken:   rpcResp.GetCredentials().GetSecurityToken(),
		},
		AssumedRoleUser: types.AssumedRoleUser{
			AssumedRoleId: rpcResp.GetAssumedRoleUser().GetAssumedRoleId(),
			Arn:           rpcResp.GetAssumedRoleUser().GetArn(),
		},
	}
	return
}
