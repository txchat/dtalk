package cdk

import (
	"time"

	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/backend/model/db"
	"github.com/txchat/dtalk/service/backend/model/types"
)

func (s *ServiceContent) CreateCdksSvc(req *types.CreateCdksReq) (res *types.CreateCdksResp, err error) {
	defer func() {
		if err != nil {
			s.log.Err(err).Interface("req", req).Msg("CreateCdksSvc")
		} else {
			s.log.Info().Interface("req", req).Msg("CreateCdksSvc")
		}
	}()

	cdkId := util.ToInt64(req.CdkId)

	// 判断 cdkId 是否存在
	_, err = s.GetCdkTypeByCdkId(cdkId)
	if err != nil {
		return nil, err
	}

	for _, c := range req.CdkContents {
		ok, err := s.dao.CheckCdkExist(cdkId, c)
		if err != nil || ok {
			continue
		}

		id, err := s.idGenRPCClient.GetID()
		if err != nil {
			continue
		}
		cdk := &db.Cdk{
			Id:         id,
			CdkId:      cdkId,
			CdkContent: c,
			UserId:     "",
			CdkStatus:  0,
			OrderId:    0,
			TimeInfo: db.TimeInfo{
				CreateTime: time.Now().UnixNano() / 1e6,
				UpdateTime: time.Now().UnixNano() / 1e6,
				DeleteTime: 0,
			},
			ExchangeTime: 0,
		}
		err = s.dao.InsertCdk(cdk)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
