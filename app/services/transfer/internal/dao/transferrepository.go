package dao

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
	xredis "github.com/txchat/dtalk/pkg/redis"
	"github.com/txchat/im/api/protocol"
)

const (
	prefixUserIncrSeq = "USER_INCR_SEQ:%s" // user incr seq
)

func keyUserSeq(uid string) string {
	return fmt.Sprintf(prefixUserIncrSeq, uid)
}

type TransferRepository struct {
	redis *redis.Pool
}

func NewTransferRepository(redisConfig xredis.Config) *TransferRepository {
	return &TransferRepository{
		redis: xredis.NewPool(redisConfig),
	}
}

func (repo *TransferRepository) IncrUserSeq(ctx context.Context, uid string) (int64, error) {
	key := keyUserSeq(uid)
	conn := repo.redis.Get()
	defer conn.Close()
	if err := conn.Send("INCR", key); err != nil {
		return 0, err
	}
	if err := conn.Flush(); err != nil {
		return 0, err
	}
	return redis.Int64(conn.Receive())
}

func (repo *TransferRepository) SaveUserChatRecord(ctx context.Context, p *protocol.Proto) error {

}

func (repo *TransferRepository) MarkUserChatRecordReceived(ctx context.Context, uid string, seq int64) error {

}
