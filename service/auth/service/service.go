package service

import (
	"net/http"

	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/service/auth/config"
	"github.com/txchat/dtalk/service/auth/dao"
	"github.com/txchat/dtalk/service/auth/model"
)

type Service struct {
	log        log15.Logger
	cfg        *config.Config
	dao        *dao.Dao
	httpClient *http.Client
}

func New(cfg *config.Config) *Service {
	s := &Service{
		log: log15.New("module", "auth/svc"),
		cfg: cfg,
		dao: dao.New(cfg),
		httpClient: &http.Client{
			Timeout: model.Timeout,
		},
	}
	return s
}

func (s *Service) Ping() error {
	return nil
}

func (s Service) Config() *config.Config {
	return s.cfg
}
