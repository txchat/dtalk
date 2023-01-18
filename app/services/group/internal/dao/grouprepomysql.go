package dao

import (
	"github.com/txchat/dtalk/app/services/group/internal/model"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/pkg/util"
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

type GroupRepositoryMysql struct {
	conn *mysql.MysqlConn
}

func NewGroupRepositoryMysql(conn *mysql.MysqlConn) *GroupRepositoryMysql {
	return &GroupRepositoryMysql{
		conn: conn,
	}
}

func (repo *GroupRepositoryMysql) NewTx() (*mysql.MysqlTx, error) {
	return repo.conn.NewTx()
}

func (repo *GroupRepositoryMysql) InsertGroupInfo(tx *mysql.MysqlTx, g *model.GroupInfo) (int64, int64, error) {
	num, lastId, err := tx.Exec(_InsertGroupInfo, g.GroupId, g.GroupMarkId, g.GroupName, g.GroupAvatar, g.GroupMemberNum, g.GroupMaximum,
		g.GroupIntroduce, g.GroupStatus, g.GroupOwnerId, g.GroupCreateTime, g.GroupUpdateTime,
		g.GroupJoinType, g.GroupMuteType, g.GroupFriendType, g.GroupAESKey, g.GroupPubName, g.GroupType)
	return num, lastId, err
}

// InsertGroupMembers 批量插入群成员
func (repo *GroupRepositoryMysql) InsertGroupMembers(tx *mysql.MysqlTx, members []*model.GroupMember, updateTime int64) (int64, int64, error) {
	var params []interface{}
	cases := ""
	if len(members) < 1 {
		return 0, 0, nil
	}
	for i, m := range members {
		if i == 0 {
			cases += "(?,?,?,?,?,?)"
		} else {
			cases += ",(?,?,?,?,?,?)"
		}
		params = append(params, m.GroupId, m.GroupMemberId, m.GroupMemberName, m.GroupMemberType, m.GroupMemberJoinTime, m.GroupMemberUpdateTime)
	}
	SQL := _InsertGroupMembersPrefix + cases + _InsertGroupMembersSuffix
	params = append(params, model.GroupMemberTypeNormal, updateTime, updateTime)

	num, lastId, err := tx.Exec(SQL, params...)
	return num, lastId, err
}

func (repo *GroupRepositoryMysql) GetGroupById(gid int64) (*model.GroupInfo, error) {
	records, err := repo.conn.Query(_GetGroupInfoByGroupId, gid)
	if err != nil {
		return nil, err
	}
	if len(records) < 1 {
		return nil, xerror.ErrGroupNotExist
	}
	record := records[0]
	return model.ConvertGroupInfo(record), nil
}

func (repo *GroupRepositoryMysql) GetGroupByMarkId(markId string) (*model.GroupInfo, error) {
	records, err := repo.conn.Query(_GetGroupInfoByGroupMarkId, markId)
	if err != nil {
		return nil, err
	}
	if len(records) < 1 {
		return nil, xerror.ErrGroupNotExist
	}
	record := records[0]
	return model.ConvertGroupInfo(record), nil
}

func (repo *GroupRepositoryMysql) GetGroupMutedNumbers(gid int64, now int64) (int32, error) {
	records, err := repo.conn.Query(_GetGroupMuteNum, gid, model.GroupMemberTypeOther, now)
	if err != nil {
		return 0, err
	}
	return util.MustToInt32(records[0]["count(*)"]), nil
}

func (repo *GroupRepositoryMysql) GetGroupManagerNumbers(gid int64) (int32, error) {
	records, err := repo.conn.Query(_GetAdminNumByGroupId, gid, model.GroupMemberTypeManager)
	if err != nil {
		return 0, err
	}
	return util.MustToInt32(records[0]["count(*)"]), nil
}

func (repo *GroupRepositoryMysql) GetMemberById(gid int64, mid string) (*model.GroupMember, error) {
	maps, err := repo.conn.Query(_GetGroupMemberWithMuteTime, gid, mid, model.GroupMemberTypeOther)
	if err != nil {
		return nil, err
	}
	if len(maps) == 0 {
		return nil, xerror.ErrGroupMemberNotExist
	}
	return model.ConvertGroupMember(maps[0]), nil
}

func (repo *GroupRepositoryMysql) GetUnLimitedMembers(gid int64) ([]*model.GroupMember, error) {
	maps, err := repo.conn.Query(_GetMembersByGroupId, gid, model.GroupMemberTypeOther)
	if err != nil {
		return nil, err
	}

	members := make([]*model.GroupMember, 0, len(maps))
	for _, m := range maps {
		members = append(members, model.ConvertGroupMember(m))
	}
	return members, nil
}

func (repo *GroupRepositoryMysql) GetLimitedMembers(gid, start, num int64) ([]*model.GroupMember, error) {
	maps, err := repo.conn.Query(_GetMembersByGroupIdWithLimit, gid, model.GroupMemberTypeOther, start, num)
	if err != nil {
		return nil, err
	}

	members := make([]*model.GroupMember, 0, len(maps))
	for _, m := range maps {
		members = append(members, model.ConvertGroupMember(m))
	}
	return members, nil
}

func (repo *GroupRepositoryMysql) GetMutedMembers(gid, time int64) ([]*model.GroupMember, error) {
	maps, err := repo.conn.Query(_GetGroupMembersMuted, gid, model.GroupMemberTypeOther, time)
	if err != nil {
		return nil, err
	}
	res := make([]*model.GroupMember, 0, len(maps))
	for _, m := range maps {
		res = append(res, model.ConvertGroupMember(m))
	}
	return res, nil
}

func (repo *GroupRepositoryMysql) JoinedGroups(uid string) ([]int64, error) {
	maps, err := repo.conn.Query(_GetGroupIdsByMemberId, uid, model.GroupMemberTypeOther)
	if err != nil {
		return nil, err
	}
	ret := make([]int64, len(maps))
	for i, m := range maps {
		ret[i] = util.MustToInt64(m["group_id"])
	}
	return ret, nil
}

func (repo *GroupRepositoryMysql) UpdateGroupStatus(tx *mysql.MysqlTx, group *model.GroupInfo) (int64, int64, error) {
	return tx.Exec(_UpdateGroupInfoStatus, group.GroupStatus, group.GroupUpdateTime, group.GroupId)
}

func (repo *GroupRepositoryMysql) UpdateGroupMemberRole(tx *mysql.MysqlTx, member *model.GroupMember) (int64, int64, error) {
	return tx.Exec(_UpdateGroupMemberType, member.GroupMemberType, member.GroupMemberUpdateTime, member.GroupId, member.GroupMemberId)
}

func (repo *GroupRepositoryMysql) UpdateGroupMembersMuteTime(tx *mysql.MysqlTx, members []*model.GroupMember) (int64, int64, error) {
	var params []interface{}
	cases := ""
	if len(members) < 1 {
		return 0, 0, nil
	}
	for i, m := range members {
		if i == 0 {
			cases += "(?,?,?,?)"
		} else {
			cases += ",(?,?,?,?)"
		}
		params = append(params, m.GroupId, m.GroupMemberId, m.GroupMemberMuteTime, m.GroupMemberMuteUpdateTime)
	}
	SQL := _UpdateGroupMemberMuteTime + cases

	num, lastId, err := tx.Exec(SQL, params...)
	return num, lastId, err
}

func (repo *GroupRepositoryMysql) UpdateGroupOwner(tx *mysql.MysqlTx, group *model.GroupInfo) (int64, int64, error) {
	return tx.Exec(_UpdateGroupInfoOwnerId, group.GroupOwnerId, group.GroupUpdateTime, group.GroupId)
}

func (repo *GroupRepositoryMysql) UpdateGroupAvatar(group *model.GroupInfo) (int64, int64, error) {
	return repo.conn.Exec(_UpdateGroupInfoAvatar, group.GroupAvatar, group.GroupUpdateTime, group.GroupId)
}

func (repo *GroupRepositoryMysql) UpdateGroupFriendlyType(group *model.GroupInfo) (int64, int64, error) {
	return repo.conn.Exec(_UpdateGroupInfoFriendType, group.GroupFriendType, group.GroupUpdateTime, group.GroupId)
}

func (repo *GroupRepositoryMysql) UpdateGroupJoinType(group *model.GroupInfo) (int64, int64, error) {
	return repo.conn.Exec(_UpdateGroupInfoJoinType, group.GroupJoinType, group.GroupUpdateTime, group.GroupId)
}

func (repo *GroupRepositoryMysql) UpdateGroupMuteType(group *model.GroupInfo) (int64, int64, error) {
	return repo.conn.Exec(_UpdateGroupInfoMuteType, group.GroupMuteType, group.GroupUpdateTime, group.GroupId)
}

func (repo *GroupRepositoryMysql) UpdateGroupName(group *model.GroupInfo) (int64, int64, error) {
	return repo.conn.Exec(_UpdateGroupInfoName, group.GroupName, group.GroupPubName, group.GroupUpdateTime, group.GroupId)
}

func (repo *GroupRepositoryMysql) UpdateGroupMembersNumber(tx *mysql.MysqlTx, group *model.GroupInfo) (int64, int64, error) {
	return tx.Exec(_UpdateGroupInfoGroupNum, group.GroupMemberNum, group.GroupUpdateTime, group.GroupId)
}

func (repo *GroupRepositoryMysql) UpdateGroupMemberName(member *model.GroupMember) (int64, int64, error) {
	return repo.conn.Exec(_UpdateGroupMemberName, member.GroupMemberName, member.GroupMemberUpdateTime, member.GroupId, member.GroupMemberId)
}
