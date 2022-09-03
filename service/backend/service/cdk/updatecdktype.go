package cdk

import (
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/backend/model/types"
)

func (s *ServiceContent) UpdateCdkTypeSvc(req *types.UpdateCdkTypeReq) (res *types.UpdateCdkTypeResp, err error) {
	defer func() {
		if err != nil {
			s.log.Err(err).Interface("req", req).Msg("UpdateCdkTypeSvc")
		} else {
			s.log.Info().Interface("req", req).Msg("UpdateCdkTypeSvc")
		}
	}()

	cdkId := util.MustToInt64(req.CdkId)
	err = s.dao.UpdateCdkType(cdkId, req.CdkName, req.CoinName, req.ExchangeRate)
	if err != nil {
		return nil, err
	}

	return &types.UpdateCdkTypeResp{}, nil
}
