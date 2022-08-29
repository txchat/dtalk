package dao

import (
	"fmt"
	"time"

	xerror "github.com/txchat/dtalk/pkg/error"

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
		&model.AddrBackup{}, &model.AddrRelate{},
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
		return nil, xerror.ErrNotFound
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
		return nil, xerror.ErrNotFound
	}
	return &record, nil
}

func (repo *BackupRepositoryMysql) createAddrBackup(r *model.AddrBackup) error {
	return repo.conn.Create(r).Error
}

func (repo *BackupRepositoryMysql) createAddrRelate(r *model.AddrRelate) error {
	return repo.conn.Create(r).Error
}

func (repo *BackupRepositoryMysql) updatePhone(r *model.AddrBackup) error {
	var records []model.AddrBackup
	tx := repo.conn.Begin()
	err := tx.Where("phone = ? AND address != ?", r.Phone, r.Address).Find(&records).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	}
	items := map[string]interface{}{"phone": "", "area": ""}
	for _, record := range records {
		err = tx.Model(model.AddrBackup{}).Where("address = ?", record.Address).Updates(items).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Model(model.AddrBackup{}).Where("address = ?", r.Address).Updates(r).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (repo *BackupRepositoryMysql) updatePhoneRelate(r *model.AddrRelate) error {
	var records []model.AddrRelate
	tx := repo.conn.Begin()
	err := tx.Where("phone = ? AND address != ?", r.Phone, r.Address).Find(&records).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	}
	items := map[string]interface{}{"phone": "", "area": ""}
	for _, record := range records {
		err = tx.Model(model.AddrRelate{}).Where("address = ?", record.Address).Updates(items).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Model(model.AddrRelate{}).Where("address = ?", r.Address).Updates(r).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (repo *BackupRepositoryMysql) updateEmail(r *model.AddrBackup) error {
	var records []model.AddrBackup
	tx := repo.conn.Begin()
	err := tx.Where("email = ? AND address != ?", r.Email, r.Address).Find(&records).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	}
	items := map[string]interface{}{"email": ""}
	for _, record := range records {
		err = tx.Model(model.AddrBackup{}).Where("address = ?", record.Address).Updates(items).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Model(model.AddrBackup{}).Where("address = ?", r.Address).Updates(r).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (repo *BackupRepositoryMysql) UpdateAddrBackup(tp int32, r *model.AddrBackup) error {
	var record model.AddrBackup
	err := repo.conn.Where("address = ?", r.Address).First(&record).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		r.CreateTime = time.Now()
		err = repo.createAddrBackup(r)
		if err != nil {
			return err
		}
	}
	switch tp {
	case model.Phone:
		return repo.updatePhone(r)
	case model.Email:
		return repo.updateEmail(r)
	default:
		return model.ErrQueryType
	}
}

func (repo *BackupRepositoryMysql) UpdateAddrRelate(tp int32, r *model.AddrRelate) error {
	var record model.AddrRelate
	err := repo.conn.Where("address = ?", r.Address).First(&record).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		r.CreateTime = time.Now()
		err = repo.createAddrRelate(r)
		if err != nil {
			return err
		}
	}
	switch tp {
	case model.Phone:
		return repo.updatePhoneRelate(r)
	case model.Email:
		return model.ErrQueryType
	default:
		return model.ErrQueryType
	}
}

func (repo *BackupRepositoryMysql) UpdateMnemonic(r *model.AddrBackup) error {
	err := repo.conn.Model(model.AddrBackup{}).Where("address = ?", r.Address).Updates(r).Error
	if err != nil {
		return err
	}
	return nil
}
