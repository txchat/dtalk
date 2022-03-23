package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/txchat/dtalk/service/backup/model"
)

func (d *Dao) QueryAddressEnrolment(btyAddr string) (*model.AddrMove, error) {
	var record model.AddrMove
	err := d.db.Where("bty_addr = ?", btyAddr).First(&record).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		d.log.Error("Query err", "err", err, "bty_addr", btyAddr)
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &record, nil
}

func (d *Dao) CreateAddressEnrolment(r *model.AddrMove) error {
	err := d.db.Create(r).Error
	if err != nil {
		d.log.Error("CreateAddressEnrolment err", "err", err, "r", r)
		return err
	}
	return nil
}
