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
	SinglePush(deviceToken string, notification Notification, extra *Extra) error
	SingleCustomPush(address string, notification Notification, extra *Extra) error
}

type Notification struct {
	Title    string
	Subtitle string
	Body     string
}

type Extra struct {
	SessionKey  string `json:"sessionKey"`
	ChannelType int32  `json:"channelType"`
	TimeOutTime int64  `json:"-"`
}

func (e *Extra) ToBytes() ([]byte, error) {
	return json.Marshal(e)
}
