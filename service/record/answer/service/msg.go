package service

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/record/answer/dao"
	"github.com/txchat/dtalk/service/record/answer/model"
	record "github.com/txchat/dtalk/service/record/proto"
	logic "github.com/txchat/im/api/logic/grpc"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"
	xproto "github.com/txchat/imparse/proto"
)

type DB struct {
	dao *dao.Dao
}

func (d *DB) GetMsg(ctx context.Context, from, msgId string) (*imparse.MsgIndex, error) {
	r, err := d.dao.GetRecordSeqIndex(from, msgId)
	if err != nil {
		return nil, err
	}
	if r == nil {
		return nil, nil
	}
	return &imparse.MsgIndex{
		Mid:        r.Mid,
		Seq:        r.Seq,
		SenderId:   r.SenderId,
		CreateTime: r.CreateTime,
	}, nil
}

func (d *DB) AddMsg(ctx context.Context, uid string, m *imparse.MsgIndex) error {
	return d.dao.AddRecordSeqIndex(uid, &model.MsgIndex{
		Mid:        m.Mid,
		Seq:        m.Seq,
		SenderId:   m.SenderId,
		CreateTime: m.CreateTime,
	})
}

func (d *DB) GetMid(ctx context.Context) (id int64, err error) {
	return d.dao.GetMid(ctx)
}

func (d *DB) GetFilters() map[imparse.FrameType][]imparse.Filter {
	//filters
	return map[imparse.FrameType][]imparse.Filter{
		chat.GroupFrameType: {
			func(ctx context.Context, frame imparse.Frame) error {
				f := frame.(*chat.GroupFrame)
				//判断群聊拦截
				if f.GetMsgType() != xproto.MsgType_Notice {
					if ok, err := d.dao.CheckInGroup(ctx, f.GetFrom(), util.MustToInt64(f.GetTarget())); !ok {
						if err != nil {
							return err
						}
						return model.ErrGroupMemberNotExists
					}
				}
				return nil
			},
		},
	}
}

//send msg callback
type Exec struct {
	appId       string
	dao         *dao.Dao
	logicClient logic.LogicClient
}

func (e *Exec) Transport(ctx context.Context, mid int64, key, from, target string, ch imparse.Channel, frameType imparse.FrameType, data []byte) error {
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

func (e *Exec) RevAck(ctx context.Context, id int64, keys []string, data []byte) error {
	keysMsg := &logic.KeysMsg{
		AppId:  e.appId,
		ToKeys: keys,
		Msg:    data,
	}

	_, err := e.logicClient.PushByKeys(ctx, keysMsg)
	return err
}

//send msg callback
type withoutAckExec struct {
	appId       string
	dao         *dao.Dao
	logicClient logic.LogicClient
}

func (e *withoutAckExec) Transport(ctx context.Context, mid int64, key, from, target string, ch imparse.Channel, frameType imparse.FrameType, data []byte) error {
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

func (e *withoutAckExec) RevAck(ctx context.Context, id int64, keys []string, data []byte) error {
	return nil
}
