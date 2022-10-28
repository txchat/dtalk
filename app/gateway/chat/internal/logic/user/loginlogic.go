package user

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
