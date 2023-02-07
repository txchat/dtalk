package bizproto

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/proto/common"
	"github.com/txchat/imparse/util"
)

//private
type PrivateFrame struct {
	*StandardFrame
	base *common.Common

	stored bool
}

func NewPrivateFrame(standardFrame *StandardFrame, bizPro *common.Common) *PrivateFrame {
	frame := &PrivateFrame{
		StandardFrame: standardFrame,
		base:          bizPro,
	}
	frame.SetBody(frame)
	bizPro.From = frame.GetFrom()
	frame.SetTarget(bizPro.GetTarget())
	frame.SetTransmissionMethod(imparse.UniCast)
	return frame
}

func (p *PrivateFrame) Type() imparse.FrameType {
	return PrivateFrameType
}

func (p *PrivateFrame) Filter(ctx context.Context, db imparse.Cache, filters ...imparse.Filter) (uint64, error) {
	//查询是否有重复消息
	msg, err := db.GetMsg(ctx, p.GetFrom(), p.base.GetSeq())
	if err != nil {
		return 0, err
	}

	if msg != nil {
		p.stored = true
		p.base.Mid, err = strconv.ParseInt(msg.Mid, 10, 64)
		if err != nil {
			return 0, err
		}
		p.base.Datetime = msg.CreateTime
	} else {
		for _, filter := range filters {
			err = filter(ctx, p)
			if err != nil {
				return 0, err
			}
		}
		p.stored = false
		p.base.Mid, err = db.GetMid(ctx)
		if err != nil {
			return 0, err
		}
		p.base.Datetime = uint64(util.TimeNowUnixNano() / int64(time.Millisecond))
		err := db.AddMsg(ctx, p.GetFrom(), &imparse.MsgIndex{
			Mid:        strconv.FormatInt(p.base.GetMid(), 10),
			Seq:        p.base.GetSeq(),
			SenderId:   p.GetFrom(),
			CreateTime: p.base.GetDatetime(),
		})
		if err != nil {
			return 0, err
		}
	}
	p.mid = p.base.Mid
	p.createTime = p.base.Datetime
	return p.base.GetDatetime(), nil
}

func (p *PrivateFrame) Transport(ctx context.Context, exec imparse.Exec) error {
	if p.stored {
		return nil
	}
	data, err := p.PushData()
	if err != nil {
		return err
	}
	return exec.Transport(ctx, p.base.GetMid(), p.GetKey(), p.GetFrom(), p.GetTarget(), p.GetTransmissionMethod(), p.Type(), data)
}

func (p *PrivateFrame) Ack(ctx context.Context, exec imparse.Exec) (int64, error) {
	ackBytes, err := p.AckData()
	if err != nil {
		return 0, err
	}
	return p.base.GetMid(), exec.RevAck(ctx, p.base.GetMid(), []string{p.GetKey()}, ackBytes)
}

func (p *PrivateFrame) AckBody() ([]byte, error) {
	body, err := proto.Marshal(&common.CommonAck{
		Mid:      p.base.GetMid(),
		Datetime: p.base.GetDatetime(),
	})
	if err != nil {
		return nil, fmt.Errorf("marshal CommonAck err: %v", err)
	}
	data, err := proto.Marshal(&common.Proto{
		EventType: common.Proto_commonAck,
		Body:      body,
	})
	if err != nil {
		return nil, fmt.Errorf("marshal Proto err: %v", err)
	}
	return data, err
}

func (p *PrivateFrame) PushBody() ([]byte, error) {
	var err error
	var data []byte
	pro := common.Proto{
		EventType: common.Proto_common,
	}
	pro.Body, err = proto.Marshal(p.base)
	if err != nil {
		return nil, fmt.Errorf("marshal Common err: %v", err)
	}
	data, err = proto.Marshal(&pro)
	if err != nil {
		return nil, fmt.Errorf("marshal Proto err: %v", err)
	}
	return data, err
}

//
func (p *PrivateFrame) GetChannelType() common.Channel {
	return p.base.ChannelType
}

func (p *PrivateFrame) GetMsgType() common.MsgType {
	return p.base.MsgType
}

func (p *PrivateFrame) GetBase() *common.Common {
	return p.base
}
