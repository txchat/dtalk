package dao

import (
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/service/group/model/db"
)

const (
	_InsertGroupApplysPrefix = `INSERT INTO dtalk_group_apply ( id, group_id, inviter_id, member_id, apply_note,
						operator_id, apply_status, reject_reason, create_time, update_time) VALUES `
	_InsertGroupApplysSuffix = ``

	_GetGroupApplyById = `SELECT * FROM dtalk_group_apply WHERE id=?`
	_GetGroupApplys    = `SELECT * FROM dtalk_group_apply WHERE group_id=? ORDER BY id desc LIMIT ?,?`

	_UpdateGroupApply = `UPDATE dtalk_group_apply SET operator_id=?, apply_status=?, reject_reason=?, update_time=? WHERE id=?`
)

// dtalk_group_apply

// InsertGroupApplys 插入加群审批
func (d *Dao) InsertGroupApplys(tx *mysql.MysqlTx, groupApplys []*db.GroupApply) error {
	var vals []interface{}
	valSql := ""
	for i, groupApply := range groupApplys {
		if i == 0 {
			valSql += "(?,?,?,?,?,?,?,?,?,?)"
		} else {
			valSql += ",(?,?,?,?,?,?,?,?,?,?)"
		}
		vals = append(vals,
			groupApply.Id,
			groupApply.GroupId,
			groupApply.InviterId,
			groupApply.MemberId,
			groupApply.ApplyNote,
			groupApply.OperatorId,
			groupApply.ApplyStatus,
			groupApply.RejectReason,
			groupApply.CreateTime,
			groupApply.UpdateTime,
		)
	}
	SQL := _InsertGroupApplysPrefix + valSql + _InsertGroupApplysSuffix

	_, _, err := tx.Exec(SQL, vals...)
	return err
}

// GetGroupApplyById 通过 ID 获得群审批详情
func (d *Dao) GetGroupApplyById(id int64) (*db.GroupApply, error) {
	maps, err := d.conn.Query(_GetGroupApplyById, id)
	if err != nil {
		return nil, err
	}

	if len(maps) == 0 {
		return nil, nil
	}

	res := maps[0]

	return db.ConvertGroupApply(res), nil
}

func (d *Dao) GetGroupApplys(groupId int64, limit, offset int32) ([]*db.GroupApply, error) {
	maps, err := d.conn.Query(_GetGroupApplys, groupId, offset, limit)
	if err != nil {
		return nil, err
	}
	res := make([]*db.GroupApply, 0)
	for _, m := range maps {
		t := db.ConvertGroupApply(m)
		res = append(res, t)
	}

	return res, nil
}

func (d *Dao) UpdateGroupApply(operatorId string, applyStatus int32, rejectReason string, id int64) error {
	nowTime := d.getNowTime()
	_, _, err := d.conn.Exec(_UpdateGroupApply, operatorId, applyStatus, rejectReason, nowTime, id)
	return err
}
