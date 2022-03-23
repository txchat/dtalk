package cdk

import (
	"fmt"
	"github.com/pkg/errors"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/backend/model/biz"
	"github.com/txchat/dtalk/service/backend/model/types"
)

// ListenOrder 监听创建订单消息队列
func (s *ServiceContent) ListenOrder() {
	// 单线程模型, 保证发号正确
	for {
		select {
		case msg := <-s.CdkOrderMessage:

			msg.Done <- s.createOrder(msg)
		}
	}
}

// createOrder 处理创建订单消息
func (s *ServiceContent) createOrder(msg *biz.CdkOrderMessage) *biz.CdkOrder {
	s.log.Info().Str("PersonId", msg.PersonId).
		Int64("CdkId", msg.CdkId).
		Int64("Number", msg.Number).
		Msg("ListenOrder")

	// 获得 orderId
	orderId, err := s.idGenRPCClient.GetID()
	if err != nil {
		return &biz.CdkOrder{
			Err:     err,
			OrderId: 0,
		}
	}

	if msg.Number == 0 {
		return &biz.CdkOrder{
			Err:     errors.New("number is zero"),
			OrderId: 0,
		}
	}

	// 判断这个人这种优惠券的数量
	count, err := s.dao.GetCdksCountByUserIdAndCdkId(msg.CdkId, msg.PersonId)
	if err != nil {
		return &biz.CdkOrder{
			Err:     err,
			OrderId: 0,
		}
	}

	if count+msg.Number > s.CdkMaxNumber {
		return &biz.CdkOrder{
			Err:     xerror.NewError(xerror.CdkMaxNumberErr).SetExtMessage(fmt.Sprintf("已有 %d 张, 将兑换 %d 张, 超出 %d 张", count, msg.Number, s.CdkMaxNumber)),
			OrderId: 0,
		}
	}

	// 得到 n 个未发放的 cdk
	cdks, err := s.GetUnusedCdks(msg.CdkId, msg.Number)
	if err != nil {
		return &biz.CdkOrder{
			Err:     err,
			OrderId: 0,
		}
	}

	// 冻结 cdk
	err = s.FrozenCdks(cdks, msg.PersonId, orderId)
	if err != nil {
		return &biz.CdkOrder{
			Err:     err,
			OrderId: 0,
		}
	}

	// 加入定时清理订单任务队列

	return &biz.CdkOrder{
		Err:     nil,
		OrderId: orderId,
	}
}

// CreateCdkOrderSvc 创建订单 Svc
func (s *ServiceContent) CreateCdkOrderSvc(req *types.CreateCdkOrderReq) (resp *types.CreateCdkOrderResp, err error) {
	personId := req.PersonId
	cdkId := util.ToInt64(req.CdkId)
	number := req.Number
	done := make(chan *biz.CdkOrder)

	defer func() {
		if err != nil {
			s.log.Err(err).Interface("req", req).Str("personId", personId).Msg("CreateCdkOrderSvc")
		} else {
			s.log.Info().Interface("req", req).Str("personId", personId).Msg("CreateCdkOrderSvc")
		}
	}()

	s.CdkOrderMessage <- &biz.CdkOrderMessage{
		PersonId: personId,
		CdkId:    cdkId,
		Number:   number,
		Done:     done,
	}

	res := <-done
	if res.Err != nil {
		return nil, res.Err
	}

	return &types.CreateCdkOrderResp{
		OrderId: util.ToString(res.OrderId),
	}, nil
}

// GetUnusedCdks 按顺序获得未发放的 cdk
func (s *ServiceContent) GetUnusedCdks(cdkId int64, number int64) ([]*biz.Cdk, error) {
	cdkPos, err := s.dao.GetUnusedCdks(cdkId, number)
	if err != nil {
		return nil, err
	}

	if len(cdkPos) < int(number) {
		return nil, xerror.NewError(xerror.CdkOutOfStock)
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

// FrozenCdks 冻结 cdk
func (s *ServiceContent) FrozenCdks(cdks []*biz.Cdk, userId string, orderId int64) error {
	ids := make([]int64, 0, len(cdks))
	for _, cdk := range cdks {
		ids = append(ids, cdk.Id)
	}

	return s.dao.FrozenCdksStatus(ids, userId, orderId)
}

// TODO : 处理超时为完成支付的订单
