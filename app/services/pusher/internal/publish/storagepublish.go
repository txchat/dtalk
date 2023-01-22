package publish

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	xkafka "github.com/txchat/dtalk/pkg/mq/kafka"
	record "github.com/txchat/dtalk/proto/record"
	"github.com/txchat/imparse"
)

type Storage struct {
	AppID string
	topic string
	conn  *xkafka.Producer
}

func NewStoragePublish(appID string, cfg xkafka.ProducerConfig) *Storage {
	conn := xkafka.NewProducer(cfg)
	return &Storage{
		AppID: appID,
		topic: fmt.Sprintf("store-%s-topic", appID),
		conn:  conn,
	}
}

func (p *Storage) BatchPush(ctx context.Context, key, fromId string) (err error) {
	pushMsg := &record.RecordDeal{
		AppId:  p.AppID,
		FromId: fromId,
		Key:    key,
		Opt:    record.Operation_BatchPush,
		Msg:    nil,
	}
	b, err := proto.Marshal(pushMsg)
	if err != nil {
		return
	}
	_, _, err = p.conn.Publish(p.topic, fromId, b)
	return
}

func (p *Storage) MarkRead(ctx context.Context, key, fromId string, tp imparse.FrameType, mids []int64) (err error) {
	msg, err := proto.Marshal(&record.Marked{
		Type: string(tp),
		Mids: mids,
	})
	if err != nil {
		return
	}
	pushMsg := &record.RecordDeal{
		AppId:  p.AppID,
		FromId: fromId,
		Key:    key,
		Opt:    record.Operation_MarkRead,
		Msg:    msg,
	}
	b, err := proto.Marshal(pushMsg)
	if err != nil {
		return
	}
	_, _, err = p.conn.Publish(p.topic, fromId, b)
	return
}

func (p *Storage) Sync(ctx context.Context, key, fromId string, mid int64) (err error) {
	msg, err := proto.Marshal(&record.Sync{
		Mid: mid,
	})
	if err != nil {
		return
	}
	pushMsg := &record.RecordDeal{
		AppId:  p.AppID,
		FromId: fromId,
		Key:    key,
		Opt:    record.Operation_SyncMsg,
		Msg:    msg,
	}
	b, err := proto.Marshal(pushMsg)
	if err != nil {
		return
	}
	_, _, err = p.conn.Publish(p.topic, fromId, b)
	return
}
