package dao

import (
	"os"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	xtime "github.com/txchat/dtalk/pkg/time"
	"github.com/txchat/dtalk/service/oss/config"
)

var (
	testLog   zerolog.Logger
	testRedis *redis.Pool
)

func TestMain(m *testing.M) {
	testLog = zlog.Logger
	c := &config.Redis{
		Network:      "tcp",
		Addr:         "127.0.0.1:6379",
		Auth:         "",
		Active:       60000,
		Idle:         1024,
		DialTimeout:  xtime.Duration(200 * time.Millisecond),
		ReadTimeout:  xtime.Duration(500 * time.Millisecond),
		WriteTimeout: xtime.Duration(500 * time.Millisecond),
		IdleTimeout:  xtime.Duration(120 * time.Second),
		Expire:       xtime.Duration(30 * time.Minute),
	}

	testRedis = &redis.Pool{
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
	os.Exit(m.Run())
}
