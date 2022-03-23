package model

type GeneralResponse struct {
	Code    int         `json:"code"`
	Message int         `json:"message"`
	Data    interface{} `json:"data"`
}

type SignInResponse struct {
	AppId      string `json:"appId"`
	Key        string `json:"key"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

type AuthResponse struct {
	Uid string `json:"uid"`
}
