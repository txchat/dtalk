package config

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/conf"
)

func TestParsePusherConfig(t *testing.T) {
	var pusherYamlData = `
Timeout: 10m
Pushers: # 友盟离线推送物料
  android:
    Env: "debug"
    AppKey: ""
    AppMasterSecret: ""
    MiActivity: ""
  ios:
    Env: "debug"
    AppKey: ""
    AppMasterSecret: ""
    MiActivity: ""
`
	type topLevel struct {
		Timeout time.Duration
		Pushers map[string]Pusher
	}
	var c topLevel
	err := conf.LoadFromYamlBytes(bytes.NewBufferString(pusherYamlData).Bytes(), &c)
	assert.Nil(t, err)
	assert.EqualValues(t, &topLevel{
		Timeout: time.Minute * 10,
		Pushers: map[string]Pusher{
			"android": {
				Env:             "debug",
				AppKey:          "",
				AppMasterSecret: "",
				MiActivity:      "",
			},
			"ios": {
				Env:             "debug",
				AppKey:          "",
				AppMasterSecret: "",
				MiActivity:      "",
			},
		},
	}, &c)
}
