package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/txchat/dtalk/api/proto/chat"
	xredis "github.com/txchat/dtalk/pkg/redis"
	"github.com/txchat/dtalk/pkg/util"
)

const (
	prefixUserIncrSeq = "transfer:global-seq:%s" // user incr seq
	prefixIndexCidMid = "transfer:cid-mid:%s"
)

var (
	ErrEmptyValue  = errors.New("error empty value")
	ErrReplyNumber = errors.New("error reply number")
)

func keyUserSeq(uid string) string {
	return fmt.Sprintf(prefixUserIncrSeq, uid)
}

func keyCidMidIndex(cid string) string {
	return fmt.Sprintf(prefixIndexCidMid, cid)
}

type TransferRepository struct {
	redis             *redis.Pool
	cidMidIndexExpire time.Duration
	//mongo *mongo.Client
}

func NewTransferRepository(redisConfig xredis.Config) *TransferRepository {
	return &TransferRepository{
		redis:             xredis.NewPool(redisConfig),
		cidMidIndexExpire: time.Hour * 72,
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

func (repo *TransferRepository) SaveUserChatRecord(ctx context.Context, chatProto *chat.Chat) error {
	//TODO
	//coll := repo.mongo.Database("dtalk").Collection("record")
	//
	//result, err := coll.InsertOne(context.TODO(), newRestaurant)
	//if err != nil {
	//	panic(err)
	//}
	return nil
}

func (repo *TransferRepository) MarkUserChatRecordReceived(ctx context.Context, uid string, seq int64) error {
	//TODO
	return nil
}

func (repo *TransferRepository) AddIndexCidMid(ctx context.Context, cid, mid string, state chat.SendMessageReply_FailedType) error {
	if cid == "" || mid == "" {
		return ErrEmptyValue
	}
	key := keyCidMidIndex(cid)
	conn := repo.redis.Get()
	defer conn.Close()
	if err := conn.Send("HMSET", key, "mid", mid, "state", int32(state)); err != nil {
		return err
	}
	if err := conn.Send("EXPIRE", key, repo.cidMidIndexExpire); err != nil {
		return err
	}
	if err := conn.Flush(); err != nil {
		return err
	}
	_, err := conn.Receive()
	return err
}

func (repo *TransferRepository) GetMidByCid(ctx context.Context, cid string) (string, chat.SendMessageReply_FailedType, error) {
	if cid == "" {
		return "", 0, ErrEmptyValue
	}
	key := keyCidMidIndex(cid)
	conn := repo.redis.Get()
	defer conn.Close()
	if err := conn.Send("HMGET", key, "mid", "state"); err != nil {
		return "", 0, err
	}
	if err := conn.Flush(); err != nil {
		return "", 0, err
	}
	rlt, err := redis.Strings(conn.Receive())
	if err != nil {
		return "", 0, err
	}
	if len(rlt) < 2 {
		return "", 0, ErrReplyNumber
	}
	return rlt[0], chat.SendMessageReply_FailedType(util.MustToInt32(rlt[1])), err
}
