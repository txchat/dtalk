package model

import "sync"

// 弃用
type User struct {
	status map[string]bool
	sync.RWMutex
}

func NewUser() *User {
	user := &User{
		status: make(map[string]bool),
	}
	return user
}

func (u *User) GetStatus(key string) bool {
	u.RLock()
	defer u.RUnlock()
	val, ok := u.status[key]
	if !ok {
		return false
	}
	return val
}

func (u *User) SetStatus(key string, isBusy bool) error {
	u.Lock()
	defer u.Unlock()

	preStatus := u.status[key]
	if preStatus == BUSY && isBusy == BUSY {
		return ErrUserBusy
	}

	u.status[key] = isBusy
	return nil
}
