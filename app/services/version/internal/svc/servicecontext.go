package svc

import (
	"fmt"

	"github.com/txchat/dtalk/app/services/version/internal/config"
	"github.com/txchat/dtalk/app/services/version/internal/dao"
	"github.com/txchat/dtalk/pkg/mysql"
)

type ServiceContext struct {
	Config config.Config
	Repo   dao.DeviceRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, err := mysql.NewMysqlConn(c.MySQL.Host, fmt.Sprintf("%v", c.MySQL.Port),
		c.MySQL.User, c.MySQL.Pwd, c.MySQL.DB, "UTF8MB4")
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		Repo:   dao.NewVersionRepositoryMysql(conn),
	}
}
