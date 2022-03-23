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
	"github.com/txchat/dtalk/service/record/store/config"
)

var (
	testLog   zerolog.Logger
	testConn  *mysql.MysqlConn
	testRedis *redis.Pool

	testConnRecord *mysql.MysqlConn
)

func TestMain(m *testing.M) {
	testConn = newDB(&config.MySQL{
		Host: "127.0.0.1",
		Port: 3306,
		User: "root",
		Pwd:  "123456",
		Db:   "dtalk",
	})
	testConnRecord = newDB(&config.MySQL{
		Host: "127.0.0.1",
		Port: 3306,
		User: "root",
		Pwd:  "123456",
		Db:   "dtalk",
	})
	testLog = zlog.Logger
	testRedis = newRedis(&config.Redis{
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
	})
	os.Exit(m.Run())
}
