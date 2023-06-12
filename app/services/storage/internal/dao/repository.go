package dao

import (
	"github.com/txchat/dtalk/app/services/storage/internal/model"
)

type Repository interface {
	AppendPrivateMsg(m *model.MsgContent, sender *model.MsgRelation, receiver *model.MsgRelation) (int64, int64, error)
	DelPrivateMsg(mid string) (int64, int64, error)
	AppendGroupMsg(m *model.MsgContent, relations []*model.MsgRelation) (int64, int64, error)
	DelGroupMsg(mid string) (int64, int64, error)

	GetPrivateMsgByMid(mid string) (*model.MsgContent, error)
	GetGroupMsgByMid(mid string) (*model.MsgContent, error)
	GetPrivateChatSessionMsg(uid, target, start string, num int32) ([]*model.MsgContent, error)
	GetGroupChatSessionMsg(uid, gid, start string, num int32) ([]*model.MsgContent, error)

	AppendSignal(m *model.SignalContent) (int64, int64, error)
	BatchAppendSignal(m []*model.SignalContent) (int64, int64, error)
	AddRecordFocus(uid string, mid string, time int64) error
	GetRecordFocusNumber(mid string) (int32, error)
}
