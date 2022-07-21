package model

import "time"

type VIPEntity struct {
	UID        string    `json:"uid" gorm:"primaryKey;column:uid;comment:用户地址"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;type:datetime;comment:更新时间"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;type:datetime;comment:创建时间"`
}

// TableName returns the table name of the DtalkVip model
func (d *VIPEntity) TableName() string {
	return "dtalk_vip"
}
