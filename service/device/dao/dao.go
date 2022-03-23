package dao

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/service/device/config"
)

type Dao struct {
	log   zerolog.Logger
	redis *redis.Pool
}

func New(c *config.Config) *Dao {
	d := &Dao{
		log:   zlog.Logger,
		redis: newRedis(c.Redis),
	}
	return d
}

func newRedis(c *config.Redis) *redis.Pool {
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
