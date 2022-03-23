package dao

import (
	"os"
	"testing"

	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/service/group/config"
)

var (
	testLog  log15.Logger
	testConn *mysql.MysqlConn
)

func TestMain(m *testing.M) {
	testConn = newDB(&config.MySQL{
		Host: "127.0.0.1",
		Port: 3306,
		User: "root",
		Pwd:  "123456",
		Db:   "dtalk",
	})
	testLog = log15.New("model", "group/test")
	os.Exit(m.Run())
}
