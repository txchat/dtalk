package types

type GeneralResponse struct {
	Result  int         `json:"result"`
	Message int         `json:"message"`
	Data    interface{} `json:"data"`
}

// http 请求和返回

// CdkType cdk 种类
type CdkType struct {
	CdkId        string `json:"cdkId"`
	CdkName      string `json:"cdkName"`
	CoinName     string `json:"coinName"`
	ExchangeRate int64  `json:"exchangeRate"`
	CdkInfo      string `json:"cdkInfo"`
	// 未发放的cdk数量
	CdkAvailable int64 `json:"cdkAvailable"`
	// 已发放的cdk数量
	CdkUsed int64 `json:"cdkUsed"`
	// 冻结状态中的cdk数量
	CdkFrozen int64 `json:"cdkFrozen"`
}

// Cdk cdk 实例
type Cdk struct {
	Id           string `json:"id"`
	CdkId        string `json:"cdkId"`
	CdkName      string `json:"cdkName"`
	CdkContent   string `json:"cdkContent"`
	UserId       string `json:"userId"`
	CdkStatus    int32  `json:"cdkStatus"`
	OrderId      string `json:"orderId"`
	CreateTime   string `json:"createTime"`
	ExchangeTime string `json:"exchangeTime"`
}

// PageInfo 分页信息
type PageInfo struct {
	// 页数
	Page int64 `json:"page"`
	// 每页数量
	PageSize int64 `json:"pageSize"`
}

// -----------------------------backend-----------------------------

// CreateCdkTypeReq 创建 cdk 种类请求
type CreateCdkTypeReq struct {
	CdkName      string `json:"cdkName"`
	CoinName     string `json:"coinName"  binding:"required"`
	ExchangeRate int64  `json:"exchangeRate"  binding:"required"`
	CdkInfo      string `json:"cdkInfo"`
}

// CreateCdkTypeResp 创建 cdk 种类响应
type CreateCdkTypeResp struct {
	CdkId string `json:"cdkId"`
}

// GetCdkTypesReq 查询 cdk 种类请求
type GetCdkTypesReq struct {
	PageInfo
	CoinName string `json:"coinName"`
}

// GetCdkTypesResp 查询 cdk 种类响应
type GetCdkTypesResp struct {
	TotalElements int64      `json:"totalElements"`
	TotalPages    int64      `json:"totalPages"`
	CdkTypes      []*CdkType `json:"cdkTypes"`
}

// CreateCdksReq 批量上传 cdk 请求
type CreateCdksReq struct {
	CdkId       string   `json:"cdkId"  binding:"required"`
	CdkContents []string `json:"cdkContents"  binding:"required"`
}

// CreateCdksResp 批量上传 cdk 响应
type CreateCdksResp struct {
}

// GetCdksReq 查询 cdk 实例请求
type GetCdksReq struct {
	PageInfo
	CdkId      string `json:"cdkId"  binding:"required"`
	CdkContent string `json:"cdkContent"`
}

// GetCdksResp 查询 cdk 实例响应
type GetCdksResp struct {
	TotalElements int64  `json:"totalElements"`
	TotalPages    int64  `json:"totalPages"`
	Cdks          []*Cdk `json:"cdks"`
}

// GetCdksWithCdkNameReq 根据 cdkName 查询 cdk 实例请求
//type GetCdksWithCdkNameReq struct {
//	PageInfo
//	CdkName string `json:"cdkName"`
//}

// GetCdksWithCdkNameResp 根据 cdkName 查询 cdk 实例响应
//type GetCdksWithCdkNameResp struct {
//	TotalElements int64  `json:"totalElements"`
//	TotalPages    int64  `json:"totalPages"`
//	Cdks          []*Cdk `json:"cdks"`
//}

// DeleteCdksReq 删除 cdk 实例请求
type DeleteCdksReq struct {
	Ids []string `json:"ids"  binding:"required"`
}

// DeleteCdksResp 删除 cdk 实例响应
type DeleteCdksResp struct {
}

// DeleteCdkTypesReq 删除 cdkType 请求
type DeleteCdkTypesReq struct {
	CdkIds []string `json:"cdkIds"  binding:"required"`
}

// DeleteCdkTypesResp 删除 cdkType 响应
type DeleteCdkTypesResp struct {
}

type UpdateCdkTypeReq struct {
	CdkId        string `json:"cdkId" binding:"required"`
	CdkName      string `json:"cdkName" binding:"required"`
	CoinName     string `json:"coinName" binding:"required"`
	ExchangeRate int64  `json:"exchangeRate" binding:"required"`
}

type UpdateCdkTypeResp struct {
}

type ExchangeCdksReq struct {
	Ids []string `json:"ids"  binding:"required"`
}

type ExchangeCdksResp struct {
}

// -----------------------------app-----------------------------

// GetCdksByUserIdReq 查询某人拥有的 cdk 实例请求
type GetCdksByUserIdReq struct {
	PageInfo
	//CdkId string `json:"cdkId" binding:"required"`
	PersonId string `json:"-"`
}

// GetCdksByUserIdResp 查询某人拥有的 cdk 实例响应
type GetCdksByUserIdResp struct {
	TotalElements int64  `json:"totalElements"`
	TotalPages    int64  `json:"totalPages"`
	Cdks          []*Cdk `json:"cdks"`
}

// GetCdkTypeByCoinNameReq 根据票券名称查询对应的 cdk 信息请求
type GetCdkTypeByCoinNameReq struct {
	CoinName string `json:"coinName"  binding:"required"`
}

// GetCdkTypeByCoinNameResp 根据票券名称查询对应的 cdk 信息响应
type GetCdkTypeByCoinNameResp struct {
	*CdkType
}

// CreateCdkOrderReq 创建兑换券订单请求
type CreateCdkOrderReq struct {
	PersonId string `json:"-"`
	// cdk 种类编号
	CdkId string `json:"cdkId"  binding:"required"`
	// 兑换数量
	Number int64 `json:"number"  binding:"required"`
}

// CreateCdkOrderResp 创建兑换券订单响应
type CreateCdkOrderResp struct {
	// 订单编号
	OrderId string `json:"orderId"`
}

// DealCdkOrderReq 处理兑换券订单请求
type DealCdkOrderReq struct {
	PersonId string `json:"-"`
	// 订单编号
	OrderId string `json:"orderId"  binding:"required"`
	// 处理结果
	Result bool `json:"result"`
	// 转账记录 hash
	TransferHash string `json:"transferHash" binding:"required"`
}

// DealCdkOrderResp 处理兑换券订单响应
type DealCdkOrderResp struct {
}
