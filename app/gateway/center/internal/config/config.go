package config

import (
	"github.com/txchat/dtalk/app/gateway/center/internal/types"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	ChatNodes     []*types.ChatNode
	ContractNodes []*types.ContractNode
	Modules       []Module
	VersionRPC    zrpc.RpcClientConf
	BackupRPC     zrpc.RpcClientConf
	Backend       struct {
		Platform string
		Users    []BackendUser
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}

type Module struct {
	Name      string   `json:"name"` // enums: wallet、oa、redpacket
	IsEnabled bool     `json:"isEnabled"`
	EndPoints []string `json:"endPoints"`
}

type BackendUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
