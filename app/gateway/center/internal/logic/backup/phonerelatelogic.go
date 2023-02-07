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

type PhoneRelateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewPhoneRelateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneRelateLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &PhoneRelateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *PhoneRelateLogic) PhoneRelate(req *types.PhoneRelateReq) (resp *types.PhoneRelateResp, err error) {
	_, err = l.svcCtx.BackupRPC.UpdateAddressRelated(l.ctx, &backupclient.UpdateAddressRelatedReq{
		Type: backup.BackupType_Phone,
		Stub: &backupclient.AddressInfo{
			Address:    l.custom.UID,
			Area:       req.Area,
			Phone:      req.Phone,
			Mnemonic:   req.Mnemonic,
			UpdateTime: time.Now().UnixMicro(),
			CreateTime: time.Now().UnixMicro(),
		},
	})
	resp = &types.PhoneRelateResp{Address: l.custom.UID}
	return
}
