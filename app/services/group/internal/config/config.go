package config

import (
	xmysql "github.com/txchat/dtalk/pkg/mysql"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MySQL          xmysql.Config
	IDGenRPC       zrpc.RpcClientConf
	AnswerRPC      zrpc.RpcClientConf
	LogicRPCClient xgrpc.RPCClient
	Group          Group
	AppID          string
}

type Group struct {
	MaxMembers  int
	MaxManagers int
}
