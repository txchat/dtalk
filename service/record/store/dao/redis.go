package dao

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/txchat/dtalk/service/record/store/model"
)

const (
	_prefixRecordCache = "record:%v"
	_prefixRecordFocus = "record-focus:%v"
)

func keyUserRecords(uid string) string {
	return fmt.Sprintf(_prefixRecordCache, uid)
}

func keyRecordFocus(mid int64) string {
	return fmt.Sprintf(_prefixRecordFocus, mid)
}

//key:version; val:json
func (d *Dao) AddRecordCache(uid string, ver uint64, m *model.MsgCache) error {
	val, err := json.Marshal(m)
	if err != nil {
		return err
	}
	key := keyUserRecords(uid)
	conn := d.redis.Get()
	defer conn.Close()
	if err := conn.Send("ZADD", key, ver, val); err != nil {
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.Send(ZADD %s,%d,%s)", key, ver, val))
		return err
	}
	if err := conn.Flush(); err != nil {
		d.log.Error().Err(err).Msg("conn.Flush()")
		return err
	}
	if _, err := conn.Receive(); err != nil {
		d.log.Error().Err(err).Msg("conn.Receive()")
		return err
	}
	return nil
}

func (d *Dao) UserRecords(uid string, ver uint64) ([]*model.MsgCache, error) {
	key := keyUserRecords(uid)
	conn := d.redis.Get()
	defer conn.Close()
	nMap, err := redis.StringMap(conn.Do("ZREVRANGEBYSCORE", key, "+inf", ver, "WITHSCORES"))
	if err != nil {
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.DO(ZREVRANGEBYSCORE %s, %s, %d, %s)", key, "+inf", ver, "WITHSCORES"))
		return nil, err
	}
	msg := make([]*model.MsgCache, 0)
	for val := range nMap {
		item := model.MsgCache{}
		err := json.Unmarshal([]byte(val), &item)
		if err != nil {
			return nil, err
		}
		msg = append(msg, &item)
	}
	return msg, nil
}

//key:version; val:json
func (d *Dao) AddRecordFocus(uid string, mid int64, time uint64) error {
	key := keyRecordFocus(mid)
	conn := d.redis.Get()
	defer conn.Close()
	if err := conn.Send("SADD", key, uid); err != nil {
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.Send(SADD %s,%s)", key, uid))
		return err
	}
	if err := conn.Flush(); err != nil {
		d.log.Error().Err(err).Msg("conn.Flush()")
		return err
	}
	if _, err := conn.Receive(); err != nil {
		d.log.Error().Err(err).Msg("conn.Receive()")
		return err
	}
	return nil
}

func (d *Dao) GetRecordFocusNumber(mid int64) (int32, error) {
	key := keyRecordFocus(mid)
	conn := d.redis.Get()
	defer conn.Close()
	num, err := redis.Int(conn.Do("SCARD", key))
	if err != nil {
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.DO(SCARD %s)", key))
		return 0, err
	}
	return int32(num), nil
}
