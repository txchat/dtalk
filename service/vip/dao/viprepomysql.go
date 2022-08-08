package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/txchat/dtalk/service/vip/config"
	"github.com/txchat/dtalk/service/vip/model"
)

type VIPRepositoryMySQL struct {
	db *gorm.DB
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
		&model.VIPEntity{},
	)
	return db
}

func NewVIPRepositoryMySQL(env string, cfg *config.MySQL) *VIPRepositoryMySQL {
	repo := &VIPRepositoryMySQL{
		db: newDB(env, cfg),
	}
	return repo
}

func (repo *VIPRepositoryMySQL) GetTx() *gorm.DB {
	return repo.db.Begin()
}

func (repo *VIPRepositoryMySQL) CloseDB() {
	defer repo.db.Close()
}

func (repo *VIPRepositoryMySQL) GetVIP(uid string) (*model.VIPEntity, error) {
	var vip model.VIPEntity
	err := repo.db.First(&vip).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &vip, nil
}

func (repo *VIPRepositoryMySQL) GetScopeVIP(offset, limit int32) (vips []*model.VIPEntity, err error) {
	err = repo.db.Offset(offset).Limit(limit).Find(&vips).Error
	return
}

func (repo *VIPRepositoryMySQL) GetVIPCount() (count int32, err error) {
	err = repo.db.Model(&model.VIPEntity{}).Count(count).Error
	return
}

func (repo *VIPRepositoryMySQL) AddVIP(v *model.VIPEntity) error {
	return repo.db.Create(v).Error
}
