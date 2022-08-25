package model

import "time"

type AddrBackup struct {
	Address    string    `json:"address" gorm:"primary_key;column:address;type:varchar(255);comment:'用户地址'"`
	Area       string    `json:"area" gorm:"column:area;type:varchar(4);comment:'区号'"`
	Phone      string    `json:"phone" gorm:"column:phone;type:varchar(11);comment:'手机号'"`
	Email      string    `json:"email" gorm:"column:email;type:varchar(30);comment:'邮箱'"`
	Mnemonic   string    `json:"mnemonic" gorm:"column:mnemonic;type:varchar(1020);comment:'助记词'"`
	PrivateKey string    `json:"private_key" gorm:"column:private_key;type:varchar(1020);comment:'加密私钥'"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;type:datetime;comment:'更新时间'"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;type:datetime;comment:'创建时间'"`
}

// TableName returns the table name of the DtalkAddrBackup model
func (d *AddrBackup) TableName() string {
	return "dtalk_addr_backup"
}

type AddrRelate struct {
	Address    string    `json:"address" gorm:"primary_key;column:address;type:varchar(255);comment:'用户地址'"`
	Area       string    `json:"area" gorm:"column:area;type:varchar(4);comment:'区号'"`
	Phone      string    `json:"phone" gorm:"column:phone;type:varchar(11);comment:'手机号'"`
	Email      string    `json:"email" gorm:"column:email;type:varchar(30);comment:'邮箱'"`
	Mnemonic   string    `json:"mnemonic" gorm:"column:mnemonic;type:varchar(1020);comment:'助记词'"`
	PrivateKey string    `json:"private_key" gorm:"column:private_key;type:varchar(1020);comment:'加密私钥'"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;type:datetime;comment:'更新时间'"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;type:datetime;comment:'创建时间'"`
}

// TableName returns the table name of the DtalkAddrBackup model
func (d *AddrRelate) TableName() string {
	return "dtalk_addr_relate"
}
