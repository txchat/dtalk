package service

import (
	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/service/backup/config"
	"github.com/txchat/dtalk/service/backup/dao"
	"github.com/txchat/dtalk/service/backup/model"
	"github.com/txchat/dtalk/service/backup/service/debug"
	"github.com/txchat/dtalk/service/backup/service/email"
	"github.com/txchat/dtalk/service/backup/service/sms"
	"github.com/txchat/dtalk/service/backup/service/whitelist"
)

type Service struct {
	log           log15.Logger
	cfg           *config.Config
	dao           *dao.Dao
	smsValidate   model.Validate
	emailValidate model.Validate
}

func New(cfg *config.Config) *Service {
	s := &Service{
		log: log15.New("module", "backup/svc"),
		cfg: cfg,
		dao: dao.New(cfg),
	}
	//SMS
	{
		var final model.Validate
		final = sms.NewSMS(cfg.SMS.Surl, cfg.SMS.AppKey, cfg.SMS.SecretKey, cfg.SMS.Msg)
		if cfg.SMS.Env == model.Debug {
			final = debug.NewDebugValidate(debug.GetMockCode(cfg.SMS.Msg), final)
		}
		s.smsValidate = whitelist.NewWhitelistValidate(cfg.Whitelist, final)
	}
	//email
	{
		var final model.Validate
		final = email.NewEmail(cfg.Email.Surl, cfg.Email.AppKey, cfg.Email.SecretKey, cfg.Email.Msg)
		if cfg.Email.Env == model.Debug {
			final = debug.NewDebugValidate(debug.GetMockCode(cfg.Email.Msg), final)
		}
		s.emailValidate = whitelist.NewWhitelistValidate(cfg.Whitelist, final)
	}
	return s
}

func (s *Service) Ping() error {
	return nil
}

func (s Service) Config() *config.Config {
	return s.cfg
}
