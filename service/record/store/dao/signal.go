package dao

import (
	"fmt"
	"strings"

	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/record/store/model"
)

const (
	_InsertSignalContent       = `INSERT INTO dtalk_signal_content(id,uid,type,state,content,create_time,update_time) VALUES(?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE update_time=?`
	_InsertSignalContentPrefix = `INSERT INTO dtalk_signal_content(id,uid,type,state,content,create_time,update_time) VALUES%s ON DUPLICATE KEY UPDATE update_time=VALUES(update_time)`
	_SetSignalContentState     = `UPDATE dtalk_signal_content SET state = ? WHERE uid = ? AND id = ?`

	_GetUnReceiveSignal = `SELECT * FROM dtalk_signal_content WHERE uid=? AND state=? ORDER BY id`
	_GetSyncSignal      = `SELECT * FROM dtalk_signal_content WHERE uid=? AND id>? ORDER BY id LIMIT ?,?`
)

func (d *Dao) AppendSignalContent(m *model.SignalContent) (int64, int64, error) {
	num, lastId, err := d.conn.Exec(_InsertSignalContent,
		m.Id, m.Uid, m.Type, m.State, m.Content, m.CreateTime, m.UpdateTime, m.UpdateTime)
	return num, lastId, err
}

func (d *Dao) BatchAppendSignalContent(m []*model.SignalContent) (int64, int64, error) {
	//element should not empty
	if len(m) == 0 {
		return 0, 0, nil
	}
	var values []interface{}
	condition := ""
	for _, row := range m {
		condition += "(?,?,?,?,?,?,?),"
		values = append(values, row.Id, row.Uid, row.Type, row.State, row.Content, row.CreateTime, row.UpdateTime)
	}
	//trim the last ,
	condition = strings.TrimSuffix(condition, ",")
	//prepare the statement and exec
	num, lastId, err := d.conn.PrepareExec(fmt.Sprintf(_InsertSignalContentPrefix, condition), values...)
	return num, lastId, err
}

func (d *Dao) MarkSignalReceived(uid string, mid int64) (int64, int64, error) {
	num, lastId, err := d.conn.Exec(_SetSignalContentState, model.Received, uid, mid)
	return num, lastId, err
}

func (d *Dao) UnReceiveSignalMsg(uid string) ([]*model.SignalContent, error) {
	maps, err := d.conn.Query(_GetUnReceiveSignal, uid, model.UnReceive)
	if err != nil {
		return nil, err
	}
	ret := make([]*model.SignalContent, len(maps))
	for i, m := range maps {
		ret[i] = &model.SignalContent{
			Id:         m["id"],
			Uid:        m["uid"],
			Type:       uint8(util.MustToUint32(m["type"])),
			State:      uint8(util.MustToUint32(m["state"])),
			Content:    m["content"],
			CreateTime: uint64(util.MustToInt64(m["create_time"])),
			UpdateTime: uint64(util.MustToInt64(m["update_time"])),
		}
	}
	return ret, err
}

func (d *Dao) SyncSignalMsg(uid string, startId, count int64) ([]*model.SignalContent, error) {
	maps, err := d.conn.Query(_GetSyncSignal, uid, startId, 0, count)
	if err != nil {
		return nil, err
	}
	ret := make([]*model.SignalContent, len(maps))
	for i, m := range maps {
		ret[i] = &model.SignalContent{
			Id:         m["id"],
			Uid:        m["uid"],
			Type:       uint8(util.MustToUint32(m["type"])),
			State:      uint8(util.MustToUint32(m["state"])),
			Content:    m["content"],
			CreateTime: uint64(util.MustToInt64(m["create_time"])),
			UpdateTime: uint64(util.MustToInt64(m["update_time"])),
		}
	}
	return ret, err
}
