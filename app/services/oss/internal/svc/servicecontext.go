package svc

import (
	"github.com/txchat/dtalk/app/services/oss/internal/config"
	"github.com/txchat/dtalk/app/services/oss/internal/dao"
	"github.com/txchat/dtalk/app/services/oss/internal/model"
	"github.com/txchat/dtalk/pkg/oss"
	"github.com/txchat/dtalk/pkg/oss/aliyun"
	"github.com/txchat/dtalk/pkg/oss/huaweiyun"
	"github.com/txchat/dtalk/pkg/oss/minio"
)

type ServiceContext struct {
	Config        config.Config
	Repo          dao.OssRepository
	AppOssEngines *model.AppOssEngine
}

func NewServiceContext(c config.Config) *ServiceContext {
	svcCtx := &ServiceContext{
		Config:        c,
		Repo:          dao.NewOssRepositoryRedis(c.RedisDB),
		AppOssEngines: model.NewAppOssManager(),
	}
	engineInit(svcCtx)
	return svcCtx
}

func engineInit(svcCtx *ServiceContext) {
	for _, o := range svcCtx.Config.Oss {
		appId := o.AppId
		ossType := o.OssType
		ossCfg := &oss.Config{
			RegionId:        o.RegionId,
			AccessKeyId:     o.AccessKeyId,
			AccessKeySecret: o.AccessKeySecret,
			Role:            o.Role,
			Policy:          o.Policy,
			DurationSeconds: o.DurationSeconds,
			Bucket:          o.Bucket,
			EndPoint:        o.EndPoint,
			PublicUrl:       o.PublicURL,
		}

		var invoker oss.Oss
		switch ossType {
		case model.OssAliyun:
			invoker = aliyun.New(ossCfg)
		case model.OssHuaweiuyn:
			invoker = huaweiyun.New(ossCfg)
		case model.OssMinio:
			invoker = minio.New(ossCfg)
		default:
			continue
		}

		svcCtx.AppOssEngines.Init(appId, ossType, invoker)
	}
}
