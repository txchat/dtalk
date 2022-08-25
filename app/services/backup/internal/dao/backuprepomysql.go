package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/txchat/dtalk/app/services/backup/internal/model"
	xmysql "github.com/txchat/dtalk/pkg/mysql"
	"github.com/zeromicro/go-zero/core/service"
)

func NewDefaultConn(mode string, cfg xmysql.Config) *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Pwd,
		cfg.Host,
		cfg.Port,
		cfg.DB))
	if err != nil {
		panic(err)
	}
	if mode != service.ProMode {
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

type BackupRepositoryMysql struct {
	conn *gorm.DB
}

func NewBackupRepositoryMysql(conn *gorm.DB) *BackupRepositoryMysql {
	return &BackupRepositoryMysql{
		conn: conn,
	}
}

func (repo *BackupRepositoryMysql) QueryBind(queryType int32, queryCase string) (*model.AddrBackup, error) {
	var record model.AddrBackup
	sqlCase := ""
	switch queryType {
	case model.Phone:
		sqlCase = "phone = ?"
	case model.Email:
		sqlCase = "email = ?"
	case model.Address:
		sqlCase = "address = ?"
	default:
		return nil, model.ErrQueryType
	}
	err := repo.conn.Where(sqlCase, queryCase).First(&record).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &record, nil
}

func (repo *BackupRepositoryMysql) QueryRelate(queryType int32, queryCase string) (*model.AddrRelate, error) {
	var record model.AddrRelate
	sqlCase := ""
	switch queryType {
	case model.Phone:
		sqlCase = "phone = ?"
	case model.Email:
		sqlCase = "email = ?"
	case model.Address:
		sqlCase = "address = ?"
	default:
		return nil, model.ErrQueryType
	}
	err := repo.conn.Where(sqlCase, queryCase).First(&record).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &record, nil
}
