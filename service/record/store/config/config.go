package config

import (
	"flag"
	"io/ioutil"
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
	flag.StringVar(&confPath, "conf", "store.toml", "default config path.")
	flag.StringVar(&regAddrs, "reg", defAddrs, "etcd register addrs. eg:127.0.0.1:2379")
	flag.StringVar(&env, "env", defEnv, "service runtime environment")
}

// Init init config.
func Init() (err error) {
	Conf = Default()
	//TODO 使用环境变量填充，临时处理
	bs, err := ioutil.ReadFile(confPath)
	if err != nil {
		return err
	}
	filled := os.ExpandEnv(string(bs))
	_, err = toml.Decode(filled, &Conf)
	//_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func Default() *Config {
	return &Config{
		AppId:     "dtalk",
		Engine:    "standard",
		Env:       env,
		SyncCache: true,
		Log: xlog.Config{
			Level:   "debug",
			Mode:    xlog.ConsoleMode,
			Path:    "",
			Display: xlog.JsonDisplay,
		},
		Trace: traceConfig.Configuration{
			ServiceName: "store",
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
		Reg: &Reg{
			Schema:   "dtalk",
			SrvName:  "store",
			RegAddrs: regAddrs,
		},
		GRPCServer: &xgrpc.ServerConfig{
			Network:                           "tcp",
			Addr:                              ":30005",
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
		GroupRPCClient: &RPCClient{
			RegAddrs: "127.0.0.1:2379",
			Schema:   "dtalk",
			SrvName:  "group",
			Dial:     xtime.Duration(time.Second),
			Timeout:  xtime.Duration(time.Second),
		},
		PusherRPCClient: &RPCClient{
			RegAddrs: "127.0.0.1:2379",
			Schema:   "dtalk",
			SrvName:  "172.16.101.126:30003",
			Dial:     xtime.Duration(time.Second),
			Timeout:  xtime.Duration(time.Second),
		},
		DeviceRPCClient: &RPCClient{
			RegAddrs: "127.0.0.1:2379",
			Schema:   "dtalk",
			SrvName:  "device",
			Dial:     xtime.Duration(time.Second),
			Timeout:  xtime.Duration(time.Second),
		},
		RevSub: &MQSubClient{
			Brokers:   []string{"127.0.0.1:9092"},
			Number:    16,
			MaxWorker: 1024,
		},
		StoreSub: &MQSubClient{
			Brokers:   []string{"127.0.0.1:9092"},
			Number:    16,
			MaxWorker: 1024,
		},
	}
}

type Config struct {
	AppId           string
	Engine          string
	Env             string
	SyncCache       bool
	Log             xlog.Config
	Trace           traceConfig.Configuration
	GRPCServer      *grpc.ServerConfig
	Reg             *Reg
	MySQL           *MySQL
	Redis           *Redis
	GroupRPCClient  *RPCClient
	PusherRPCClient *RPCClient
	DeviceRPCClient *RPCClient
	RevSub          *MQSubClient
	StoreSub        *MQSubClient
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
