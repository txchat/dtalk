package service

import (
	"context"

	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/service/discovery/config"
	"github.com/txchat/dtalk/service/discovery/dao"
)

type Service struct {
	log log15.Logger
	cfg *config.Config
	dao *dao.Dao
}

func New(cfg *config.Config) *Service {
	s := &Service{
		log: log15.New("module", "discovery/svc"),
		cfg: cfg,
		dao: dao.New(cfg),
	}
	s.StoreNodes(cfg.CNodes, cfg.DNodes)
	return s
}

func Ping(c context.Context) error {
	return nil
}

func (s *Service) Config() *config.Config {
	return s.cfg
}
