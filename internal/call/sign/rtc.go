package sign

import (
	"github.com/txchat/dtalk/internal/call"
	"github.com/txchat/dtalk/pkg/util"
)

type CloudSDK struct {
	rtc call.TLSSig
}

func NewCloudSDK(rtc call.TLSSig) *CloudSDK {
	return &CloudSDK{
		rtc: rtc,
	}
}

func (t *CloudSDK) GetTicket(user string, roomId int64) (*call.Ticket, error) {
	sdkAppId := t.rtc.GetAppId()

	// 生成接收方的 userSig 和 privateMapKey
	userSig, err := t.rtc.GetUserSig(user)
	if err != nil {
		return nil, err
	}
	privateMapKey, err := t.rtc.GenPrivateMapKey(user, util.MustToInt32(roomId), 255)
	if err != nil {
		return nil, err
	}
	ticket := &call.Ticket{
		RoomId:        roomId,
		UserSig:       userSig,
		PrivateMapKey: privateMapKey,
		SDKAppID:      sdkAppId,
	}
	return ticket, nil
}
