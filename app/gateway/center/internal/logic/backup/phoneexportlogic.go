package backup

import (
	"context"

	"github.com/txchat/dtalk/internal/notify"
	"github.com/txchat/dtalk/internal/notify/phpserverclient"

	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneExportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneExportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneExportLogic {
	return &PhoneExportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneExportLogic) PhoneExport(req *types.PhoneExportReq) (resp *types.PhoneExportResp, err error) {
	// 通过短信服务验证
	params := map[string]string{
		notify.Account:                req.Phone,
		notify.Code:                   req.Code,
		phpserverclient.ParamCodeType: l.svcCtx.Config.SMS.CodeTypes[phpserverclient.Export],
	}
	err = l.svcCtx.SmsValidate.ValidateCode(params)
	if err != nil {
		err = xerror.ErrExportAddressPhoneInconsistent
		return
	}

	return
}
