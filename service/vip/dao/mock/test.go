package mock

import (
	"time"

	"github.com/txchat/dtalk/service/vip/model"
)

type AllowMockUsers struct {
	users map[string]bool
}

func NewAllowMockUsers(wls []string) *AllowMockUsers {
	users := make(map[string]bool)
	for _, u := range wls {
		users[u] = true
	}
	return &AllowMockUsers{
		users: users,
	}
}

func (po *AllowMockUsers) GetVIP(uid string) (*model.VIPEntity, error) {
	if b, ok := po.users[uid]; b && ok {
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
