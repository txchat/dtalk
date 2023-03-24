package dao

import (
	"fmt"
	"strings"

	"github.com/txchat/dtalk/app/services/storage/internal/model"
)

const (
	_InsertPrivateMsgContent  = `INSERT INTO dtalk_msg_content(mid,cid,sender_id,receiver_id,msg_type,content,create_time,source,reference) VALUES(?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE mid=mid`
	_InsertPrivateMsgRelation = `INSERT INTO dtalk_msg_relation(mid,owner_uid,other_uid,type,create_time) VALUES(?,?,?,?,?) ON DUPLICATE KEY UPDATE mid=mid`
	_DelPrivateMsgContent     = `DELETE FROM dtalk_msg_content WHERE mid = ?`
	_DelPrivateMsgRelation    = `DELETE FROM dtalk_msg_relation WHERE mid = ?`

	_InsertGroupMsgContent        = `INSERT INTO dtalk_group_msg_content(mid,cid,sender_id,receiver_id,msg_type,content,create_time,source,reference) VALUES(?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE mid=mid`
	_InsertGroupMsgRelationPrefix = "INSERT INTO dtalk_group_msg_relation(mid,owner_uid,other_uid,type,create_time) VALUES%s ON DUPLICATE KEY UPDATE mid=mid"
	_DelGroupMsgContent           = `DELETE FROM dtalk_group_msg_content WHERE mid = ?`
	_DelGroupMsgRelation          = `DELETE FROM dtalk_group_msg_relation WHERE mid = ?`

	_GetPrivateRecordByMid    = `SELECT co.mid as mid,co.cid as cid,co.sender_id as sender_id,co.receiver_id as receiver_id,co.msg_type as msg_type,co.content as content,co.create_time as create_time,co.source as source,co.reference as reference FROM dtalk_msg_relation AS re RIGHT JOIN dtalk_msg_content AS co ON re.mid=co.mid WHERE re.mid = ?`
	_GetGroupRecordByMid      = `SELECT co.mid as mid,co.cid as cid,co.sender_id as sender_id,co.receiver_id as receiver_id,co.msg_type as msg_type,co.content as content,co.create_time as create_time,co.source as source,co.reference as reference FROM dtalk_group_msg_relation AS re RIGHT JOIN dtalk_group_msg_content AS co ON re.mid=co.mid WHERE re.mid = ?`
	_GetPrivateChatSessionMsg = `SELECT co.mid as mid,co.cid as cid,co.sender_id as sender_id,co.receiver_id as receiver_id,co.msg_type as msg_type,co.content as content,co.create_time as create_time,co.source as source,co.reference as reference FROM dtalk_msg_relation AS re RIGHT JOIN dtalk_msg_content AS co ON re.mid=co.mid WHERE re.owner_uid=? AND re.other_uid=? AND co.mid>? ORDER BY re.mid DESC LIMIT ?,?`
	_GetGroupChatSessionMsg   = `SELECT co.mid as mid,co.cid as cid,co.sender_id as sender_id,co.receiver_id as receiver_id,co.msg_type as msg_type,co.content as content,co.create_time as create_time,co.source as source,co.reference as reference FROM dtalk_group_msg_relation AS re RIGHT JOIN dtalk_group_msg_content AS co ON re.mid=co.mid WHERE re.owner_uid=? AND re.other_uid=? AND co.mid>? ORDER BY re.mid DESC LIMIT ?,?`
)

func (repo *StorageRepository) GetPrivateMsgByMid(mid string) (*model.MsgContent, error) {
	var msg model.MsgContent
	err := repo.db.Raw(_GetPrivateRecordByMid, mid).Take(&msg).Error
	return &msg, err
}

func (repo *StorageRepository) AppendPrivateMsg(m *model.MsgContent, sender *model.MsgRelation, receiver *model.MsgRelation) (int64, int64, error) {
	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, 0, err
	}

	if err := tx.Exec(_InsertPrivateMsgContent,
		m.Mid, m.Cid, m.SenderId, m.ReceiverId, m.MsgType, m.Content, m.CreateTime, m.Source, m.Reference).Error; err != nil {
		tx.Rollback()
		return 0, 0, err
	}

	if err := tx.Exec(_InsertPrivateMsgRelation,
		sender.Mid, sender.OwnerUid, sender.OtherUid, sender.Type, sender.CreateTime).Error; err != nil {
		tx.Rollback()
		return 0, 0, err
	}

	if err := tx.Exec(_InsertPrivateMsgRelation,
		receiver.Mid, receiver.OwnerUid, receiver.OtherUid, receiver.Type, receiver.CreateTime).Error; err != nil {
		tx.Rollback()
		return 0, 0, err
	}

	return tx.RowsAffected, 0, tx.Commit().Error
}

func (repo *StorageRepository) DelPrivateMsg(mid string) (int64, int64, error) {
	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, 0, err
	}

	if err := tx.Exec(_DelPrivateMsgContent, mid).Error; err != nil {
		tx.Rollback()
		return 0, 0, err
	}

	if err := tx.Exec(_DelPrivateMsgRelation, mid).Error; err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	return tx.RowsAffected, 0, tx.Commit().Error
}

func (repo *StorageRepository) GetGroupMsgByMid(mid string) (*model.MsgContent, error) {
	var msg model.MsgContent
	err := repo.db.Raw(_GetGroupRecordByMid, mid).Take(&msg).Error
	return &msg, err
}

func (repo *StorageRepository) AppendGroupMsg(m *model.MsgContent, relations []*model.MsgRelation) (int64, int64, error) {
	//element should not empty
	if len(relations) == 0 {
		return 0, 0, nil
	}

	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, 0, err
	}

	if err := tx.Exec(_InsertGroupMsgContent,
		m.Mid, m.Cid, m.SenderId, m.ReceiverId, m.MsgType, m.Content, m.CreateTime, m.Source, m.Reference).Error; err != nil {
		tx.Rollback()
		return 0, 0, err
	}

	var values []interface{}
	condition := ""
	for _, row := range relations {
		condition += "(?,?,?,?,?),"
		values = append(values, row.Mid, row.OwnerUid, row.OtherUid, row.Type, row.CreateTime)
	}
	//trim the last ,
	condition = strings.TrimSuffix(condition, ",")
	//prepare the statement and exec

	if err := tx.Exec(fmt.Sprintf(_InsertGroupMsgRelationPrefix, condition), values...).Error; err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	return tx.RowsAffected, 0, tx.Commit().Error
}

func (repo *StorageRepository) DelGroupMsg(mid string) (int64, int64, error) {
	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, 0, err
	}

	if err := tx.Exec(_DelGroupMsgContent, mid).Error; err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	if err := tx.Exec(_DelGroupMsgRelation, mid).Error; err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	return tx.RowsAffected, 0, tx.Commit().Error
}

func (repo *StorageRepository) GetPrivateChatSessionMsg(uid, target, start string, num int32) ([]*model.MsgContent, error) {
	ret := make([]*model.MsgContent, 0)
	rows, err := repo.db.Raw(_GetPrivateChatSessionMsg, uid, target, start, 0, num).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var msg model.MsgContent
		if err = repo.db.ScanRows(rows, &msg); err != nil {
			continue
		}
		ret = append(ret, &msg)
	}
	return ret, nil
}

func (repo *StorageRepository) GetGroupChatSessionMsg(uid, gid, start string, num int32) ([]*model.MsgContent, error) {
	ret := make([]*model.MsgContent, 0)
	rows, err := repo.db.Raw(_GetGroupChatSessionMsg, uid, gid, start, 0, num).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var msg model.MsgContent
		if err = repo.db.ScanRows(rows, &msg); err != nil {
			continue
		}
		ret = append(ret, &msg)
	}
	return ret, nil
}
