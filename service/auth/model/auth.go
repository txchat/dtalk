package model

type AuthInfo struct {
	AppId      string
	Token      string
	Digest     string
	CreateTime int64
}
