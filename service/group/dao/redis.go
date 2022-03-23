package dao

import (
	"fmt"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
	"time"
)

const (
	_prefixGroup       = "x_group:%d"
	_prefixGroupMember = "x_group:%d;member:%s"
)

func keyGroup(groupId int64) string {
	return fmt.Sprintf(_prefixGroup, groupId)
}

// GetGroupCache get group from groupId
func (d *Dao) GetGroupCache(groupId int64) (*biz.GroupInfo, error) {
	key := keyGroup(groupId)
	if ok, err := d.redis.Exists(key); err != nil {
		return nil, err
	} else if !ok {
		return nil, db.ErrGroupNotExist
	}

	res := &biz.GroupInfo{}
	if err := d.redis.Read(key, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (d *Dao) SaveGroup(group *biz.GroupInfo, duration time.Duration) error {
	key := keyGroup(group.GroupId)
	if err := d.redis.Write(key, group, int(duration)); err != nil {
		return err
	}
	return nil
}

func (d *Dao) DeleteGroupCache(groupId int64) error {
	key := keyGroup(groupId)
	_, err := d.redis.Del(key)
	if err != nil {
		return err
	}

	return nil
}

//func keyGroupMember(groupId int64, memberId string) string {
//	return fmt.Sprintf(_prefixGroupMember, groupId, memberId)
//}
//
//func (d *Dao) GetGroupMemberWithMuteTime(groupId int64, memberId string) (*db.GroupMemberWithMute, error) {
//	key := keyGroupMember(groupId, memberId)
//	if ok, err := d.redis.Exists(key); err != nil {
//		return nil, err
//	} else if !ok {
//		return nil, db.ErrGroupMemberNotExist
//	}
//
//	res := &db.GroupMemberWithMute{}
//	if err := d.redis.Read(key, res); err != nil {
//		return nil, err
//	}
//
//	return res, nil
//}
//
//func (d *Dao) SaveGroupMemberWithMuteTime(member *db.GroupMemberWithMute, duration time.Duration) error {
//	key := keyGroup(member.GroupId)
//	if err := d.redis.Write(key, member, int(duration)); err != nil {
//		return err
//	}
//	return nil
//}
