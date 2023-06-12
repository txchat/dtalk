package dao

import (
	"database/sql/driver"
	"errors"

	"github.com/txchat/dtalk/app/services/group/internal/model"
	xerror "github.com/txchat/dtalk/pkg/error"
	xmysql "github.com/txchat/dtalk/pkg/mysql"
	"github.com/zeromicro/go-zero/core/service"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//nolint:deadcode,varcheck
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
	_GetAdminNumByGroupId              = `SELECT count(*) AS group_admin_num FROM dtalk_group_member WHERE group_id=? AND group_member_type=?`
	_GetMemberNumByGroupId             = `SELECT count(*) FROM dtalk_group_member WHERE group_id=? AND group_member_type<?`

	// dtalk_group_member_mute
	_UpdateGroupMemberMuteTime = `REPLACE INTO dtalk_group_member_mute (group_id, group_member_id, group_member_mute_time, group_member_mute_update_time) VALUES`
	_GetGroupMemberMuteTime    = `SELECT group_member_mute_time FROM dtalk_group_member_mute WHERE group_id=? AND group_member_id=?`
	_GetGroupMuteNum           = `SELECT count(*) AS group_mute_count From dtalk_group_member AS mem LEFT JOIN dtalk_group_member_mute AS mute 
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
mem.group_member_update_time as group_member_update_time,
mute.group_member_mute_time as group_member_mute_time
From dtalk_group_member AS mem LEFT JOIN dtalk_group_member_mute AS mute 
ON mem.group_id=mute.group_id AND mem.group_member_id=mute.group_member_id 
WHERE mem.group_id=? AND mem.group_member_id=? AND mem.group_member_type<?`
)

type GormTx struct {
	*gorm.DB
}

func (tx *GormTx) Commit() error {
	return tx.DB.Commit().Error
}

func (tx *GormTx) Rollback() error {
	return tx.DB.Rollback().Error
}

type GroupRepositoryMysql struct {
	conn *gorm.DB
}

func NewGroupRepositoryMysql(mode string, mysqlConfig xmysql.Config) *GroupRepositoryMysql {
	mysqlConfig.ParseTime = true
	mysqlConfig.SetParam("charset", "UTF8MB4")

	defaultLogger := logger.Default
	switch mode {
	case service.TestMode, service.DevMode, service.RtMode:
		defaultLogger.LogMode(logger.Info)
	case service.ProMode, service.PreMode:
		defaultLogger.LogMode(logger.Warn)
	}

	dsn := mysqlConfig.GetSQLDriverConfig().FormatDSN()
	db, err := gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &GroupRepositoryMysql{
		conn: db,
	}
}

func (repo *GroupRepositoryMysql) NewTx() (driver.Tx, error) {
	tx := repo.conn.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &GormTx{DB: tx}, nil
}

func (repo *GroupRepositoryMysql) InsertGroupInfo(tx driver.Tx, g *model.GroupInfo) (int64, int64, error) {
	txx, ok := tx.(*GormTx)
	if !ok {
		return 0, 0, model.ErrMysqlTxType
	}
	err := txx.Exec(_InsertGroupInfo, g.GroupId, g.GroupMarkId, g.GroupName, g.GroupAvatar, g.GroupMemberNum, g.GroupMaximum,
		g.GroupIntroduce, g.GroupStatus, g.GroupOwnerId, g.GroupCreateTime, g.GroupUpdateTime,
		g.GroupJoinType, g.GroupMuteType, g.GroupFriendType, g.GroupAESKey, g.GroupPubName, g.GroupType).Error
	return txx.RowsAffected, 0, err
}

// InsertGroupMembers 批量插入群成员
func (repo *GroupRepositoryMysql) InsertGroupMembers(tx driver.Tx, members []*model.GroupMember, updateTime int64) (int64, int64, error) {
	txx, ok := tx.(*GormTx)
	if !ok {
		return 0, 0, model.ErrMysqlTxType
	}

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

	err := txx.Exec(SQL, params...).Error
	return txx.RowsAffected, 0, err
}

func (repo *GroupRepositoryMysql) GetGroupById(gid int64) (*model.GroupInfo, error) {
	var group model.GroupInfo
	err := repo.conn.Raw(_GetGroupInfoByGroupId, gid).Take(&group).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.ErrGroupNotExist
		}
		return nil, err
	}
	return &group, nil
}

func (repo *GroupRepositoryMysql) GetGroupByMarkId(markId string) (*model.GroupInfo, error) {
	var group model.GroupInfo
	err := repo.conn.Raw(_GetGroupInfoByGroupMarkId, markId).Take(&group).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.ErrGroupNotExist
		}
		return nil, err
	}
	return &group, nil
}

func (repo *GroupRepositoryMysql) GetGroupMutedNumbers(gid int64, now int64) (int32, error) {
	var groupMuteCount int32
	err := repo.conn.Raw(_GetGroupMuteNum, gid, model.GroupMemberTypeOther, now).Scan(&groupMuteCount).Error
	if err != nil {
		return 0, err
	}
	return groupMuteCount, nil
}

func (repo *GroupRepositoryMysql) GetGroupManagerNumbers(gid int64) (int32, error) {
	var groupAdminNum int32
	err := repo.conn.Raw(_GetAdminNumByGroupId, gid, model.GroupMemberTypeManager).Scan(&groupAdminNum).Error
	if err != nil {
		return 0, err
	}
	return groupAdminNum, nil
}

func (repo *GroupRepositoryMysql) GetMemberById(gid int64, mid string) (*model.GroupMember, error) {
	var member model.GroupMember
	err := repo.conn.Raw(_GetGroupMemberWithMuteTime, gid, mid, model.GroupMemberTypeOther).Take(&member).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.ErrGroupMemberNotExist
		}
		return nil, err
	}
	return &member, nil
}

func (repo *GroupRepositoryMysql) GetUnLimitedMembers(gid int64) ([]*model.GroupMember, error) {
	var members []*model.GroupMember
	err := repo.conn.Raw(_GetMembersByGroupId, gid, model.GroupMemberTypeOther).Scan(&members).Error
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (repo *GroupRepositoryMysql) GetLimitedMembers(gid, start, num int64) ([]*model.GroupMember, error) {
	var members []*model.GroupMember
	err := repo.conn.Raw(_GetMembersByGroupIdWithLimit, gid, model.GroupMemberTypeOther, start, num).Scan(&members).Error
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (repo *GroupRepositoryMysql) GetMutedMembers(gid, time int64) ([]*model.GroupMember, error) {
	var members []*model.GroupMember
	err := repo.conn.Raw(_GetGroupMembersMuted, gid, model.GroupMemberTypeOther, time).Scan(&members).Error
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (repo *GroupRepositoryMysql) JoinedGroups(uid string) ([]int64, error) {
	var groupId []int64
	err := repo.conn.Raw(_GetGroupIdsByMemberId, uid, model.GroupMemberTypeOther).Scan(&groupId).Error
	if err != nil {
		return nil, err
	}
	return groupId, nil
}

func (repo *GroupRepositoryMysql) UpdateGroupStatus(tx driver.Tx, group *model.GroupInfo) (int64, int64, error) {
	txx, ok := tx.(*GormTx)
	if !ok {
		return 0, 0, model.ErrMysqlTxType
	}
	return 0, 0, txx.Exec(_UpdateGroupInfoStatus, group.GroupStatus, group.GroupUpdateTime, group.GroupId).Error
}

func (repo *GroupRepositoryMysql) UpdateGroupMemberRole(tx driver.Tx, member *model.GroupMember) (int64, int64, error) {
	txx, ok := tx.(*GormTx)
	if !ok {
		return 0, 0, model.ErrMysqlTxType
	}
	return 0, 0, txx.Exec(_UpdateGroupMemberType, member.GroupMemberType, member.GroupMemberUpdateTime, member.GroupId, member.GroupMemberId).Error
}

func (repo *GroupRepositoryMysql) UpdateGroupMembersMuteTime(tx driver.Tx, members []*model.GroupMember) (int64, int64, error) {
	txx, ok := tx.(*GormTx)
	if !ok {
		return 0, 0, model.ErrMysqlTxType
	}

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
	return 0, 0, txx.Exec(SQL, params...).Error
}

func (repo *GroupRepositoryMysql) UpdateGroupOwner(tx driver.Tx, group *model.GroupInfo) (int64, int64, error) {
	txx, ok := tx.(*GormTx)
	if !ok {
		return 0, 0, model.ErrMysqlTxType
	}
	return 0, 0, txx.Exec(_UpdateGroupInfoOwnerId, group.GroupOwnerId, group.GroupUpdateTime, group.GroupId).Error
}

func (repo *GroupRepositoryMysql) UpdateGroupAvatar(group *model.GroupInfo) (int64, int64, error) {
	return 0, 0, repo.conn.Exec(_UpdateGroupInfoAvatar, group.GroupAvatar, group.GroupUpdateTime, group.GroupId).Error
}

func (repo *GroupRepositoryMysql) UpdateGroupFriendlyType(group *model.GroupInfo) (int64, int64, error) {
	return 0, 0, repo.conn.Exec(_UpdateGroupInfoFriendType, group.GroupFriendType, group.GroupUpdateTime, group.GroupId).Error
}

func (repo *GroupRepositoryMysql) UpdateGroupJoinType(group *model.GroupInfo) (int64, int64, error) {
	return 0, 0, repo.conn.Exec(_UpdateGroupInfoJoinType, group.GroupJoinType, group.GroupUpdateTime, group.GroupId).Error
}

func (repo *GroupRepositoryMysql) UpdateGroupMuteType(group *model.GroupInfo) (int64, int64, error) {
	return 0, 0, repo.conn.Exec(_UpdateGroupInfoMuteType, group.GroupMuteType, group.GroupUpdateTime, group.GroupId).Error
}

func (repo *GroupRepositoryMysql) UpdateGroupName(group *model.GroupInfo) (int64, int64, error) {
	return 0, 0, repo.conn.Exec(_UpdateGroupInfoName, group.GroupName, group.GroupPubName, group.GroupUpdateTime, group.GroupId).Error
}

func (repo *GroupRepositoryMysql) UpdateGroupMembersNumber(tx driver.Tx, group *model.GroupInfo) (int64, int64, error) {
	txx, ok := tx.(*GormTx)
	if !ok {
		return 0, 0, model.ErrMysqlTxType
	}
	return 0, 0, txx.Exec(_UpdateGroupInfoGroupNum, group.GroupMemberNum, group.GroupUpdateTime, group.GroupId).Error
}

func (repo *GroupRepositoryMysql) UpdateGroupMemberName(member *model.GroupMember) (int64, int64, error) {
	return 0, 0, repo.conn.Exec(_UpdateGroupMemberName, member.GroupMemberName, member.GroupMemberUpdateTime, member.GroupId, member.GroupMemberId).Error
}
