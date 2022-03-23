package cdk

import (
	"github.com/txchat/dtalk/service/backend/model/types"
	"github.com/txchat/dtalk/pkg/util"
)

func (s *ServiceContent) DeleteCdkTypesSvc(req *types.DeleteCdkTypesReq) (res *types.DeleteCdkTypesResp, err error) {
	defer func() {
		if err != nil {
			s.log.Err(err).Interface("req", req).Msg("DeleteCdkTypesSvc")
		} else {
			s.log.Info().Interface("req", req).Msg("DeleteCdkTypesSvc")
		}
	}()

	cdkIds := make([]int64, len(req.CdkIds), len(req.CdkIds))
	for i, id := range req.CdkIds {
		cdkIds[i] = util.ToInt64(id)
	}

	err = s.dao.DeleteCdkTypes(cdkIds)
	if err != nil {
		return nil, err
	}

	err = s.dao.DeleteCdksByCdkIds(cdkIds)
	if err != nil {
		return nil, err
	}

	res = &types.DeleteCdkTypesResp{}

	return res, nil
}
