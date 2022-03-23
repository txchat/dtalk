package db

// 数据库模型
// 持久化模型 PO

type TimeInfo struct {
	CreateTime int64
	UpdateTime int64
	DeleteTime int64
}

type CdkType struct {
	CdkId        int64
	CdkName      string
	CdkInfo      string
	CoinName     string
	ExchangeRate int64
	TimeInfo
}

type Cdk struct {
	Id         int64
	CdkId      int64
	CdkContent string
	UserId     string
	CdkStatus  int32
	OrderId    int64
	TimeInfo
	ExchangeTime int64
}

type CdkOrder struct {
	Id           int64
	OrderId      int64
	TransferHash string
	UserId       string
	OrderStatus  int32
	TimeInfo
	ExchangeTime int64
}

const (
	CdkUnused   = 0
	CdkFrozen   = 1
	CdkUsed     = 2
	CdkExchange = 3
)
