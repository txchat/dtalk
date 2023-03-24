package dao

import (
	"fmt"
	"strings"

	"github.com/txchat/dtalk/app/services/storage/internal/model"
)

const (
	_InsertSignalContent       = `INSERT INTO dtalk_signal_content(uid,seq,type,content,create_time,update_time) VALUES(?,?,?,?,?,?) ON DUPLICATE KEY UPDATE update_time=?`
	_InsertSignalContentPrefix = `INSERT INTO dtalk_signal_content(uid,seq,type,content,create_time,update_time) VALUES%s ON DUPLICATE KEY UPDATE update_time=VALUES(update_time)`
)

func (repo *StorageRepository) AppendSignal(m *model.SignalContent) (int64, int64, error) {
	err := repo.db.Exec(_InsertSignalContent,
		m.Uid, m.Seq, m.Type, m.Content, m.CreateTime, m.UpdateTime, m.UpdateTime).Error
	return 0, 0, err
}

func (repo *StorageRepository) BatchAppendSignal(m []*model.SignalContent) (int64, int64, error) {
	//element should not empty
	if len(m) == 0 {
		return 0, 0, nil
	}
	var values []interface{}
	condition := ""
	for _, row := range m {
		condition += "(?,?,?,?,?,?),"
		values = append(values, row.Uid, row.Seq, row.Type, row.Content, row.CreateTime, row.UpdateTime)
	}
	//trim the last ,
	condition = strings.TrimSuffix(condition, ",")
	//prepare the statement and exec
	err := repo.db.Exec(fmt.Sprintf(_InsertSignalContentPrefix, condition), values...).Error
	return 0, 0, err
}
