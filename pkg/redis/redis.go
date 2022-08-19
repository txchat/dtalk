// From https://gitlab.33.cn/proof/backend-micro/blob/dev/pkg/gredis/gredis.go

package redis

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	xtime "github.com/txchat/dtalk/pkg/time"
)

// Config .
type Config struct {
	Network      string
	Addr         string
	Auth         string
	Active       int
	Idle         int
	DialTimeout  xtime.Duration
	ReadTimeout  xtime.Duration
	WriteTimeout xtime.Duration
	IdleTimeout  xtime.Duration
	Expire       xtime.Duration
}

func NewPool(c Config) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     c.Idle,
		MaxActive:   c.Active,
		IdleTimeout: time.Duration(c.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(c.Network, c.Addr,
				redis.DialConnectTimeout(time.Duration(c.DialTimeout)),
				redis.DialReadTimeout(time.Duration(c.ReadTimeout)),
				redis.DialWriteTimeout(time.Duration(c.WriteTimeout)),
				redis.DialPassword(c.Auth),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

// Pool Redis缓存池结构
type Pool struct {
	pool *redis.Pool
	sync.RWMutex
}

func New(c *Config) *Pool {
	pool := &redis.Pool{
		MaxIdle:     c.Idle,
		MaxActive:   c.Active,
		IdleTimeout: time.Duration(c.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(c.Network, c.Addr,
				redis.DialConnectTimeout(time.Duration(c.DialTimeout)),
				redis.DialReadTimeout(time.Duration(c.ReadTimeout)),
				redis.DialWriteTimeout(time.Duration(c.WriteTimeout)),
				redis.DialPassword(c.Auth),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
	return &Pool{pool: pool}
}

// Do 向Redis服务发送命令并返回收到的答复
func (p *Pool) Do(cmd string, args ...interface{}) (interface{}, error) {
	conn := p.pool.Get()
	if err := conn.Err(); err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.Do(cmd, args...)
}

// Set 将数据data关联到给定key，time为key的超时时间（秒）
func (p *Pool) Set(key string, data interface{}, time ...int) error {
	var err error
	if len(time) != 0 {
		_, err = p.Do("SET", key, data, "EX", time[0])
	} else {
		_, err = p.Do("SET", key, data)
	}
	return err
}

// Exists 检查给定key是否存在
func (p *Pool) Exists(key string) (bool, error) {
	return redis.Bool(p.Do("EXISTS", key))
}

// GetBytes 返回给定key所关联的[]byte值
func (p *Pool) GetBytes(key string) ([]byte, error) {
	return redis.Bytes(p.Do("GET", key))
}

// GetString 返回给定key所关联的string值
func (p *Pool) GetString(key string) (string, error) {
	return redis.String(p.Do("GET", key))
}

// GetInt 返回给定key所关联的int值
func (p *Pool) GetInt(key string) (int, error) {
	return redis.Int(p.Do("GET", key))
}

// GetInt64 返回给定key所关联的int64值
func (p *Pool) GetInt64(key string) (int64, error) {
	return redis.Int64(p.Do("GET", key))
}

// Read 将给定key所关联的值反序列化到obj对象
func (p *Pool) Read(key string, obj interface{}) error {
	if data, err := p.GetBytes(key); err == nil {
		return json.Unmarshal(data, obj)
	} else {
		return err
	}
}

// Write 将数据data序列化后关联到给定key，time为key的超时时间（秒）
func (p *Pool) Write(key string, obj interface{}, time ...int) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	if len(time) != 0 {
		_, err = p.Do("SET", key, data, "EX", time[0])
	} else {
		_, err = p.Do("SET", key, data)
	}
	return err
}

// Del 删除给定key
func (p *Pool) Del(key string) (bool, error) {
	return redis.Bool(p.Do("DEL", key))
}

// LikeDel 模糊删除给定key
//func (p *Pool) LikeDel(key string) error {
//	keys, err := redis.Strings(p.Do("KEYS", "*"+key+"*"))
//	if err != nil {
//		return err
//	}
//	_, err = redis.Bool(p.Do("DEL", sliceUtil.StringsToInterfaces(keys)...))
//	return err
//}

// TTL 返回给定key的剩余生存时间（秒）
func (p *Pool) TTL(key string) (int, error) {
	return redis.Int(p.Do("TTL", key))
}

// Incr 将key所储存的数字值增一
func (p *Pool) Incr(key string) (int64, error) {
	return redis.Int64(p.Do("INCR", key))
}

// IncrBy 将key所储存的数字值加上增量increment
func (p *Pool) IncrBy(key string, increment int64) (int64, error) {
	return redis.Int64(p.Do("INCRBY", key, increment))
}

// Decr 将key所储存的数字值减一
func (p *Pool) Decr(key string) (int64, error) {
	return redis.Int64(p.Do("DECR", key))
}

// DecrBy 将key所储存的数字值减去减量decrement
func (p *Pool) DecrBy(key string, decrement int64) (int64, error) {
	return redis.Int64(p.Do("DECRBY", key, decrement))
}
