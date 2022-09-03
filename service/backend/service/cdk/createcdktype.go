package cdk

import (
	"time"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/backend/model/db"
	"github.com/txchat/dtalk/service/backend/model/types"
)

func (s *ServiceContent) CreateCdkTypeSvc(req *types.CreateCdkTypeReq) (res *types.CreateCdkTypeResp, err error) {
	defer func() {
		if err != nil {
			s.log.Err(err).Interface("req", req).Msg("CreateCdkTypeSvc")
		} else {
			s.log.Info().Interface("req", req).Msg("CreateCdkTypeSvc")
		}
	}()

	cdkId, err := s.idGenRPCClient.GetID()
	if err != nil {
		return nil, err
	}

	_, err = s.GetCdkTypeByCoinName(req.CoinName)
	if err != nil && err.Error() != xerror.NewError(xerror.CdkCoinNameErr).Error() {
		return nil, err
	}

	cdkTypePo := &db.CdkType{
		CdkId:        cdkId,
		CdkName:      req.CdkName,
		CdkInfo:      req.CdkInfo,
		CoinName:     req.CoinName,
		ExchangeRate: req.ExchangeRate,
		TimeInfo: db.TimeInfo{
			CreateTime: time.Now().UnixNano() / 1e6,
			UpdateTime: time.Now().UnixNano() / 1e6,
			DeleteTime: 0,
		},
	}
	err = s.dao.InsertCdkType(cdkTypePo)
	if err != nil {
		return nil, err
	}
	res = &types.CreateCdkTypeResp{
		CdkId: util.MustToString(cdkId),
	}

	return res, nil
}
