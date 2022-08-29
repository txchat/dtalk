package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/backup/backup"
	"github.com/txchat/dtalk/app/services/backup/internal/model"
	"github.com/txchat/dtalk/app/services/backup/internal/svc"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAddressRelatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAddressRelatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAddressRelatedLogic {
	return &UpdateAddressRelatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAddressRelatedLogic) UpdateAddressRelated(in *backup.UpdateAddressRelatedReq) (*backup.UpdateAddressRelatedResp, error) {
	err := l.svcCtx.Repo.UpdateAddrRelate(int32(in.GetType()), &model.AddrRelate{
		Address:    in.GetStub().Address,
		Area:       in.GetStub().Area,
		Phone:      in.GetStub().Phone,
		Email:      in.GetStub().Email,
		Mnemonic:   in.GetStub().Mnemonic,
		PrivateKey: in.GetStub().PrivateKey,
		UpdateTime: util.UnixToTime(in.GetStub().UpdateTime),
	})
	return &backup.UpdateAddressRelatedResp{}, err
}
