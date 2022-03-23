package config

import (
	"flag"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/txchat/dtalk/pkg/net/grpc"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	xtime "github.com/txchat/dtalk/pkg/time"
)

var (
	confPath string
	regAddrs string

	// Conf config
	Conf *Config
)

func init() {
	var (
		defAddrs = os.Getenv("REGADDRS")
	)
	flag.StringVar(&confPath, "conf", "generator.toml", "default config path.")
	flag.StringVar(&regAddrs, "reg", defAddrs, "etcd register addrs. eg:127.0.0.1:2379")
}

// Init init config.
func Init() (err error) {
	Conf = Default()
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

// Default new a config with specified defualt value.
func Default() *Config {
	return &Config{
		Node: 1,
		GRPCServer: &xgrpc.ServerConfig{
			Network:                           "tcp",
			Addr:                              ":30002",
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
	Node int64
	//reg
	Reg *Reg
	//gRPC server
	GRPCServer *grpc.ServerConfig
}

// Reg is service register/discovery config
type Reg struct {
	Schema   string
	SrvName  string
	RegAddrs string // etcd addrs, seperate by ','
}
