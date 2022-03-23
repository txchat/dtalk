package dao

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/txchat/dtalk/service/record/pusher/logH"
)

const (
	_prefixConnSeq = "conn-seq:%v"
)

func keyConnection(cid string) string {
	return fmt.Sprintf(_prefixConnSeq, cid)
}

func connSeqIndexValConvert(tp string, logs []string) (string, error) {
	val := fmt.Sprintf("type=%s;", tp)
	val += strings.Join(logs, ",")
	//check
	reg := regexp.MustCompile(`^type=(\S+);`)
	if !reg.MatchString(val) {
		return val, errors.New("conn seq index convert err")
	}
	return val, nil
}

func connSeqIndexValParse(val string) ([]string, error) {
	reg := regexp.MustCompile(`^type=(\S+);`)
	if !reg.MatchString(val) {
		return nil, errors.New("conn seq index parse err")
	}
	//找出type子串
	typeCnt := reg.FindStringSubmatch(val)
	if len(typeCnt) < 2 {
		return nil, errors.New("conn seq index parse err: cnt len < 2")
	}
	//找出;所在的尾部下标
	typeIdx := reg.FindStringIndex(val)
	if len(typeIdx) < 2 {
		return nil, errors.New("conn seq index parse err: idx len < 2")
	}
	log := strings.Split(val[typeIdx[1]:], ",")
	var ret = make([]string, len(log)+1)
	copy(ret[1:], log)
	ret[0] = typeCnt[1]
	return ret, nil
}

//key:connect id; val:logs id
func (d *Dao) AddConnSeqIndex(cid string, seq int32, item *logH.ConnSeqItem) error {
	key := keyConnection(cid)
	conn := d.redis.Get()
	defer conn.Close()
	val, err := json.Marshal(item)
	if err != nil {
		return err
	}
	if err := conn.Send("HSET", key, seq, val); err != nil {
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.Send(HSET %s,%d,%s)", key, seq, val))
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

func (d *Dao) GetConnSeqIndex(cid string, seq int32) (*logH.ConnSeqItem, error) {
	key := keyConnection(cid)
	conn := d.redis.Get()
	defer conn.Close()
	data, err := redis.String(conn.Do("HGET", key, seq))
	if err != nil {
		if err == redis.ErrNil {
			return nil, nil
		}
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.DO(HGET %s, %d)", key, seq))
		return nil, err
	}
	var item logH.ConnSeqItem
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

func (d *Dao) ClearConnSeq(cid string) error {
	key := keyConnection(cid)
	conn := d.redis.Get()
	defer conn.Close()

	if err := conn.Send("DEL", key); err != nil {
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.Send(DEL %s)", key))
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
