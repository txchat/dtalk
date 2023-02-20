package dao

import (
	"fmt"
	"strings"

	"github.com/txchat/dtalk/app/services/storage/internal/model"
)

const (
	_InsertSignalContent       = `INSERT INTO dtalk_signal_content(seq,uid,type,content,create_time,update_time) VALUES(?,?,?,?,?,?) ON DUPLICATE KEY UPDATE update_time=?`
	_InsertSignalContentPrefix = `INSERT INTO dtalk_signal_content(seq,uid,type,content,create_time,update_time) VALUES%s ON DUPLICATE KEY UPDATE update_time=VALUES(update_time)`
)

func (repo *StorageRepository) AppendSignalContent(m *model.SignalContent) (int64, int64, error) {
	num, lastId, err := repo.mysql.Exec(_InsertSignalContent,
		m.Seq, m.Uid, m.Type, m.Content, m.CreateTime, m.UpdateTime, m.UpdateTime)
	return num, lastId, err
}

func (repo *StorageRepository) BatchAppendSignalContent(m []*model.SignalContent) (int64, int64, error) {
	//element should not empty
	if len(m) == 0 {
		return 0, 0, nil
	}
	var values []interface{}
	condition := ""
	for _, row := range m {
		condition += "(?,?,?,?,?,?),"
		values = append(values, row.Seq, row.Uid, row.Type, row.Content, row.CreateTime, row.UpdateTime)
	}
	//trim the last ,
	condition = strings.TrimSuffix(condition, ",")
	//prepare the statement and exec
	num, lastId, err := repo.mysql.PrepareExec(fmt.Sprintf(_InsertSignalContentPrefix, condition), values...)
	return num, lastId, err
}
