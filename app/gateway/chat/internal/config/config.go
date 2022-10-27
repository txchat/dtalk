package config

import (
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	xtime "github.com/txchat/dtalk/pkg/time"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Revoke     Revoke
	CallRPC    zrpc.RpcClientConf
	AnswerRPC  zrpc.RpcClientConf
	StorageRPC zrpc.RpcClientConf

	GroupRPCClient xgrpc.RPCClient
}

type Revoke struct {
	Expire xtime.Duration
}
