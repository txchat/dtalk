package cdk

import (
	"time"

	"github.com/txchat/dtalk/service/backend/model/biz"
)

func (s *ServiceContent) CleanFrozenOrder() {
	for range time.Tick(10 * time.Minute) {
		// get frozen cdk
		cdks, err := s.getFrozenCdks()
		if err != nil {
			continue
		}

		nowTime := time.Now().UnixNano() / 1e6
		subTime := time.Now().Add(5*time.Minute).UnixNano()/1e6 - nowTime
		OrderSet := make(map[int64]bool, 0)
		for _, cdk := range cdks {
			if nowTime-cdk.UpdateTime > subTime {
				OrderSet[cdk.OrderId] = true
			}
		}

		for k := range OrderSet {
			err := s.dao.CleanFrozenCdks(k)
			if err != nil {
				s.log.Err(err).Int64("orderId", k).Msg("CleanFrozenCdks Error")
			}
		}
	}
}

func (s *ServiceContent) getFrozenCdks() ([]*biz.Cdk, error) {
	cdkPos, err := s.dao.GetFrozenCdks()
	if err != nil {
		return nil, err
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
			UpdateTime:   cdkPo.UpdateTime,
			ExchangeTime: cdkPo.ExchangeTime,
		}
		cdks = append(cdks, cdk)
	}

	return cdks, nil
}
