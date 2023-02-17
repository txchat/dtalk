package record

import (
	"context"
	"time"

	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/api/proto/signal"
	"github.com/txchat/dtalk/app/gateway/chat/internal/model"
	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"
	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/groupclient"
	"github.com/txchat/dtalk/app/services/storage/storageclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/dtalk/pkg/util"
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
		Tp:  message.Channel_Private,
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
	err = l.svcCtx.SignalHub.FocusPrivateMessage(l.ctx, []string{sender, target}, action)
	return err
}

func (l *FocusLogic) focusGroup(operator string, mid int64) error {
	//查找消息
	rd, err := l.svcCtx.StorageRPC.GetRecord(l.ctx, &storageclient.GetRecordReq{
		Tp:  message.Channel_Group,
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
	gid := util.MustToInt64(target)

	//群成员判断
	memOpt, err := l.svcCtx.GroupRPC.MemberInfo(l.ctx, &groupclient.MemberInfoReq{
		Gid: gid,
		Uid: operator,
	})
	if err != nil || memOpt.GetMember() == nil {
		return err
	}
	switch memOpt.GetMember().GetRole() {
	case group.RoleType_Owner, group.RoleType_Manager, group.RoleType_NormalMember:
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
	err = l.svcCtx.SignalHub.FocusGroupMessage(l.ctx, gid, action)
	return err
}
