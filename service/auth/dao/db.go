package dao

const (
	_InsertSignInInfo = `INSERT IGNORE INTO dtalk_ver_auth ( app_id, app_config, app_key, create_time, update_time)VALUES( ?, ?, ?, ?, ?)`
	_LoadConfig       = `SELECT app_config FROM dtalk_ver_auth WHERE app_id=?`
	_GetKey           = `SELECT app_key FROM dtalk_ver_auth WHERE app_id=?`
)

func (d *Dao) SignIn(appId string, config *string, key string, createTime, updateTime int64) (int64, error) {
	num, _, err := d.conn.Exec(_InsertSignInInfo, appId, *config, key, createTime, updateTime)
	if err != nil {
		return num, err
	}

	return num, nil
}

func (d *Dao) LoadConfig(appId string) (string, error) {
	records, err := d.conn.Query(_LoadConfig, appId)
	if err != nil {
		return "", err
	}
	if len(records) == 0 {
		return "", nil
	}

	config := records[0]["app_config"]

	return config, nil
}

func (d *Dao) GetKey(appId string) (string, error) {
	records, err := d.conn.Query(_GetKey, appId)
	if err != nil {
		return "", err
	}
	if len(records) == 0 {
		return "", nil
	}
	key := records[0]["app_key"]

	return key, nil
}
