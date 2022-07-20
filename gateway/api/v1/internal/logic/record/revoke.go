package record

import (
	"context"
	"time"

	"github.com/txchat/dtalk/service/group/model/biz"

	xproto "github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/txchat/dtalk/gateway/api/v1/internal/model"
	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/imparse/proto"
)

type Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogic(ctx context.Context, svcCtx *svc.ServiceContext) Logic {
	return Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Logic) RevokePersonal(Operator string, mid int64) error {
	//查找消息
	rd, err := l.svcCtx.StoreClient.GetRecord(l.ctx, proto.Channel_ToUser, mid)
	if err != nil {
		return err
	}
	target := rd.ReceiverId
	if rd.SenderId != Operator || time.Since(util.UnixToTime(int64(rd.CreateTime))) > time.Duration(l.svcCtx.Config().Revoke.Expire) {
		return model.ErrPermission
	}
	action := &proto.SignalRevoke{
		Mid:      mid,
		Operator: Operator,
		Self:     rd.SenderId == Operator,
	}
	body, err := xproto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}
	err = l.svcCtx.AnswerClient.UniCastSignal(l.ctx, proto.SignalType_Revoke, target, body)
	if err != nil {
		return err
	}
	if err := l.svcCtx.StoreClient.DelRecord(l.ctx, proto.Channel_ToUser, mid); err != nil {
		return err
	}
	return l.svcCtx.AnswerClient.UniCastSignal(l.ctx, proto.SignalType_Revoke, Operator, body)
}

func (l *Logic) RevokeGroup(Operator string, mid int64) error {
	//查找消息
	rd, err := l.svcCtx.StoreClient.GetRecord(l.ctx, proto.Channel_ToGroup, mid)
	if err != nil {
		return err
	}
	target := rd.ReceiverId
	if rd.SenderId == Operator && time.Since(util.UnixToTime(int64(rd.CreateTime))) > time.Duration(l.svcCtx.Config().Revoke.Expire) {
		return model.ErrPermission
	}
	if rd.SenderId != Operator {
		//执行者
		memOpt, err := l.svcCtx.GroupClient.GetMember(l.ctx, util.ToInt64(target), Operator)
		if err != nil || memOpt == nil {
			return err
		}
		//消息所有者
		memOwn, err := l.svcCtx.GroupClient.GetMember(l.ctx, util.ToInt64(target), rd.SenderId)
		if err != nil || memOwn == nil {
			return err
		}
		switch memOpt.GroupMemberType {
		case biz.GroupMemberTypeOwner:
		case biz.GroupMemberTypeAdmin:
			if memOwn.GroupMemberType == biz.GroupMemberTypeOwner {
				return model.ErrPermission
			}
		default:
			return model.ErrPermission
		}
	}
	action := &proto.SignalRevoke{
		Mid:      mid,
		Operator: Operator,
		Self:     rd.SenderId == Operator,
	}
	body, err := xproto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}
	if err := l.svcCtx.StoreClient.DelRecord(l.ctx, proto.Channel_ToGroup, mid); err != nil {
		return err
	}
	return l.svcCtx.AnswerClient.GroupCastSignal(l.ctx, proto.SignalType_Revoke, target, body)
}
