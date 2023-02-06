package bizproto

import (
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/im/api/protocol"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/proto/common"
	"github.com/txchat/imparse/proto/signal"
)

type _handleEvent func(*StandardFrame, []byte) (imparse.Frame, error)

var events = map[common.Proto_EventType]_handleEvent{
	common.Proto_common: func(s *StandardFrame, body []byte) (imparse.Frame, error) {
		var pro common.Common
		err := proto.Unmarshal(body, &pro)
		if err != nil {
			return nil, err
		}
		switch pro.ChannelType {
		case common.Channel_ToUser:
			return NewPrivateFrame(s, &pro), nil
		case common.Channel_ToGroup:
			return NewGroupFrame(s, &pro), nil
		default:
			return nil, imparse.ErrExecSupport
		}
	},
	common.Proto_Signal: func(s *StandardFrame, body []byte) (imparse.Frame, error) {
		var pro signal.Signal
		err := proto.Unmarshal(body, &pro)
		if err != nil {
			return nil, err
		}
		return NewNoticeFrame(s, &pro), nil
	},
}

//标准解析器 定义了解析标准协议的方法
type StandardParse struct {
}

func (s *StandardParse) NewFrame(key, from string, in io.Reader, opts ...Option) (imparse.Frame, error) {
	data, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	//
	var p protocol.Proto
	err = proto.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}

	switch p.GetVer() {
	case 0, 1:
	default:
	}

	switch p.GetOp() {
	case int32(protocol.Op_SendMsg):
	case int32(protocol.Op_Auth):

	}
	//业务服务协议解析
	var pro common.Proto
	err = proto.Unmarshal(p.Body, &pro)
	if err != nil {
		return nil, err
	}

	//解析event事件
	event, ok := events[pro.EventType]
	if !ok || event == nil {
		return nil, imparse.ErrorEnvType
	}
	frame, err := event(NewStandardFrame(&p, key, from, opts...), pro.Body)
	if err != nil {
		return nil, err
	}
	return frame, err
}
