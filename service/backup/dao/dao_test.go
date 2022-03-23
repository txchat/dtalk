package dao

import (
	"os"
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/inconshreveable/log15"
	"github.com/jinzhu/gorm"
	"github.com/txchat/dtalk/service/backup/config"
)

var (
	testLog   log15.Logger
	testConn  *gorm.DB
	testRedis *redis.Pool
)

func TestMain(m *testing.M) {
	testConn = newDB("debug", &config.MySQL{
		Host: "172.16.101.107",
		Port: 3306,
		User: "root",
		Pwd:  "123456",
		Db:   "dtalk",
	})
	testLog = log15.New("model", "test")
	os.Exit(m.Run())
}
