package model

type AddrMove struct {
	BtyAddr string `json:"bty_addr" gorm:"primary_key;column:bty_addr;type:varchar(255);comment:'bty地址'"`
	BtcAddr string `json:"btc_addr" gorm:"column:btc_addr;type:varchar(255);not null;comment:'btc地址'"`
	State   int32  `json:"state" gorm:"column:state;type:tinyint(255);comment:'0->对应关系已建立；1-> 好友关系已迁移'"`
}

// TableName returns the table name of the DtalkAddrMove model
func (d *AddrMove) TableName() string {
	return "dtalk_addr_move"
}
