package dao

import (
	"os"
	"testing"

	xmysql "github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/pkg/redis"
	"github.com/zeromicro/go-zero/core/service"
)

var (
	mysqlRootPassword string
	repo              *StorageRepository
)

func TestMain(m *testing.M) {
	mysqlRootPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
	repo = NewStorageRepository(service.TestMode, redis.Config{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
		Auth:    "",
		Active:  60000,
		Idle:    1024,
	}, xmysql.Config{
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		User:   "root",
		Passwd: mysqlRootPassword,
		DBName: "dtalk_record",
	})
	os.Exit(m.Run())
}
