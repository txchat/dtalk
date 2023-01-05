package sign

import (
	"encoding/json"

	"github.com/txchat/dtalk/internal/call"
	"github.com/txchat/dtalk/pkg/util"
)

type Ticket struct {
	RoomId        int64
	UserSig       string
	PrivateMapKey string
	SDKAppID      int32
}

func (t *Ticket) ToBytes() (call.Ticket, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func FromBytes(data []byte) (*Ticket, error) {
	var ticket *Ticket
	err := json.Unmarshal(data, ticket)
	return ticket, err
}

type CloudSDK struct {
	rtc TLSSig
}

func NewCloudSDK(rtc TLSSig) *CloudSDK {
	return &CloudSDK{
		rtc: rtc,
	}
}

func (t *CloudSDK) GetTicket(user string, roomId int64) (call.Ticket, error) {
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
	ticket := Ticket{
		RoomId:        roomId,
		UserSig:       userSig,
		PrivateMapKey: privateMapKey,
		SDKAppID:      sdkAppId,
	}
	data, err := ticket.ToBytes()
	if err != nil {
		return nil, err
	}
	return data, nil
}
