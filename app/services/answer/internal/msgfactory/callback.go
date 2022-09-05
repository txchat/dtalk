package msgfactory

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/app/services/answer/internal/dao"
	record "github.com/txchat/dtalk/service/record/proto"
	logic "github.com/txchat/im/api/logic/grpc"
	"github.com/txchat/imparse"
)

// send msg callback
type withCometLevelAckCallback struct {
	appId       string
	dao         dao.AnswerRepository
	logicClient logic.LogicClient
}

func NewWithCometLevelAckCallback() *withCometLevelAckCallback {
	return &withCometLevelAckCallback{}
}

func (e *withCometLevelAckCallback) Transport(ctx context.Context, mid int64, key, from, target string, ch imparse.Channel, frameType imparse.FrameType, data []byte) error {
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
	return e.dao.PublishToSend(ctx, from, b)
}

func (e *withCometLevelAckCallback) RevAck(ctx context.Context, id int64, keys []string, data []byte) error {
	keysMsg := &logic.KeysMsg{
		AppId:  e.appId,
		ToKeys: keys,
		Msg:    data,
	}

	_, err := e.logicClient.PushByKeys(ctx, keysMsg)
	return err
}

// inner send msg callback
type withoutAckCallback struct {
	appId       string
	dao         *dao.Dao
	logicClient logic.LogicClient
}

func NewWithoutAckCallback() *withoutAckCallback {
	return &withoutAckCallback{}
}

func (e *withoutAckCallback) Transport(ctx context.Context, mid int64, key, from, target string, ch imparse.Channel, frameType imparse.FrameType, data []byte) error {
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
	return e.dao.PublishToSend(ctx, from, b)
}

func (e *withoutAckCallback) RevAck(ctx context.Context, id int64, keys []string, data []byte) error {
	return nil
}
