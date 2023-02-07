package logic

import (
	"context"
	"fmt"

	"github.com/txchat/dtalk/app/services/backup/backup"
	"github.com/txchat/dtalk/app/services/backup/internal/model"
	"github.com/txchat/dtalk/app/services/backup/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRelatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRelatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRelatedLogic {
	return &QueryRelatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryRelatedLogic) QueryRelated(in *backup.QueryRelatedReq) (*backup.QueryRelatedResp, error) {
	switch x := in.Params.(type) {
	case *backup.QueryRelatedReq_BindPhone:
		item, err := l.svcCtx.Repo.QueryRelate(model.Phone, x.BindPhone.GetPhone())
		if err != nil {
			return &backup.QueryRelatedResp{}, err
		}
		return &backup.QueryRelatedResp{
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
		return &backup.QueryRelatedResp{}, fmt.Errorf("QueryBindReq.Params is not set")
	default:
		return &backup.QueryRelatedResp{}, fmt.Errorf("QueryBindReq.Params has unexpected type %T", x)
	}
}
