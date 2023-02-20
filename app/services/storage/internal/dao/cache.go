package dao

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

const (
	_prefixRecordFocus = "record-focus:%v"
)

func keyRecordFocus(mid string) string {
	return fmt.Sprintf(_prefixRecordFocus, mid)
}

//key:version; val:json
func (repo *StorageRepository) AddRecordFocus(uid string, mid string, time int64) error {
	key := keyRecordFocus(mid)
	conn := repo.redis.Get()
	defer conn.Close()
	if err := conn.Send("SADD", key, uid); err != nil {
		return err
	}
	if err := conn.Flush(); err != nil {
		return err
	}
	if _, err := conn.Receive(); err != nil {
		return err
	}
	return nil
}

func (repo *StorageRepository) GetRecordFocusNumber(mid string) (int32, error) {
	key := keyRecordFocus(mid)
	conn := repo.redis.Get()
	defer conn.Close()
	num, err := redis.Int(conn.Do("SCARD", key))
	if err != nil {
		return 0, err
	}
	return int32(num), nil
}
