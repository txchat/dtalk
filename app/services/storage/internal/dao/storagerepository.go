package dao

import (
	"github.com/txchat/dtalk/app/services/storage/internal/model"
	xmysql "github.com/txchat/dtalk/pkg/mysql"
)

type StorageRepository interface {
	NewTx() (*xmysql.MysqlTx, error)

	AppendPrivateMsgContent(tx *xmysql.MysqlTx, m *model.PrivateMsgContent) (int64, int64, error)
	AppendPrivateMsgRelation(tx *xmysql.MysqlTx, m *model.PrivateMsgRelation) (int64, int64, error)
	DelPrivateMsgContent(mid int64) (int64, int64, error)
	AppendGroupMsgContent(m *model.GroupMsgContent) (int64, int64, error)
	DelGroupMsgContent(mid int64) (int64, int64, error)

	UserLastMsg(uid string, num int64) ([]*model.PrivateMsgContent, error)
	UserMsgAfter(uid string, startMid, count int64) ([]*model.PrivateMsgContent, error)
	GetMsgBySeq(senderUID, seq string) (*model.PrivateMsgContent, error)
	//group about
	AppendGroupMsgRelation(tx *xmysql.MysqlTx, m []*model.PrivateMsgRelation) (int64, int64, error)
	MarkGroupMsgReceived(uid string, mid int64) (int64, int64, error)
	UnReceiveGroupMsg(uid string) ([]*model.PrivateMsgContent, error)
	GroupMsgAfter(uid string, startMid, count int64) ([]*model.PrivateMsgContent, error)
	// record
	GetPriRecord(fromId, targetId string, mid int64, recordCount int64) ([]*model.PrivateMsgContent, error)
	GetPrivateRecordByMid(mid int64) (*model.PrivateMsgContent, error)
	GetGroupRecordByMid(mid int64) (*model.PrivateMsgContent, error)
	//cache
	AddRecordFocus(uid string, mid int64, time uint64) error
	GetRecordFocusNumber(mid int64) (int32, error)
	//signal
	AppendSignalContent(m *model.SignalContent) (int64, int64, error)
	BatchAppendSignalContent(m []*model.SignalContent) (int64, int64, error)
	MarkSignalReceived(uid string, mid int64) (int64, int64, error)
	UnReceiveSignalMsg(uid string) ([]*model.SignalContent, error)
	SyncSignalMsg(uid string, startId, count int64) ([]*model.SignalContent, error)
}
