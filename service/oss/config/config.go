package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	xtime "github.com/txchat/dtalk/pkg/time"
	"time"
)

var (
	confPath string

	Conf *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "oss.toml", "default config path.")
}

// Init init config.
func Init() (err error) {
	Conf = Default()
	_, err = toml.DecodeFile(confPath, &Conf)
	err = check()
	return
}

func Default() *Config {
	return &Config{
		Env: "debug",
		Server: &HttpServer{
			Addr: "0.0.0.0:18005",
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
	}
}

func check() error {
	// appId+ossType 唯一
	ossMap := make(map[string]string)
	for _, oss := range Conf.Oss {
		key := oss.AppId + "-" + oss.OssType
		if _, ok := ossMap[key]; ok {
			return errors.New("有重复的 AppId 和 OssType")
		}
		ossMap[key] = key
	}
	return nil
}

// Config config.
type Config struct {
	Env    string
	Server *HttpServer
	Redis  *Redis
	Oss    []*Oss
}

type HttpServer struct {
	Addr string
}

// MySQL .
type MySQL struct {
	Host string
	Port int32
	User string
	Pwd  string
	Db   string
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

type Oss struct {
	AppId   string
	OssType string
	// RegionId 决定获得 tempKey 的速度快不快
	RegionId        string
	AccessKeyId     string
	AccessKeySecret string
	Role            string
	Policy          string
	DurationSeconds int
	Bucket          string
	// EndPoint 资源终端
	EndPoint  string
	PublicUrl string
}
