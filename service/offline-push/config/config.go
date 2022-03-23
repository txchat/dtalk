package config

import (
	"flag"
	"github.com/BurntSushi/toml"
)

var (
	confPath string

	// Conf config
	Conf *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "offline-push.toml", "default config path.")
}

// Init init config.
func Init() (err error) {
	Conf = Default()
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

// Default new a config with specified defualt value.
func Default() *Config {
	return &Config{
		MQSub: &MQSubClient{
			Brokers: nil,
			Number:  0,
		},
	}
}

type Config struct {
	//push config
	AppId   string
	Pushers map[string]*Pusher
	MQSub   *MQSubClient
}

type Pusher struct {
	Env             string
	AppKey          string
	AppMasterSecret string
	MiActivity      string
}

type MQSubClient struct {
	Brokers []string
	Number  uint32
}
