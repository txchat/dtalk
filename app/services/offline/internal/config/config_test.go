package config

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/conf"
)

func TestParsePusherConfig(t *testing.T) {
	var pusherYamlData = `
Pushers: # 友盟离线推送物料
  Android:
    Env: "debug"
    AppKey: ""
    AppMasterSecret: ""
    MiActivity: ""
  iOS:
    Env: "debug"
    AppKey: ""
    AppMasterSecret: ""
    MiActivity: ""
`
	type topLevel struct {
		Pushers map[string]Pusher
	}
	var c topLevel
	err := conf.LoadFromYamlBytes(bytes.NewBufferString(pusherYamlData).Bytes(), &c)
	assert.Nil(t, err)
	assert.EqualValues(t, &topLevel{
		Pushers: map[string]Pusher{
			"Android": {
				Env:             "debug",
				AppKey:          "",
				AppMasterSecret: "",
				MiActivity:      "",
			},
			"iOS": {
				Env:             "debug",
				AppKey:          "",
				AppMasterSecret: "",
				MiActivity:      "",
			},
		},
	}, &c)
}
