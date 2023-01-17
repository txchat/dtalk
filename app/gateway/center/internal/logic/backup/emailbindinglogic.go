package backup

import (
	"context"
	"time"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/notify"

	"github.com/txchat/dtalk/app/services/backup/backup"
	"github.com/txchat/dtalk/app/services/backup/backupclient"
	xhttp "github.com/txchat/dtalk/pkg/net/http"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailBindingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewEmailBindingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailBindingLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &EmailBindingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *EmailBindingLogic) EmailBinding(req *types.EmailBindingReq) (resp *types.EmailBindingResp, err error) {
	// 通过邮箱验证
	params := map[string]string{
		notify.ParamEmail:    req.Email,
		notify.ParamCode:     req.Code,
		notify.ParamCodeType: l.svcCtx.Config.Email.CodeTypes[notify.Bind],
	}
	err = l.svcCtx.EmailValidate.ValidateCode(params)
	if err != nil {
		err = xerror.ErrCodeError
		return
	}
	_, err = l.svcCtx.BackupRPC.UpdateAddressBackup(l.ctx, &backupclient.UpdateAddressBackupReq{
		Type: backup.BackupType_Email,
		Stub: &backupclient.AddressInfo{
			Address:    l.custom.UID,
			Email:      req.Email,
			Mnemonic:   req.Mnemonic,
			UpdateTime: time.Now().UnixMicro(),
		},
	})
	resp = &types.EmailBindingResp{Address: l.custom.UID}
	return
}
