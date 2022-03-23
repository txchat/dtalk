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
	"github.com/uber/jaeger-client-go"
	traceConfig "github.com/uber/jaeger-client-go/config"
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
	flag.StringVar(&confPath, "conf", "pusher.toml", "default config path.")
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
		AppId:  "dtalk",
		Engine: "standard",
		Env:    env,
		Log: xlog.Config{
			Level:   "debug",
			Mode:    xlog.ConsoleMode,
			Path:    "",
			Display: xlog.JsonDisplay,
		},
		Trace: traceConfig.Configuration{
			ServiceName: "pusher",
			Gen128Bit:   true,
			Sampler: &traceConfig.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &traceConfig.ReporterConfig{
				LogSpans:           true,
				LocalAgentHostPort: "127.0.0.1:6831",
			},
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
			SrvName:  "pusher",
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
		GroupRPCClient: &RPCClient{
			RegAddrs: "127.0.0.1:2379",
			Schema:   "dtalk",
			SrvName:  "group",
			Dial:     xtime.Duration(time.Second),
			Timeout:  xtime.Duration(time.Second),
		},
		AnswerRPCClient: &RPCClient{
			RegAddrs: "127.0.0.1:2379",
			Schema:   "dtalk",
			SrvName:  "answer",
			Dial:     xtime.Duration(time.Second),
			Timeout:  xtime.Duration(time.Second),
		}, DeviceRPCClient: &RPCClient{
			RegAddrs: "127.0.0.1:2379",
			Schema:   "dtalk",
			SrvName:  "device",
			Dial:     xtime.Duration(time.Second),
			Timeout:  xtime.Duration(time.Second),
		},
		IMSub: &MQSubClient{
			Brokers:   []string{"127.0.0.1:9092"},
			Number:    16,
			MaxWorker: 1024,
		},
		RevSub: &MQSubClient{
			Brokers:   []string{"127.0.0.1:9092"},
			Number:    16,
			MaxWorker: 1024,
		},
		StorePub: &MQPubServer{
			Brokers: []string{"127.0.0.1:9092"},
		},
		OffPush: &OffPush{
			IsEnabled: true,
			OffPushPub: &MQPubServer{
				Brokers: []string{"127.0.0.1:9092"},
			},
		},
	}
}

type Config struct {
	AppId           string
	Engine          string
	Env             string
	Log             xlog.Config
	Trace           traceConfig.Configuration
	LogicRPCClient  *RPCClient
	IdGenRPCClient  *RPCClient
	GroupRPCClient  *RPCClient
	AnswerRPCClient *RPCClient
	DeviceRPCClient *RPCClient
	//gRPC server
	GRPCServer *grpc.ServerConfig
	Reg        *Reg
	Redis      *Redis
	IMSub      *MQSubClient
	RevSub     *MQSubClient
	StorePub   *MQPubServer
	OffPush    *OffPush
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

// RPCClient is RPC client config.
type RPCClient struct {
	RegAddrs string // etcd addrs, seperate by ','
	Schema   string
	SrvName  string // call
	Dial     xtime.Duration
	Timeout  xtime.Duration
}

type MQSubClient struct {
	Brokers   []string
	Number    uint32
	MaxWorker int
}

type OffPush struct {
	IsEnabled  bool
	OffPushPub *MQPubServer
}

// Kafka .
type MQPubServer struct {
	Brokers []string
}
