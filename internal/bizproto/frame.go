package bizproto

import (
	"github.com/golang/protobuf/proto"
	"github.com/txchat/im/api/protocol"
	"github.com/txchat/imparse"
)

const (
	PrivateFrameType imparse.FrameType = "private"
	GroupFrameType   imparse.FrameType = "group"
	SignalFrameType  imparse.FrameType = "signal"
)

type Option func(*Options)
type Options struct {
	mid                int64
	createTime         uint64
	target             string
	transmissionMethod imparse.Channel
}

func WithMid(t int64) Option {
	return func(o *Options) {
		o.mid = t
	}
}

func WithCreateTime(t uint64) Option {
	return func(o *Options) {
		o.createTime = t
	}
}

func WithTarget(t string) Option {
	return func(o *Options) {
		o.target = t
	}
}

func WithTransmissionMethod(t imparse.Channel) Option {
	return func(o *Options) {
		o.transmissionMethod = t
	}
}

//标准帧 定义了标准协议的Ack数据和Push数据格式
type StandardFrame struct {
	body imparse.BizProto
	base *protocol.Proto

	mid                int64
	createTime         uint64
	key                string
	from               string
	target             string
	transmissionMethod imparse.Channel
}

func NewStandardFrame(base *protocol.Proto, key, from string, opts ...Option) *StandardFrame {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	return &StandardFrame{base: base, key: key, from: from, mid: options.mid, createTime: options.createTime, target: options.target, transmissionMethod: options.transmissionMethod}
}

func (f *StandardFrame) Data() ([]byte, error) {
	p := protocol.Proto{
		Ver: f.base.GetVer(),
		Op:  f.base.GetOp(),
		Seq: f.base.GetSeq(),
		Ack: f.base.GetAck(),
	}
	var err error
	p.Body, err = f.body.PushBody()
	if err != nil {
		return nil, err
	}
	//send transfer msg
	return proto.Marshal(&p)
}

func (f *StandardFrame) AckData() ([]byte, error) {
	p := protocol.Proto{
		Ver: f.base.GetVer(),
		Op:  int32(protocol.Op_SendMsgReply),
		Seq: f.base.GetSeq(),
		Ack: f.base.GetAck(),
	}
	var err error
	p.Body, err = f.body.AckBody()
	if err != nil {
		return nil, err
	}
	//send msg ack
	return proto.Marshal(&p)
}

func (f *StandardFrame) PushData() ([]byte, error) {
	p := protocol.Proto{
		Ver: f.base.GetVer(),
		Op:  int32(protocol.Op_ReceiveMsg),
		Seq: f.base.GetSeq(),
		Ack: f.base.GetAck(),
	}
	var err error
	p.Body, err = f.body.PushBody()
	if err != nil {
		return nil, err
	}
	//send to client B
	return proto.Marshal(&p)
}

func (f *StandardFrame) GetMid() int64 {
	return f.mid
}

func (f *StandardFrame) SetMid(mid int64) {
	f.mid = mid
}

func (f *StandardFrame) GetCreateTime() uint64 {
	return f.createTime
}

func (f *StandardFrame) SetCreateTime(createTime uint64) {
	f.createTime = createTime
}

func (f *StandardFrame) GetTarget() string {
	return f.target
}

func (f *StandardFrame) SetTarget(target string) {
	f.target = target
}

func (f *StandardFrame) GetTransmissionMethod() imparse.Channel {
	return f.transmissionMethod
}

func (f *StandardFrame) SetTransmissionMethod(transmissionMethod imparse.Channel) {
	f.transmissionMethod = transmissionMethod
}

func (f *StandardFrame) GetFrom() string {
	return f.from
}

func (f *StandardFrame) SetFrom(from string) {
	f.from = from
}

func (f *StandardFrame) GetKey() string {
	return f.key
}

func (f *StandardFrame) SetKey(key string) {
	f.key = key
}

func (f *StandardFrame) GetBody() imparse.BizProto {
	return f.body
}

func (f *StandardFrame) SetBody(body imparse.BizProto) {
	f.body = body
}
