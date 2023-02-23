package dao

import (
	"context"

	"github.com/txchat/dtalk/api/proto/chat"
)

type Repository interface {
	IncrUserSeq(ctx context.Context, uid string) (int64, error)
	SaveUserChatRecord(ctx context.Context, chatProto *chat.Chat) error
	MarkUserChatRecordReceived(ctx context.Context, uid string, seq int64) error

	AddIndexCidMid(ctx context.Context, cid, mid string) error
	GetMidByCid(ctx context.Context, cid string) (string, error)
}
