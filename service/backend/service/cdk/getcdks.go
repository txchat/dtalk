package cdk

import (
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/backend/model/types"
)

func (s *ServiceContent) GetCdksSvc(req *types.GetCdksReq) (res *types.GetCdksResp, err error) {
	defer func() {
		if err != nil {
			s.log.Err(err).Interface("req", req).Msg("GetCdksSvc")
		} else {
			s.log.Info().Interface("req", req).Msg("GetCdksSvc")
		}
	}()

	cdks, totalElements, totalPages, err := s.GetCdks(util.MustToInt64(req.CdkId), req.CdkContent, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	typesCdks := make([]*types.Cdk, len(cdks), len(cdks))
	for i, cdk := range cdks {
		typesCdk := cdk.ToTypes()
		typesCdks[i] = typesCdk
	}
	res = &types.GetCdksResp{
		TotalElements: totalElements,
		TotalPages:    totalPages,
		Cdks:          typesCdks,
	}

	return res, nil
}
