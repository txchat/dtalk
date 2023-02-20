package dao

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	xmysql "github.com/txchat/dtalk/pkg/mysql"
	xredis "github.com/txchat/dtalk/pkg/redis"
)

type StorageRepository struct {
	redis *redis.Pool
	mysql *xmysql.MysqlConn
}

func NewStorageRepository(redisConfig xredis.Config, mysqlConfig xmysql.Config) *StorageRepository {
	mysqlConn, err := xmysql.NewMysqlConn(mysqlConfig.Host, fmt.Sprintf("%v", mysqlConfig.Port),
		mysqlConfig.User, mysqlConfig.Pwd, mysqlConfig.DB, "UTF8MB4")
	if err != nil {
		panic(err)
	}
	return &StorageRepository{
		redis: xredis.NewPool(redisConfig),
		mysql: mysqlConn,
	}
}

func (repo *StorageRepository) NewTx() (*xmysql.MysqlTx, error) {
	return repo.mysql.NewTx()
}
