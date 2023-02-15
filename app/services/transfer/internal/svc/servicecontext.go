package svc

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/device/deviceclient"
	"github.com/txchat/dtalk/app/services/generator/generatorclient"
	"github.com/txchat/dtalk/app/services/group/groupclient"
	"github.com/txchat/dtalk/app/services/pusher/pusherclient"
	"github.com/txchat/dtalk/app/services/transfer/internal/config"
	"github.com/txchat/dtalk/app/services/transfer/internal/dao"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/proto/record"
	"github.com/txchat/im/api/protocol"
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
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		Repo:         nil,
		PusherClient: nil,
		GroupClient:  nil,
		Producer:     xkafka.NewProducer(c.Producer),
	}
}

func (s *ServiceContext) saveMessageToStorage(ctx context.Context, p *protocol.Proto) error {
	//TODO publish store
	return nil
}

func (s *ServiceContext) asyncPushMessage(ctx context.Context, from string, p *protocol.Proto) error {
	msg, err := proto.Marshal(p)
	if err != nil {
		return err
	}

	v, err := proto.Marshal(&record.PushMsgMQ{
		AppId:   s.Config.AppID,
		From:    from,
		Target:  p.GetTarget(),
		Channel: p.GetChannel(),
		Msg:     msg,
	})
	if err != nil {
		return err
	}
	_, _, err = s.Producer.Publish(fmt.Sprintf("biz-%s-push", s.Config.AppID), from, v)
	return err
}

func (s *ServiceContext) TransferMessage(ctx context.Context, from string, p *protocol.Proto) error {
	p.Op = int32(protocol.Op_ReceiveMsg)
	switch p.GetChannel() {
	case protocol.Channel_Group:
		members, err := s.GroupClient.MembersInfo(ctx, &groupclient.MembersInfoReq{
			Gid: util.MustToInt64(p.GetTarget()),
			Uid: nil,
		})
		if err != nil {
			return err
		}
		for _, member := range members.GetMembers() {
			p = deepCopy(p)
			// 1, seq增加
			var seq int64
			seq, err = s.Repo.IncrUserSeq(ctx, member.GetUid())
			if err != nil {
				continue
			}
			// 2. 持久化
			p.Seq = seq
			// 写同步库
			err = s.Repo.SaveUserChatRecord(ctx, p)
			if err != nil {
				continue
			}
			// 异步写存储库
			err = s.saveMessageToStorage(ctx, p)
			if err != nil {
				continue
			}
		}
		// 推送
		_, err = s.PusherClient.PushGroup(ctx, &pusherclient.PushGroupReq{
			App:  s.Config.AppID,
			Gid:  p.GetTarget(),
			Body: nil,
		})
		if err != nil {
			//异步处理推送
			err = s.asyncPushMessage(ctx, from, p)
			if err != nil {
				//TODO log
			}
		}
	case protocol.Channel_Private:
		// 1, seq增加
		seq, err := s.Repo.IncrUserSeq(ctx, p.GetTarget())
		if err != nil {
			return err
		}
		// 2. 持久化
		p.Seq = seq
		// 写同步库
		err = s.Repo.SaveUserChatRecord(ctx, p)
		if err != nil {
			return err
		}
		// 异步写存储库
		err = s.saveMessageToStorage(ctx, p)
		if err != nil {
			return err
		}
		// 3. 推送
		_, err = s.PusherClient.PushList(ctx, &pusherclient.PushListReq{
			App:  s.Config.AppID,
			From: from,
			Uid:  []string{p.GetTarget()},
			Body: nil,
		})
		if err != nil {
			//异步处理推送
			err = s.asyncPushMessage(ctx, from, p)
			if err != nil {
				//TODO log
			}
		}
	default:
		//TODO return error
	}
	return nil
}

func deepCopy(p *protocol.Proto) *protocol.Proto {
	newP := new(protocol.Proto)
	newP.Ver = p.Ver
	newP.Op = p.Op
	newP.Seq = p.Seq
	newP.Ack = p.Ack
	newP.Mid = p.Mid
	newP.Channel = p.Channel
	newP.Target = p.Target
	newP.Time = p.Time
	newP.Body = make([]byte, len(p.Body))
	copy(newP.Body, p.Body)
	return newP
}
