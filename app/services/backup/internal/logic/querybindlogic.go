package logic

import (
	"context"
	"fmt"

	"github.com/txchat/dtalk/app/services/backup/backup"
	"github.com/txchat/dtalk/app/services/backup/internal/model"
	"github.com/txchat/dtalk/app/services/backup/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type QueryBindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryBindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBindLogic {
	return &QueryBindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryBindLogic) QueryBind(in *backup.QueryBindReq) (*backup.QueryBindResp, error) {
	switch x := in.Params.(type) {
	case *backup.QueryBindReq_BindPhone:
		item, err := l.svcCtx.Repo.QueryBind(model.Phone, x.BindPhone.GetPhone())
		if err != nil {
			return &backup.QueryBindResp{}, err
		}
		return &backup.QueryBindResp{
			Info: &backup.AddressInfo{
				Address:    item.Address,
				Area:       item.Area,
				Phone:      item.Phone,
				Email:      item.Email,
				Mnemonic:   item.Mnemonic,
				PrivateKey: item.PrivateKey,
				UpdateTime: item.UpdateTime.UnixMicro(),
				CreateTime: item.CreateTime.UnixMicro(),
			},
		}, err
	case *backup.QueryBindReq_BindEmail:
		item, err := l.svcCtx.Repo.QueryBind(model.Email, x.BindEmail.GetEmail())
		if err != nil {
			return &backup.QueryBindResp{}, err
		}
		return &backup.QueryBindResp{
			Info: &backup.AddressInfo{
				Address:    item.Address,
				Area:       item.Area,
				Phone:      item.Phone,
				Email:      item.Email,
				Mnemonic:   item.Mnemonic,
				PrivateKey: item.PrivateKey,
				UpdateTime: item.UpdateTime.UnixMicro(),
				CreateTime: item.CreateTime.UnixMicro(),
			},
		}, err
	case *backup.QueryBindReq_BindAddress:
		item, err := l.svcCtx.Repo.QueryBind(model.Address, x.BindAddress.GetAddr())
		if err != nil {
			return &backup.QueryBindResp{}, err
		}
		return &backup.QueryBindResp{
			Info: &backup.AddressInfo{
				Address:    item.Address,
				Area:       item.Area,
				Phone:      item.Phone,
				Email:      item.Email,
				Mnemonic:   item.Mnemonic,
				PrivateKey: item.PrivateKey,
				UpdateTime: item.UpdateTime.UnixMicro(),
				CreateTime: item.CreateTime.UnixMicro(),
			},
		}, err
	case nil:
		return &backup.QueryBindResp{}, fmt.Errorf("QueryBindReq.Params is not set")
	default:
		return &backup.QueryBindResp{}, fmt.Errorf("QueryBindReq.Params has unexpected type %T", x)
	}
}
