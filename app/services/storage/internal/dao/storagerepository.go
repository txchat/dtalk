package dao

import (
	"github.com/txchat/dtalk/app/services/storage/internal/model"
	xmysql "github.com/txchat/dtalk/pkg/mysql"
)

type StorageRepository interface {
	NewTx() (*xmysql.MysqlTx, error)

	AppendMsgContent(tx *xmysql.MysqlTx, m *model.MsgContent) (int64, int64, error)
	AppendMsgRelation(tx *xmysql.MysqlTx, m *model.MsgRelation) (int64, int64, error)
	MarkMsgReceived(uid string, mid int64) (int64, int64, error)
	DelMsgContent(mid int64) (int64, int64, error)
	UnReceiveMsg(uid string) ([]*model.MsgContent, error)
	UserLastMsg(uid string, num int64) ([]*model.MsgContent, error)
	UserMsgAfter(uid string, startMid, count int64) ([]*model.MsgContent, error)
	GetMsgBySeq(senderUID, seq string) (*model.MsgContent, error)
	IncMsgVersion(uid string) (uint64, error)
	//group about
	AppendGroupMsgContent(tx *xmysql.MysqlTx, m *model.MsgContent) (int64, int64, error)
	AppendGroupMsgRelation(tx *xmysql.MysqlTx, m []*model.MsgRelation) (int64, int64, error)
	MarkGroupMsgReceived(uid string, mid int64) (int64, int64, error)
	DelGroupMsgContent(mid int64) (int64, int64, error)
	UnReceiveGroupMsg(uid string) ([]*model.MsgContent, error)
	GroupMsgAfter(uid string, startMid, count int64) ([]*model.MsgContent, error)
	// record
	GetPriRecord(fromId, targetId string, mid int64, recordCount int64) ([]*model.MsgContent, error)
	GetSpecifyRecord(mid int64) (*model.MsgContent, error)
	GetSpecifyGroupRecord(mid int64) (*model.MsgContent, error)
	//cache
	AddRecordCache(uid string, ver uint64, m *model.MsgCache) error
	UserRecords(uid string, ver uint64) ([]*model.MsgCache, error)
	AddRecordFocus(uid string, mid int64, time uint64) error
	GetRecordFocusNumber(mid int64) (int32, error)
	//signal
	AppendSignalContent(m *model.SignalContent) (int64, int64, error)
	BatchAppendSignalContent(m []*model.SignalContent) (int64, int64, error)
	MarkSignalReceived(uid string, mid int64) (int64, int64, error)
	UnReceiveSignalMsg(uid string) ([]*model.SignalContent, error)
	SyncSignalMsg(uid string, startId, count int64) ([]*model.SignalContent, error)
}
