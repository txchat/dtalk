package dao

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	xmysql "github.com/txchat/dtalk/pkg/mysql"
	xredis "github.com/txchat/dtalk/pkg/redis"
)

type UniRepository struct {
	redis *redis.Pool
	mysql *xmysql.MysqlConn
}

func NewUniRepository(redisConfig xredis.Config, mysqlConfig xmysql.Config) *UniRepository {
	mysqlConn, err := xmysql.NewMysqlConn(mysqlConfig.Host, fmt.Sprintf("%v", mysqlConfig.Port),
		mysqlConfig.User, mysqlConfig.Pwd, mysqlConfig.DB, "UTF8MB4")
	if err != nil {
		panic(err)
	}
	return &UniRepository{
		redis: xredis.NewPool(redisConfig),
		mysql: mysqlConn,
	}
}

func (repo *UniRepository) NewTx() (*xmysql.MysqlTx, error) {
	return repo.mysql.NewTx()
}
