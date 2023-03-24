package dao

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/mohae/deepcopy"
	"github.com/stretchr/testify/assert"
	"github.com/txchat/dtalk/app/services/version/internal/model"
)

var (
	mysqlRootPassword string
	repo              *VersionRepositoryMysql
)

func TestMain(m *testing.M) {
	mysqlRootPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
	repo = NewVersionRepositoryMysql(mysql.Config{
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		User:   "root",
		Passwd: mysqlRootPassword,
		DBName: "dtalk",
	})
	os.Exit(m.Run())
}

var (
	now = time.Now()
	v1  = model.VersionForm{
		Id:          1,
		Platform:    "chat",
		Status:      1,
		DeviceType:  "Android",
		VersionName: "test version 1",
		VersionCode: 1,
		URL:         "test.com",
		Force:       false,
		Description: model.Description{
			"1.描述one",
			"2.表述two",
		},
		OpeUser:    "",
		Md5:        "123",
		Size:       123,
		UpdateTime: now.UnixMilli(),
		CreateTime: now.UnixMilli(),
	}
	v2 = model.VersionForm{
		Id:          2,
		Platform:    "chat",
		Status:      0,
		DeviceType:  "Android",
		VersionName: "test version 2",
		VersionCode: 2,
		URL:         "test.com",
		Force:       false,
		Description: nil,
		OpeUser:     "",
		Md5:         "456",
		Size:        456,
		UpdateTime:  now.UnixMilli(),
		CreateTime:  now.UnixMilli(),
	}
)

func TestVersionRepositoryMysql_AddVersionInfo(t *testing.T) {
	_, _, err := repo.AddVersionInfo(context.Background(), &v1)
	assert.Nil(t, err)
	v1Info, err := repo.GetVersionInfo(context.Background(), 1)
	assert.Nil(t, err)
	assert.EqualValues(t, &v1, v1Info)
	_, _, err = repo.AddVersionInfo(context.Background(), &v2)
	assert.Nil(t, err)
	lastV, err := repo.LastReleaseVersion(context.Background(), "chat", "Android")
	assert.Nil(t, err)
	assert.EqualValues(t, &v1, lastV)
	err = repo.ReleaseSpecificVersion(context.Background(), 2, now.UnixMilli(), "test user1")
	assert.Nil(t, err)
	newV2 := deepcopy.Copy(v2).(model.VersionForm)
	newV2.Status = 1
	newV2.OpeUser = "test user1"
	v2Info, err := repo.LastReleaseVersion(context.Background(), "chat", "Android")
	assert.Nil(t, err)
	assert.EqualValues(t, &newV2, v2Info)
	num, err := repo.ForceNumberBetween(context.Background(), "chat", "Android", 0, 100)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), num)
	v2.Force = true
	_, _, err = repo.UpdateVersionInfo(context.Background(), &v2)
	assert.Nil(t, err)
	num, err = repo.ForceNumberBetween(context.Background(), "chat", "Android", 0, 100)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), num)
	num, err = repo.SpecificPlatformAndDeviceTypeCount(context.Background(), "chat", "Android")
	assert.Nil(t, err)
	assert.Equal(t, int64(2), num)
	versions, err := repo.SpecificPlatformAndDeviceTypeVersions(context.Background(), "chat", "Android", 0, 100)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(versions))
}
