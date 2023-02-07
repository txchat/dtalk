package backenduser

import "github.com/txchat/dtalk/app/gateway/center/internal/config"

type UserManager struct {
	storage map[string]string
}

func NewUserManager(backendUsers []config.BackendUser) *UserManager {
	users := make(map[string]string)
	for _, user := range backendUsers {
		users[user.Username] = user.Password
	}
	return &UserManager{
		storage: users,
	}
}

func (um *UserManager) IsMatch(username, password string) bool {
	if pwd, ok := um.storage[username]; ok && pwd == password {
		return true
	}
	return false
}
