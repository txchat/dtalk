package ios

import (
	"errors"
	"strconv"

	"github.com/txchat/dtalk/pkg/util"

	push "github.com/oofpgDLD/u-push"
	ios_push "github.com/oofpgDLD/u-push/ios"
	pusher "github.com/txchat/dtalk/pkg/third-part-push"
)

type iOSPusher struct {
	AppKey          string
	AppMasterSecret string
	MiActivity      string
	environment     string
}

func (t *iOSPusher) SinglePush(deviceToken string, notification pusher.Notification, extra *pusher.Extra) error {
	var client push.PushClient
	unicast := ios_push.NewIOSUnicast(t.AppKey, t.AppMasterSecret)

	//fmt.Println(t.AppKey, t.AppMasterSecret, t.DeviceToken, title, text)
	unicast.SetDeviceToken(deviceToken)

	//unicast.SetAlert("IOS 单播测试")
	unicast.SetAlertJson(push.IOSAlert{
		Title:    notification.Title,
		Subtitle: notification.Subtitle,
		Body:     notification.Body,
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
	unicast.SetCustomizedField("sessionKey", extra.SessionKey)
	unicast.SetCustomizedField("channelType", strconv.FormatInt(int64(extra.ChannelType), 10))
	return client.Send(unicast)
}

func (t *iOSPusher) SingleCustomPush(address string, notification pusher.Notification, extra *pusher.Extra) error {
	var client push.PushClient
	unicast := ios_push.NewIOSCustomizedcast(t.AppKey, t.AppMasterSecret)

	//fmt.Println(t.AppKey, t.AppMasterSecret, t.DeviceToken, title, text)
	unicast.SetAlias(address, "ADDRESS")

	//unicast.SetAlert("IOS 单播测试")
	unicast.SetAlertJson(push.IOSAlert{
		Title:    notification.Title,
		Subtitle: notification.Subtitle,
		Body:     notification.Body,
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
	unicast.SetCustomizedField("sessionKey", extra.SessionKey)
	unicast.SetCustomizedField("channelType", strconv.FormatInt(int64(extra.ChannelType), 10))
	return client.Send(unicast)
}
