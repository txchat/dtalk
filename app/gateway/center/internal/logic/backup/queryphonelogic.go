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

type QueryPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPhoneLogic {
	return &QueryPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPhoneLogic) QueryPhone(req *types.QueryPhoneReq) (resp *types.QueryPhoneResp, err error) {
	var queryBindResp *backupclient.QueryBindResp
	var queryRelatedResp *backupclient.QueryRelatedResp
	queryBindResp, err = l.svcCtx.BackupRPC.QueryBind(l.ctx, &backupclient.QueryBindReq{
		Params: &backup.QueryBindReq_BindPhone{
			BindPhone: &backupclient.QueryBindReqPhone{
				Area:  req.Area,
				Phone: req.Phone,
			},
		},
	})
	isExists := queryBindResp != nil && queryBindResp.GetInfo() != nil
	if queryBindResp == nil || queryBindResp.GetInfo() == nil {
		queryRelatedResp, err = l.svcCtx.BackupRPC.QueryRelated(l.ctx, &backupclient.QueryRelatedReq{
			Params: &backup.QueryRelatedReq_BindPhone{
				BindPhone: &backupclient.QueryRelatedReqPhone{
					Area:  req.Area,
					Phone: req.Phone,
				},
			},
		})
		isExists = queryRelatedResp != nil && queryRelatedResp.GetInfo() != nil
	}
	resp = &types.QueryPhoneResp{Exists: isExists}
	if xerror.ErrNotFound.Equal(err) {
		err = nil
	}
	return
}
