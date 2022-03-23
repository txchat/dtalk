package cdk

import "github.com/txchat/dtalk/service/backend/model/types"

func (s *ServiceContent) GetCdkTypeByCoinNameSvc(req *types.GetCdkTypeByCoinNameReq) (res *types.GetCdkTypeByCoinNameResp, err error) {
	coinName := req.CoinName

	defer func() {
		if err != nil {
			s.log.Error().Err(err).Str("coinName", coinName).Msg("GetCdkTypeByCoinNameSvc")
		}
	}()

	cdkType, err := s.GetCdkTypeByCoinName(coinName)
	if err != nil {
		return nil, err
	}

	return &types.GetCdkTypeByCoinNameResp{
		CdkType: cdkType.ToTypes(),
	}, nil
}
