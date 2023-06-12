package dao

import (
	"github.com/gomodule/redigo/redis"
	xmysql "github.com/txchat/dtalk/pkg/mysql"
	xredis "github.com/txchat/dtalk/pkg/redis"
	"github.com/zeromicro/go-zero/core/service"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type StorageRepository struct {
	redis *redis.Pool
	db    *gorm.DB
}

func NewStorageRepository(mode string, redisConfig xredis.Config, mysqlConfig xmysql.Config) *StorageRepository {
	mysqlConfig.ParseTime = true
	mysqlConfig.SetParam("charset", "UTF8MB4")

	defaultLogger := logger.Default
	switch mode {
	case service.TestMode, service.DevMode, service.RtMode:
		defaultLogger.LogMode(logger.Info)
	case service.ProMode, service.PreMode:
		defaultLogger.LogMode(logger.Warn)
	}

	dsn := mysqlConfig.GetSQLDriverConfig().FormatDSN()
	db, err := gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &StorageRepository{
		redis: xredis.NewPool(redisConfig),
		db:    db,
	}
}
