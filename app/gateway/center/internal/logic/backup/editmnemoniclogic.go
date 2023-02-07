package backup

import (
	"context"
	"time"

	"github.com/txchat/dtalk/app/services/backup/backupclient"
	xhttp "github.com/txchat/dtalk/pkg/net/http"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditMnemonicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewEditMnemonicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditMnemonicLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &EditMnemonicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *EditMnemonicLogic) EditMnemonic(req *types.EditMnemonicReq) (resp *types.EditMnemonicResp, err error) {
	_, err = l.svcCtx.BackupRPC.UpdateMnemonic(l.ctx, &backupclient.UpdateMnemonicReq{
		Stub: &backupclient.AddressInfo{
			Address:    l.custom.UID,
			Mnemonic:   req.Mnemonic,
			UpdateTime: time.Now().UnixMicro(),
		},
	})
	return
}
