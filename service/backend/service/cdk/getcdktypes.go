package cdk

import "github.com/txchat/dtalk/service/backend/model/types"

func (s *ServiceContent) GetCdkTypesSvc(req *types.GetCdkTypesReq) (res *types.GetCdkTypesResp, err error) {
	defer func() {
		if err != nil {
			s.log.Err(err).Interface("req", req).Msg("GetCdkTypesSvc")
		} else {
			s.log.Info().Interface("req", req).Msg("GetCdkTypesSvc")
		}
	}()

	cdkTypes, totalElements, totalPages, err := s.GetCdkTypes(req.CoinName, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	typesCdks := make([]*types.CdkType, len(cdkTypes), len(cdkTypes))
	for i, cdkType := range cdkTypes {
		typesCdkType := cdkType.ToTypes()
		typesCdks[i] = typesCdkType
	}
	res = &types.GetCdkTypesResp{
		TotalElements: totalElements,
		TotalPages:    totalPages,
		CdkTypes:      typesCdks,
	}

	return res, nil
}
