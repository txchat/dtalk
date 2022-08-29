package backup

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/services/backup/backup"
	"github.com/txchat/dtalk/app/services/backup/backupclient"
	xhttp "github.com/txchat/dtalk/pkg/net/http"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewAddressRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressRetrieveLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &AddressRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *AddressRetrieveLogic) AddressRetrieve(req *types.AddressRetrieveReq) (resp *types.AddressRetrieveResp, err error) {
	var queryBindResp *backupclient.QueryBindResp
	queryBindResp, err = l.svcCtx.BackupRPC.QueryBind(l.ctx, &backupclient.QueryBindReq{
		Params: &backup.QueryBindReq_BindAddress{
			BindAddress: &backupclient.QueryBindReqAddress{
				Addr: l.custom.UID,
			},
		},
	})
	if queryBindResp != nil && queryBindResp.GetInfo() != nil {
		resp = &types.AddressRetrieveResp{
			AddressInfo: types.AddressInfo{
				Address:    queryBindResp.GetInfo().GetAddress(),
				Area:       queryBindResp.GetInfo().GetArea(),
				Phone:      queryBindResp.GetInfo().GetPhone(),
				Email:      queryBindResp.GetInfo().GetEmail(),
				Mnemonic:   queryBindResp.GetInfo().GetMnemonic(),
				PrivateKey: queryBindResp.GetInfo().GetPrivateKey(),
				UpdateTime: queryBindResp.GetInfo().GetUpdateTime(),
				CreateTime: queryBindResp.GetInfo().GetCreateTime(),
			},
		}
	}
	if xerror.ErrNotFound.Equal(err) {
		err = nil
	}
	return
}
