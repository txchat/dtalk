package cdk

import (
	"github.com/txchat/dtalk/service/backend/model/types"
	"github.com/txchat/dtalk/pkg/util"
	"time"
)

func (s *ServiceContent) DeleteCdksSvc(req *types.DeleteCdksReq) (res *types.DeleteCdksResp, err error) {
	defer func() {
		if err != nil {
			s.log.Err(err).Interface("req", req).Msg("DeleteCdksSvc")
		} else {
			s.log.Info().Interface("req", req).Msg("DeleteCdksSvc")
		}
	}()

	ids := make([]int64, len(req.Ids), len(req.Ids))
	for i, id := range req.Ids {
		ids[i] = util.ToInt64(id)
	}

	err = s.dao.DeleteCdks(ids, time.Now().UnixNano()/1e6)
	if err != nil {
		return nil, err
	}

	res = &types.DeleteCdksResp{}

	return res, nil
}
