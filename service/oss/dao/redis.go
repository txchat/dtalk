package dao

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/txchat/dtalk/pkg/oss"
	"github.com/txchat/dtalk/service/oss/model"
)

const (
	_prefixOssConfigCache = "oss-config:%v-%v"
)

func keyOssConfig(app, ossType string) string {
	return fmt.Sprintf(_prefixOssConfigCache, app, ossType)
}

//key:name; val:json
func (d *Dao) SaveAssumeRole(appId, ossType string, cfg *oss.Config, data *oss.AssumeRoleResp) error {
	if cfg == nil {
		return model.ErrNilPoint
	}
	val, err := json.Marshal(data)
	if err != nil {
		return err
	}
	key := keyOssConfig(appId, ossType)
	conn := d.redis.Get()
	defer conn.Close()
	if err := conn.Send("SET", key, val); err != nil {
		d.log.Error().Err(err).Str("key", key).Interface("val", val).Msg("conn.Send(SET)")
		return err
	}
	if err := conn.Send("EXPIRE", key, cfg.DurationSeconds); err != nil {
		d.log.Error().Err(err).Str("key", key).Interface("val", val).Msg("conn.Send(EXPIRE)")
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

func (d *Dao) GetAssumeRole(appId, ossType string, cfg *oss.Config) (*oss.AssumeRoleResp, error) {
	if cfg == nil {
		return nil, model.ErrNilPoint
	}
	key := keyOssConfig(appId, ossType)
	conn := d.redis.Get()
	defer conn.Close()
	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		d.log.Error().Err(err).Str("key", key).Msg("conn.DO(GET)")
		return nil, err
	}
	var data oss.AssumeRoleResp
	err = json.Unmarshal([]byte(reply), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
