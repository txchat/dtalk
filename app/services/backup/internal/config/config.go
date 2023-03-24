package config

import (
	"github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MySQL mysql.Config
}
