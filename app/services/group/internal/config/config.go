package config

import (
	"github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MySQL       mysql.Config
	IDGenRPC    zrpc.RpcClientConf
	TransferRPC zrpc.RpcClientConf
	LogicRPC    zrpc.RpcClientConf
	Group       Group
	AppID       string
}

type Group struct {
	MaxMembers  int
	MaxManagers int
}
