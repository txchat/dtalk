package record

import (
	"context"
	"time"

	xproto "github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/txchat/dtalk/app/gateway/chat/internal/model"
	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"
	"github.com/txchat/dtalk/app/services/answer/answerclient"
	"github.com/txchat/dtalk/app/services/storage/storageclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/dtalk/pkg/util"
	groupApi "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/imparse/proto/common"
	"github.com/txchat/imparse/proto/signal"
	"github.com/zeromicro/go-zero/core/logx"
)

type FocusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewFocusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FocusLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &FocusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *FocusLogic) Focus(req *types.FocusReq) (resp *types.FocusResp, err error) {
	operator := l.custom.UID
	switch req.Type {
	case model.Private:
		err = l.focusPersonal(operator, req.Mid)
	case model.Group:
		err = l.focusGroup(operator, req.Mid)
	default:
		err = xerror.ErrInvalidParams
		return
	}
	return
}

func (l *FocusLogic) focusPersonal(operator string, mid int64) error {
	//查找消息
	rd, err := l.svcCtx.StorageRPC.GetRecord(l.ctx, &storageclient.GetRecordReq{
		Tp:  common.Channel_ToUser,
		Mid: mid,
	})
	if err != nil {
		return err
	}
	target := rd.ReceiverId
	sender := rd.SenderId
	if operator == sender || operator != target {
		return model.ErrPermission
	}
	now := uint64(util.TimeNowUnixNano() / int64(time.Millisecond))
	addFocusResp, err := l.svcCtx.StorageRPC.AddRecordFocus(l.ctx, &storageclient.AddRecordFocusReq{
		Uid:  operator,
		Mid:  mid,
		Time: now,
	})
	if err != nil {
		return err
	}
	action := &signal.SignalFocusMessage{
		Mid:        mid,
		Uid:        operator,
		CurrentNum: addFocusResp.GetCurrentNum(),
		Time:       now,
	}
	body, err := xproto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}
	_, err = l.svcCtx.AnswerRPC.UniCastSignal(l.ctx, &answerclient.UniCastSignalReq{
		Type:   signal.SignalType_FocusMessage,
		Target: sender,
		Body:   body,
	})
	if err != nil {
		return err
	}
	_, err = l.svcCtx.AnswerRPC.UniCastSignal(l.ctx, &answerclient.UniCastSignalReq{
		Type:   signal.SignalType_FocusMessage,
		Target: target,
		Body:   body,
	})
	return err
}

func (l *FocusLogic) focusGroup(operator string, mid int64) error {
	//查找消息
	rd, err := l.svcCtx.StorageRPC.GetRecord(l.ctx, &storageclient.GetRecordReq{
		Tp:  common.Channel_ToGroup,
		Mid: mid,
	})
	if err != nil {
		return err
	}
	target := rd.ReceiverId
	sender := rd.SenderId
	if operator == sender {
		return model.ErrPermission
	}
	//群成员判断
	memOpt, err := l.svcCtx.GroupRPC.GetMember(l.ctx, &groupApi.GetMemberReq{
		MemberId: operator,
		GroupId:  util.MustToInt64(target),
	})
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
	addFocusResp, err := l.svcCtx.StorageRPC.AddRecordFocus(l.ctx, &storageclient.AddRecordFocusReq{
		Uid:  operator,
		Mid:  mid,
		Time: now,
	})
	if err != nil {
		return err
	}
	action := &signal.SignalFocusMessage{
		Mid:        mid,
		Uid:        operator,
		CurrentNum: addFocusResp.GetCurrentNum(),
		Time:       now,
	}
	body, err := xproto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}
	_, err = l.svcCtx.AnswerRPC.GroupCastSignal(l.ctx, &answerclient.GroupCastSignalReq{
		Type:   signal.SignalType_FocusMessage,
		Target: target,
		Body:   body,
	})
	return err
}
