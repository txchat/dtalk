package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/txchat/dtalk/pkg/redis"
	xtime "github.com/txchat/dtalk/pkg/time"
	"time"
)

var (
	confPath string

	Conf *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "call.toml", "default config path.")
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
			Addr: "0.0.0.0:18013",
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
	}
}

type Config struct {
	Env             string
	AppId           string
	Node            int32
	HttpServer      *HttpServer
	Redis           *redis.Config
	IdGenRPCClient  *RPCClient
	AnswerRPCClient *RPCClient
	//gRPC server
	TCRTCConfig *RTCConfig
}

type HttpServer struct {
	Addr string
}

// RPCClient is RPC client config.
type RPCClient struct {
	RegAddrs string // etcd addrs, seperate by ','
	Schema   string
	SrvName  string // call
	Dial     xtime.Duration
	Timeout  xtime.Duration
}

// RTCConfig 腾讯云音视频控制台配置
type RTCConfig struct {
	SDKAppId  int
	SecretKey string
	Expire    int
}
