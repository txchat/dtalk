package android

import (
	"errors"
	"github.com/txchat/dtalk/pkg/util"
	"strconv"

	push "github.com/oofpgDLD/u-push"
	android_push "github.com/oofpgDLD/u-push/android"
	"github.com/txchat/dtalk/service/offline-push/pusher"
)

type androidPusher struct {
	AppKey          string
	AppMasterSecret string
	MiActivity      string
	environment     string
}

func (t *androidPusher) SinglePush(deviceToken, title, text string, extra *pusher.Extra) error {
	var client push.PushClient
	unicast := android_push.NewAndroidUnicast(t.AppKey, t.AppMasterSecret)

	//fmt.Println(t.AppKey, t.AppMasterSecret, t.DeviceToken, title, text)
	unicast.SetDeviceToken(deviceToken)
	unicast.SetTitle(title)
	unicast.SetText(text)
	unicast.GoCustomAfterOpen("")
	unicast.SetDisplayType(push.NOTIFICATION)
	unicast.SetMipush(true, t.MiActivity)
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
	unicast.SetExtraField("address", extra.Address)
	unicast.SetExtraField("channelType", strconv.FormatInt(int64(extra.ChannelType), 10))
	return client.Send(unicast)
}

func (t *androidPusher) SingleCustomPush(address, title, text string, extra *pusher.Extra) error {
	var client push.PushClient
	unicast := android_push.NewAndroidCustomizedcast(t.AppKey, t.AppMasterSecret)

	//fmt.Println(t.AppKey, t.AppMasterSecret, t.DeviceToken, title, text)
	unicast.SetAlias(address, "ADDRESS")
	unicast.SetTitle(title)
	unicast.SetText(text)
	unicast.GoCustomAfterOpen("")
	unicast.SetDisplayType(push.NOTIFICATION)
	unicast.SetMipush(true, t.MiActivity)
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
	unicast.SetExtraField("address", extra.Address)
	unicast.SetExtraField("channelType", strconv.FormatInt(int64(extra.ChannelType), 10))
	return client.Send(unicast)
}
