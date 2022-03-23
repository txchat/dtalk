package model

// 用户地址登录请求结果
type AddressLoginResp struct {
	// 用户地址
	Address string `json:"address" example:"123"`
}

type AddressLoginReq struct {
	ConnType int32 `json:"connType" example:"0"`
}

type AddressLoginNotAllowedErr struct {
	UUid       string `json:"uuid"`
	Device     int32  `json:"device"`
	DeviceName string `json:"deviceName"`
	Datetime   uint64 `json:"datetime"`
}
