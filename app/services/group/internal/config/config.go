package config

import (
	xmysql "github.com/txchat/dtalk/pkg/mysql"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MySQL     xmysql.Config
	IDGenRPC  zrpc.RpcClientConf
	AnswerRPC zrpc.RpcClientConf
	LogicRPC  zrpc.RpcClientConf
	Group     Group
	AppID     string
}

type Group struct {
	MaxMembers  int
	MaxManagers int
}
