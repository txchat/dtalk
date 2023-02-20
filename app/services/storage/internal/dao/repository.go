package dao

import (
	"github.com/txchat/dtalk/app/services/storage/internal/model"
	xmysql "github.com/txchat/dtalk/pkg/mysql"
)

type Repository interface {
	NewTx() (*xmysql.MysqlTx, error)

	AppendPrivateMsgContent(tx *xmysql.MysqlTx, m *model.MsgContent) (int64, int64, error)
	AppendPrivateMsgRelation(tx *xmysql.MysqlTx, m *model.MsgRelation) (int64, int64, error)
	DelPrivateMsgContent(mid string) (int64, int64, error)
	AppendGroupMsgContent(tx *xmysql.MysqlTx, m *model.MsgContent) (int64, int64, error)
	AppendGroupMsgRelation(tx *xmysql.MysqlTx, m []*model.MsgRelation) (int64, int64, error)
	DelGroupMsgContent(mid string) (int64, int64, error)

	GetPrivateRecordByMid(mid string) (*model.MsgContent, error)
	GetGroupRecordByMid(mid string) (*model.MsgContent, error)
	GetPrivateChatSessionMsg(uid, target, start string, num int32) ([]*model.MsgContent, error)
	GetGroupChatSessionMsg(uid, gid, start string, num int32) ([]*model.MsgContent, error)

	AppendSignalContent(m *model.SignalContent) (int64, int64, error)
	BatchAppendSignalContent(m []*model.SignalContent) (int64, int64, error)
	AddRecordFocus(uid string, mid string, time int64) error
	GetRecordFocusNumber(mid string) (int32, error)
}
