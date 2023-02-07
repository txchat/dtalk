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

type QueryEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryEmailLogic {
	return &QueryEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryEmailLogic) QueryEmail(req *types.QueryEmailReq) (resp *types.QueryEmailResp, err error) {
	var queryBindResp *backupclient.QueryBindResp
	queryBindResp, err = l.svcCtx.BackupRPC.QueryBind(l.ctx, &backupclient.QueryBindReq{
		Params: &backup.QueryBindReq_BindEmail{
			BindEmail: &backupclient.QueryBindReqEmail{
				Email: req.Email,
			},
		},
	})
	isExists := queryBindResp != nil && queryBindResp.GetInfo() != nil
	resp = &types.QueryEmailResp{Exists: isExists}
	if xerror.ErrNotFound.Equal(err) {
		err = nil
	}
	return
}
