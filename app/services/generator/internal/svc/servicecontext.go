package svc

import (
	"github.com/txchat/dtalk/app/services/generator/internal/config"
	"github.com/txchat/dtalk/pkg/util"
)

type ServiceContext struct {
	Config      config.Config
	IDGenerator *util.Snowflake
}

func NewServiceContext(c config.Config) *ServiceContext {
	//TODO 节点id根据横向扩容去设置，暂时支持1个节点
	g, err := util.NewSnowflake(1)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:      c,
		IDGenerator: g,
	}
}
