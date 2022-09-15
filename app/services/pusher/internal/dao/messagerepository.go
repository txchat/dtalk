package dao

import (
	"github.com/txchat/dtalk/app/services/pusher/internal/recordhelper"
)

type MessageRepository interface {
	AddConnSeqIndex(cid string, seq int32, item *recordhelper.ConnSeqItem) error
	GetConnSeqIndex(cid string, seq int32) (*recordhelper.ConnSeqItem, error)
	ClearConnSeq(cid string) error
	GetGroupSession(cid string, seq int32) (session string, err error)
}
