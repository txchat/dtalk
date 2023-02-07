package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/backup/backup"
	"github.com/txchat/dtalk/app/services/backup/internal/model"
	"github.com/txchat/dtalk/app/services/backup/internal/svc"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAddressBackupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAddressBackupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAddressBackupLogic {
	return &UpdateAddressBackupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAddressBackupLogic) UpdateAddressBackup(in *backup.UpdateAddressBackupReq) (*backup.UpdateAddressBackupResp, error) {
	err := l.svcCtx.Repo.UpdateAddrBackup(int32(in.GetType()), &model.AddrBackup{
		Address:    in.GetStub().Address,
		Area:       in.GetStub().Area,
		Phone:      in.GetStub().Phone,
		Email:      in.GetStub().Email,
		Mnemonic:   in.GetStub().Mnemonic,
		PrivateKey: in.GetStub().PrivateKey,
		UpdateTime: util.UnixToTime(in.GetStub().UpdateTime),
	})
	return &backup.UpdateAddressBackupResp{}, err
}
