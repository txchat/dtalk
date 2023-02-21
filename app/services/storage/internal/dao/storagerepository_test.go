package dao

import (
	"os"
	"testing"

	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/pkg/redis"
)

var (
	mysqlRootPassword string
	repo              *StorageRepository
)

func TestMain(m *testing.M) {
	mysqlRootPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
	repo = NewStorageRepository(redis.Config{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
		Auth:    "",
		Active:  60000,
		Idle:    1024,
	}, mysql.Config{
		Host: "127.0.0.1",
		Port: 3306,
		User: "root",
		Pwd:  mysqlRootPassword,
		DB:   "dtalk_record",
	})
	os.Exit(m.Run())
}
