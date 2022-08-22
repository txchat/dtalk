package config

import (
	"github.com/txchat/dtalk/app/gateway/center/internal/types"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	ChatNodes     []*types.ChatNode
	ContractNodes []*types.ContractNode
	Modules       []Module
}

type Module struct {
	Name      string   `json:"name"` // enums: wallet、oa、redpacket
	IsEnabled bool     `json:"isEnabled"`
	EndPoints []string `json:"endPoints"`
}
