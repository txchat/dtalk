package mock

import (
	"time"

	"github.com/txchat/dtalk/service/vip/model"
)

type AllowMockUsers struct {
}

func NewAllowMockUsers() *AllowMockUsers {
	return &AllowMockUsers{}
}

func (po *AllowMockUsers) GetVIP(uid string) (*model.VIPEntity, error) {
	if uid == "12quUKnXMaHfxYvUB9bePW3k4eSj6H4ADo" {
		return &model.VIPEntity{
			UID:        "12quUKnXMaHfxYvUB9bePW3k4eSj6H4ADo",
			UpdateTime: time.Time{},
			CreateTime: time.Time{},
		}, nil
	}
	return nil, nil
}

func (po *AllowMockUsers) GetScopeVIP(start, limit int32) ([]*model.VIPEntity, error) {
	return nil, nil
}

func (po *AllowMockUsers) GetVIPCount() (int32, error) {
	return 0, nil
}

func (po *AllowMockUsers) AddVIP(*model.VIPEntity) error {
	return nil
}
