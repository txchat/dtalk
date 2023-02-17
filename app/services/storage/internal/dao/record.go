package dao

import (
	"fmt"
	"strings"

	"github.com/txchat/dtalk/app/services/storage/internal/model"
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/pkg/util"
)

const (
	_InsertMsgContent  = `INSERT INTO dtalk_msg_content(mid,cid,sender_id,receiver_id,msg_type,content,create_time,source,reference) VALUES(?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE mid=mid`
	_InsertMsgRelation = `INSERT INTO dtalk_msg_relation(mid,owner_uid,other_uid,type,create_time) VALUES(?,?,?,?,?) ON DUPLICATE KEY UPDATE mid=mid`
	_DelMsgContent     = `DELETE FROM dtalk_msg_content WHERE mid = ?`

	_InsertGroupMsgContent        = `INSERT INTO dtalk_group_msg_content(mid,cid,sender_id,receiver_id,group_id,msg_type,content,create_time,source,reference) VALUES(?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE mid=mid`
	_InsertGroupMsgRelationPrefix = "INSERT INTO dtalk_group_msg_relation(mid,owner_uid,other_uid,type,state,create_time) VALUES%s ON DUPLICATE KEY UPDATE mid=mid"
	_SetGroupRecordState          = `UPDATE dtalk_group_msg_relation SET state = ? WHERE owner_uid = ? AND mid = ?`
	_DelGroupMsgContent           = `DELETE FROM dtalk_group_msg_content WHERE mid = ?`

	//_GetUnReceiveMsg      = `SELECT co.mid as mid,co.seq as seq,co.sender_id as sender_id,co.receiver_id as receiver_id,co.msg_type as msg_type,co.content as content,co.create_time as create_time,co.source as source,co.reference as reference FROM dtalk_msg_relation AS re RIGHT JOIN dtalk_msg_content AS co ON re.mid=co.mid WHERE re.owner_uid=? AND re.type=? AND re.state=? ORDER BY re.mid`
	_GetUnReceiveMsg = `SELECT co.mid as mid,co.seq as seq,co.sender_id as sender_id,co.receiver_id as receiver_id,co.msg_type as msg_type,co.content as content,co.create_time as create_time,co.source as source,co.reference as reference FROM dtalk_msg_relation AS re RIGHT JOIN dtalk_msg_content AS co ON re.mid=co.mid WHERE re.owner_uid=? AND re.state=? ORDER BY re.mid`
	//_GetUnReceiveGroupMsg = `SELECT co.mid as mid,co.seq as seq,co.sender_id as sender_id,co.receiver_id as receiver_id,co.msg_type as msg_type,co.content as content,co.create_time as create_time,co.source as source,co.reference as reference FROM dtalk_group_msg_relation AS re RIGHT JOIN dtalk_group_msg_content AS co ON re.mid=co.mid WHERE re.owner_uid=? AND re.type=? AND re.state=? ORDER BY re.mid`
	_GetUnReceiveGroupMsg = `SELECT co.mid as mid,co.seq as seq,co.sender_id as sender_id,co.receiver_id as receiver_id,co.msg_type as msg_type,co.content as content,co.create_time as create_time,co.source as source,co.reference as reference FROM dtalk_group_msg_relation AS re RIGHT JOIN dtalk_group_msg_content AS co ON re.mid=co.mid WHERE re.owner_uid=? AND re.state=? ORDER BY re.mid`
	_GetUserMsg           = `SELECT co.mid as mid,co.seq as seq,co.sender_id as sender_id,co.receiver_id as receiver_id,co.msg_type as msg_type,co.content as content,co.create_time as create_time,co.source as source,co.reference as reference FROM dtalk_msg_relation AS re RIGHT JOIN dtalk_msg_content AS co ON re.mid=co.mid WHERE re.owner_uid=? ORDER BY re.mid LIMIT ?,?`
	_GetUserMsgAfter      = `SELECT co.mid as mid,co.seq as seq,co.sender_id as sender_id,co.receiver_id as receiver_id,co.msg_type as msg_type,co.content as content,co.create_time as create_time,co.source as source,co.reference as reference FROM dtalk_msg_relation AS re RIGHT JOIN dtalk_msg_content AS co ON re.mid=co.mid WHERE re.owner_uid=? AND co.mid>? ORDER BY re.mid DESC LIMIT ?,?`
	_GetGroupMsgAfter     = `SELECT co.mid as mid,co.seq as seq,co.sender_id as sender_id,co.receiver_id as receiver_id,co.msg_type as msg_type,co.content as content,co.create_time as create_time,co.source as source,co.reference as reference FROM dtalk_group_msg_relation AS re RIGHT JOIN dtalk_group_msg_content AS co ON re.mid=co.mid WHERE re.owner_uid=? AND co.mid>? ORDER BY re.mid DESC LIMIT ?,?`

	_GetSpecifyRecords      = `SELECT co.mid as mid,co.seq as seq,co.sender_id as sender_id,co.receiver_id as receiver_id,co.msg_type as msg_type,co.content as content,co.create_time as create_time,co.source as source,co.reference as reference FROM dtalk_msg_relation AS re RIGHT JOIN dtalk_msg_content AS co ON re.mid=co.mid WHERE re.mid = ?`
	_GetSpecifyGroupRecords = `SELECT co.mid as mid,co.seq as seq,co.sender_id as sender_id,co.receiver_id as receiver_id,co.msg_type as msg_type,co.content as content,co.create_time as create_time,co.source as source,co.reference as reference FROM dtalk_group_msg_relation AS re RIGHT JOIN dtalk_group_msg_content AS co ON re.mid=co.mid WHERE re.mid = ?`

	_GetMsgBySeq = `SELECT mid,seq,sender_id,receiver_id,msg_type,content,create_time,source,reference FROM dtalk_msg_content WHERE sender_id=? AND seq=?`

	_IncMsgVersion = `INSERT INTO dtalk_msg_version(uid,version) VALUES(?,?) ON DUPLICATE KEY UPDATE version=version+1`
	_GetMsgVersion = `SELECT version FROM dtalk_msg_version WHERE uid = ?`

	_GetPriRecords = `SELECT co.mid as mid,co.seq as seq,co.sender_id as sender_id,co.receiver_id as receiver_id,co.msg_type as msg_type,co.content as content,co.create_time as create_time FROM dtalk_msg_relation AS re RIGHT JOIN dtalk_msg_content AS co ON re.mid=co.mid WHERE re.owner_uid=? AND re.other_uid=? AND re.mid<? AND re.state=1 ORDER BY re.mid desc LIMIT 0,?`
)

func convertMsgContent(m map[string]string) *model.PrivateMsgContent {
	return &model.PrivateMsgContent{
		Mid:        m["mid"],
		Cid:        m["seq"],
		SenderId:   m["sender_id"],
		ReceiverId: m["receiver_id"],
		MsgType:    util.MustToUint32(m["msg_type"]),
		Content:    m["content"],
		CreateTime: uint64(util.MustToInt64(m["create_time"])),
		Source:     m["source"],
		Reference:  m["reference"],
	}
}

func (repo *UniRepository) AppendPrivateMsgContent(tx *mysql.MysqlTx, m *model.PrivateMsgContent) (int64, int64, error) {
	num, lastId, err := tx.Exec(_InsertMsgContent,
		m.Mid, m.Cid, m.SenderId, m.ReceiverId, m.MsgType, m.Content, m.CreateTime, m.Source, m.Reference)
	return num, lastId, err
}

func (repo *UniRepository) AppendPrivateMsgRelation(tx *mysql.MysqlTx, m *model.PrivateMsgRelation) (int64, int64, error) {
	num, lastId, err := tx.Exec(_InsertMsgRelation,
		m.Mid, m.OwnerUid, m.OtherUid, m.Type, m.CreateTime)
	return num, lastId, err
}

func (repo *UniRepository) DelPrivateMsgContent(mid int64) (int64, int64, error) {
	num, lastId, err := repo.mysql.Exec(_DelMsgContent, mid)
	return num, lastId, err
}

func (repo *UniRepository) UserLastMsg(uid string, num int64) ([]*model.PrivateMsgContent, error) {
	maps, err := repo.mysql.Query(_GetUserMsg, uid, 0, num)
	if err != nil {
		return nil, err
	}
	ret := make([]*model.PrivateMsgContent, len(maps))
	for i, m := range maps {
		ret[i] = convertMsgContent(m)
	}
	return ret, err
}

func (repo *UniRepository) UserMsgAfter(uid string, startMid, count int64) ([]*model.PrivateMsgContent, error) {
	maps, err := repo.mysql.Query(_GetUserMsgAfter, uid, startMid, 0, count)
	if err != nil {
		return nil, err
	}
	ret := make([]*model.PrivateMsgContent, len(maps))
	for i, m := range maps {
		ret[i] = convertMsgContent(m)
	}
	return ret, err
}

func (repo *UniRepository) GetMsgBySeq(senderUID, seq string) (*model.PrivateMsgContent, error) {
	maps, err := repo.mysql.Query(_GetMsgBySeq, senderUID, seq)
	if err != nil {
		return nil, err
	}
	if len(maps) < 1 {
		return nil, nil
	}
	m := maps[0]
	return convertMsgContent(m), err
}

func (repo *UniRepository) AppendGroupMsgContent(m *model.GroupMsgContent) (int64, int64, error) {
	num, lastId, err := repo.mysql.Exec(_InsertGroupMsgContent,
		m.Mid, m.Cid, m.SenderId, m.ReceiverId, m.GroupId, m.MsgType, m.Content, m.CreateTime, m.Source, m.Reference)
	return num, lastId, err
}

func (repo *UniRepository) AppendGroupMsgRelation(tx *mysql.MysqlTx, m []*model.PrivateMsgRelation) (int64, int64, error) {
	//element should not empty
	if len(m) == 0 {
		return 0, 0, nil
	}
	var values []interface{}
	condition := ""
	for _, row := range m {
		condition += "(?,?,?,?,?,?),"
		values = append(values, row.Mid, row.OwnerUid, row.OtherUid, row.Type, row.State, row.CreateTime)
	}
	//trim the last ,
	condition = strings.TrimSuffix(condition, ",")
	//prepare the statement and exec
	num, lastId, err := tx.PrepareExec(fmt.Sprintf(_InsertGroupMsgRelationPrefix, condition), values...)
	return num, lastId, err
}

func (repo *UniRepository) MarkGroupMsgReceived(uid string, mid int64) (int64, int64, error) {
	num, lastId, err := repo.mysql.Exec(_SetGroupRecordState, model.Received, uid, mid)
	return num, lastId, err
}

func (repo *UniRepository) DelGroupMsgContent(mid int64) (int64, int64, error) {
	num, lastId, err := repo.mysql.Exec(_DelGroupMsgContent, mid)
	return num, lastId, err
}

func (repo *UniRepository) UnReceiveGroupMsg(uid string) ([]*model.PrivateMsgContent, error) {
	maps, err := repo.mysql.Query(_GetUnReceiveGroupMsg, uid, model.UnReceive)
	if err != nil {
		return nil, err
	}
	ret := make([]*model.PrivateMsgContent, len(maps))
	for i, m := range maps {
		ret[i] = convertMsgContent(m)
	}
	return ret, err
}

func (repo *UniRepository) GroupMsgAfter(uid string, startMid, count int64) ([]*model.PrivateMsgContent, error) {
	maps, err := repo.mysql.Query(_GetGroupMsgAfter, uid, startMid, 0, count)
	if err != nil {
		return nil, err
	}
	ret := make([]*model.PrivateMsgContent, len(maps))
	for i, m := range maps {
		ret[i] = convertMsgContent(m)
	}
	return ret, err
}

// get record

// GetPriRecord 获得聊天记录
func (repo *UniRepository) GetPriRecord(fromId, targetId string, mid int64, recordCount int64) ([]*model.PrivateMsgContent, error) {
	maps, err := repo.mysql.Query(_GetPriRecords, fromId, targetId, mid, recordCount)
	if err != nil {
		return nil, err
	}
	res := make([]*model.PrivateMsgContent, len(maps))
	for i, m := range maps {
		res[i] = &model.PrivateMsgContent{
			Mid:        m["mid"],
			Cid:        m["seq"],
			SenderId:   m["sender_id"],
			ReceiverId: m["receiver_id"],
			MsgType:    util.MustToUint32(m["msg_type"]),
			Content:    m["content"],
			CreateTime: uint64(util.MustToInt64(m["create_time"])),
		}
	}

	return res, nil
}

// GetRecord 获得指定聊天记录
func (repo *UniRepository) GetPrivateRecordByMid(mid int64) (*model.PrivateMsgContent, error) {
	maps, err := repo.mysql.Query(_GetSpecifyRecords, mid)
	if err != nil {
		return nil, err
	}
	if len(maps) < 1 {
		return nil, model.ErrRecordNotFind
	}
	return convertMsgContent(maps[0]), nil
}

func (repo *UniRepository) GetGroupRecordByMid(mid int64) (*model.PrivateMsgContent, error) {
	maps, err := repo.mysql.Query(_GetSpecifyGroupRecords, mid)
	if err != nil {
		return nil, err
	}
	if len(maps) < 1 {
		return nil, model.ErrRecordNotFind
	}
	return convertMsgContent(maps[0]), nil
}
