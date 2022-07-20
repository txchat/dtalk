package dao

import (
	"math"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/backend/model"
)

const (
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

func (d *Dao) InsertVersion(form *model.VersionForm) (int64, int64, error) {
	num, lastId, err := d.conn.Exec(_InsertVersionInfo, form.Platform, form.Status, form.DeviceType, form.VersionCode, form.VersionName, form.Url, form.Force, form.Size, form.Md5, form.Description.ToString(), form.OpeUser, form.UpdateTime, form.CreateTime)
	return num, lastId, err
}

func (d *Dao) UpdateVersion(form *model.VersionForm) (*model.VersionForm, error) {
	_, _, err := d.conn.Exec(_UpdateVersionInfo, form.VersionName, form.VersionCode, form.Url, form.Size, form.Md5, form.Description.ToString(), form.Force, form.UpdateTime, form.OpeUser, form.Id)
	if err != nil {
		return nil, err
	}
	records, err := d.conn.Query(_QueryVersionInfoById, form.Id)
	if err != nil {
		return nil, err
	}
	if len(records) < 1 {
		return nil, xerror.NewError(xerror.ParamsError).SetExtMessage("无此id的记录")
	}
	record := records[0]
	version, err := model.ConvertVersionForm(&record)
	if err != nil {
		return nil, err
	}
	return version, nil
}

func (d *Dao) ChangeVersionStatus(form *model.VersionForm) error {
	tx, err := d.conn.NewTx()
	if err != nil {
		return err
	}
	_, _, err = tx.Exec(_ChangeVersionStatus, form.UpdateTime, form.OpeUser, form.Id)
	if err != nil {
		return err
	}

	records, err := tx.Query(_QueryVersionInfoById, form.Id)
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return xerror.NewError(xerror.ParamsError).SetExtMessage("无此id的记录")
	}
	deviceType := records[0]["device_type"]
	platform := records[0]["platform"]

	records, err = tx.Query(_QueryOtherReleaseVersionInfo, form.Id, deviceType, platform)
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
	_, _, err = tx.Exec(_ChangeOtherReleaseVersionStatus, form.UpdateTime, form.OpeUser, id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) GetVersionList(form *model.VersionForm, page int64, size int64) (*[]model.VersionForm, int64, int64, error) {
	var records []map[string]string
	offset := page * size
	records, err := d.conn.Query(_GetVersionList, form.Platform, form.DeviceType, offset, size)
	if err != nil {
		return nil, 0, 0, err
	}
	response, err := d.ConvertVersionList(&records)
	if err != nil {
		return nil, 0, 0, err
	}
	countRecords, err := d.conn.Query(_GetVersionCount, form.Platform, form.DeviceType)
	if err != nil {
		return nil, 0, 0, err
	}
	totalElements := util.ToInt64(countRecords[0]["COUNT(*)"])
	totalPages := int64(math.Ceil(float64(totalElements) / float64(size)))
	return response, totalElements, totalPages, nil
}

func (d *Dao) CheckAndUpdateVersion(form *model.VersionForm) (*model.VersionForm, error) {
	records, err := d.conn.Query(_GetReleaseVersionInfo, form.DeviceType, form.Platform)
	if err != nil {
		return nil, err
	}

	if len(records) < 1 {
		return &model.VersionForm{}, nil
	}
	record, err := model.ConvertVersionForm(&records[0])
	if err != nil {
		return nil, err
	}
	if !record.Force {
		//判断期间是否包含强制更新版本
		record.Force, _ = d.checkHaveForce(form, util.ToInt(record.VersionCode))
	}
	return record, nil
}

func (d *Dao) ConvertVersionList(records *[]map[string]string) (*[]model.VersionForm, error) {
	//var versionList []model.VersionForm
	versionList := make([]model.VersionForm, len(*records), len(*records))
	for i, record := range *records {
		version, err := model.ConvertVersionForm(&record)
		if err != nil {
			return nil, err
		}
		versionList[i] = *version
	}
	return &versionList, nil
}

func (d *Dao) checkHaveForce(form *model.VersionForm, newVer int) (bool, error) {
	oldVer := util.ToInt(form.VersionCode)
	if oldVer >= newVer {
		return false, nil
	}
	records, err := d.conn.Query(_GetForceVersionCount, form.DeviceType, form.Platform, oldVer, newVer)
	if err != nil {
		return false, err
	}

	if len(records) < 1 {
		return false, nil
	}
	num := util.ToInt64(records[0]["force_num"])
	return num > 0, nil
}
