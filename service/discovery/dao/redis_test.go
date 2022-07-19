package dao

import (
	"os"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/inconshreveable/log15"
	xtime "github.com/txchat/dtalk/pkg/time"
	"github.com/txchat/dtalk/service/discovery/config"
	"github.com/txchat/dtalk/service/discovery/model"
)

var test_redis_pool *redis.Pool

func TestMain(m *testing.M) {
	c := &config.Redis{
		Network:      "tcp",
		Addr:         "172.24.143.30:6379",
		Auth:         "",
		Active:       60000,
		Idle:         1024,
		DialTimeout:  xtime.Duration(200 * time.Millisecond),
		ReadTimeout:  xtime.Duration(500 * time.Millisecond),
		WriteTimeout: xtime.Duration(500 * time.Millisecond),
		IdleTimeout:  xtime.Duration(120 * time.Second),
		Expire:       xtime.Duration(30 * time.Minute),
	}

	test_redis_pool = &redis.Pool{
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
	os.Exit(m.Run())
}

func TestDao_GetCNodes(t *testing.T) {
	type fields struct {
		log   log15.Logger
		redis *redis.Pool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.CNode
		wantErr bool
	}{
		{
			name: "test get all",
			fields: fields{
				log:   log15.New(""),
				redis: test_redis_pool,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				log:   tt.fields.log,
				redis: tt.fields.redis,
			}
			got, err := d.GetCNodes()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, node := range got {
				t.Logf("item %v = %v", i, node)
			}
		})
	}
}
