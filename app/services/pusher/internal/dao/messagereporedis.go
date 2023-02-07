package dao

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/txchat/dtalk/internal/recordhelper"
	xredis "github.com/txchat/dtalk/pkg/redis"
)

const (
	_prefixConnSeq = "conn-seq:%v"
)

func keyConnection(cid string) string {
	return fmt.Sprintf(_prefixConnSeq, cid)
}

type MessageRepositoryRedis struct {
	redis *redis.Pool
}

func NewMessageRepositoryRedis(c xredis.Config) *MessageRepositoryRedis {
	return &MessageRepositoryRedis{
		redis: xredis.NewPool(c),
	}
}

//key:connect id; val:logs id
func (repo *MessageRepositoryRedis) AddConnSeqIndex(cid string, seq int32, item *recordhelper.ConnSeqItem) error {
	key := keyConnection(cid)
	conn := repo.redis.Get()
	defer conn.Close()
	val, err := json.Marshal(item)
	if err != nil {
		return err
	}
	if err := conn.Send("HSET", key, seq, val); err != nil {
		return fmt.Errorf("conn.Send(HSET %s,%d,%s) failed:%v", key, seq, val, err)
	}
	if err := conn.Flush(); err != nil {
		return fmt.Errorf("conn.Flush() failed:%v", err)
	}
	if _, err := conn.Receive(); err != nil {
		return fmt.Errorf("conn.Receive() failed:%v", err)
	}
	return nil
}

func (repo *MessageRepositoryRedis) GetConnSeqIndex(cid string, seq int32) (*recordhelper.ConnSeqItem, error) {
	key := keyConnection(cid)
	conn := repo.redis.Get()
	defer conn.Close()
	data, err := redis.String(conn.Do("HGET", key, seq))
	if err != nil {
		if err == redis.ErrNil {
			return nil, nil
		}
		return nil, fmt.Errorf("conn.DO(HGET %s, %d) failed:%v", key, seq, err)
	}
	var item recordhelper.ConnSeqItem
	err = json.Unmarshal([]byte(data), &item)
	if err != nil {
		return nil, err
	}
	//logsStr := ret[1:]
	//tp := ret[0]
	//var logs = make([]uint64, len(logsStr))
	//for i, l := range logsStr {
	//	log, err := strconv.ParseInt(l, 10, 64)
	//	if err != nil {
	//		return "", nil, err
	//	}
	//	logs[i] = uint64(log)
	//}
	//return bizroto.EventType(tp), logs, nil
	return &item, nil
}

func (repo *MessageRepositoryRedis) ClearConnSeq(cid string) error {
	key := keyConnection(cid)
	conn := repo.redis.Get()
	defer conn.Close()

	if err := conn.Send("DEL", key); err != nil {
		return fmt.Errorf("conn.Send(DEL %s) failed:%v", key, err)
	}
	if err := conn.Flush(); err != nil {
		return fmt.Errorf("conn.Flush() failed:%v", err)
	}
	if _, err := conn.Receive(); err != nil {
		return fmt.Errorf("conn.Receive() failed:%v", err)
	}
	return nil
}

func (repo *MessageRepositoryRedis) GetGroupSession(cid string, seq int32) (session string, err error) {
	//TODO call logic get log mark by connect seq
	return "", nil
}
