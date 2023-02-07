package backup

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/notify"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type SendEmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewSendEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailCodeLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &SendEmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *SendEmailCodeLogic) SendEmailCode(req *types.SendEmailCodeReq) (resp *types.SendEmailCodeResp, err error) {
	// 发送邮箱验证码
	params := map[string]string{
		notify.ParamEmail:    req.Email,
		notify.ParamCodeType: l.svcCtx.Config.SMS.CodeTypes[req.CodeType],
	}
	_, err = l.svcCtx.EmailValidate.Send(params)
	if err != nil {
		l.Error("call email validate instance failed", "err", err)
		err = xerror.ErrCodeError
		return
	}
	return
}
