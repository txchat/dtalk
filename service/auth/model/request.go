package model

import "mime/multipart"

type SignInRequest struct {
	AppId      string                `json:"appId" form:"appId"`
	ConfigFile *multipart.FileHeader `json:"-" form:"-"`
}

type AuthRequest struct {
	AppId string `json:"appId"`
	Token string `json:"token"`
}
