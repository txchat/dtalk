package config

import (
	"bytes"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/conf"
)

func TestReadYamlConfig(t *testing.T) {
	var yamlConfiglData = `
MySQL:
  Net: "tcp"
  Addr: "${MYSQL_HOST}:3306"
  User: "root"
  Passwd: "${MYSQL_PASSWORD}"
  DBName: "dtalk"
`
	type topLevel struct {
		MySQL mysql.Config `json:",optional"`
	}
	var c topLevel
	err := conf.LoadFromYamlBytes(bytes.NewBufferString(yamlConfiglData).Bytes(), &c)
	assert.Nil(t, err)
	assert.EqualValues(t, &topLevel{
		MySQL: mysql.Config{
			User:   "root",
			Passwd: "${MYSQL_PASSWORD}",
			Net:    "tcp",
			Addr:   "${MYSQL_HOST}:3306",
			DBName: "dtalk",
			Params: map[string]string{},
		},
	}, &c)
}
