// From https://gitlab.33.cn/proof/backend-micro/blob/dev/pkg/gredis/gredis.go

package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// Config Redis配置项
type Config struct {
	Network      string
	Addr         string
	Auth         string
	Active       int
	Idle         int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	Expire       time.Duration
}

func NewPool(c Config) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     c.Idle,
		MaxActive:   c.Active,
		IdleTimeout: c.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(c.Network, c.Addr,
				redis.DialConnectTimeout(c.DialTimeout),
				redis.DialReadTimeout(c.ReadTimeout),
				redis.DialWriteTimeout(c.WriteTimeout),
				redis.DialPassword(c.Auth),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}
