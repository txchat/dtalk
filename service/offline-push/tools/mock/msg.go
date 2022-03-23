package mock

import (
	"github.com/golang/protobuf/proto"
	offlinepush "github.com/txchat/dtalk/service/offline-push/api"
	xproto "github.com/txchat/imparse/proto"
	"time"
)

type Msg struct {
	AppId       string
	DeviceType  offlinepush.Device
	Nickname    string
	TargetId    string
	DeviceToken string
}

func (m *Msg) Data() ([]byte, error) {
	//需要推送
	pushMsg := &offApi.OffPushMsg{
		AppId:       m.AppId,
		Device:      m.DeviceType,
		Title:       m.Nickname,
		Content:     "[你收到一条消息]",
		Token:       m.DeviceToken,
		ChannelType: int32(xproto.Channel_ToUser),
		Target:      m.TargetId,
		Timeout:     time.Now().Add(time.Minute * 7).Unix(),
	}
	return proto.Marshal(pushMsg)
}
