package svc

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	xkafka "github.com/oofpgDLD/kafka-go"
	"github.com/txchat/dtalk/api/proto/chat"
	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/pusher/internal/config"
	"github.com/txchat/dtalk/app/services/pusher/internal/dao"
	"github.com/txchat/dtalk/internal/proto/offline"
	"github.com/txchat/dtalk/internal/recordutil"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/im/app/logic/logicclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Repo      dao.MessageRepository
	DeviceRPC deviceclient.Device
	LogicRPC  logicclient.Logic
	Producer  *xkafka.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	repo := dao.NewMessageRepositoryRedis(c.RedisDB)
	return &ServiceContext{
		Config: c,
		Repo:   repo,
		DeviceRPC: deviceclient.NewDevice(zrpc.MustNewClient(c.DeviceRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		LogicRPC: logicclient.NewLogic(zrpc.MustNewClient(c.LogicRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		Producer: xkafka.NewProducer(c.Producer),
	}
}

func (s *ServiceContext) PublishThirdPartyPushMQ(ctx context.Context, from string, targets []string, body []byte) error {
	var chatProto *chat.Chat
	err := proto.Unmarshal(body, chatProto)
	if err != nil {
		return err
	}
	if chatProto.GetType() != chat.Chat_message {
		return nil
	}
	var msg *message.Message
	err = proto.Unmarshal(chatProto.GetBody(), msg)
	if err != nil {
		return err
	}
	v, err := proto.Marshal(&offline.ThirdPartyPushMQ{
		AppId:       s.Config.AppID,
		ChannelType: msg.GetChannelType(),
		Session:     msg.GetTarget(),
		From:        from,
		Target:      targets,
		Content:     recordutil.MessageAlterContent(msg),
		Datetime:    util.TimeNowUnixMilli(),
	})
	if err != nil {
		return err
	}
	_, _, err = s.Producer.Publish(fmt.Sprintf("biz-%s-offlinepush", s.Config.AppID), from, v)
	return err
}
