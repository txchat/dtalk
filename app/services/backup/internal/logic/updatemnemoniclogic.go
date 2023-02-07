package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/backup/internal/model"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/services/backup/backup"
	"github.com/txchat/dtalk/app/services/backup/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMnemonicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMnemonicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMnemonicLogic {
	return &UpdateMnemonicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMnemonicLogic) UpdateMnemonic(in *backup.UpdateMnemonicReq) (*backup.UpdateMnemonicResp, error) {
	err := l.svcCtx.Repo.UpdateMnemonic(&model.AddrBackup{
		Address:    in.GetStub().Address,
		Mnemonic:   in.GetStub().Mnemonic,
		UpdateTime: util.UnixToTime(in.GetStub().UpdateTime),
	})
	return &backup.UpdateMnemonicResp{}, err
}
