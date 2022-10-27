package ios

import (
	"errors"
	"strconv"

	"github.com/txchat/dtalk/pkg/util"

	push "github.com/oofpgDLD/u-push"
	ios_push "github.com/oofpgDLD/u-push/ios"
	"github.com/txchat/dtalk/service/offline-push/pusher"
)

type iOSPusher struct {
	AppKey          string
	AppMasterSecret string
	MiActivity      string
	environment     string
}

func (t *iOSPusher) SinglePush(deviceToken, title, text string, extra *pusher.Extra) error {
	var client push.PushClient
	unicast := ios_push.NewIOSUnicast(t.AppKey, t.AppMasterSecret)

	//fmt.Println(t.AppKey, t.AppMasterSecret, t.DeviceToken, title, text)
	unicast.SetDeviceToken(deviceToken)

	//unicast.SetAlert("IOS 单播测试")
	unicast.SetAlertJson(push.IOSAlert{
		Title: title,
		Body:  text,
	})
	//unicast.SetBadge(0)
	unicast.SetSound("default")
	unicast.SetExpireTime(util.UnixToTime(extra.TimeOutTime).In(util.Shanghai()).Format("2006-01-02 15:04:05"))

	switch t.environment {
	case "debug":
		// 测试模式
		unicast.SetTestMode()
	case "release":
		// 线上模式
		unicast.SetReleaseMode()
	default:
		return errors.New("unknown environment")
	}
	// Set customized fields
	unicast.SetCustomizedField("address", extra.Address)
	unicast.SetCustomizedField("channelType", strconv.FormatInt(int64(extra.ChannelType), 10))
	return client.Send(unicast)
}

func (t *iOSPusher) SingleCustomPush(address, title, text string, extra *pusher.Extra) error {
	var client push.PushClient
	unicast := ios_push.NewIOSCustomizedcast(t.AppKey, t.AppMasterSecret)

	//fmt.Println(t.AppKey, t.AppMasterSecret, t.DeviceToken, title, text)
	unicast.SetAlias(address, "ADDRESS")

	//unicast.SetAlert("IOS 单播测试")
	unicast.SetAlertJson(push.IOSAlert{
		Title: title,
		Body:  text,
	})
	//unicast.SetBadge(0)
	unicast.SetSound("default")
	unicast.SetExpireTime(util.UnixToTime(extra.TimeOutTime).In(util.Shanghai()).Format("2006-01-02 15:04:05"))

	switch t.environment {
	case "debug":
		// 测试模式
		unicast.SetTestMode()
	case "release":
		// 线上模式
		unicast.SetReleaseMode()
	default:
		return errors.New("unknown environment")
	}
	// Set customized fields
	unicast.SetCustomizedField("address", extra.Address)
	unicast.SetCustomizedField("channelType", strconv.FormatInt(int64(extra.ChannelType), 10))
	return client.Send(unicast)
}
