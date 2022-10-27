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

type RevokeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewRevokeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RevokeLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &RevokeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *RevokeLogic) Revoke(req *types.RevokeReq) (resp *types.RevokeResp, err error) {
	// todo: add your logic here and delete this line
	operator := l.custom.UID
	switch req.Type {
	case model.Private:
		err = l.revokePersonal(operator, req.Mid)
	case model.Group:
		err = l.revokeGroup(operator, req.Mid)
	default:
		err = xerror.ErrInvalidParams
		return
	}
	return
}

func (l *RevokeLogic) revokePersonal(operator string, mid int64) error {
	//查找消息
	record, err := l.svcCtx.StorageRPC.GetRecord(l.ctx, &storageclient.GetRecordReq{
		Tp:  common.Channel_ToUser,
		Mid: mid,
	})
	if err != nil {
		return err
	}
	target := record.ReceiverId
	if record.SenderId != operator || time.Since(util.UnixToTime(int64(record.CreateTime))) > time.Duration(l.svcCtx.Config.Revoke.Expire) {
		return model.ErrPermission
	}
	action := &signal.SignalRevoke{
		Mid:      mid,
		Operator: operator,
		Self:     record.SenderId == operator,
	}
	body, err := xproto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}
	_, err = l.svcCtx.AnswerRPC.UniCastSignal(l.ctx, &answerclient.UniCastSignalReq{
		Type:   signal.SignalType_Revoke,
		Target: target,
		Body:   body,
	})
	if err != nil {
		return err
	}
	if _, err := l.svcCtx.StorageRPC.DelRecord(l.ctx, &storageclient.DelRecordReq{
		Tp:  common.Channel_ToUser,
		Mid: mid,
	}); err != nil {
		return err
	}
	_, err = l.svcCtx.AnswerRPC.UniCastSignal(l.ctx, &answerclient.UniCastSignalReq{
		Type:   signal.SignalType_Revoke,
		Target: operator,
		Body:   body,
	})
	return err
}

func (l *RevokeLogic) revokeGroup(Operator string, mid int64) error {
	//查找消息
	record, err := l.svcCtx.StorageRPC.GetRecord(l.ctx, &storageclient.GetRecordReq{
		Tp:  common.Channel_ToGroup,
		Mid: mid,
	})
	if err != nil {
		return err
	}
	target := record.ReceiverId
	if record.SenderId == Operator && time.Since(util.UnixToTime(int64(record.CreateTime))) > time.Duration(l.svcCtx.Config.Revoke.Expire) {
		return model.ErrPermission
	}
	if record.SenderId != Operator {
		//执行者
		memOpt, err := l.svcCtx.GroupRPC.GetMember(l.ctx, &groupApi.GetMemberReq{
			MemberId: Operator,
			GroupId:  util.MustToInt64(target),
		})
		if err != nil || memOpt == nil {
			return err
		}
		//消息所有者
		memOwn, err := l.svcCtx.GroupRPC.GetMember(l.ctx, &groupApi.GetMemberReq{
			MemberId: record.SenderId,
			GroupId:  util.MustToInt64(target),
		})
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
	action := &signal.SignalRevoke{
		Mid:      mid,
		Operator: Operator,
		Self:     record.SenderId == Operator,
	}
	body, err := xproto.Marshal(action)
	if err != nil {
		return errors.WithMessagef(err, "proto.Marshal, action=%+v", action)
	}
	if _, err := l.svcCtx.StorageRPC.DelRecord(l.ctx, &storageclient.DelRecordReq{
		Tp:  common.Channel_ToGroup,
		Mid: mid,
	}); err != nil {
		return err
	}
	_, err = l.svcCtx.AnswerRPC.GroupCastSignal(l.ctx, &answerclient.GroupCastSignalReq{
		Type:   signal.SignalType_Revoke,
		Target: target,
		Body:   body,
	})
	return err
}
