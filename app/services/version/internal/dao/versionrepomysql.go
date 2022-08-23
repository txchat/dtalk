package dao

import (
	"context"

	"github.com/txchat/dtalk/app/services/version/internal/model"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/pkg/util"
)

const (
	_GetVersionInfo                  = `SELECT * FROM dtalk_ver_backend WHERE id=?`
	_InsertVersionInfo               = `INSERT INTO dtalk_ver_backend ( platform, state, device_type, version_code, version_name, download_url, force_update,size,md5, description, ope_user, update_time, create_time)VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ,?,?)`
	_UpdateVersionInfo               = `UPDATE dtalk_ver_backend SET version_name=?,version_code=?,download_url=?,size=?,md5=?,description=?,force_update=?,update_time=?,ope_user=? WHERE id=?`
	_ChangeVersionStatus             = `UPDATE dtalk_ver_backend SET state=1,update_time=?,ope_user=? WHERE id=?`
	_QueryOtherReleaseVersionInfo    = `SELECT * FROM dtalk_ver_backend WHERE id!=? AND device_type=? AND platform=? AND state=1`
	_ChangeOtherReleaseVersionStatus = `UPDATE dtalk_ver_backend SET state=0,update_time=?,ope_user=? WHERE id=?`
	_GetVersionList                  = `SELECT * FROM dtalk_ver_backend WHERE platform LIKE ? AND device_type LIKE ? ORDER BY create_time DESC LIMIT ?,?`
	_GetVersionCount                 = `SELECT COUNT(*) FROM dtalk_ver_backend WHERE platform LIKE ? AND device_type LIKE ?`
	_GetForceVersionCount            = `SELECT COUNT(*) AS force_num FROM dtalk_ver_backend WHERE device_type=? AND platform=? AND force_update = 1 AND version_code > ? AND version_code < ?`
	_GetReleaseVersionInfo           = `SELECT * FROM dtalk_ver_backend WHERE device_type=? AND platform=? AND state = 1 ORDER BY id DESC LIMIT 1`
	_QueryVersionInfoById            = `SELECT * FROM dtalk_ver_backend WHERE id=?`
)

type VersionRepositoryMysql struct {
	conn *mysql.MysqlConn
}

func NewVersionRepositoryMysql(conn *mysql.MysqlConn) *VersionRepositoryMysql {
	return &VersionRepositoryMysql{
		conn: conn,
	}
}

func (repo *VersionRepositoryMysql) GetVersionInfo(ctx context.Context, vid int64) (*model.VersionForm, error) {
	records, err := repo.conn.Query(_GetVersionInfo, vid)
	if err != nil {
		return nil, err
	}
	if len(records) < 1 {
		return nil, xerror.NewError(xerror.ParamsError).SetExtMessage("无此id的记录")
	}
	record := records[0]
	version, err := model.ConvertVersionForm(record)
	if err != nil {
		return nil, err
	}
	return version, nil
}

func (repo *VersionRepositoryMysql) AddVersionInfo(ctx context.Context, form *model.VersionForm) (int64, int64, error) {
	num, lastId, err := repo.conn.Exec(_InsertVersionInfo, form.Platform, form.Status, form.DeviceType, form.VersionCode, form.VersionName, form.URL, form.Force, form.Size, form.Md5, form.Description.ToString(), form.OpeUser, form.UpdateTime, form.CreateTime)
	return num, lastId, err
}

func (repo *VersionRepositoryMysql) UpdateVersionInfo(ctx context.Context, form *model.VersionForm) (int64, int64, error) {
	num, lastId, err := repo.conn.Exec(_UpdateVersionInfo, form.VersionName, form.VersionCode, form.URL, form.Size, form.Md5, form.Description.ToString(), form.Force, form.UpdateTime, form.OpeUser, form.Id)
	return num, lastId, err
}

// ReleaseSpecificVersion 将指定版本id的记录状态改为上线，同时将与之同平台和设备类型的版本号的状态置为下线。
func (repo *VersionRepositoryMysql) ReleaseSpecificVersion(ctx context.Context, vid, updateTime int64, operator string) error {
	tx, err := repo.conn.NewTx()
	if err != nil {
		return err
	}
	_, _, err = tx.Exec(_ChangeVersionStatus, updateTime, operator, vid)
	if err != nil {
		return err
	}

	records, err := tx.Query(_QueryVersionInfoById, vid)
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return xerror.NewError(xerror.ParamsError).SetExtMessage("无此id的记录")
	}
	deviceType := records[0]["device_type"]
	platform := records[0]["platform"]

	records, err = tx.Query(_QueryOtherReleaseVersionInfo, vid, deviceType, platform)
	if err != nil {
		return err
	}
	if len(records) == 0 {
		err = tx.Commit()
		if err != nil {
			return err
		}
		return nil
	}

	id := util.ToInt64(records[0]["id"])
	_, _, err = tx.Exec(_ChangeOtherReleaseVersionStatus, updateTime, operator, id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repo *VersionRepositoryMysql) SpecificPlatformAndDeviceTypeVersions(ctx context.Context, platform, deviceType string, page int64, size int64) ([]*model.VersionForm, error) {
	offset := page * size
	records, err := repo.conn.Query(_GetVersionList, platform, deviceType, offset, size)
	if err != nil {
		return nil, err
	}

	versions := make([]*model.VersionForm, len(records))
	for i, record := range records {
		version, err := model.ConvertVersionForm(record)
		if err != nil {
			return nil, err
		}
		versions[i] = version
	}
	return versions, nil
}

func (repo *VersionRepositoryMysql) SpecificPlatformAndDeviceTypeCount(ctx context.Context, platform, deviceType string) (int64, error) {
	countRecords, err := repo.conn.Query(_GetVersionCount, platform, deviceType)
	if err != nil {
		return 0, err
	}
	if len(countRecords) == 0 {
		return 0, xerror.NewError(xerror.ParamsError).SetExtMessage("无此记录")
	}
	totalCount := util.ToInt64(countRecords[0]["COUNT(*)"])
	return totalCount, nil
}

func (repo *VersionRepositoryMysql) LastReleaseVersion(ctx context.Context, platform, deviceType string) (*model.VersionForm, error) {
	records, err := repo.conn.Query(_GetReleaseVersionInfo, deviceType, platform)
	if err != nil {
		return nil, err
	}

	if len(records) < 1 {
		return &model.VersionForm{}, nil
	}
	record, err := model.ConvertVersionForm(records[0])
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (repo *VersionRepositoryMysql) ForceNumberBetween(ctx context.Context, platform, deviceType string, begin, end int64) (int64, error) {
	if begin >= end {
		return 0, nil
	}
	records, err := repo.conn.Query(_GetForceVersionCount, deviceType, platform, begin, end)
	if err != nil {
		return 0, err
	}

	if len(records) < 1 {
		return 0, nil
	}
	return util.ToInt64(records[0]["force_num"]), nil
}
