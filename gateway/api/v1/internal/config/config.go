package config

import (
	"flag"
	"time"

	"github.com/txchat/dtalk/pkg/slg"

	"github.com/BurntSushi/toml"
	xtime "github.com/txchat/dtalk/pkg/time"
	"github.com/uber/jaeger-client-go"
	traceConfig "github.com/uber/jaeger-client-go/config"
)

var (
	confPath string

	Conf *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "gateway.toml", "default config path.")
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
			Addr: "0.0.0.0:18002",
		},
		Trace: traceConfig.Configuration{
			ServiceName: "gateway",
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
		Revoke: &Revoke{
			Expire: xtime.Duration(time.Hour * 24),
		},
		AnswerRPCClient: &RPCClient{
			RegAddrs: "127.0.0.1:2379",
			Schema:   "dtalk",
			SrvName:  "answer",
			Dial:     xtime.Duration(time.Second),
			Timeout:  xtime.Duration(time.Second),
		},
		StoreRPCClient: &RPCClient{
			RegAddrs: "127.0.0.1:2379",
			Schema:   "dtalk",
			SrvName:  "store",
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
		DeviceRPCClient: &RPCClient{
			RegAddrs: "127.0.0.1:2379",
			Schema:   "dtalk",
			SrvName:  "device",
			Dial:     xtime.Duration(time.Second),
			Timeout:  xtime.Duration(time.Second),
		},
		VIPRPCClient: &RPCClient{
			RegAddrs: "127.0.0.1:2379",
			Schema:   "dtalk",
			SrvName:  "vip",
			Dial:     xtime.Duration(time.Second),
			Timeout:  xtime.Duration(time.Second),
		},
		SlgHTTPClient: &slg.HTTPClientConfig{
			Host: "",
			Salt: "",
		},
	}
}

type Config struct {
	Env             string
	Server          *HttpServer
	Trace           traceConfig.Configuration
	Modules         []Module
	Revoke          *Revoke
	AnswerRPCClient *RPCClient
	StoreRPCClient  *RPCClient
	GroupRPCClient  *RPCClient
	DeviceRPCClient *RPCClient
	VIPRPCClient    *RPCClient
	SlgHTTPClient   *slg.HTTPClientConfig
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

type Module struct {
	Name      string   `json:"name"` // enums: wallet、oa、redpacket
	IsEnabled bool     `json:"isEnabled"`
	EndPoints []string `json:"endPoints"`
}

type Revoke struct {
	Expire xtime.Duration
}
