package config

import (
	"flag"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/txchat/dtalk/pkg/net/grpc"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	xtime "github.com/txchat/dtalk/pkg/time"
	xlog "github.com/txchat/im-pkg/log"
)

var (
	confPath string
	regAddrs string
	env      string

	Conf *Config
)

func init() {
	var (
		defAddrs = os.Getenv("REGADDRS")
		defEnv   = os.Getenv("DTALKENV")
	)
	flag.StringVar(&confPath, "conf", "device.toml", "default config path.")
	flag.StringVar(&regAddrs, "reg", defAddrs, "etcd register addrs. eg:127.0.0.1:2379")
	flag.StringVar(&env, "env", defEnv, "service runtime environment")
}

// Init init config.
func Init() (err error) {
	Conf = Default()
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func Default() *Config {
	return &Config{
		Env: env,
		Log: xlog.Config{
			Level:   "debug",
			Mode:    xlog.ConsoleMode,
			Path:    "",
			Display: xlog.JsonDisplay,
		},
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
		Reg: &Reg{
			Schema:   "dtalk",
			SrvName:  "device",
			RegAddrs: regAddrs,
		},
		Redis: &Redis{
			Network:      "tcp",
			Addr:         "127.0.0.1:6379",
			Auth:         "",
			Active:       60000,
			Idle:         1024,
			DialTimeout:  xtime.Duration(200 * time.Millisecond),
			ReadTimeout:  xtime.Duration(500 * time.Millisecond),
			WriteTimeout: xtime.Duration(500 * time.Millisecond),
			IdleTimeout:  xtime.Duration(120 * time.Second),
			Expire:       xtime.Duration(30 * time.Minute),
		},
	}
}

type Config struct {
	Env string
	Log xlog.Config
	//gRPC server
	GRPCServer *grpc.ServerConfig
	Reg        *Reg
	Redis      *Redis
}

// Redis .
type Redis struct {
	Network      string
	Addr         string
	Auth         string
	Active       int
	Idle         int
	DialTimeout  xtime.Duration
	ReadTimeout  xtime.Duration
	WriteTimeout xtime.Duration
	IdleTimeout  xtime.Duration
	Expire       xtime.Duration
}

// Reg is service register/discovery config
type Reg struct {
	Schema   string
	SrvName  string // call
	RegAddrs string // etcd addrs, seperate by ','
}
