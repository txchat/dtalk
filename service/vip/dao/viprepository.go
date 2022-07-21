package dao

import "github.com/txchat/dtalk/service/vip/model"

type VIPRepository interface {
	GetVIP() (*model.VIPEntity, error)
	GetScopeVIP(start, limit int32) ([]*model.VIPEntity, error)
	GetVIPCount() (int32, error)
	AddVIP(*model.VIPEntity) error
}
