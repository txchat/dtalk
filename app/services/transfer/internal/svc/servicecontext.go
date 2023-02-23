package svc

import (
	"context"
	"fmt"

	checker "github.com/txchat/dtalk/internal/recordutil/dtalk"

	"github.com/txchat/dtalk/internal/proto/record"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/api/proto/chat"
	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/generator/generatorclient"
	"github.com/txchat/dtalk/app/services/group/groupclient"
	"github.com/txchat/dtalk/app/services/pusher/pusherclient"
	"github.com/txchat/dtalk/app/services/transfer/internal/config"
	"github.com/txchat/dtalk/app/services/transfer/internal/dao"
	"github.com/txchat/dtalk/pkg/util"
	xkafka "github.com/txchat/pkg/mq/kafka"
)

type ServiceContext struct {
	Config config.Config

	Repo         dao.Repository
	DeviceClient deviceclient.Device
	PusherClient pusherclient.Pusher
	GroupClient  groupclient.Group
	IDGenClient  generatorclient.Generator
	Producer     *xkafka.Producer
	Filters      map[message.Channel]checker.Filter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		Repo:         nil,
		DeviceClient: nil,
		PusherClient: nil,
		GroupClient:  nil,
		IDGenClient:  nil,
		Producer:     xkafka.NewProducer(c.Producer),
		Filters:      nil,
	}
}

func (s *ServiceContext) saveMessageToStorage(ctx context.Context, from string, target []string, chatProto *chat.Chat) error {
	v, err := proto.Marshal(&record.StoreMsgMQ{
		AppId:  s.Config.AppID,
		From:   from,
		Target: target,
		Chat:   chatProto,
	})
	if err != nil {
		return err
	}
	_, _, err = s.Producer.Publish(fmt.Sprintf("biz-%s-store", s.Config.AppID), from, v)
	return err
}

func (s *ServiceContext) asyncPushMessage(ctx context.Context, channel message.Channel, from, target string, body []byte) error {
	v, err := proto.Marshal(&record.PushMsgMQ{
		AppId:   s.Config.AppID,
		From:    from,
		Target:  target,
		Channel: channel,
		Msg:     body,
	})
	if err != nil {
		return err
	}
	_, _, err = s.Producer.Publish(fmt.Sprintf("biz-%s-push", s.Config.AppID), from, v)
	return err
}

func (s *ServiceContext) TransferMessage(ctx context.Context, channel message.Channel, from, target string, chatProto *chat.Chat) error {
	switch channel {
	case message.Channel_Group:
		members, err := s.GroupClient.MembersInfo(ctx, &groupclient.MembersInfoReq{
			Gid: util.MustToInt64(target),
			Uid: nil,
		})
		if err != nil {
			return err
		}
		var memId = make([]string, 0, len(members.GetMembers()))
		for _, member := range members.GetMembers() {
			memId = append(memId, member.GetUid())
			chatProto = deepCopy(chatProto)
			// 1, seq增加
			var seq int64
			seq, err = s.Repo.IncrUserSeq(ctx, member.GetUid())
			if err != nil {
				continue
			}
			chatProto.Seq = seq
			// 2. 持久化
			// 写同步库
			err = s.Repo.SaveUserChatRecord(ctx, chatProto)
			if err != nil {
				continue
			}
			// 推送
			_, err = s.PusherClient.PushList(ctx, &pusherclient.PushListReq{
				App:  s.Config.AppID,
				From: from,
				Uid:  []string{member.GetUid()},
				Body: nil,
			})
			if err != nil {
				//异步处理推送
				err = s.asyncPushMessage(ctx, message.Channel_Private, from, target, nil)
				if err != nil {
					//TODO log
				}
			}
		}
		// 异步写存储库
		err = s.saveMessageToStorage(ctx, from, memId, chatProto)
		if err != nil {
			return err
		}
	case message.Channel_Private:
		// 1, seq增加
		seq, err := s.Repo.IncrUserSeq(ctx, target)
		if err != nil {
			return err
		}
		chatProto.Seq = seq
		// 2. 持久化
		// 写同步库
		err = s.Repo.SaveUserChatRecord(ctx, chatProto)
		if err != nil {
			return err
		}
		// 异步写存储库
		err = s.saveMessageToStorage(ctx, from, []string{target}, chatProto)
		if err != nil {
			return err
		}
		// 3. 推送
		body, err := proto.Marshal(chatProto)
		if err != nil {
			return err
		}
		_, err = s.PusherClient.PushList(ctx, &pusherclient.PushListReq{
			App:  s.Config.AppID,
			From: from,
			Uid:  []string{target},
			Body: body,
		})
		if err != nil {
			//异步处理推送
			err = s.asyncPushMessage(ctx, channel, from, target, body)
			if err != nil {
				//TODO log
			}
		}
	default:
		//TODO return error
	}
	return nil
}

func deepCopy(p *chat.Chat) *chat.Chat {
	newP := new(chat.Chat)
	newP.Type = p.Type
	newP.Seq = p.Seq
	newP.Body = make([]byte, len(p.Body))
	copy(newP.Body, p.Body)
	return newP
}
