package auth

type Reply struct {
	Result  int                    `json:"result"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type ErrorDataReconnectNotAllowed struct {
	Code    int `json:"code"`
	Message struct {
		Datetime   int64  `json:"datetime"`
		Device     int    `json:"device"`
		DeviceName string `json:"deviceName"`
		UUID       string `json:"uuid"`
	} `json:"message"`
	Service string `json:"service"`
}

type SuccessData struct {
	Address string `json:"address"`
}
