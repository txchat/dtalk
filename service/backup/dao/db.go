package dao

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/txchat/dtalk/service/backup/model"
)

func (d *Dao) Query(tp uint32, query string) (*model.AddrBackup, error) {
	var record model.AddrBackup
	sqlCase := ""
	switch tp {
	case model.Phone:
		sqlCase = "phone = ?"
	case model.Email:
		sqlCase = "email = ?"
	case model.Address:
		sqlCase = "address = ?"
	default:
		return nil, model.ErrQueryType
	}
	err := d.db.Where(sqlCase, query).First(&record).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		d.log.Error("Query err", "err", err, "type", tp, "query", query)
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &record, nil
}

func (d *Dao) QueryRelate(tp uint32, query string) (*model.AddrRelate, error) {
	var record model.AddrRelate
	sqlCase := ""
	switch tp {
	case model.Phone:
		sqlCase = "phone = ?"
	case model.Email:
		sqlCase = "email = ?"
	case model.Address:
		sqlCase = "address = ?"
	default:
		return nil, model.ErrQueryType
	}
	err := d.db.Where(sqlCase, query).First(&record).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		d.log.Error("Query err", "err", err, "type", tp, "query", query)
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &record, nil
}

func (d *Dao) createAddrBackup(r *model.AddrBackup) error {
	err := d.db.Create(r).Error
	if err != nil {
		d.log.Error("Create AddrBackup err", "err", err, "r", r)
	}
	return err
}

func (d *Dao) updatePhone(r *model.AddrBackup) error {
	var records []model.AddrBackup
	tx := d.db.Begin()
	err := tx.Where("phone = ? AND address != ?", r.Phone, r.Address).Find(&records).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	}
	new := map[string]interface{}{"phone": "", "area": ""}
	for _, record := range records {
		err = tx.Model(model.AddrBackup{}).Where("address = ?", record.Address).Updates(new).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Model(model.AddrBackup{}).Where("address = ?", r.Address).Updates(r).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (d *Dao) updatePhoneRelate(r *model.AddrRelate) error {
	var records []model.AddrRelate
	tx := d.db.Begin()
	err := tx.Where("phone = ? AND address != ?", r.Phone, r.Address).Find(&records).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	}
	new := map[string]interface{}{"phone": "", "area": ""}
	for _, record := range records {
		err = tx.Model(model.AddrRelate{}).Where("address = ?", record.Address).Updates(new).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Model(model.AddrRelate{}).Where("address = ?", r.Address).Updates(r).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (d *Dao) updateEmail(r *model.AddrBackup) error {
	var records []model.AddrBackup
	tx := d.db.Begin()
	err := tx.Where("email = ? AND address != ?", r.Email, r.Address).Find(&records).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	}
	new := map[string]interface{}{"email": ""}
	for _, record := range records {
		err = tx.Model(model.AddrBackup{}).Where("address = ?", record.Address).Updates(new).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Model(model.AddrBackup{}).Where("address = ?", r.Address).Updates(r).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (d *Dao) UpdateAddrBackup(tp uint32, r *model.AddrBackup) error {
	var record model.AddrBackup
	err := d.db.Where("address = ?", r.Address).First(&record).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		r.CreateTime = time.Now()
		err = d.createAddrBackup(r)
		if err != nil {
			return err
		}
	}
	switch tp {
	case model.Phone:
		return d.updatePhone(r)
	case model.Email:
		return d.updateEmail(r)
	default:
		return model.ErrQueryType
	}
}

func (d *Dao) UpdateAddrRelate(tp uint32, r *model.AddrRelate) error {
	var record model.AddrRelate
	err := d.db.Where("address = ?", r.Address).First(&record).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		r.CreateTime = time.Now()
		err := d.db.Create(r).Error
		if err != nil {
			d.log.Error("Create AddrRelate err", "err", err, "r", r)
		}
	}
	switch tp {
	case model.Phone:
		return d.updatePhoneRelate(r)
	case model.Email:
		return model.ErrQueryType
	default:
		return model.ErrQueryType
	}
}

func (d *Dao) UpdateMnemonic(r *model.AddrBackup) error {
	err := d.db.Model(model.AddrBackup{}).Where("address = ?", r.Address).Updates(r).Error
	if err != nil {
		d.log.Error("UpdateMnemonic err", "err", err, "r", r)
		return err
	}
	return nil
}
