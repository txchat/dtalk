package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/txchat/dtalk/pkg/net/grpc"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	xtime "github.com/txchat/dtalk/pkg/time"
	"time"
)

var (
	confPath string

	Conf *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "auth.toml", "default config path.")
}

// Init init config.
func Init() (err error) {
	Conf = Default()
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func Default() *Config {
	return &Config{
		Env: "debug",
		Server: &HttpServer{
			Addr: "127.0.0.1:8080",
		},
		Reg: &Reg{
			Schema:   "dtalk",
			SrvName:  "auth",
			RegAddrs: "127.0.0.1:2379",
		},
		GRPCServer: &xgrpc.ServerConfig{
			Network:                           "tcp",
			Addr:                              ":18012",
			Timeout:                           xtime.Duration(time.Second),
			KeepAliveMaxConnectionIdle:        xtime.Duration(time.Second * 60),
			KeepAliveMaxConnectionAge:         xtime.Duration(time.Hour * 2),
			KeepAliveMaxMaxConnectionAgeGrace: xtime.Duration(time.Second * 20),
			KeepAliveTime:                     xtime.Duration(time.Second * 60),
			KeepAliveTimeout:                  xtime.Duration(time.Second * 20),
		},
	}
}

type Config struct {
	Env    string
	Server *HttpServer
	MySQL  *MySQL
	Reg    *Reg
	//gRPC server
	GRPCServer *grpc.ServerConfig
	Key        *Key
}

type HttpServer struct {
	Addr string
}

type MySQL struct {
	Host string
	Port int32
	User string
	Pwd  string
	Db   string
}

// Reg is service register/discovery config
type Reg struct {
	Schema   string
	SrvName  string // call
	RegAddrs string // etcd addrs, seperate by ','
}

type Key struct {
	KeyExpireDuration time.Duration
}
