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
)

var (
	confPath string

	Conf *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "vip.toml", "default config path.")
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
		Env: "release",
		Log: xlog.Config{
			Level:   "debug",
			Mode:    xlog.ConsoleMode,
			Path:    "",
			Display: xlog.JsonDisplay,
		},
		GRPCServer: &xgrpc.ServerConfig{
			Network:                           "tcp",
			Addr:                              ":30006",
			Timeout:                           xtime.Duration(time.Second),
			KeepAliveMaxConnectionIdle:        xtime.Duration(time.Second * 60),
			KeepAliveMaxConnectionAge:         xtime.Duration(time.Hour * 2),
			KeepAliveMaxMaxConnectionAgeGrace: xtime.Duration(time.Second * 20),
			KeepAliveTime:                     xtime.Duration(time.Second * 60),
			KeepAliveTimeout:                  xtime.Duration(time.Second * 20),
		},
		Reg: &Reg{
			Schema:   "dtalk",
			SrvName:  "vip",
			RegAddrs: "127.0.0.1:2379",
		},
		MySQL: &MySQL{
			Host: "127.0.0.1",
			Port: 3306,
			User: "root",
			Pwd:  "123456",
			Db:   "dtalk",
		},
	}
}

type Config struct {
	Env string
	Log xlog.Config
	Reg *Reg
	//gRPC server
	GRPCServer *grpc.ServerConfig
	MySQL      *MySQL
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
	RegAddrs string // etcd address, be separated from ','
}
