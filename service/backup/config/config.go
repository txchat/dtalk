package config

import (
	"flag"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/txchat/dtalk/pkg/net/grpc"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	xtime "github.com/txchat/dtalk/pkg/time"
)

var (
	confPath string

	Conf *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "backup.toml", "default config path.")
}

// Init init config.
func Init() (err error) {
	Conf = Default()
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func Default() *Config {
	return &Config{
		Server: &HttpServer{
			Addr: "0.0.0.0:18004",
		},
		Reg: &Reg{
			Schema:   "dtalk",
			SrvName:  "group",
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
		MySQL: &MySQL{
			Host: "127.0.0.1",
			Port: 3306,
			User: "root",
			Pwd:  "123456",
			Db:   "dtalk",
		},
		SMS: SMS{
			Surl:      "",
			AppKey:    "",
			SecretKey: "",
			Msg:       "",
			Env:       "debug",
			CodeTypes: map[string]string{},
		},
		Email: Email{
			Surl:      "",
			AppKey:    "",
			SecretKey: "",
			Msg:       "",
			Env:       "debug",
			CodeTypes: map[string]string{},
		},
		Whitelist: []SMSEmailWhitelist{
			{},
		},
	}
}

type Config struct {
	Env    string
	Server *HttpServer
	Reg    *Reg
	//gRPC server
	GRPCServer *grpc.ServerConfig
	MySQL      *MySQL
	SMS        SMS
	Email      Email
	Whitelist  []SMSEmailWhitelist
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

type SMS struct {
	Surl      string
	AppKey    string
	SecretKey string
	Msg       string
	Env       string
	CodeTypes map[string]string
}

type Email struct {
	Surl      string
	AppKey    string
	SecretKey string
	Msg       string
	Env       string
	CodeTypes map[string]string
}

type SMSEmailWhitelist struct {
	Account string
	Code    string
	Enable  bool
}
