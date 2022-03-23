package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	xtime "github.com/txchat/dtalk/pkg/time"
	"time"
)

var (
	confPath string

	Conf *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "backend.toml", "default config path.")
}

// Init init config.
func Init() (err error) {
	Conf = Default()
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func Default() *Config {
	return &Config{
		Env:      "debug",
		Platform: "chat33pro",
		Server: &HttpServer{
			Addr: "0.0.0.0:18202",
		},
		MySQL: &MySQL{
			Host: "127.0.0.1",
			Port: 3306,
			User: "root",
			Pwd:  "123456",
			Db:   "dtalk",
		},
		Debug: &Debug{
			Flag: false,
		},
		Release: &Release{
			Key:                 "123321",
			Issuer:              "Bob",
			TokenExpireDuration: 86400000000000,
			UserName:            "root",
			Password:            "root",
		},
		IdGenRPCClient: &RPCClient{
			RegAddrs: "127.0.0.1:2379",
			Schema:   "dtalk",
			SrvName:  "generator",
			Dial:     xtime.Duration(time.Second),
			Timeout:  xtime.Duration(time.Second),
		},
		CdkMaxNumber: 10,
		Chain33Client: Chain33Client{
			BlockChainAddr: "",
			Title:          "",
		},
		CdkMod: false,
	}
}

type Config struct {
	Env            string
	Platform       string
	Server         *HttpServer
	MySQL          *MySQL
	Debug          *Debug
	Release        *Release
	IdGenRPCClient *RPCClient
	CdkMaxNumber   int64
	Chain33Client  Chain33Client
	CdkMod         bool
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

type Debug struct {
	Flag bool
}

type Release struct {
	Key                 string
	Issuer              string
	TokenExpireDuration time.Duration
	UserName            string
	Password            string
}

// RPCClient is RPC client config.
type RPCClient struct {
	RegAddrs string // etcd addrs, seperate by ','
	Schema   string
	SrvName  string // call
	Dial     xtime.Duration
	Timeout  xtime.Duration
}

type Chain33Client struct {
	BlockChainAddr string
	Title          string
}
