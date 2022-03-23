package cdk

import (
	"github.com/txchat/dtalk/service/backend/model/biz"
	"github.com/txchat/dtalk/service/backend/model/types"
)

func (s *ServiceContent) GetCdksByUserIdSvc(req *types.GetCdksByUserIdReq) (res *types.GetCdksByUserIdResp, err error) {
	personId := req.PersonId
	page := req.Page
	pageSize := req.PageSize

	defer func() {
		if err != nil {
			s.log.Error().Err(err).Str("personId", personId).Msg("GetCdksByUserIdSvc")
		}
	}()

	cdks, totalElements, totalPages, err := s.GetCdksByUserId(personId, page, pageSize)
	if err != nil {
		return nil, err
	}

	cdkVos := make([]*types.Cdk, 0, len(cdks))
	cdkNameMap := make(map[int64]string, 0)
	for _, cdk := range cdks {
		tcdk := cdk.ToTypes()
		cdkName, ok := cdkNameMap[cdk.CdkId]
		if !ok {
			cdkType, err := s.GetCdkTypeByCdkId(cdk.CdkId)
			if err != nil {
				continue
			}
			cdkName = cdkType.CdkName
			cdkNameMap[cdk.CdkId] = cdkName
		}
		tcdk.CdkName = cdkName
		cdkVos = append(cdkVos, tcdk)
	}

	return &types.GetCdksByUserIdResp{
		TotalElements: totalElements,
		TotalPages:    totalPages,
		Cdks:          cdkVos,
	}, nil
}

func (s *ServiceContent) GetCdksByUserId(userId string, page, pageSize int64) ([]*biz.Cdk, int64, int64, error) {
	cdkPos, totalElements, totalPages, err := s.dao.GetCdksWithUserId(userId, page, pageSize)
	if err != nil {
		return nil, 0, 0, err
	}

	cdks := make([]*biz.Cdk, 0, len(cdkPos))
	for _, cdkPo := range cdkPos {
		cdk := &biz.Cdk{
			Id:           cdkPo.Id,
			CdkId:        cdkPo.CdkId,
			CdkContent:   cdkPo.CdkContent,
			UserId:       cdkPo.UserId,
			CdkStatus:    cdkPo.CdkStatus,
			OrderId:      cdkPo.OrderId,
			CreateTime:   cdkPo.CreateTime,
			ExchangeTime: cdkPo.ExchangeTime,
		}
		cdks = append(cdks, cdk)
	}

	return cdks, totalElements, totalPages, nil
}
