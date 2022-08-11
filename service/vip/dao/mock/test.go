package mock

import (
	"time"

	"github.com/txchat/dtalk/service/vip/model"
)

var users = map[string]bool{
	"12quUKnXMaHfxYvUB9bePW3k4eSj6H4ADo": true,
	"1EwiqdTK68Wgp5geDhqRf9ocrhBg9dJCqX": true,
	"1cMZ4qdSn9erZVmYf2wtsWxKVU6ZyUkSD":  true,
}

type AllowMockUsers struct {
}

func NewAllowMockUsers() *AllowMockUsers {
	return &AllowMockUsers{}
}

func (po *AllowMockUsers) GetVIP(uid string) (*model.VIPEntity, error) {
	if b, ok := users[uid]; b && ok {
		return &model.VIPEntity{
			UID:        uid,
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
