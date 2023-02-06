package bizproto

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/proto/common"
	"github.com/txchat/imparse/proto/signal"
	"github.com/txchat/imparse/util"
)

//private
type SignalFrame struct {
	*StandardFrame
	base *signal.Signal

	mid        int64
	createTime uint64
}

func NewNoticeFrame(standardFrame *StandardFrame, bizPro *signal.Signal) *SignalFrame {
	frame := &SignalFrame{
		StandardFrame: standardFrame,
		base:          bizPro,
	}
	frame.SetBody(frame)
	return frame
}

func (p *SignalFrame) Type() imparse.FrameType {
	return SignalFrameType
}

func (p *SignalFrame) Filter(ctx context.Context, db imparse.Cache, filters ...imparse.Filter) (uint64, error) {
	var err error
	for _, filter := range filters {
		err = filter(ctx, p)
		if err != nil {
			return 0, err
		}
	}
	p.mid, err = db.GetMid(ctx)
	if err != nil {
		return 0, err
	}
	p.createTime = uint64(util.TimeNowUnixNano() / int64(time.Millisecond))
	p.base.Sid = p.mid
	return p.createTime, nil
}

func (p *SignalFrame) Transport(ctx context.Context, exec imparse.Exec) error {
	data, err := p.PushData()
	if err != nil {
		return err
	}
	return exec.Transport(ctx, p.mid, p.GetKey(), p.GetFrom(), p.GetTarget(), p.GetTransmissionMethod(), p.Type(), data)
}

func (p *SignalFrame) Ack(ctx context.Context, exec imparse.Exec) (int64, error) {
	return p.mid, nil
}

func (p *SignalFrame) AckBody() ([]byte, error) {
	body, err := proto.Marshal(p.base)
	if err != nil {
		return nil, fmt.Errorf("marshal NotifyMsg err: %v", err)
	}
	if err != nil {
		return nil, fmt.Errorf("marshal NotifyMsgAck err: %v", err)
	}
	data, err := proto.Marshal(&common.Proto{
		EventType: common.Proto_Signal,
		Body:      body,
	})
	if err != nil {
		return nil, fmt.Errorf("marshal Proto err: %v", err)
	}
	return data, err
}

func (p *SignalFrame) PushBody() ([]byte, error) {
	var err error
	var data []byte
	pro := common.Proto{
		EventType: common.Proto_Signal,
	}
	pro.Body, err = proto.Marshal(p.base)
	if err != nil {
		return nil, fmt.Errorf("marshal NotifyMsg err: %v", err)
	}
	data, err = proto.Marshal(&pro)
	if err != nil {
		return nil, fmt.Errorf("marshal Proto err: %v", err)
	}
	return data, err
}

func (p *SignalFrame) GetBase() *signal.Signal {
	return p.base
}
