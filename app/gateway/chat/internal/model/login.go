package model

type LoginNotAllowedErr struct {
	Code    int64 `json:"code"`
	Message struct {
		Datetime   uint64 `json:"datetime"`
		Device     int32  `json:"device"`
		DeviceName string `json:"deviceName"`
		Uuid       string `json:"uuid"`
	} `json:"message"`
	Service string `json:"service"`
}
