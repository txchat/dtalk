package dao

import (
	"context"

	"github.com/pkg/errors"
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
	"github.com/txchat/dtalk/pkg/util"
)

// 数据读写层，数据库和缓存全部在这层统一处理，包括 cache miss 处理。

// GetAllGroupInfo .
func (d *Dao) GetAllGroupInfo() ([]*biz.GroupInfo, error) {
	groupPos, err := d.getAllGroupInfo()
	if err != nil {
		return nil, err
	}
	groups := make([]*biz.GroupInfo, 0, len(groupPos))
	for i := range groupPos {
		groups = append(groups, groupPos[i].ToBiz())
	}
	return groups, nil
}

// GetGroupInfoByGroupMarkId 查询群信息
func (d *Dao) GetGroupInfoByGroupMarkId(groupMarkId string) (*biz.GroupInfo, error) {
	// redis get
	groupInfo, err := d.getGroupInfoByGroupMarkId(groupMarkId)
	if err != nil {
		return nil, err
	}
	// redis put
	return groupInfo.ToBiz(), nil
}

// GetGroupInfoByGroupId 查询群信息
func (d *Dao) GetGroupInfoByGroupId(ctx context.Context, groupId int64) (*biz.GroupInfo, error) {
	log := d.GetLogWithTrace(ctx)
	// 加缓存 get
	group, err := d.GetGroupCache(groupId)
	if err != nil {
		//log.Warn().Err(err).Int64("groupId", groupId).Msg("GetGroupCache err")
	} else {
		return group, nil
	}

	groupInfo, err := d.getGroupInfoByGroupId(groupId)
	if err != nil {
		return nil, err
	}
	res := groupInfo.ToBiz()

	muteNum, err := d.GetGroupMuteNum(groupId)
	if err != nil {
		return nil, err
	}
	res.MuteNum = muteNum

	adminNum, err := d.getAdminNumByGroupId(groupId)
	if err != nil {
		return nil, err
	}
	res.AdminNum = adminNum

	// 加缓存 set 设置固定值+随机的过期时间
	expire := d.GetRandRedisExpire()
	err = d.SaveGroup(res, expire)
	if err != nil {
		log.Warn().Err(err).Int64("groupId", groupId).Msg("SaveGroup err")
	}
	return res, nil
}

// GetGroupInfosByGroupIds 查询群信息s
func (d *Dao) GetGroupInfosByGroupIds(ctx context.Context, groupIds []int64) ([]*biz.GroupInfo, error) {

	groupInfos, err := d.getGroupInfosByGroupIds(groupIds)
	if err != nil {
		return nil, err
	}

	ress := make([]*biz.GroupInfo, 0, len(groupInfos))
	for _, groupInfo := range groupInfos {
		ress = append(ress, groupInfo.ToBiz())
	}

	return ress, nil
}

// GetGroupIdsByMemberId 查询用户所有加入的群 ID
func (d *Dao) GetGroupIdsByMemberId(memberId string) ([]int64, error) {
	groupIds, err := d.getGroupIdsByMemberId(memberId)
	if err != nil {
		return nil, errors.WithMessagef(err, "getGroupIdsByMemberId memberId=%s", memberId)
	}
	return groupIds, nil
}

// GetMembersByGroupId 查询群里的所有成员信息
func (d *Dao) GetMembersByGroupId(groupId int64) ([]*biz.GroupMember, error) {
	res, err := d.getMembersByGroupId(groupId)
	if err != nil {
		return nil, err
	}

	groupMembers := make([]*biz.GroupMember, 0)
	for _, groupMember := range res {
		groupMembers = append(groupMembers, groupMember.ToBiz())
	}
	return groupMembers, nil
}

// GetGroupMembersByGroupIdWithLimit 查询群内前 n 个群成员信息
func (d *Dao) GetGroupMembersByGroupIdWithLimit(groupId, n, m int64) ([]*biz.GroupMember, error) {
	members, err := d.getMembersByGroupIdWithLimit(groupId, n, m)
	if err != nil {
		return nil, err
	}

	groupMembers := make([]*biz.GroupMember, 0)
	for _, groupMember := range members {
		groupMembers = append(groupMembers, groupMember.ToBiz())
	}
	return groupMembers, nil
}

// UpdateGroupInfoMemberNum 更新群成员数量
func (d *Dao) UpdateGroupInfoMemberNum(ctx context.Context, groupId int64) (int32, error) {
	log := d.GetLogWithTrace(ctx)
	members, err := d.GetMembersByGroupId(groupId)
	if err != nil {
		return 0, errors.WithMessagef(err, "GetMembersByGroupId, groupId=%d", groupId)
	}
	newNum := util.ToInt32(len(members))
	nowTime := d.getNowTime()

	groupInfo := &db.GroupInfo{
		GroupId:         groupId,
		GroupMemberNum:  newNum,
		GroupUpdateTime: nowTime,
	}
	_, _, err = d.updateGroupInfoMemberNum(groupInfo)
	if err != nil {
		return 0, errors.WithMessagef(err, "UpdateGroupInfoMemberNum, groupId=%d", groupId)
	}

	// 删 group 缓存
	err = d.DeleteGroupCache(groupId)
	if err != nil {
		log.Err(err).Int64("groupId", groupId).Msg("DeleteGroupCache err")
	}

	return newNum, nil
}

// UpdateGroupInfoName 更新群名称
func (d *Dao) UpdateGroupInfoName(ctx context.Context, groupId int64, name, publicName string) error {
	log := d.GetLogWithTrace(ctx)
	nowTime := d.getNowTime()
	groupInfo := &db.GroupInfo{
		GroupId:         groupId,
		GroupName:       name,
		GroupUpdateTime: nowTime,
		GroupPubName:    publicName,
	}
	if _, _, err := d.updateGroupInfoName(groupInfo); err != nil {
		return errors.WithMessagef(err, "UpdateGroupInfoName, groupInfo=%+v", groupInfo)
	}

	// 删 group 缓存
	err := d.DeleteGroupCache(groupId)
	if err != nil {
		log.Warn().Err(err).Int64("groupId", groupId).Msg("DeleteGroupCache err")
	}
	return nil
}

// UpdateGroupInfoAvatar 更新群头像
func (d *Dao) UpdateGroupInfoAvatar(ctx context.Context, groupId int64, avatar string) error {
	log := d.GetLogWithTrace(ctx)
	nowTime := d.getNowTime()
	groupInfo := &db.GroupInfo{
		GroupId:         groupId,
		GroupAvatar:     avatar,
		GroupUpdateTime: nowTime,
	}
	if _, _, err := d.updateGroupInfoAvatar(groupInfo); err != nil {
		return errors.WithMessagef(err, "updateGroupInfoAvatar, groupInfo=%+v", groupInfo)
	}

	// 删 group 缓存
	err := d.DeleteGroupCache(groupId)
	if err != nil {
		log.Warn().Err(err).Int64("groupId", groupId).Msg("DeleteGroupCache err")
	}
	return nil
}

// UpdateGroupInfoJoinType 更新加群设置
func (d *Dao) UpdateGroupInfoJoinType(ctx context.Context, groupId int64, joinType int32) error {
	log := d.GetLogWithTrace(ctx)
	nowTime := d.getNowTime()
	groupInfo := &db.GroupInfo{
		GroupId:         groupId,
		GroupJoinType:   joinType,
		GroupUpdateTime: nowTime,
	}
	if _, _, err := d.updateGroupInfoJoinType(groupInfo); err != nil {
		return errors.WithMessagef(err, "updateGroupInfoJoinType, groupInfo=%+v", groupInfo)
	}

	// 删 group 缓存
	err := d.DeleteGroupCache(groupId)
	if err != nil {
		log.Warn().Err(err).Int64("groupId", groupId).Msg("DeleteGroupCache err")
	}

	return nil
}

// UpdateGroupInfoFriendType 更新群内加好友设置
func (d *Dao) UpdateGroupInfoFriendType(ctx context.Context, groupId int64, friendType int32) error {
	log := d.GetLogWithTrace(ctx)
	nowTime := d.getNowTime()
	groupInfo := &db.GroupInfo{
		GroupId:         groupId,
		GroupFriendType: friendType,
		GroupUpdateTime: nowTime,
	}
	if _, _, err := d.updateGroupInfoFriendType(groupInfo); err != nil {
		return errors.WithMessagef(err, "UpdateGroupInfoFriendType, groupInfo=%+v", groupInfo)
	}

	// 删 group 缓存
	err := d.DeleteGroupCache(groupId)
	if err != nil {
		log.Warn().Err(err).Int64("groupId", groupId).Msg("DeleteGroupCache err")
	}

	return nil
}

// UpdateGroupInfoMuteType 更新群禁言设置
func (d *Dao) UpdateGroupInfoMuteType(ctx context.Context, groupId int64, muteType int32) error {
	log := d.GetLogWithTrace(ctx)
	nowTime := d.getNowTime()
	groupInfo := &db.GroupInfo{
		GroupId:         groupId,
		GroupMuteType:   muteType,
		GroupUpdateTime: nowTime,
	}
	if _, _, err := d.updateGroupInfoMuteType(groupInfo); err != nil {
		return errors.WithMessagef(err, "UpdateGroupInfoMuteType, groupInfo=%+v", groupInfo)
	}

	// 删 group 缓存
	err := d.DeleteGroupCache(groupId)
	if err != nil {
		log.Warn().Err(err).Int64("groupId", groupId).Msg("DeleteGroupCache err")
	}

	return nil
}

//UpdateGroupMemberName 更新群成员昵称
func (d *Dao) UpdateGroupMemberName(groupId int64, memberId, memberName string) error {
	nowTime := d.getNowTime()
	groupMember := &db.GroupMember{
		GroupId:               groupId,
		GroupMemberId:         memberId,
		GroupMemberName:       memberName,
		GroupMemberUpdateTime: nowTime,
	}
	if _, _, err := d.updateGroupMemberName(groupMember); err != nil {
		return errors.WithMessagef(err, "UpdateGroupMemberName, groupMember=%+v", groupMember)
	}
	return nil
}

// UpdateGroupMemberType 更新群成员类型
func (d *Dao) UpdateGroupMemberType(ctx context.Context, groupId int64, mnemberId string, memberType int32) error {
	log := d.GetLogWithTrace(ctx)

	nowTime := d.getNowTime()
	groupMember := &db.GroupMember{
		GroupId:               groupId,
		GroupMemberId:         mnemberId,
		GroupMemberType:       memberType,
		GroupMemberUpdateTime: nowTime,
	}
	if _, _, err := d.updateGroupMemberType(groupMember); err != nil {
		return errors.WithMessagef(err, "UpdateGroupMemberType, groupMember=%+v", groupMember)
	}

	err := d.DeleteGroupCache(groupId)
	if err != nil {
		log.Warn().Err(err).Int64("groupId", groupId).Msg("DeleteGroupCache err")
	}
	return nil
}

// GetAdminNumByGroupId 查询群内管理员数量
func (d *Dao) GetAdminNumByGroupId(groupId int64) (int32, error) {
	num, err := d.getAdminNumByGroupId(groupId)
	if err != nil {
		return 0, errors.WithMessagef(err, "GetAdminNumByGroupId, groupId=%d", groupId)
	}
	return num, nil
}

// InsertGroupMembers 批量插入群成员
func (d *Dao) InsertGroupMembers(tx *mysql.MysqlTx, groupMembers []*db.GroupMember) error {
	if len(groupMembers) == 0 {
		return nil
	}

	_, _, err := d.insertGroupMembers(tx, groupMembers)
	if err != nil {
		return errors.WithMessagef(err, "InsertGroupMembers, groupMembers=%v", groupMembers)
	}

	_ = d.DeleteGroupCache(groupMembers[0].GroupId)
	return nil
}

// GetGroupMemberMuteTime 查询群成员禁言时间
func (d *Dao) GetGroupMemberMuteTime(groupId int64, memberId string) (int64, error) {
	muteTime, err := d.getGroupMemberMuteTime(groupId, memberId)
	if err != nil {
		return 0, errors.WithMessagef(err, "GetGroupMemberMuteTime groupId=%d, memberId=%s", groupId, memberId)
	}
	return muteTime, nil
}

// GetGroupMuteNum 查询群内禁言人数
func (d *Dao) GetGroupMuteNum(groupId int64) (int32, error) {
	nowTime := d.getNowTime()
	muteNum, err := d.getGroupMuteNum(groupId, nowTime)
	if err != nil {
		return 0, errors.WithMessagef(err, "GetGroupMuteNum groupId=%d, nowTime=%d", groupId, nowTime)
	}
	return muteNum, nil
}

// GetGroupMemberByGroupIdAndMemberId 查询一个群成员信息
func (d *Dao) GetGroupMemberByGroupIdAndMemberId(ctx context.Context, groupId int64, memberId string) (*biz.GroupMember, error) {
	//log := d.GetLogWithTrace(ctx)
	//加缓存 get
	//groupMember, err := d.GetGroupMemberWithMuteTime(groupId, memberId)
	//if err != nil {
	//	log.Warn().Err(err).Int64("groupId", groupId).Str("memberId", memberId).Msg("GetGroupMemberWithMuteTime")
	//} else {
	//	return groupMember.ToBiz(), nil
	//}

	groupMember, err := d.getGroupMemberWithMuteTime(groupId, memberId)
	if err != nil {
		return nil, err
	}

	// 加缓存 set 设置固定值+随机的过期时间
	//expire := d.GetRandRedisExpire()
	//err = d.SaveGroupMemberWithMuteTime(groupMember, expire)
	//if err != nil {
	//	log.Warn().Err(err).Int64("groupId", groupId).Str("memberId", memberId).Msg("SaveGroupMemberWithMuteTime")
	//}
	return groupMember.ToBiz(), nil
}

// GetGroupMembersMutedByGroupId 查询群内被禁言的群成员信息
func (d *Dao) GetGroupMembersMutedByGroupId(groupId int64) ([]*biz.GroupMember, error) {
	nowTime := d.getNowTime()
	muteList, err := d.getGroupMuteList(groupId, nowTime)
	if err != nil {
		return nil, err
	}

	groupMembers := make([]*biz.GroupMember, 0)
	for _, groupMember := range muteList {
		groupMembers = append(groupMembers, groupMember.ToBiz())
	}
	return groupMembers, nil
}

// GetGroupMemberWithMuteTimeByGroupIdAndMemberId 查询一个带禁言时间的群成员信息
// 没用 删掉
func (d *Dao) GetGroupMemberWithMuteTimeByGroupIdAndMemberId(groupId int64, memberId string) (*biz.GroupMember, error) {
	member, err := d.getGroupMemberWithMuteTime(groupId, memberId)
	if err != nil {
		err = errors.WithMessagef(err, "GetGroupMemberWithMuteTimeByGroupIdAndMemberId GetGroupMemberWithMuteTime groupId=%d, memberId=%s", groupId, memberId)
		return nil, err
	}
	return member.ToBiz(), nil
}

func (d *Dao) UpdateGroupMemberMuteTimes(ctx context.Context, tx *mysql.MysqlTx, groupMemberMutes []*db.GroupMemberMute) error {
	if len(groupMemberMutes) == 0 {
		return nil
	}

	log := d.GetLogWithTrace(ctx)

	if _, _, err := d.updateGroupMemberMuteTimes(tx, groupMemberMutes); err != nil {
		return errors.WithMessagef(err, "UpdateGroupMemberMuteTimes, groupMemberMutes=%+v", groupMemberMutes)
	}

	// 删 group 缓存
	groupId := groupMemberMutes[0].GroupId
	err := d.DeleteGroupCache(groupId)
	if err != nil {
		log.Warn().Err(err).Int64("groupId", groupId).Msg("DeleteGroupCache err")
	}
	return nil
}
