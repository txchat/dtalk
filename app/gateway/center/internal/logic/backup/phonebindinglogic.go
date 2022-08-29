package backup

import (
	"context"
	"time"

	"github.com/txchat/dtalk/app/services/backup/backup"
	"github.com/txchat/dtalk/app/services/backup/backupclient"
	xhttp "github.com/txchat/dtalk/pkg/net/http"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneBindingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewPhoneBindingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneBindingLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &PhoneBindingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *PhoneBindingLogic) PhoneBinding(req *types.PhoneBindingReq) (resp *types.PhoneBindingResp, err error) {
	// todo: 通过短信服务验证
	_, err = l.svcCtx.BackupRPC.UpdateAddressBackup(l.ctx, &backupclient.UpdateAddressBackupReq{
		Type: backup.BackupType_Phone,
		Stub: &backupclient.AddressInfo{
			Address:    l.custom.UID,
			Area:       req.Area,
			Phone:      req.Phone,
			Mnemonic:   req.Mnemonic,
			UpdateTime: time.Now().UnixMicro(),
		},
	})
	resp = &types.PhoneBindingResp{Address: l.custom.UID}
	return
}
