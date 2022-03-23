package cdk

import (
	chainTypes "github.com/33cn/chain33/types"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/backend/model/biz"
	"github.com/txchat/dtalk/service/backend/model/db"
	"github.com/txchat/dtalk/service/backend/model/types"
	"time"
)

func (s *ServiceContent) DealCdkOrderSvc(req *types.DealCdkOrderReq) (res *types.DealCdkOrderResp, err error) {
	personId := req.PersonId
	orderId := util.ToInt64(req.OrderId)
	//result := req.Result
	transferHash := req.TransferHash

	defer func() {
		if err != nil {
			s.log.Err(err).Interface("req", req).Str("personId", personId).Msg("DealCdkOrderSvc")
		} else {
			s.log.Info().Interface("req", req).Str("personId", personId).Msg("DealCdkOrderSvc")
		}
	}()

	// 根据 orderId 查询 cdks
	cdks, err := s.getCdksByOrderId(orderId)
	if err != nil {
		return nil, err
	}

	// 判断 personId 是否一致
	if cdks[0].CheckUserId(personId) == false {
		return nil, xerror.NewError(xerror.CdkOrderError)
	}

	// 判断 cdk 是否处于冻结状态
	if cdks[0].CheckFrozen() == false {
		return nil, xerror.NewError(xerror.CdkStatusNotFrozen)
	}

	// 异步处理订单
	ids := make([]int64, 0, len(cdks))
	for _, cdk := range cdks {
		ids = append(ids, cdk.Id)
	}
	go s.dealOrder(ids, transferHash, orderId)

	// 根据 result 判断处理流程

	//if result {
	//	err = s.dealSuccessCdkOrder(ids, req.TransferHash)
	//} else {
	//	err = s.dealFailedCdkOrder(ids)
	//}
	//if err != nil {
	//	return nil, err
	//}

	return &types.DealCdkOrderResp{}, nil
}

// dealOrder 查询交易 hash 情况并处理订单
func (s *ServiceContent) dealOrder(ids []int64, hash string, orderId int64) {
	i := 0
	maxTryTimes := 20
	for i = 0; i <= maxTryTimes; i++ {
		res, err := s.checkBlockTxResult(hash)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		if res.Success == true {
			err = s.dealSuccessCdkOrder(ids)
			if err != nil {
				s.log.Error().Err(err).Int64("orderId", orderId).Str("hash", hash).Msg("dealSuccessCdkOrder err")
			}
			return
		} else {
			break
		}
	}

	if i > maxTryTimes {
		s.log.Error().Int64("orderId", orderId).Str("hash", hash).Msg("transfer hash not exist err")
	} else {
		s.log.Error().Int64("orderId", orderId).Str("hash", hash).Msg("transfer hash failed")
	}

	err := s.dealFailedCdkOrder(ids)
	if err != nil {
		s.log.Error().Err(err).Int64("orderId", orderId).Str("hash", hash).Msg("dealFailedCdkOrder err")
	}
	return
}

func (s *ServiceContent) dealSuccessCdkOrder(ids []int64) error {
	//

	return s.dao.UpdateCdksStatus(ids, db.CdkUsed)
}

func (s *ServiceContent) dealFailedCdkOrder(ids []int64) error {
	return s.dao.UpdateCdksStatus(ids, db.CdkUnused)
}

func (s *ServiceContent) getCdksByOrderId(orderId int64) ([]*biz.Cdk, error) {
	cdkPos, err := s.dao.GetCdksByOrderId(orderId)
	if err != nil {
		return nil, err
	}

	if len(cdkPos) == 0 {
		return nil, xerror.NewError(xerror.CdkOrderError)
	}

	cdks := make([]*biz.Cdk, 0, len(cdkPos))
	for _, cdkPo := range cdkPos {
		cdk := &biz.Cdk{
			Id:         cdkPo.Id,
			CdkId:      cdkPo.CdkId,
			CdkContent: cdkPo.CdkContent,
			UserId:     cdkPo.UserId,
			CdkStatus:  cdkPo.CdkStatus,
			OrderId:    cdkPo.OrderId,
		}
		cdks = append(cdks, cdk)
	}

	return cdks, nil
}

func (s *ServiceContent) checkBlockTxResult(hash string) (biz.TxResult, error) {
	tx, err := s.chain33Client.GetRealTx(hash)
	if err != nil {
		return biz.TxResult{}, err
	}

	switch tx.Receipt.GetTy() {
	case chainTypes.ExecOk:
		to := ""
		amount := int64(0)
		symbol := ""
		if tx.Tx != nil {
			to = tx.Tx.To
		}
		if len(tx.Assets) > 0 {
			assert1 := tx.Assets[0]
			amount = assert1.Amount
			symbol = assert1.GetSymbol()
		}

		return biz.TxResult{
			Success: true,
			To:      to,
			Amount:  amount,
			Symbol:  symbol,
		}, nil
	case chainTypes.ExecErr:
	case chainTypes.ExecPack:
	}
	return biz.TxResult{Success: false}, nil
}
