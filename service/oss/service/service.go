package service

import (
	"github.com/rs/zerolog"
	"github.com/txchat/dtalk/pkg/logger"
	"github.com/txchat/dtalk/pkg/oss"
	"github.com/txchat/dtalk/pkg/oss/aliyun"
	"github.com/txchat/dtalk/pkg/oss/huaweiyun"
	"github.com/txchat/dtalk/pkg/oss/minio"
	"github.com/txchat/dtalk/service/oss/config"
	"github.com/txchat/dtalk/service/oss/dao"
	"github.com/txchat/dtalk/service/oss/model"
)

type Service struct {
	log       zerolog.Logger
	cfg       *config.Config
	dao       *dao.Dao
	ossEngine map[string]*model.App
}

func New(cfg *config.Config) *Service {
	s := &Service{
		log:       logger.New(cfg.Env, "oss"),
		cfg:       cfg,
		dao:       dao.New(cfg),
		ossEngine: make(map[string]*model.App),
	}
	s.engineInit()
	return s
}

func (s *Service) Ping() error {
	return nil
}

func (s Service) Config() *config.Config {
	return s.cfg
}

func (s *Service) engineInit() {
	for _, cfg := range s.cfg.Oss {
		var invoker oss.Oss
		ossCfg := convert2Oss(cfg)
		switch cfg.OssType {
		case model.Oss_Aliyun:
			invoker = aliyun.New(ossCfg)
		case model.Oss_Huaweiuyn:
			invoker = huaweiyun.New(ossCfg)
		case model.Oss_Minio:
			invoker = minio.New(ossCfg)
		default:
			s.log.Error().Str("cfg.OssType", cfg.OssType).Msg("oss plugin not find")
			continue
		}

		app, ok := s.ossEngine[cfg.AppId]
		if !ok {
			app = model.NewApp(cfg.AppId)
		}
		app.DefaultOssType = cfg.OssType
		app.Register(cfg.OssType, invoker)
		s.ossEngine[cfg.AppId] = app
	}
}

func (s *Service) GetEngine(appId string) *model.App {
	return s.ossEngine[appId]
}

func convert2Oss(ossCfg *config.Oss) *oss.Config {
	ossConfigs := &oss.Config{}
	ossConfigs.RegionId = ossCfg.RegionId
	ossConfigs.AccessKeyId = ossCfg.AccessKeyId
	ossConfigs.AccessKeySecret = ossCfg.AccessKeySecret
	ossConfigs.Role = ossCfg.Role
	ossConfigs.Policy = ossCfg.Policy
	ossConfigs.DurationSeconds = ossCfg.DurationSeconds
	ossConfigs.Bucket = ossCfg.Bucket
	ossConfigs.EndPoint = ossCfg.EndPoint
	ossConfigs.PublicUrl = ossCfg.PublicUrl

	return ossConfigs
}
