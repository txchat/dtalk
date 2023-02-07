package dao

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/oss"
	xredis "github.com/txchat/dtalk/pkg/redis"
)

const (
	_prefixOssConfigCache = "oss-config:%v-%v"
)

func keyOssConfig(app, ossType string) string {
	return fmt.Sprintf(_prefixOssConfigCache, app, ossType)
}

type OssRepositoryRedis struct {
	redis *redis.Pool
}

func NewOssRepositoryRedis(c xredis.Config) *OssRepositoryRedis {
	return &OssRepositoryRedis{
		redis: xredis.NewPool(c),
	}
}

//key:name; val:json
func (repo *OssRepositoryRedis) SaveAssumeRole(appId, ossType string, cfg *oss.Config, data *oss.AssumeRoleResp) error {
	if cfg == nil {
		return xerror.ErrOssEndpointNotExist
	}
	val, err := json.Marshal(data)
	if err != nil {
		return err
	}
	key := keyOssConfig(appId, ossType)
	conn := repo.redis.Get()
	defer conn.Close()
	if err := conn.Send("SET", key, val); err != nil {
		return fmt.Errorf("conn.Send(SET %s,%s) failed:%v", key, val, err)
	}
	if err := conn.Send("EXPIRE", key, cfg.DurationSeconds); err != nil {
		return fmt.Errorf("conn.Send(EXPIRE %s,%s) failed:%v", key, cfg.DurationSeconds, err)
	}
	if err := conn.Flush(); err != nil {
		return fmt.Errorf("conn.Flush() failed:%v", err)
	}
	if _, err := conn.Receive(); err != nil {
		return fmt.Errorf("conn.Receive() failed:%v", err)
	}
	return nil
}

func (repo *OssRepositoryRedis) GetAssumeRole(appId, ossType string, cfg *oss.Config) (*oss.AssumeRoleResp, error) {
	if cfg == nil {
		return nil, xerror.ErrOssEndpointNotExist
	}
	key := keyOssConfig(appId, ossType)
	conn := repo.redis.Get()
	defer conn.Close()
	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return nil, fmt.Errorf("conn.Do(GET %s) failed:%v", key, err)
	}
	var data oss.AssumeRoleResp
	err = json.Unmarshal([]byte(reply), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
