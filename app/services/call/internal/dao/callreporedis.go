package dao

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/txchat/dtalk/app/services/call/internal/model"
	xredis "github.com/txchat/dtalk/pkg/redis"
)

const (
	_prefixCallSession = "session:%d"
)

func keySession(traceId int64) string {
	return fmt.Sprintf(_prefixCallSession, traceId)
}

type CallRepositoryRedis struct {
	redis *redis.Pool
}

func NewCallRepositoryRedis(c xredis.Config) *CallRepositoryRedis {
	return &CallRepositoryRedis{
		redis: xredis.NewPool(c),
	}
}

// GetSession get session from traceId
func (repo *CallRepositoryRedis) GetSession(traceId int64) (*model.Session, error) {
	key := keySession(traceId)
	conn := repo.redis.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return nil, fmt.Errorf("conn.DO(EXISTS %s ) failed:%v", key, err)
	}
	if !ok {
		return nil, model.ErrSessionNotExist
	}

	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, fmt.Errorf("conn.DO(GET %s) failed:%v", key, err)
	}

	item := model.Session{}
	err = json.Unmarshal(data, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (repo *CallRepositoryRedis) SaveSession(session model.Session) error {
	key := keySession(session.TaskID)

	data, err := json.Marshal(session)
	if err != nil {
		return err
	}
	conn := repo.redis.Get()
	defer conn.Close()
	if err := conn.Send("SET", key, data, "EX", model.SessionTimeout); err != nil {
		return fmt.Errorf("conn.Send(SET %s,%s, EX, %s) failed:%v", key, data, model.SessionTimeout, err)
	}
	n := 1
	for i := 0; i < n; i++ {
		if _, err := conn.Receive(); err != nil {
			return fmt.Errorf("conn.Receive() failed:%v", err)
		}
	}
	return nil
}
