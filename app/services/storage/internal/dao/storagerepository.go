package dao

import (
	"github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	xredis "github.com/txchat/dtalk/pkg/redis"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type StorageRepository struct {
	redis *redis.Pool
	db    *gorm.DB
}

func NewStorageRepository(redisConfig xredis.Config, mysqlConfig mysql.Config) *StorageRepository {
	if mysqlConfig.Params == nil {
		mysqlConfig.Params = make(map[string]string)
	}
	mysqlConfig.Params["charset"] = "UTF8MB4"
	dsn := mysqlConfig.FormatDSN()
	db, err := gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &StorageRepository{
		redis: xredis.NewPool(redisConfig),
		db:    db,
	}
}
