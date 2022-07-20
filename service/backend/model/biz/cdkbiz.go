package biz

import (
	"time"

	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/backend/model/db"
	"github.com/txchat/dtalk/service/backend/model/types"
)

// 业务模型

// CdkType cdk 种类
type CdkType struct {
	CdkId        int64  `json:"cdkId,omitempty"`
	CdkName      string `json:"cdkName,omitempty"`
	CoinName     string `json:"coinName,omitempty"`
	ExchangeRate int64  `json:"exchangeRate,omitempty"`
	CdkInfo      string `json:"cdkInfo,omitempty"`
	// 未发放的cdk数量
	CdkAvailable int64 `json:"cdkAvailable,omitempty"`
	// 已发放的cdk数量
	CdkUsed int64 `json:"cdkUsed,omitempty"`
	// 冻结状态中的cdk数量
	CdkFrozen int64 `json:"cdkFrozen,omitempty"`
}

func (cdkType *CdkType) ToTypes() *types.CdkType {
	return &types.CdkType{
		CdkId:        util.ToString(cdkType.CdkId),
		CdkName:      cdkType.CdkName,
		CoinName:     cdkType.CoinName,
		ExchangeRate: cdkType.ExchangeRate,
		CdkInfo:      cdkType.CdkInfo,
		CdkAvailable: cdkType.CdkAvailable,
		CdkUsed:      cdkType.CdkUsed,
		CdkFrozen:    cdkType.CdkFrozen,
	}
}

// Cdk cdk 实例
type Cdk struct {
	Id           int64  `json:"id,omitempty"`
	CdkId        int64  `json:"cdkId,omitempty"`
	CdkContent   string `json:"cdkContent,omitempty"`
	UserId       string `json:"userId,omitempty"`
	CdkStatus    int32  `json:"cdkStatus,omitempty"`
	OrderId      int64  `json:"orderId,omitempty"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64
	ExchangeTime int64 `json:"exchangeTime"`
}

func (cdk *Cdk) ToTypes() *types.Cdk {
	return &types.Cdk{
		Id:           util.ToString(cdk.Id),
		CdkId:        util.ToString(cdk.CdkId),
		CdkContent:   cdk.CdkContent,
		UserId:       cdk.UserId,
		CdkStatus:    cdk.CdkStatus,
		OrderId:      util.ToString(cdk.OrderId),
		CreateTime:   util.ToString(cdk.CreateTime),
		ExchangeTime: util.ToString(cdk.ExchangeTime),
	}
}

func (cdk *Cdk) CheckUserId(userId string) bool {
	return cdk.UserId == userId
}

func (cdk *Cdk) CheckFrozen() bool {
	return cdk.CdkStatus == db.CdkFrozen
}

type CdkOrderMessage struct {
	PersonId string
	CdkId    int64
	Number   int64
	Done     chan *CdkOrder
}

func NewCdkOrderMessage(personId string, cdkId, number int64) *CdkOrderMessage {
	return &CdkOrderMessage{
		PersonId: personId,
		CdkId:    cdkId,
		Number:   number,
		Done:     make(chan *CdkOrder),
	}
}

type CdkOrder struct {
	Err     error
	OrderId int64
}

type ClearFrozenOrderMessage struct {
	OrderId  int64
	Deadline time.Duration
}
