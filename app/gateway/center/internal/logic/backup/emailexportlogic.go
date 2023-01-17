package backup

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/notify"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

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
		notify.ParamEmail:    req.Email,
		notify.ParamCode:     req.Code,
		notify.ParamCodeType: l.svcCtx.Config.Email.CodeTypes[notify.Export],
	}
	err = l.svcCtx.EmailValidate.ValidateCode(params)
	if err != nil {
		err = xerror.ErrExportAddressEmailInconsistent
		return
	}
	return
}
