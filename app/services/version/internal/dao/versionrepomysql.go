package dao

import (
	"context"

	"github.com/go-sql-driver/mysql"
	"github.com/txchat/dtalk/app/services/version/internal/model"
	"github.com/txchat/dtalk/pkg/util"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	_GetVersionInfo                  = `SELECT * FROM dtalk_ver_backend WHERE id=?`
	_InsertVersionInfo               = `INSERT INTO dtalk_ver_backend ( platform, state, device_type, version_code, version_name, download_url, force_update,size,md5, description, ope_user, update_time, create_time)VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ,?,?)`
	_UpdateVersionInfo               = `UPDATE dtalk_ver_backend SET version_name=?,version_code=?,download_url=?,size=?,md5=?,description=?,force_update=?,update_time=?,ope_user=? WHERE id=?`
	_ChangeVersionStatus             = `UPDATE dtalk_ver_backend SET state=1,update_time=?,ope_user=? WHERE id=?`
	_QueryOtherReleaseVersionInfo    = `SELECT * FROM dtalk_ver_backend WHERE id!=? AND device_type=? AND platform=? AND state=1`
	_ChangeOtherReleaseVersionStatus = `UPDATE dtalk_ver_backend SET state=0,update_time=?,ope_user=? WHERE id=?`
	_GetVersionList                  = `SELECT * FROM dtalk_ver_backend WHERE platform LIKE ? AND device_type LIKE ? ORDER BY create_time DESC LIMIT ?,?`
	_GetVersionCount                 = `SELECT COUNT(*) AS version_count FROM dtalk_ver_backend WHERE platform LIKE ? AND device_type LIKE ?`
	_GetForceVersionCount            = `SELECT COUNT(*) AS force_num FROM dtalk_ver_backend WHERE device_type=? AND platform=? AND force_update = 1 AND version_code > ? AND version_code < ?`
	_GetReleaseVersionInfo           = `SELECT * FROM dtalk_ver_backend WHERE device_type=? AND platform=? AND state = 1 ORDER BY id DESC LIMIT 1`
	_QueryVersionInfoById            = `SELECT * FROM dtalk_ver_backend WHERE id=?`
)

type VersionRepositoryMysql struct {
	db *gorm.DB
}

func NewVersionRepositoryMysql(mysqlConfig mysql.Config) *VersionRepositoryMysql {
	if mysqlConfig.Params == nil {
		mysqlConfig.Params = make(map[string]string)
	}
	mysqlConfig.Params["charset"] = "UTF8MB4"
	dsn := mysqlConfig.FormatDSN()
	db, err := gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &VersionRepositoryMysql{
		db: db,
	}
}

func (repo *VersionRepositoryMysql) GetVersionInfo(ctx context.Context, vid int64) (*model.VersionForm, error) {
	var version model.VersionForm
	err := repo.db.Raw(_GetVersionInfo, vid).Take(&version).Error
	return &version, err
}

func (repo *VersionRepositoryMysql) AddVersionInfo(ctx context.Context, form *model.VersionForm) (int64, int64, error) {
	err := repo.db.Exec(_InsertVersionInfo, form.Platform, form.Status, form.DeviceType, form.VersionCode, form.VersionName, form.URL, form.Force, form.Size, form.Md5, form.Description.ToString(), form.OpeUser, form.UpdateTime, form.CreateTime).Error
	return repo.db.RowsAffected, 0, err
}

func (repo *VersionRepositoryMysql) UpdateVersionInfo(ctx context.Context, form *model.VersionForm) (int64, int64, error) {
	err := repo.db.Exec(_UpdateVersionInfo, form.VersionName, form.VersionCode, form.URL, form.Size, form.Md5, form.Description.ToString(), form.Force, form.UpdateTime, form.OpeUser, form.Id).Error
	return repo.db.RowsAffected, 0, err
}

// ReleaseSpecificVersion 将指定版本id的记录状态改为上线，同时将与之同平台和设备类型的版本号的状态置为下线。
func (repo *VersionRepositoryMysql) ReleaseSpecificVersion(ctx context.Context, vid, updateTime int64, operator string) error {
	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Exec(_ChangeVersionStatus, updateTime, operator, vid).Error; err != nil {
		tx.Rollback()
		return err
	}

	var versionInfo model.VersionForm
	if err := tx.Raw(_QueryVersionInfoById, vid).Take(&versionInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	var otherVersions []model.VersionForm
	if err := tx.Raw(_QueryOtherReleaseVersionInfo, vid, versionInfo.DeviceType, versionInfo.Platform).Scan(&otherVersions).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(otherVersions) < 1 {
		return tx.Commit().Error
	}
	for _, version := range otherVersions {
		id := util.MustToInt64(version.Id)
		if err := tx.Exec(_ChangeOtherReleaseVersionStatus, updateTime, operator, id).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (repo *VersionRepositoryMysql) SpecificPlatformAndDeviceTypeVersions(ctx context.Context, platform, deviceType string, page int64, size int64) ([]*model.VersionForm, error) {
	offset := page * size
	var versions []*model.VersionForm
	if err := repo.db.Raw(_GetVersionList, platform, deviceType, offset, size).Scan(&versions).Error; err != nil {
		return nil, err
	}
	return versions, nil
}

func (repo *VersionRepositoryMysql) SpecificPlatformAndDeviceTypeCount(ctx context.Context, platform, deviceType string) (int64, error) {
	var versionCount int64
	err := repo.db.Raw(_GetVersionCount, platform, deviceType).Scan(&versionCount).Error
	if err != nil {
		return 0, err
	}
	totalCount := util.MustToInt64(versionCount)
	return totalCount, nil
}

func (repo *VersionRepositoryMysql) LastReleaseVersion(ctx context.Context, platform, deviceType string) (*model.VersionForm, error) {
	var version model.VersionForm
	err := repo.db.Raw(_GetReleaseVersionInfo, deviceType, platform).Scan(&version).Error
	if err != nil {
		return nil, err
	}
	return &version, nil
}

func (repo *VersionRepositoryMysql) ForceNumberBetween(ctx context.Context, platform, deviceType string, begin, end int64) (int64, error) {
	if begin >= end {
		return 0, nil
	}
	var forceNum int64
	err := repo.db.Raw(_GetForceVersionCount, deviceType, platform, begin, end).Scan(&forceNum).Error
	if err != nil {
		return 0, err
	}
	return util.MustToInt64(forceNum), nil
}
