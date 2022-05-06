package dao

import (
	"os"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/pkg/mysql"
	xtime "github.com/txchat/dtalk/pkg/time"
	"github.com/txchat/dtalk/service/device/config"
)

var (
	testLog   zerolog.Logger
	testConn  *mysql.MysqlConn
	testRedis *redis.Pool
)

func TestMain(m *testing.M) {
	testLog = zlog.Logger
	testRedis = newRedis(&config.Redis{
		Network:      "tcp",
		Addr:         "172.16.101.127:6379",
		Auth:         "",
		Active:       60000,
		Idle:         1024,
		DialTimeout:  xtime.Duration(200 * time.Millisecond),
		ReadTimeout:  xtime.Duration(500 * time.Millisecond),
		WriteTimeout: xtime.Duration(500 * time.Millisecond),
		IdleTimeout:  xtime.Duration(120 * time.Second),
		Expire:       xtime.Duration(30 * time.Minute),
	})
	os.Exit(m.Run())
}
