package service

import (
	"time"

	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/service/backend/config"
	"github.com/txchat/dtalk/service/backend/dao"
	"github.com/txchat/dtalk/service/backend/service/cdk"
	idgen "github.com/txchat/dtalk/service/generator/api"
)

type Service struct {
	log            log15.Logger
	cfg            *config.Config
	dao            *dao.Dao
	Platform       string
	idGenRPCClient *idgen.Client

	CdkService *cdk.ServiceContent
}

func New(cfg *config.Config) *Service {
	s := &Service{
		log:            log15.New("module", "backend/svc"),
		cfg:            cfg,
		dao:            dao.New(cfg),
		Platform:       cfg.Platform,
		idGenRPCClient: idgen.New(cfg.IdGenRPCClient.RegAddrs, cfg.IdGenRPCClient.Schema, cfg.IdGenRPCClient.SrvName, time.Duration(cfg.IdGenRPCClient.Dial)),
	}

	if cfg.CdkMod {
		s.CdkService = cdk.NewServiceContent(cfg.Env, s.dao, s.idGenRPCClient, cfg.CdkMaxNumber, cfg.Chain33Client)
	}

	return s
}

func (s *Service) Ping() error {
	return nil
}

func (s Service) Config() *config.Config {
	return s.cfg
}
