package dao

import (
	"time"

	"github.com/txchat/dtalk/app/services/backup/internal/model"
	xerror "github.com/txchat/dtalk/pkg/error"
	xmysql "github.com/txchat/dtalk/pkg/mysql"
	"github.com/zeromicro/go-zero/core/service"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDefaultConn(mode string, mysqlConfig xmysql.Config) *gorm.DB {
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
	db, err := gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      defaultLogger,
	})
	if err != nil {
		panic(err)
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB AUTO_INCREMENT=1 CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci").AutoMigrate(
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
