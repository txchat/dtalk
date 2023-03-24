package dao

import (
	"os"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/mohae/deepcopy"
	"github.com/stretchr/testify/assert"
	"github.com/txchat/dtalk/app/services/backup/internal/model"
	"github.com/zeromicro/go-zero/core/service"
	"gorm.io/gorm/utils"
)

var (
	mysqlRootPassword string
	repo              *BackupRepositoryMysql
)

func TestMain(m *testing.M) {
	mysqlRootPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
	mysqlRootPassword = "123456"
	conn := NewDefaultConn(service.TestMode, mysql.Config{
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		User:   "root",
		Passwd: mysqlRootPassword,
		DBName: "dtalk",
	})
	repo = NewBackupRepositoryMysql(conn)
	os.Exit(m.Run())
}

var (
	now              = time.Now()
	addrBackupRecord = model.AddrBackup{
		Address:    "138WycPfu4JketDH9aL54Xsq1Ds4cEQfRa",
		Area:       "86",
		Phone:      "15700000000",
		Email:      "0000@gmail.com",
		Mnemonic:   "饭 吨 纺 毕 再 他 寺 降 易 溜 诚 静 依 演 每",
		PrivateKey: "02ae3d82a3d0042f0e786f546d7f9c5d6d25a35393c4b232f2ca8114bccfcd3fcd",
		UpdateTime: now,
		CreateTime: now,
	}
	addrRelateRecord = model.AddrRelate{
		Address:    "138WycPfu4JketDH9aL54Xsq1Ds4cEQfRa",
		Area:       "86",
		Phone:      "15700000000",
		Email:      "0000@gmail.com",
		Mnemonic:   "饭 吨 纺 毕 再 他 寺 降 易 溜 诚 静 依 演 每",
		PrivateKey: "02ae3d82a3d0042f0e786f546d7f9c5d6d25a35393c4b232f2ca8114bccfcd3fcd",
		UpdateTime: now,
		CreateTime: now,
	}
)

func TestBackupRepositoryMysql_UpdateAddrBackup(t *testing.T) {
	err := repo.UpdateAddrBackup(model.Phone, &addrBackupRecord)
	assert.Nil(t, err)
	err = repo.UpdateAddrBackup(model.Email, &addrBackupRecord)
	assert.Nil(t, err)
	backup, err := repo.QueryBind(model.Phone, "15700000000")
	assert.Nil(t, err)
	utils.AssertEqual(addrBackupRecord, backup)
	backup, err = repo.QueryBind(model.Email, "0000@gmail.com")
	assert.Nil(t, err)
	utils.AssertEqual(addrBackupRecord, backup)
}

func TestBackupRepositoryMysql_UpdateAddrRelate(t *testing.T) {
	err := repo.UpdateAddrRelate(model.Phone, &addrRelateRecord)
	assert.Nil(t, err)
	err = repo.UpdateAddrRelate(model.Email, &addrRelateRecord)
	assert.Equal(t, model.ErrQueryType, err)
	backup, err := repo.QueryRelate(model.Phone, "15700000000")
	assert.Nil(t, err)
	utils.AssertEqual(addrRelateRecord, backup)
	backup, err = repo.QueryRelate(model.Email, "0000@gmail.com")
	assert.Nil(t, err)
	utils.AssertEqual(addrRelateRecord, backup)
}

func TestBackupRepositoryMysql_UpdateMnemonic(t *testing.T) {
	updateMne := deepcopy.Copy(addrBackupRecord).(model.AddrBackup)
	updateMne.Mnemonic = "new mne"
	err := repo.UpdateMnemonic(&updateMne)
	assert.Nil(t, err)
	backup, err := repo.QueryRelate(model.Phone, "15700000000")
	assert.Nil(t, err)
	utils.AssertEqual(updateMne, backup)
}
