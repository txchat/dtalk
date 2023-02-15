package dao

import (
	"context"

	"github.com/txchat/dtalk/app/services/transfer/internal/model"
	"github.com/txchat/im/api/protocol"
)

type Repository interface {
	IncrUserSeq(ctx context.Context, uid string) (int64, error)
	SaveUserChatRecord(ctx context.Context, p *protocol.Proto) error
	MarkUserChatRecordReceived(ctx context.Context, uid string, seq int64) error
	MappingClientSeq(ctx context.Context, uid, uuid string, seq int64) error
	GetMidByClientSeq(ctx context.Context, uuid string, seq int64) (int64, error)
	GetUserClientSeqByMid(ctx context.Context, mid int64) (*model.MessageIndex, error)
	UpdateUserLatestRev(ctx context.Context, uuid, sender, receiver string, lastSeq int64) (int64, error)
}
