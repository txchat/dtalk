package config

import (
	"github.com/txchat/dtalk/app/gateway/center/internal/types"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	ChatNodes     []*types.ChatNode
	ContractNodes []*types.ContractNode
}
