package cdk

import (
	"github.com/txchat/dtalk/service/backend/model/db"
	"github.com/txchat/dtalk/service/backend/model/types"
	"github.com/txchat/dtalk/pkg/util"
)

func (s *ServiceContent) ExchangeCdksSvc(req *types.ExchangeCdksReq) (res *types.ExchangeCdksResp, err error) {
	defer func() {
		if err != nil {
			s.log.Err(err).Interface("req", req).Msg("ExchangeCdksSvc")
		} else {
			s.log.Info().Interface("req", req).Msg("ExchangeCdksSvc")
		}
	}()

	ids := make([]int64, len(req.Ids), len(req.Ids))
	for i, id := range req.Ids {
		ids[i] = util.ToInt64(id)
	}

	err = s.dao.UpdateCdksStatus(ids, db.CdkExchange)
	if err != nil {
		return nil, err
	}

	res = &types.ExchangeCdksResp{}

	return res, nil
}
