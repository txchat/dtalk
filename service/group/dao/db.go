package dao

import (
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
	"strings"
)

const (
	// dtalk_group_info
	_InsertGroupInfo = `INSERT INTO dtalk_group_info ( group_id, group_mark_id, group_name, group_avatar, group_member_num, group_maximum,
						group_introduce, group_status, group_owner_id, group_create_time, group_update_time, group_join_type,
						group_mute_type, group_friend_type, group_aes_key, group_pub_name, group_type ) 
						VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`
	_UpdateGroupInfoName = `UPDATE dtalk_group_info SET group_name=?, group_pub_name=?, group_update_time=? 
							WHERE group_id=?`
	_UpdateGroupInfoAvatar = `UPDATE dtalk_group_info SET group_avatar=?, group_update_time=? 
							WHERE group_id=?`
	_UpdateGroupInfoGroupNum = `UPDATE dtalk_group_info SET group_member_num=?, group_update_time=? 
							WHERE group_id=?`
	_UpdateGroupInfoMaximum = `UPDATE dtalk_group_info SET group_maximum=?, group_update_time=? 
							WHERE group_id=?`
	_UpdateGroupInfoIntroduce = `UPDATE dtalk_group_info SET group_introduce=?, group_update_time=? 
							WHERE group_id=?`
	_UpdateGroupInfoStatus = `UPDATE dtalk_group_info SET group_status=?, group_update_time=? 
							WHERE group_id=?`
	_UpdateGroupInfoOwnerId = `UPDATE dtalk_group_info SET group_owner_id=?, group_update_time=? 
							WHERE group_id=?`
	_UpdateGroupInfoJoinType = `UPDATE dtalk_group_info SET group_join_type=?, group_update_time=? 
							WHERE group_id=?`
	_UpdateGroupInfoMuteType = `UPDATE dtalk_group_info SET group_mute_type=?, group_update_time=? 
							WHERE group_id=?`
	_UpdateGroupInfoFriendType = `UPDATE dtalk_group_info SET group_friend_type=?, group_update_time=? 
							WHERE group_id=?`
	_GetGroupInfoByGroupId     = `SELECT * FROM dtalk_group_info WHERE group_id=?`
	_GetGroupInfosByGroupIds   = `SELECT * FROM dtalk_group_info WHERE group_id IN (?)`
	_GetGroupInfoByGroupMarkId = `SELECT * FROM dtalk_group_info WHERE group_mark_id=?`
	_GetAllGroupInfo           = `SELECT * FROM dtalk_group_info`
	_MaintainAESKeyAndPubName  = `UPDATE dtalk_group_info SET group_aes_key=?, group_pub_name=? WHERE group_id=?`
	_MaintainGroupType         = `UPDATE dtalk_group_info SET group_type=?, group_join_type=? WHERE group_id=?`

	// dtalk_group_member
	_InsertGroupMember = `INSERT INTO dtalk_group_member ( group_id, group_member_id, group_member_name, group_member_type,
						group_member_join_time, group_member_update_time) VALUES ( ?, ?, ?, ?, ?, ? ) 
						ON DUPLICATE KEY UPDATE group_member_type=?, group_member_join_time=?,
						group_member_update_time=?`
	_InsertGroupMembersPrefix = `INSERT INTO dtalk_group_member ( group_id, group_member_id, group_member_name, group_member_type,
						group_member_join_time, group_member_update_time) VALUES `
	_InsertGroupMembersSuffix = `ON DUPLICATE KEY UPDATE group_member_name='', group_member_type=?, group_member_join_time=?,
						group_member_update_time=?`
	_UpdateGroupMemberName = `UPDATE dtalk_group_member SET group_member_name=?, group_member_update_time=? 
							WHERE group_id=? AND group_member_id=?`
	_UpdateGroupMemberType = `UPDATE dtalk_group_member SET group_member_type=?, group_member_update_time=? 
							WHERE group_id=? AND group_member_id=?`
	_GetGroupIdsByMemberId             = `SELECT group_id FROM dtalk_group_member WHERE group_member_id=? AND group_member_type<?`
	_GetMemberInfoByMemberIdAndGroupId = `SELECT * FROM dtalk_group_member WHERE group_id=? AND group_member_id=? AND group_member_type<?`
	_GetMembersByGroupIdWithLimit      = `
SELECT 
mem.group_id as group_id, 
mem.group_member_id as group_member_id, 
mem.group_member_type as group_member_type, 
mem.group_member_name as group_member_name, 
mem.group_member_join_time as group_member_join_time,
mute.group_member_mute_time as group_member_mute_time
From dtalk_group_member AS mem LEFT JOIN dtalk_group_member_mute AS mute 
ON mem.group_id=mute.group_id AND mem.group_member_id=mute.group_member_id
WHERE mem.group_id=? AND mem.group_member_type<? ORDER BY group_member_type desc, group_member_join_time ASC limit ?, ?
`
	_GetMembersByGroupId = `
SELECT 
mem.group_id as group_id, 
mem.group_member_id as group_member_id, 
mem.group_member_type as group_member_type, 
mem.group_member_name as group_member_name, 
mem.group_member_join_time as group_member_join_time,
mute.group_member_mute_time as group_member_mute_time
From dtalk_group_member AS mem LEFT JOIN dtalk_group_member_mute AS mute 
ON mem.group_id=mute.group_id AND mem.group_member_id=mute.group_member_id
WHERE mem.group_id=? AND mem.group_member_type<? ORDER BY group_member_type desc, group_member_join_time ASC
`

	_GetMemberTypeByMemberIdAndGroupId = `SELECT group_member_type FROM dtalk_group_member WHERE group_id=? AND group_member_id=?`
	_GetAdminNumByGroupId              = `SELECT count(*) FROM dtalk_group_member WHERE group_id=? AND group_member_type=?`
	_GetMemberNumByGroupId             = `SELECT count(*) FROM dtalk_group_member WHERE group_id=? AND group_member_type<?`

	// dtalk_group_member_mute
	_UpdateGroupMemberMuteTime = `REPLACE INTO dtalk_group_member_mute (group_id, group_member_id, group_member_mute_time, group_member_mute_update_time) VALUES`
	_GetGroupMemberMuteTime    = `SELECT group_member_mute_time FROM dtalk_group_member_mute WHERE group_id=? AND group_member_id=?`
	_GetGroupMuteNum           = `SELECT count(*) From dtalk_group_member AS mem LEFT JOIN dtalk_group_member_mute AS mute 
ON mem.group_id=mute.group_id AND mem.group_member_id=mute.group_member_id 
WHERE mem.group_id=? AND mem.group_member_type<? AND mute.group_member_mute_time>?`
	_GetGroupMembersMuted = `SELECT 
mem.group_id as group_id, 
mem.group_member_id as group_member_id, 
mem.group_member_type as group_member_type, 
mem.group_member_name as group_member_name, 
mute.group_member_mute_time as group_member_mute_time
From dtalk_group_member AS mem LEFT JOIN dtalk_group_member_mute AS mute 
ON mem.group_id=mute.group_id AND mem.group_member_id=mute.group_member_id 
WHERE mem.group_id=? AND mem.group_member_type<? AND mute.group_member_mute_time>?`
	_GetGroupMemberWithMuteTime = `SELECT 
mem.group_id as group_id, 
mem.group_member_id as group_member_id, 
mem.group_member_type as group_member_type, 
mem.group_member_name as group_member_name, 
mem.group_member_join_time as group_member_join_time,
mute.group_member_mute_time as group_member_mute_time
From dtalk_group_member AS mem LEFT JOIN dtalk_group_member_mute AS mute 
ON mem.group_id=mute.group_id AND mem.group_member_id=mute.group_member_id 
WHERE mem.group_id=? AND mem.group_member_id=? AND mem.group_member_type<?`
)

// dtalk_group_info

func (d *Dao) InsertGroupInfo(tx *mysql.MysqlTx, form *db.GroupInfo) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := tx.Exec(_InsertGroupInfo, form.GroupId, form.GroupMarkId, form.GroupName, form.GroupAvatar, form.GroupMemberNum, form.GroupMaximum,
		form.GroupIntroduce, form.GroupStatus, form.GroupOwnerId, nowTime, nowTime,
		form.GroupJoinType, form.GroupMuteType, form.GroupFriendType, form.GroupAESKey, form.GroupPubName, form.GroupType)
	return num, lastId, err
}

func (d *Dao) updateGroupInfoName(form *db.GroupInfo) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := d.conn.Exec(_UpdateGroupInfoName, form.GroupName, form.GroupPubName, nowTime, form.GroupId)
	return num, lastId, err
}

func (d *Dao) updateGroupInfoAvatar(form *db.GroupInfo) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := d.conn.Exec(_UpdateGroupInfoAvatar, form.GroupAvatar, nowTime, form.GroupId)
	return num, lastId, err
}

func (d *Dao) updateGroupInfoMemberNum(form *db.GroupInfo) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := d.conn.Exec(_UpdateGroupInfoGroupNum, form.GroupMemberNum, nowTime, form.GroupId)
	return num, lastId, err
}

// UpdateGroupInfoMaximum .
// no usage
func (d *Dao) UpdateGroupInfoMaximum(form *db.GroupInfo) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := d.conn.Exec(_UpdateGroupInfoMaximum, form.GroupMaximum, nowTime, form.GroupId)
	return num, lastId, err
}

// UpdateGroupInfoIntroduce .
// no usage
func (d *Dao) UpdateGroupInfoIntroduce(form *db.GroupInfo) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := d.conn.Exec(_UpdateGroupInfoIntroduce, form.GroupIntroduce, nowTime, form.GroupId)
	return num, lastId, err
}

// UpdateGroupInfoStatus .
// no usage
func (d *Dao) UpdateGroupInfoStatus(form *db.GroupInfo) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := d.conn.Exec(_UpdateGroupInfoStatus, form.GroupStatus, nowTime, form.GroupId)
	return num, lastId, err
}

func (d *Dao) UpdateGroupInfoStatusWithTx(tx *mysql.MysqlTx, form *db.GroupInfo) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := tx.Exec(_UpdateGroupInfoStatus, form.GroupStatus, nowTime, form.GroupId)
	_ = d.DeleteGroupCache(form.GroupId)
	return num, lastId, err
}

func (d *Dao) UpdateGroupInfoOwnerIdWithTx(tx *mysql.MysqlTx, form *db.GroupInfo) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := tx.Exec(_UpdateGroupInfoOwnerId, form.GroupOwnerId, nowTime, form.GroupId)
	_ = d.DeleteGroupCache(form.GroupId)
	return num, lastId, err
}

func (d *Dao) updateGroupInfoJoinType(form *db.GroupInfo) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := d.conn.Exec(_UpdateGroupInfoJoinType, form.GroupJoinType, nowTime, form.GroupId)
	return num, lastId, err
}

func (d *Dao) updateGroupInfoMuteType(form *db.GroupInfo) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := d.conn.Exec(_UpdateGroupInfoMuteType, form.GroupMuteType, nowTime, form.GroupId)
	return num, lastId, err
}

func (d *Dao) updateGroupInfoFriendType(form *db.GroupInfo) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := d.conn.Exec(_UpdateGroupInfoFriendType, form.GroupFriendType, nowTime, form.GroupId)
	return num, lastId, err
}

func (d *Dao) getGroupInfoByGroupMarkId(groupMarkId string) (*db.GroupInfo, error) {
	maps, err := d.conn.Query(_GetGroupInfoByGroupMarkId, groupMarkId)
	if err != nil {
		return nil, err
	}
	if len(maps) == 0 {
		return nil, model.ErrRecordNotExist
	}
	res := maps[0]
	return db.ConvertGroupInfo(res), nil
}

func (d *Dao) getGroupInfoByGroupId(groupId int64) (*db.GroupInfo, error) {
	maps, err := d.conn.Query(_GetGroupInfoByGroupId, groupId)
	if err != nil {
		return nil, err
	}
	if len(maps) == 0 {
		return nil, model.ErrRecordNotExist
	}
	res := maps[0]
	return db.ConvertGroupInfo(res), nil
}

func (d *Dao) getGroupInfosByGroupIds(groupIds []int64) ([]*db.GroupInfo, error) {
	if len(groupIds) == 0 {
		return nil, nil
	}

	groupIdsSQL := ""
	for _, groupId := range groupIds {
		groupIdsSQL += util.ToString(groupId) + ","
	}
	groupIdsSQL = strings.TrimSuffix(groupIdsSQL, ",")

	maps, err := d.conn.Query(_GetGroupInfosByGroupIds, groupIdsSQL)
	if err != nil {
		return nil, err
	}

	return db.ConvertGroupInfos(maps), nil
}

func (d *Dao) getAllGroupInfo() ([]*db.GroupInfo, error) {
	maps, err := d.conn.Query(_GetAllGroupInfo)
	if err != nil {
		return nil, err
	}
	if len(maps) == 0 {
		return nil, model.ErrRecordNotExist
	}
	res := make([]*db.GroupInfo, len(maps))
	for i := range maps {
		res[i] = db.ConvertGroupInfo(maps[i])
	}
	return res, nil
}

func (d *Dao) MaintainAESKeyAndPubName(form *db.GroupInfo) (int64, int64, error) {
	num, lastId, err := d.conn.Exec(_MaintainAESKeyAndPubName, form.GroupAESKey, form.GroupPubName, form.GroupId)
	_ = d.DeleteGroupCache(form.GroupId)
	return num, lastId, err
}

func (d *Dao) MaintainGroupType(form *db.GroupInfo) (int64, int64, error) {
	num, lastId, err := d.conn.Exec(_MaintainGroupType, form.GroupType, biz.GroupJoinTypeAdmin, form.GroupId)
	_ = d.DeleteGroupCache(form.GroupId)
	return num, lastId, err
}

// dtalk_group_member

// InsertGroupMember .
// no usage
func (d *Dao) InsertGroupMember(tx *mysql.MysqlTx, form *db.GroupMember) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := tx.Exec(_InsertGroupMember, form.GroupId, form.GroupMemberId, form.GroupMemberName,
		form.GroupMemberType, nowTime, nowTime, form.GroupMemberType,
		nowTime, nowTime)
	return num, lastId, err
}

func (d *Dao) insertGroupMembers(tx *mysql.MysqlTx, groupMembers []*db.GroupMember) (int64, int64, error) {
	nowTime := d.getNowTime()
	var vals []interface{}
	valSql := ""
	for i, groupMember := range groupMembers {
		if i == 0 {
			valSql += "(?,?,?,?,?,?)"
		} else {
			valSql += ",(?,?,?,?,?,?)"
		}
		vals = append(vals, groupMember.GroupId, groupMember.GroupMemberId, groupMember.GroupMemberName, groupMember.GroupMemberType,
			nowTime, nowTime)
	}
	SQL := _InsertGroupMembersPrefix + valSql + _InsertGroupMembersSuffix
	vals = append(vals, biz.GroupMemberTypeNormal, nowTime, nowTime)

	num, lastId, err := tx.Exec(SQL, vals...)
	return num, lastId, err
}

func (d *Dao) updateGroupMemberName(form *db.GroupMember) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := d.conn.Exec(_UpdateGroupMemberName, form.GroupMemberName, nowTime, form.GroupId, form.GroupMemberId)
	return num, lastId, err
}

func (d *Dao) updateGroupMemberType(form *db.GroupMember) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := d.conn.Exec(_UpdateGroupMemberType, form.GroupMemberType, nowTime, form.GroupId, form.GroupMemberId)
	return num, lastId, err
}

func (d *Dao) UpdateGroupMemberTypeWithTx(tx *mysql.MysqlTx, form *db.GroupMember) (int64, int64, error) {
	nowTime := d.getNowTime()
	num, lastId, err := tx.Exec(_UpdateGroupMemberType, form.GroupMemberType, nowTime, form.GroupId, form.GroupMemberId)
	_ = d.DeleteGroupCache(form.GroupId)
	return num, lastId, err
}

func (d *Dao) getGroupIdsByMemberId(memberId string) ([]int64, error) {
	maps, err := d.conn.Query(_GetGroupIdsByMemberId, memberId, biz.GroupMemberTypeOther)
	if err != nil {
		return nil, err
	}

	ret := make([]int64, len(maps))
	for i, m := range maps {
		ret[i] = util.ToInt64(m["group_id"])
	}
	return ret, err
}

func (d *Dao) getMemberInfoByMemberIdAndGroupId(groupMemberId string, groupId int64) (*db.GroupMember, error) {
	maps, err := d.conn.Query(_GetMemberInfoByMemberIdAndGroupId, groupId, groupMemberId, biz.GroupMemberTypeOther)
	if err != nil {
		return nil, err
	}
	if len(maps) == 0 {
		return nil, model.ErrRecordNotExist
	}
	res := maps[0]
	return db.ConvertGroupMember(res), nil
}

func (d *Dao) getMembersByGroupId(groupId int64) ([]*db.GroupMemberWithMute, error) {
	maps, err := d.conn.Query(_GetMembersByGroupId, groupId, biz.GroupMemberTypeOther)
	if err != nil {
		return nil, err
	}
	if len(maps) == 0 {
		return nil, nil
	}
	res := make([]*db.GroupMemberWithMute, len(maps))

	for i := range maps {
		res[i] = db.ConvertGroupMemberWithMute(maps[i])
	}
	return res, nil
}

func (d *Dao) getMembersByGroupIdWithLimit(groupId, n, m int64) ([]*db.GroupMemberWithMute, error) {
	maps, err := d.conn.Query(_GetMembersByGroupIdWithLimit, groupId, biz.GroupMemberTypeOther, n, m)
	if err != nil {
		return nil, err
	}
	if len(maps) == 0 {
		return nil, nil
	}
	res := make([]*db.GroupMemberWithMute, len(maps))

	for i := range maps {
		res[i] = db.ConvertGroupMemberWithMute(maps[i])
	}
	return res, nil
}

func (d *Dao) GetMemberTypeMemberIdAndGroupId(groupMemberId string, groupId int64) (int32, error) {
	maps, err := d.conn.Query(_GetMemberTypeByMemberIdAndGroupId, groupId, groupMemberId)
	if err != nil {
		return -1, err
	}
	if len(maps) == 0 {
		return biz.GroupMemberTypeOther, nil
	}
	return util.ToInt32(maps[0]["group_member_type"]), nil
}

func (d *Dao) getAdminNumByGroupId(groupId int64) (int32, error) {
	res, err := d.conn.Query(_GetAdminNumByGroupId, groupId, biz.GroupMemberTypeAdmin)
	if err != nil {
		return 0, err
	}
	return util.ToInt32(res[0]["count(*)"]), nil
}

func (d *Dao) getMemberNumByGroupId(groupId int64) (int32, error) {
	res, err := d.conn.Query(_GetMemberNumByGroupId, groupId, biz.GroupMemberTypeOther)
	if err != nil {
		return 0, err
	}
	return util.ToInt32(res[0]["count(*)"]), nil
}

// dtalk_group_member_mute
func (d *Dao) updateGroupMemberMuteTimes(tx *mysql.MysqlTx, groupMemberMutes []*db.GroupMemberMute) (int64, int64, error) {
	nowTime := d.getNowTime()
	var vals []interface{}
	valSql := ""
	for i, groupMemberMute := range groupMemberMutes {
		if i == 0 {
			valSql += "(?,?,?,?)"
		} else {
			valSql += ",(?,?,?,?)"
		}
		vals = append(vals, groupMemberMute.GroupId, groupMemberMute.GroupMemberId, groupMemberMute.GroupMemberMuteTime, nowTime)
	}
	SQL := _UpdateGroupMemberMuteTime + valSql

	num, lastId, err := tx.Exec(SQL, vals...)
	return num, lastId, err
}

func (d *Dao) getGroupMemberMuteTime(groupId int64, memberId string) (int64, error) {
	maps, err := d.conn.Query(_GetGroupMemberMuteTime, groupId, memberId)
	if err != nil {
		return 0, err
	}
	if len(maps) == 0 {
		return 0, nil
	}
	return util.ToInt64(maps[0]["group_member_mute_time"]), nil
}

func (d *Dao) getGroupMuteNum(groupId int64, nowTime int64) (int32, error) {
	res, err := d.conn.Query(_GetGroupMuteNum, groupId, biz.GroupMemberTypeOther, nowTime)
	if err != nil {
		return 0, err
	}
	return util.ToInt32(res[0]["count(*)"]), nil
}

func (d *Dao) getGroupMuteList(groupId int64, nowTime int64) ([]*db.GroupMemberWithMute, error) {
	maps, err := d.conn.Query(_GetGroupMembersMuted, groupId, biz.GroupMemberTypeOther, nowTime)
	if err != nil {
		return nil, err
	}
	if len(maps) == 0 {
		return nil, nil
	}
	res := make([]*db.GroupMemberWithMute, len(maps))

	for i := range maps {
		res[i] = db.ConvertGroupMemberWithMute(maps[i])
	}
	return res, nil
}

func (d *Dao) getGroupMemberWithMuteTime(groupId int64, memberId string) (*db.GroupMemberWithMute, error) {
	maps, err := d.conn.Query(_GetGroupMemberWithMuteTime, groupId, memberId, biz.GroupMemberTypeOther)
	if err != nil {
		return nil, err
	}
	if len(maps) == 0 {
		return nil, model.ErrRecordNotExist
	}
	res := db.ConvertGroupMemberWithMute(maps[0])

	return res, nil
}
