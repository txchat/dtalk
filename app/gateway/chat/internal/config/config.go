package config

import (
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
	GroupRPC   zrpc.RpcClientConf
}

type Revoke struct {
	Expire xtime.Duration
}
