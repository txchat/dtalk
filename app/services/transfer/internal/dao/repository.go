package dao

import (
	"context"

	"github.com/txchat/dtalk/api/proto/chat"

	"github.com/txchat/dtalk/app/services/transfer/internal/model"
)

type Repository interface {
	IncrUserSeq(ctx context.Context, uid string) (int64, error)
	SaveUserChatRecord(ctx context.Context, chatProto *chat.Chat) error
	MarkUserChatRecordReceived(ctx context.Context, uid string, seq int64) error
	GetChatRecordSeqByMid(ctx context.Context, mid string) (int64, error)

	MappingClientSeq(ctx context.Context, index *model.MessageIndex) error
	GetMidByClientSeq(ctx context.Context, uuid string, seq int64) (int64, error)
	GetUserClientSeqByMid(ctx context.Context, mid int64) (*model.MessageIndex, error)
	UpdateLastReceiveMid(ctx context.Context, receiver, mid string) error
}
