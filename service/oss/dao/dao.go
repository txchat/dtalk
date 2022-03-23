package dao

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog"
	"github.com/txchat/dtalk/pkg/logger"
	conf "github.com/txchat/dtalk/service/oss/config"
)

type Dao struct {
	log   zerolog.Logger
	redis *redis.Pool
}

func New(c *conf.Config) *Dao {
	d := &Dao{
		log:   logger.New(c.Env, "oss/dao"),
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
