package config

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txchat/dtalk/pkg/notify"
	"github.com/zeromicro/go-zero/core/conf"
)

func TestParseWhitelistConfig(t *testing.T) {
	var yamlData = `
Whitelist:
  -
    Account: "11111111111"
    Code: "12345"
    Enable: true
`
	type topLevel struct {
		Whitelist []notify.Whitelist
	}
	var c topLevel
	err := conf.LoadFromYamlBytes(bytes.NewBufferString(yamlData).Bytes(), &c)
	assert.Nil(t, err)
	assert.EqualValues(t, &topLevel{
		Whitelist: []notify.Whitelist{
			{
				Account: "11111111111",
				Code:    "12345",
				Enable:  true,
			},
		},
	}, &c)
}

func TestParseSMSConfig(t *testing.T) {
	var yamlData = `
SMS:
  Surl: ""
  AppKey: ""
  SecretKey: ""
  Msg: ""
  Env: ""
  CodeTypes:
    quick: "quick"
    bind: "bind"
    export: "import"
Email:
  Surl: ""
  AppKey: ""
  SecretKey: ""
  Msg: ""
  Env: ""
  CodeTypes:
    quick: "quick"
    bind: "bind"
    export: "import"
`
	type topLevel struct {
		SMS   notify.Config
		Email notify.Config
	}
	var c topLevel
	err := conf.LoadFromYamlBytes(bytes.NewBufferString(yamlData).Bytes(), &c)
	assert.Nil(t, err)
	assert.EqualValues(t, &topLevel{
		SMS: notify.Config{
			Surl:      "",
			AppKey:    "",
			SecretKey: "",
			Msg:       "",
			Env:       "",
			CodeTypes: map[string]string{
				"quick":  "quick",
				"bind":   "bind",
				"export": "import",
			},
		},
		Email: notify.Config{
			Surl:      "",
			AppKey:    "",
			SecretKey: "",
			Msg:       "",
			Env:       "",
			CodeTypes: map[string]string{
				"quick":  "quick",
				"bind":   "bind",
				"export": "import",
			},
		},
	}, &c)
}
