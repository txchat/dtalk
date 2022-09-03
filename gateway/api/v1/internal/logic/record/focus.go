package record

import (
	"time"

	xproto "github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/txchat/dtalk/gateway/api/v1/internal/model"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/imparse/proto"
)

func (l *Logic) FocusPersonal(Operator string, logId int64) error {
	//查找消息
	rd, err := l.svcCtx.StoreClient.GetRecord(l.ctx, proto.Channel_ToUser, logId)
	if err != nil {
		return err
	}
	target := rd.ReceiverId
	sender := rd.SenderId
	if Operator == sender || Operator != target {
		return model.ErrPermission
	}
	now := uint64(util.TimeNowUnixNano() / int64(time.Millisecond))
	num, err := l.svcCtx.StoreClient.AddRecordFocus(l.ctx, Operator, logId, now)
	if err != nil {
		return err
	}
	action := &proto.SignalFocusMessage{
		Mid:        logId,
		Uid:        Operator,
		CurrentNum: num,
		Time:       now,
	}
	body, err := xproto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}
	err = l.svcCtx.AnswerClient.UniCastSignal(l.ctx, proto.SignalType_FocusMessage, sender, body)
	if err != nil {
		return err
	}
	return l.svcCtx.AnswerClient.UniCastSignal(l.ctx, proto.SignalType_FocusMessage, target, body)
}

func (l *Logic) FocusGroup(Operator string, logId int64) error {
	//查找消息
	rd, err := l.svcCtx.StoreClient.GetRecord(l.ctx, proto.Channel_ToGroup, logId)
	if err != nil {
		return err
	}
	target := rd.ReceiverId
	sender := rd.SenderId
	if Operator == sender {
		return model.ErrPermission
	}
	//群成员判断
	memOpt, err := l.svcCtx.GroupClient.GetMember(l.ctx, util.MustToInt64(target), Operator)
	if err != nil || memOpt == nil {
		return err
	}
	switch memOpt.GroupMemberType {
	case biz.GroupMemberTypeOwner:
	case biz.GroupMemberTypeAdmin:
	case biz.GroupMemberTypeNormal:
	default:
		return model.ErrPermission
	}

	now := uint64(util.TimeNowUnixNano() / int64(time.Millisecond))
	num, err := l.svcCtx.StoreClient.AddRecordFocus(l.ctx, Operator, logId, now)
	if err != nil {
		return err
	}
	action := &proto.SignalFocusMessage{
		Mid:        logId,
		Uid:        Operator,
		CurrentNum: num,
		Time:       now,
	}
	body, err := xproto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}
	return l.svcCtx.AnswerClient.GroupCastSignal(l.ctx, proto.SignalType_FocusMessage, target, body)
}
