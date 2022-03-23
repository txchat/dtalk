package model

type GeneralResponse struct {
	Code    int         `json:"code"`
	Message int         `json:"message"`
	Data    interface{} `json:"data"`
}

type GetAddressRequest struct {
	Query string `json:"query" binding:"required"`
}

type GetAddressResponse struct {
	Address string `json:"address"`
}
