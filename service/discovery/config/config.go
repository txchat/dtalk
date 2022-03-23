package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	xtime "github.com/txchat/dtalk/pkg/time"
	"github.com/txchat/dtalk/service/discovery/model"
	"time"
)

var (
	confPath string

	Conf *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "discovery.toml", "default config path.")
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
			Addr: "0.0.0.0:18001",
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
		CNodes: []*model.CNode{
			{
				Name:    "聊天节点1",
				Address: "",
			},
		},
		DNodes: []*model.DNode{
			{
				Name:    "合约节点1",
				Address: "",
			},
		},
	}
}

// Config config.
type Config struct {
	Server *HttpServer
	Redis  *Redis
	CNodes []*model.CNode
	DNodes []*model.DNode
}

type HttpServer struct {
	Addr string
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
