package model

type GeneralResponse struct {
	Result  int         `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
