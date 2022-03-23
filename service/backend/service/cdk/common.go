package cdk

import (
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/backend/model/biz"
)

// GetCdkTypeByCoinName 通过 coinName 查询 cdkType
func (s *ServiceContent) GetCdkTypeByCoinName(coinName string) (*biz.CdkType, error) {
	// 通过 coinName 查询 cdkType 信息
	cdkTypePos, err := s.dao.GetCdkTypesWithCoinName(coinName)
	if err != nil {
		return nil, err
	}
	if len(cdkTypePos) == 0 {
		return nil, xerror.NewError(xerror.CdkCoinNameErr)
	}
	cdkTypePo := cdkTypePos[0]

	// 通过 cdkId 查询各种状态的 cdk 数量
	unused, frozen, used, err := s.dao.GetCdksCount(cdkTypePo.CdkId)
	if err != nil {
		return nil, err
	}

	return &biz.CdkType{
		CdkId:        cdkTypePo.CdkId,
		CdkName:      cdkTypePo.CdkName,
		CoinName:     cdkTypePo.CoinName,
		ExchangeRate: cdkTypePo.ExchangeRate,
		CdkInfo:      cdkTypePo.CdkInfo,
		CdkAvailable: unused,
		CdkUsed:      used,
		CdkFrozen:    frozen,
	}, nil
}

// GetCdkTypes 分页查询 CdkType 列表
func (s *ServiceContent) GetCdkTypes(coinName string, page, pageSize int64) ([]*biz.CdkType, int64, int64, error) {
	cdkTypePos, totalElements, totalPages, err := s.dao.GetCdkTypes(coinName, page, pageSize)
	if err != nil {
		return nil, 0, 0, err
	}

	cdkTypes := make([]*biz.CdkType, 0, len(cdkTypePos))
	for _, cdkTypePo := range cdkTypePos {
		// 通过 cdkId 查询各种状态的 cdk 数量
		unused, frozen, used, err := s.dao.GetCdksCount(cdkTypePo.CdkId)
		if err != nil {
			return nil, 0, 0, err
		}

		cdkType := &biz.CdkType{
			CdkId:        cdkTypePo.CdkId,
			CdkName:      cdkTypePo.CdkName,
			CoinName:     cdkTypePo.CoinName,
			ExchangeRate: cdkTypePo.ExchangeRate,
			CdkInfo:      cdkTypePo.CdkInfo,
			CdkAvailable: unused,
			CdkUsed:      used,
			CdkFrozen:    frozen,
		}

		cdkTypes = append(cdkTypes, cdkType)
	}
	return cdkTypes, totalElements, totalPages, nil
}

// GetCdks 分页查询 cdk 列表
func (s *ServiceContent) GetCdks(cdkId int64, cdkContent string, page, pageSize int64) ([]*biz.Cdk, int64, int64, error) {
	cdkPos, totalElements, totalPages, err := s.dao.GetCdks(cdkId, cdkContent, page, pageSize)
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

// GetCdkTypeByCdkId 通过 cdkId 查询 cdkType
func (s *ServiceContent) GetCdkTypeByCdkId(cdkId int64) (*biz.CdkType, error) {
	cdkTypePo, err := s.dao.GetCdkType(cdkId)
	if err != nil {
		return nil, err
	}

	if cdkTypePo == nil {
		return nil, xerror.NewError(xerror.CodeInnerError)
	}

	unused, frozen, used, err := s.dao.GetCdksCount(cdkTypePo.CdkId)
	if err != nil {
		return nil, err
	}

	cdkType := &biz.CdkType{
		CdkId:        cdkTypePo.CdkId,
		CdkName:      cdkTypePo.CdkName,
		CoinName:     cdkTypePo.CoinName,
		ExchangeRate: cdkTypePo.ExchangeRate,
		CdkInfo:      cdkTypePo.CdkInfo,
		CdkAvailable: unused,
		CdkUsed:      used,
		CdkFrozen:    frozen,
	}

	return cdkType, nil
}

// GetCdksWithCdkName 根据cdkName分页查询 cdk 列表
//func (s *ServiceContent) GetCdksWithCdkName(cdkName string, page, pageSize int64) ([]*biz.Cdk, int64, int64, error) {
//	cdkPos, totalElements, totalPages, err := s.dao.GetCdksByCdkName(cdkName, page, pageSize)
//	if err != nil {
//		return nil, 0, 0, err
//	}
//
//	cdks := make([]*biz.Cdk, 0, len(cdkPos))
//	for _, cdkPo := range cdkPos {
//		cdk := &biz.Cdk{
//			Id:           cdkPo.Id,
//			CdkId:        cdkPo.CdkId,
//			CdkContent:   cdkPo.CdkContent,
//			UserId:       cdkPo.UserId,
//			CdkStatus:    cdkPo.CdkStatus,
//			OrderId:      cdkPo.OrderId,
//			CreateTime:   cdkPo.CreateTime,
//			ExchangeTime: cdkPo.ExchangeTime,
//		}
//		cdks = append(cdks, cdk)
//	}
//
//	return cdks, totalElements, totalPages, nil
//}
