package dao

import (
	"fmt"

	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/service/auth/config"
)

type Dao struct {
	log  log15.Logger
	conn *mysql.MysqlConn
}

func New(c *config.Config) *Dao {
	d := &Dao{
		log:  log15.New("module", "auth/dao"),
		conn: newDB(c.MySQL),
	}
	return d
}

func newDB(cfg *config.MySQL) *mysql.MysqlConn {
	c, err := mysql.NewMysqlConn(cfg.Host, fmt.Sprintf("%v", cfg.Port),
		cfg.User, cfg.Pwd, cfg.Db, "UTF8MB4")
	if err != nil {
		log15.Error("mysql init failed", "err", err)
		panic(err)
	}
	return c
}
