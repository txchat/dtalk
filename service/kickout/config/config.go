package config

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/txchat/dtalk/pkg/slg"
	xtime "github.com/txchat/dtalk/pkg/time"
	xlog "github.com/txchat/im-pkg/log"
)

var (
	confPath string

	Conf *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "kickout.toml", "default config path.")
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
		Env:      "release",
		TaskSpec: "",
		Log: xlog.Config{
			Level:   "debug",
			Mode:    xlog.ConsoleMode,
			Path:    "",
			Display: xlog.JsonDisplay,
		},
	}
}

type Config struct {
	Env      string
	TaskSpec string
	Log      xlog.Config

	//gRPC Client
	GroupRPCClient *RPCClient
	SlgHTTPClient  *slg.HTTPClientConfig
}

// RPCClient is RPC client config.
type RPCClient struct {
	RegAddrs string // etcd address
	Schema   string
	SrvName  string // call
	Dial     xtime.Duration
	Timeout  xtime.Duration
}
