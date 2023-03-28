package backup

import (
	"context"

	"github.com/txchat/dtalk/internal/notify"
	"github.com/txchat/dtalk/internal/notify/phpserverclient"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"
	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailExportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailExportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailExportLogic {
	return &EmailExportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailExportLogic) EmailExport(req *types.EmailExportReq) (resp *types.EmailExportResp, err error) {
	// 通过邮箱验证
	params := map[string]string{
		notify.Account:                req.Email,
		notify.Code:                   req.Code,
		phpserverclient.ParamCodeType: l.svcCtx.Config.Email.CodeTypes[phpserverclient.Export],
	}
	err = l.svcCtx.EmailValidate.ValidateCode(params)
	if err != nil {
		err = xerror.ErrExportAddressEmailInconsistent
		return
	}
	return
}
