package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/txchat/dtalk/pkg/net/grpc"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	"github.com/txchat/dtalk/pkg/redis"
	xtime "github.com/txchat/dtalk/pkg/time"
	"time"
)

var (
	confPath string

	Conf *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "group.toml", "default config path.")
}

// Init init config.
func Init() (err error) {
	Conf = Default()
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func Default() *Config {
	return &Config{
		HttpServer: &HttpServer{
			Addr: "0.0.0.0:18011",
		},
		MySQL: &MySQL{
			Host: "127.0.0.1",
			Port: 3306,
			User: "root",
			Pwd:  "123456",
			Db:   "dtalk",
		},
		Reg: &Reg{
			Schema:   "dtalk",
			SrvName:  "group",
			RegAddrs: "127.0.0.1:2379",
		},
		LogicRPCClient: &RPCClient{
			RegAddrs: "127.0.0.1:2379",
			Schema:   "im",
			SrvName:  "logic",
			Dial:     xtime.Duration(time.Second),
			Timeout:  xtime.Duration(time.Second),
		},
		IdGenRPCClient: &RPCClient{
			RegAddrs: "127.0.0.1:2379",
			Schema:   "dtalk",
			SrvName:  "generator",
			Dial:     xtime.Duration(time.Second),
			Timeout:  xtime.Duration(time.Second),
		},
		AnswerRPCClient: &RPCClient{
			RegAddrs: "127.0.0.1:2379",
			Schema:   "dtalk",
			SrvName:  "answer",
			Dial:     xtime.Duration(time.Second),
			Timeout:  xtime.Duration(time.Second),
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
		GroupInfoConfig: &GroupDefault{
			GroupMaximum: 200,
			AdminNum:     10,
		},
		Redis: &redis.Config{
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
	Env            string
	AppId          string
	HttpServer     *HttpServer
	MySQL          *MySQL
	Reg            *Reg
	LogicRPCClient *RPCClient
	IdGenRPCClient *RPCClient
	//PusherRPCClient *RPCClient
	AnswerRPCClient *RPCClient
	//gRPC server
	GRPCServer      *grpc.ServerConfig
	GroupInfoConfig *GroupDefault
	Redis           *redis.Config
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

// RPCClient is RPC client config.
type RPCClient struct {
	RegAddrs string
	Schema   string
	SrvName  string // call
	Dial     xtime.Duration
	Timeout  xtime.Duration
}

type GroupDefault struct {
	// 群人数上限 [200, 2000]
	GroupMaximum int32
	// 群管理员人数上限 [10, 10]
	AdminNum int32
}
