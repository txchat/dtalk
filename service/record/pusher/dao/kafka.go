package dao

import (
	"context"
	"fmt"

	"github.com/txchat/imparse"

	"github.com/golang/protobuf/proto"
	record "github.com/txchat/dtalk/service/record/proto"
	"gopkg.in/Shopify/sarama.v1"
)

// PushMsg push a message to databus.
func (d *Dao) BatchPushPublish(ctx context.Context, key, fromId string) (err error) {
	pushMsg := &record.RecordDeal{
		AppId:  d.appId,
		FromId: fromId,
		Key:    key,
		Opt:    record.Operation_BatchPush,
		Msg:    nil,
	}
	b, err := proto.Marshal(pushMsg)
	if err != nil {
		return
	}
	appTopic := fmt.Sprintf("store-%s-topic", d.appId)
	m := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(fromId),
		Topic: appTopic,
		Value: sarama.ByteEncoder(b),
	}
	if _, _, err = d.storePub.SendMessage(m); err != nil {
		d.log.Error().Err(err).Msg("kafkaPub.SendMessage error")
	}
	return
}

// PushMsg push a message to databus.
func (d *Dao) MarkReadPublish(ctx context.Context, key, fromId string, tp imparse.FrameType, mids []int64) (err error) {
	msg, err := proto.Marshal(&record.Marked{
		Type: string(tp),
		Mids: mids,
	})
	if err != nil {
		return
	}
	pushMsg := &record.RecordDeal{
		AppId:  d.appId,
		FromId: fromId,
		Key:    key,
		Opt:    record.Operation_MarkRead,
		Msg:    msg,
	}
	b, err := proto.Marshal(pushMsg)
	if err != nil {
		return
	}
	appTopic := fmt.Sprintf("store-%s-topic", d.appId)
	m := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(fromId),
		Topic: appTopic,
		Value: sarama.ByteEncoder(b),
	}
	if _, _, err = d.storePub.SendMessage(m); err != nil {
		d.log.Error().Err(err).Msg("kafkaPub.SendMessage error")
	}
	return
}

// PushMsg push a message to databus.
func (d *Dao) SyncPublish(ctx context.Context, key, fromId string, mid int64) (err error) {
	msg, err := proto.Marshal(&record.Sync{
		Mid: mid,
	})
	if err != nil {
		return
	}
	pushMsg := &record.RecordDeal{
		AppId:  d.appId,
		FromId: fromId,
		Key:    key,
		Opt:    record.Operation_SyncMsg,
		Msg:    msg,
	}
	b, err := proto.Marshal(pushMsg)
	if err != nil {
		return
	}
	appTopic := fmt.Sprintf("store-%s-topic", d.appId)
	m := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(fromId),
		Topic: appTopic,
		Value: sarama.ByteEncoder(b),
	}
	if _, _, err = d.storePub.SendMessage(m); err != nil {
		d.log.Error().Err(err).Msg("kafkaPub.SendMessage error")
	}
	return
}

func (d *Dao) PublishOfflineMsg(ctx context.Context, fromId string, b []byte) (err error) {
	appTopic := fmt.Sprintf("offpush-%s-topic", d.appId)
	m := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(fromId),
		Topic: appTopic,
		Value: sarama.ByteEncoder(b),
	}
	if _, _, err = d.offPushPub.SendMessage(m); err != nil {
		d.log.Error().Err(err).Bytes("pushMsg", b).Msg("kafkaPub.SendMessage error")
	}
	return
}
