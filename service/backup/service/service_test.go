package service

import (
	"os"
	"testing"

	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/service/backup/config"
	"github.com/txchat/dtalk/service/backup/dao"
	"github.com/txchat/dtalk/service/backup/model"
	"github.com/txchat/dtalk/service/backup/service/debug"
	"github.com/txchat/dtalk/service/backup/service/email"
	"github.com/txchat/dtalk/service/backup/service/sms"
)

var (
	testLog           log15.Logger
	testCfg           *config.Config
	testDao           *dao.Dao
	testSmsValidate   model.Validate
	testEmailValidate model.Validate
)

func TestMain(m *testing.M) {
	cfg := &config.Config{
		Server: &config.HttpServer{
			Addr: "0.0.0.0:18004",
		},
		Env: "debug",
		MySQL: &config.MySQL{
			Host: "127.0.0.1",
			Port: 3306,
			User: "root",
			Pwd:  "123456",
			Db:   "dtalk",
		},
		SMS: config.SMS{
			Surl:      "",
			AppKey:    "",
			SecretKey: "",
			Msg:       "",
			Env:       "",
		},
		Email: config.Email{
			Surl:      "",
			AppKey:    "",
			SecretKey: "",
			Msg:       "",
			Env:       "",
		},
	}
	testLog = log15.New("test", "backup/svc")
	testCfg = cfg
	testDao = dao.New(cfg)
	testSmsValidate = debug.NewDebugValidate("111111", sms.NewSMS(cfg.SMS.Surl, cfg.SMS.AppKey, cfg.SMS.SecretKey, cfg.SMS.Msg))
	testEmailValidate = debug.NewDebugValidate("111111", email.NewEmail(cfg.Email.Surl, cfg.Email.AppKey, cfg.Email.SecretKey, cfg.Email.Msg))
	os.Exit(m.Run())
}
