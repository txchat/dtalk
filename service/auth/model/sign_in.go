package model

import "mime/multipart"

type SignInInfo struct {
	AppId      string
	ConfigFile *multipart.FileHeader
	CreateTime int64
	UpdateTime int64
}
