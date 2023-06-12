package backup

import (
	"context"

	"github.com/txchat/dtalk/internal/notify"
	"github.com/txchat/dtalk/internal/notify/phpserverclient"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"
	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type SendPhoneCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewSendPhoneCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPhoneCodeLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &SendPhoneCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *SendPhoneCodeLogic) SendPhoneCode(req *types.SendPhoneCodeReq) (resp *types.SendPhoneCodeResp, err error) {
	// 发送短信验证码
	params := map[string]string{
		notify.Account:                req.Phone,
		phpserverclient.ParamCodeType: l.svcCtx.Config.SMS.CodeTypes[req.CodeType],
	}
	_, err = l.svcCtx.SmsValidate.Send(params)
	if err != nil {
		l.Error("call email validate instance failed", "err", err)
		err = xerror.ErrCodeError
		return
	}
	return
}
