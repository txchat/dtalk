package config

import (
	"github.com/BurntSushi/toml"
)

var (
	ConfPath string
	Conf     *Config
)

// Init init config.
func Init() (err error) {
	Conf = Default()
	_, err = toml.DecodeFile(ConfPath, &Conf)
	return
}

func Default() *Config {
	return &Config{
		Client: Client{
			AppId:   "",
			FromId:  "",
			Brokers: nil,
		},
		Msg: Msg{
			AppId:       "",
			DeviceType:  0,
			Nickname:    "",
			TargetId:    "",
			DeviceToken: "",
		},
	}
}

type Config struct {
	Client Client
	Msg    Msg
}

type Client struct {
	AppId   string
	FromId  string
	Brokers []string
}

type Msg struct {
	AppId       string
	DeviceType  int32
	Nickname    string
	TargetId    string
	DeviceToken string
}
