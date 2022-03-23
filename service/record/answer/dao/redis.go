package dao

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/txchat/dtalk/service/record/answer/model"
)

const (
	_prefixRecordSeq = "record-seq:%v"
)

func keyUserRecordSeq(uid string) string {
	return fmt.Sprintf(_prefixRecordSeq, uid)
}

func (d *Dao) AddRecordSeqIndex(uid string, m *model.MsgIndex) error {
	val, err := json.Marshal(m)
	if err != nil {
		return err
	}
	key := keyUserRecordSeq(uid)
	conn := d.redis.Get()
	defer conn.Close()
	if err := conn.Send("HSET", key, m.Seq, val); err != nil {
		return err
	}
	if err := conn.Flush(); err != nil {
		return err
	}
	if _, err := conn.Receive(); err != nil {
		return err
	}
	return nil
}

func (d *Dao) GetRecordSeqIndex(uid, seq string) (*model.MsgIndex, error) {
	key := keyUserRecordSeq(uid)
	conn := d.redis.Get()
	defer conn.Close()
	val, err := redis.String(conn.Do("HGET", key, seq))
	if err != nil {
		if err == redis.ErrNil {
			return nil, nil
		}
		return nil, err
	}
	item := model.MsgIndex{}
	err = json.Unmarshal([]byte(val), &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
