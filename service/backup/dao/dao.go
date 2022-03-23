package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/inconshreveable/log15"
	"github.com/jinzhu/gorm"
	"github.com/txchat/dtalk/service/backup/config"
	"github.com/txchat/dtalk/service/backup/model"
)

type Dao struct {
	log log15.Logger
	db  *gorm.DB
}

func New(c *config.Config) *Dao {
	d := &Dao{
		log: log15.New("module", "backup/dao"),
		db:  newDB(c.Env, c.MySQL),
	}
	return d
}

func newDB(env string, cfg *config.MySQL) *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Pwd,
		cfg.Host,
		cfg.Port,
		cfg.Db))
	if err != nil {
		panic(err)
	}
	if env == model.Debug {
		db.LogMode(true)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		//return setting.DatabaseSetting.TablePrefix + defaultTableName
		return defaultTableName
	}

	db.SingularTable(true)
	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.Set("gorm:table_options",
		"ENGINE=InnoDB AUTO_INCREMENT=1 CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci").AutoMigrate(
		&model.AddrBackup{}, &model.AddrRelate{}, &model.AddrMove{},
	)
	return db
}

func (d *Dao) GetTx() *gorm.DB {
	return d.db.Begin()
}

func (d *Dao) CloseDB() {
	defer d.db.Close()
}
