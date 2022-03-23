package service

import (
	"github.com/txchat/dtalk/pkg/sign/tencentyun"
	"github.com/txchat/dtalk/service/call/model"
	idgen "github.com/txchat/dtalk/service/generator/api"
	answer "github.com/txchat/dtalk/service/record/answer/api"
	"time"

	"github.com/rs/zerolog"
	"github.com/txchat/dtalk/pkg/logger"
	"github.com/txchat/dtalk/service/call/config"
	"github.com/txchat/dtalk/service/call/dao"
)

type Service struct {
	log zerolog.Logger
	cfg *config.Config
	dao *dao.Dao

	idGenRPCClient *idgen.Client
	answerClient   *answer.Client

	room   *model.Room
	tlsSig tencentyun.TLSSig
}

const srvName = "call/srv"

func New(cfg *config.Config) *Service {
	s := &Service{
		log:            logger.New(cfg.Env, srvName),
		cfg:            cfg,
		dao:            dao.New(cfg),
		idGenRPCClient: idgen.New(cfg.IdGenRPCClient.RegAddrs, cfg.IdGenRPCClient.Schema, cfg.IdGenRPCClient.SrvName, time.Duration(cfg.IdGenRPCClient.Dial)),
		answerClient:   answer.New(cfg.AnswerRPCClient.RegAddrs, cfg.AnswerRPCClient.Schema, cfg.AnswerRPCClient.SrvName, time.Duration(cfg.AnswerRPCClient.Dial)),
		room:           model.NewRoom(cfg.Node),
		tlsSig:         tencentyun.NewTCTLSSig(cfg.TCRTCConfig.SDKAppId, cfg.TCRTCConfig.SecretKey, cfg.TCRTCConfig.Expire),
	}
	return s
}

func (s *Service) Ping() error {
	return nil
}

func (s *Service) Config() *config.Config {
	return s.cfg
}
