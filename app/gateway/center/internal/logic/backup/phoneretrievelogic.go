package backup

import (
	"context"

	"github.com/txchat/dtalk/pkg/notify"

	"github.com/txchat/dtalk/app/services/backup/backup"
	"github.com/txchat/dtalk/app/services/backup/backupclient"
	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneRetrieveLogic {
	return &PhoneRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneRetrieveLogic) PhoneRetrieve(req *types.PhoneRetrieveReq) (resp *types.PhoneRetrieveResp, err error) {
	// 通过短信服务验证
	params := map[string]string{
		notify.ParamMobile:   req.Phone,
		notify.ParamCode:     req.Code,
		notify.ParamCodeType: l.svcCtx.Config.SMS.CodeTypes[notify.Bind],
	}
	err = l.svcCtx.SmsValidate.ValidateCode(params)
	if err != nil {
		err = xerror.ErrCodeError
		return
	}

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
	if queryBindResp == nil || queryBindResp.GetInfo() == nil {
		// query
		queryRelatedResp, err = l.svcCtx.BackupRPC.QueryRelated(l.ctx, &backupclient.QueryRelatedReq{
			Params: &backup.QueryRelatedReq_BindPhone{
				BindPhone: &backupclient.QueryRelatedReqPhone{
					Area:  req.Area,
					Phone: req.Phone,
				},
			},
		})
		if queryRelatedResp != nil && queryRelatedResp.GetInfo() != nil {
			resp = &types.PhoneRetrieveResp{
				AddressInfo: types.AddressInfo{
					Address:    queryRelatedResp.GetInfo().GetAddress(),
					Area:       queryRelatedResp.GetInfo().GetArea(),
					Phone:      queryRelatedResp.GetInfo().GetPhone(),
					Email:      queryRelatedResp.GetInfo().GetEmail(),
					Mnemonic:   queryRelatedResp.GetInfo().GetMnemonic(),
					PrivateKey: queryRelatedResp.GetInfo().GetPrivateKey(),
					UpdateTime: queryRelatedResp.GetInfo().GetUpdateTime(),
					CreateTime: queryRelatedResp.GetInfo().GetCreateTime(),
				},
			}
			// and update
			_, err = l.svcCtx.BackupRPC.UpdateAddressBackup(l.ctx, &backupclient.UpdateAddressBackupReq{
				Type: backup.BackupType_Phone,
				Stub: queryRelatedResp.GetInfo(),
			})
		}
		if xerror.ErrNotFound.Equal(err) {
			err = nil
		}
		return
	}
	resp = &types.PhoneRetrieveResp{
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
	return
}
