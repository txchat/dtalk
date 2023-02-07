package backup

import (
	"context"
	"regexp"

	"github.com/txchat/dtalk/app/services/backup/backup"
	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/services/backup/backupclient"

	"github.com/txchat/dtalk/app/gateway/center/internal/svc"
	"github.com/txchat/dtalk/app/gateway/center/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressLogic {
	return &GetAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAddressLogic) GetAddress(req *types.GetAddressReq) (resp *types.GetAddressResp, err error) {
	var backupType backup.BackupType
	backupType, err = checkBackupType(req.Query)
	if err != nil {
		return
	}
	params := &backupclient.QueryBindReq{}
	switch backupType {
	case backup.BackupType_Email:
		params = &backupclient.QueryBindReq{
			Params: &backup.QueryBindReq_BindEmail{BindEmail: &backupclient.QueryBindReqEmail{
				Email: req.Query,
			}},
		}
	case backup.BackupType_Phone:
		params = &backupclient.QueryBindReq{
			Params: &backup.QueryBindReq_BindPhone{BindPhone: &backupclient.QueryBindReqPhone{
				Phone: req.Query,
			}},
		}
	case backup.BackupType_Addr:
		params = &backupclient.QueryBindReq{
			Params: &backup.QueryBindReq_BindAddress{BindAddress: &backupclient.QueryBindReqAddress{
				Addr: req.Query,
			}},
		}
	}

	rpcResp, err := l.svcCtx.BackupRPC.QueryBind(l.ctx, params)
	if rpcResp != nil && rpcResp.GetInfo() != nil {
		resp = &types.GetAddressResp{
			Address: rpcResp.GetInfo().Address,
		}
	}
	if xerror.ErrNotFound.Equal(err) {
		err = nil
	}
	return
}

func checkBackupType(query string) (backup.BackupType, error) {
	if checkEmail(query) {
		return backup.BackupType_Email, nil
	}
	if checkPhone(query) {
		return backup.BackupType_Phone, nil
	}
	if checkAddress(query) {
		return backup.BackupType_Addr, nil
	}
	return backup.BackupType_Phone, xerror.ErrOutOfRange
}

func checkEmail(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func checkPhone(phone string) bool {
	//regular := "^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$"
	//
	//reg := regexp.MustCompile(regular)
	//return reg.MatchString(phone)
	return len(phone) == 11
}

func checkAddress(address string) bool {
	return len(address) == 34
}
