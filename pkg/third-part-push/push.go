package pusher

import (
	"encoding/json"
	"errors"
)

var execPusher = make(map[string]CreateFunc)

type CreateFunc func(cfg Config) IPusher

func Register(name string, exec CreateFunc) {
	execPusher[name] = exec
}

func Load(name string) (CreateFunc, error) {
	exec, ok := execPusher[name]
	if !ok {
		return nil, errors.New("pusher not find")
	}
	return exec, nil
}

type Config struct {
	AppKey          string
	AppMasterSecret string
	MiActivity      string
	Environment     string
}

type IPusher interface {
	SinglePush(deviceToken, title, text string, extra *Extra) error
	SingleCustomPush(address, title, text string, extra *Extra) error
}

type Extra struct {
	Address     string `json:"address"`
	ChannelType int32  `json:"channelType"`
	TimeOutTime int64  `json:"-"`
}

func (e *Extra) ToBytes() ([]byte, error) {
	return json.Marshal(e)
}
