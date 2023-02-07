package redis

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/conf"
)

var redisYamlContentExample = `
RedisDB:
  Network: "tcp"
  Addr: "127.0.0.1:6379"
  Auth: ""
  Active: 60000
  Idle: 1024
  DialTimeout: "200ms"
  ReadTimeout: "500ms"
  WriteTimeout: "500ms"
  IdleTimeout: "120s"
  Expire: "30m"
`

func TestParseYamlConfig(t *testing.T) {
	type topLevel struct {
		RedisDB Config
	}
	var c topLevel
	err := conf.LoadFromYamlBytes(bytes.NewBufferString(redisYamlContentExample).Bytes(), &c)
	assert.Nil(t, err)
	assert.EqualValues(t, &topLevel{
		RedisDB: Config{
			Network:      "tcp",
			Addr:         "127.0.0.1:6379",
			Auth:         "",
			Active:       60000,
			Idle:         1024,
			DialTimeout:  200 * time.Millisecond,
			ReadTimeout:  500 * time.Millisecond,
			WriteTimeout: 500 * time.Millisecond,
			IdleTimeout:  120 * time.Second,
			Expire:       30 * time.Minute,
		},
	}, &c)
}
