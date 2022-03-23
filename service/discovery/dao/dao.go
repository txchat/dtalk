package dao

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/inconshreveable/log15"
	conf "github.com/txchat/dtalk/service/discovery/config"
)

type Dao struct {
	log   log15.Logger
	redis *redis.Pool
}

func New(c *conf.Config) *Dao {
	d := &Dao{
		log:   log15.New("module", "discovery/dao"),
		redis: newRedis(c.Redis),
	}
	return d
}

func newRedis(c *conf.Redis) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     c.Idle,
		MaxActive:   c.Active,
		IdleTimeout: time.Duration(c.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(c.Network, c.Addr,
				redis.DialConnectTimeout(time.Duration(c.DialTimeout)),
				redis.DialReadTimeout(time.Duration(c.ReadTimeout)),
				redis.DialWriteTimeout(time.Duration(c.WriteTimeout)),
				redis.DialPassword(c.Auth),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}
