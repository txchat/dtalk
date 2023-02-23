package config

import (
	"time"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Revoke      Revoke
	CallRPC     zrpc.RpcClientConf
	LogicRPC    zrpc.RpcClientConf
	TransferRPC zrpc.RpcClientConf
	StorageRPC  zrpc.RpcClientConf
	GroupRPC    zrpc.RpcClientConf
	OssRPC      zrpc.RpcClientConf
	DeviceRPC   zrpc.RpcClientConf
}

type Revoke struct {
	Expire time.Duration
}
