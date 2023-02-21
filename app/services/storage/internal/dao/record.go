package dao

import (
	"fmt"
	"strings"

	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/pkg/util"
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

func convertMsgContent(m map[string]string) *model.MsgContent {
	return &model.MsgContent{
		Mid:        m["mid"],
		Cid:        m["cid"],
		SenderId:   m["sender_id"],
		ReceiverId: m["receiver_id"],
		MsgType:    util.MustToInt32(m["msg_type"]),
		Content:    m["content"],
		CreateTime: util.MustToInt64(m["create_time"]),
		Source:     m["source"],
		Reference:  m["reference"],
	}
}

func (repo *StorageRepository) GetPrivateMsgByMid(mid string) (*model.MsgContent, error) {
	maps, err := repo.mysql.Query(_GetPrivateRecordByMid, mid)
	if err != nil {
		return nil, err
	}
	if len(maps) < 1 {
		return nil, model.ErrRecordNotFind
	}
	return convertMsgContent(maps[0]), nil
}

func (repo *StorageRepository) AppendPrivateMsgContent(tx *mysql.MysqlTx, m *model.MsgContent) (int64, int64, error) {
	num, lastId, err := tx.Exec(_InsertPrivateMsgContent,
		m.Mid, m.Cid, m.SenderId, m.ReceiverId, m.MsgType, m.Content, m.CreateTime, m.Source, m.Reference)
	return num, lastId, err
}

func (repo *StorageRepository) AppendPrivateMsgRelation(tx *mysql.MysqlTx, m *model.MsgRelation) (int64, int64, error) {
	num, lastId, err := tx.Exec(_InsertPrivateMsgRelation,
		m.Mid, m.OwnerUid, m.OtherUid, m.Type, m.CreateTime)
	return num, lastId, err
}

func (repo *StorageRepository) DelPrivateMsg(mid string) (int64, int64, error) {
	tx, err := repo.mysql.NewTx()
	if err != nil {
		return 0, 0, err
	}
	num, lastId, err := tx.Exec(_DelPrivateMsgContent, mid)
	if err != nil {
		tx.RollBack()
		return 0, 0, err
	}
	num, lastId, err = tx.Exec(_DelPrivateMsgRelation, mid)
	if err != nil {
		tx.RollBack()
		return 0, 0, err
	}
	err = tx.Commit()
	return num, lastId, err
}

func (repo *StorageRepository) GetGroupMsgByMid(mid string) (*model.MsgContent, error) {
	maps, err := repo.mysql.Query(_GetGroupRecordByMid, mid)
	if err != nil {
		return nil, err
	}
	if len(maps) < 1 {
		return nil, model.ErrRecordNotFind
	}
	return convertMsgContent(maps[0]), nil
}

func (repo *StorageRepository) AppendGroupMsgContent(tx *mysql.MysqlTx, m *model.MsgContent) (int64, int64, error) {
	num, lastId, err := tx.Exec(_InsertGroupMsgContent,
		m.Mid, m.Cid, m.SenderId, m.ReceiverId, m.MsgType, m.Content, m.CreateTime, m.Source, m.Reference)
	return num, lastId, err
}

func (repo *StorageRepository) AppendGroupMsgRelation(tx *mysql.MysqlTx, m []*model.MsgRelation) (int64, int64, error) {
	//element should not empty
	if len(m) == 0 {
		return 0, 0, nil
	}
	var values []interface{}
	condition := ""
	for _, row := range m {
		condition += "(?,?,?,?,?),"
		values = append(values, row.Mid, row.OwnerUid, row.OtherUid, row.Type, row.CreateTime)
	}
	//trim the last ,
	condition = strings.TrimSuffix(condition, ",")
	//prepare the statement and exec
	num, lastId, err := tx.PrepareExec(fmt.Sprintf(_InsertGroupMsgRelationPrefix, condition), values...)
	return num, lastId, err
}

func (repo *StorageRepository) DelGroupMsg(mid string) (int64, int64, error) {
	tx, err := repo.mysql.NewTx()
	if err != nil {
		return 0, 0, err
	}
	num, lastId, err := tx.Exec(_DelGroupMsgContent, mid)
	if err != nil {
		tx.RollBack()
		return 0, 0, err
	}
	num, lastId, err = tx.Exec(_DelGroupMsgRelation, mid)
	if err != nil {
		tx.RollBack()
		return 0, 0, err
	}
	err = tx.Commit()
	return num, lastId, err
}

func (repo *StorageRepository) GetPrivateChatSessionMsg(uid, target, start string, num int32) ([]*model.MsgContent, error) {
	maps, err := repo.mysql.Query(_GetPrivateChatSessionMsg, uid, target, start, 0, num)
	if err != nil {
		return nil, err
	}
	ret := make([]*model.MsgContent, len(maps))
	for i, m := range maps {
		ret[i] = convertMsgContent(m)
	}
	return ret, err
}

func (repo *StorageRepository) GetGroupChatSessionMsg(uid, gid, start string, num int32) ([]*model.MsgContent, error) {
	maps, err := repo.mysql.Query(_GetGroupChatSessionMsg, uid, gid, start, 0, num)
	if err != nil {
		return nil, err
	}
	ret := make([]*model.MsgContent, len(maps))
	for i, m := range maps {
		ret[i] = convertMsgContent(m)
	}
	return ret, err
}
