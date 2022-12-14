package msgfactory

import (
	"context"

	xkafka "github.com/txchat/dtalk/pkg/mq/kafka"

	"github.com/golang/protobuf/proto"
	record "github.com/txchat/dtalk/proto/record"
	logic "github.com/txchat/im/api/logic/grpc"
	"github.com/txchat/imparse"
)

// WithCometLevelAckCallback send msg callback
type WithCometLevelAckCallback struct {
	appId       string
	mqPub       *xkafka.Producer
	logicClient logic.LogicClient
}

func NewWithCometLevelAckCallback(appId string, mqPub *xkafka.Producer, logicClient logic.LogicClient) *WithCometLevelAckCallback {
	return &WithCometLevelAckCallback{
		appId:       appId,
		mqPub:       mqPub,
		logicClient: logicClient,
	}
}

func (e *WithCometLevelAckCallback) Transport(ctx context.Context, mid int64, key, from, target string, ch imparse.Channel, frameType imparse.FrameType, data []byte) error {
	pushMsg := &record.PushMsg{
		AppId:     e.appId,
		FromId:    from,
		Mid:       mid,
		Key:       key,
		Target:    target,
		Msg:       data,
		Type:      int32(ch),
		FrameType: string(frameType),
	}
	b, err := proto.Marshal(pushMsg)
	if err != nil {
		return err
	}

	_, _, err = e.mqPub.Publish(from, b)
	return err
}

func (e *WithCometLevelAckCallback) RevAck(ctx context.Context, id int64, keys []string, data []byte) error {
	keysMsg := &logic.KeysMsg{
		AppId:  e.appId,
		ToKeys: keys,
		Msg:    data,
	}

	_, err := e.logicClient.PushByKeys(ctx, keysMsg)
	return err
}

// WithoutAckCallback inner send msg callback
type WithoutAckCallback struct {
	appId string
	mqPub *xkafka.Producer
}

func NewWithoutAckCallback(appId string, mqPub *xkafka.Producer) *WithoutAckCallback {
	return &WithoutAckCallback{
		appId: appId,
		mqPub: mqPub,
	}
}

func (e *WithoutAckCallback) Transport(ctx context.Context, mid int64, key, from, target string, ch imparse.Channel, frameType imparse.FrameType, data []byte) error {
	pushMsg := &record.PushMsg{
		AppId:     e.appId,
		FromId:    from,
		Mid:       mid,
		Key:       key,
		Target:    target,
		Msg:       data,
		Type:      int32(ch),
		FrameType: string(frameType),
	}
	b, err := proto.Marshal(pushMsg)
	if err != nil {
		return err
	}

	_, _, err = e.mqPub.Publish(from, b)
	return err
}

func (e *WithoutAckCallback) RevAck(ctx context.Context, id int64, keys []string, data []byte) error {
	return nil
}