package backup

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/services/backup/backup"
	"github.com/txchat/dtalk/app/services/backup/backupclient"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailRetrieveLogic {
	return &EmailRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailRetrieveLogic) EmailRetrieve(req *types.EmailRetrieveReq) (resp *types.EmailRetrieveResp, err error) {
	// todo: 通过短信服务验证

	var queryBindResp *backupclient.QueryBindResp
	queryBindResp, err = l.svcCtx.BackupRPC.QueryBind(l.ctx, &backupclient.QueryBindReq{
		Params: &backup.QueryBindReq_BindEmail{
			BindEmail: &backupclient.QueryBindReqEmail{
				Email: req.Email,
			},
		},
	})
	if queryBindResp != nil && queryBindResp.GetInfo() != nil {
		resp = &types.EmailRetrieveResp{
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
