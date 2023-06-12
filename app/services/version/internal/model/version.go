package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Description []string

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Description
func (d *Description) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case []byte:
		if len(v) == 0 {
			return nil
		}
		*d = strings.Split(string(v), ";")
		return nil
	case string:
		if len(v) == 0 {
			return nil
		}
		*d = strings.Split(v, ";")
		return nil
	default:
		return fmt.Errorf("unexpected type %T for Description", value)
	}
}

// Value 实现 driver.Valuer 接口，Value 返回 序列化后的Description
func (d Description) Value() (driver.Value, error) {
	return strings.Join(d, ";"), nil
}

func (d Description) ToString() string {
	return strings.Join(d, ";")
}

type VersionForm struct {
	Id          int64       `json:"id" gorm:"column:id;"`
	Platform    string      `json:"platform" gorm:"column:platform;"`
	Status      int32       `json:"status" gorm:"column:state;"`
	DeviceType  string      `json:"deviceType" gorm:"column:device_type;"`
	VersionName string      `json:"versionName" gorm:"column:version_name;"`
	VersionCode int64       `json:"versionCode" gorm:"column:version_code;"`
	URL         string      `json:"url" gorm:"column:download_url;"`
	Force       bool        `json:"force" gorm:"column:force_update;"`
	Description Description `json:"description" gorm:"column:description;"`
	OpeUser     string      `json:"opeUser" gorm:"column:ope_user;"`
	Md5         string      `json:"md5" gorm:"column:md5;"`
	Size        int64       `json:"size" gorm:"column:size;"`
	UpdateTime  int64       `json:"updateTime" gorm:"column:update_time;"`
	CreateTime  int64       `json:"createTime" gorm:"column:create_time;"`
}
