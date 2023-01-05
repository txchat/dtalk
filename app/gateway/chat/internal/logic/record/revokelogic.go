package record

import (
	"context"
	"time"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/groupclient"

	"github.com/txchat/dtalk/app/gateway/chat/internal/model"
	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"
	"github.com/txchat/dtalk/app/services/storage/storageclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/dtalk/pkg/util"
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

	if _, err := l.svcCtx.StorageRPC.DelRecord(l.ctx, &storageclient.DelRecordReq{
		Tp:  common.Channel_ToUser,
		Mid: mid,
	}); err != nil {
		return err
	}

	action := &signal.SignalRevoke{
		Mid:      mid,
		Operator: operator,
		Self:     record.SenderId == operator,
	}
	err = l.svcCtx.SignalHub.RevokePrivateMessage(l.ctx, []string{operator, target}, action)
	return err
}

func (l *RevokeLogic) revokeGroup(operator string, mid int64) error {
	//查找消息
	record, err := l.svcCtx.StorageRPC.GetRecord(l.ctx, &storageclient.GetRecordReq{
		Tp:  common.Channel_ToGroup,
		Mid: mid,
	})
	if err != nil {
		return err
	}
	target := record.ReceiverId
	if record.SenderId == operator && time.Since(util.UnixToTime(int64(record.CreateTime))) > time.Duration(l.svcCtx.Config.Revoke.Expire) {
		return model.ErrPermission
	}
	gid := util.MustToInt64(target)
	if record.SenderId != operator {
		//执行者
		memOpt, err := l.svcCtx.GroupRPC.MemberInfo(l.ctx, &groupclient.MemberInfoReq{
			Gid: gid,
			Uid: operator,
		})
		if err != nil || memOpt.GetMember() == nil {
			return err
		}
		//消息所有者
		memOwn, err := l.svcCtx.GroupRPC.MemberInfo(l.ctx, &groupclient.MemberInfoReq{
			Gid: gid,
			Uid: record.SenderId,
		})
		if err != nil || memOwn.GetMember() == nil {
			return err
		}
		switch memOpt.GetMember().GetRole() {
		case group.RoleType_Owner:
		case group.RoleType_Manager:
			if memOwn.GetMember().GetRole() == group.RoleType_Owner {
				return model.ErrPermission
			}
		default:
			return model.ErrPermission
		}
	}
	if _, err := l.svcCtx.StorageRPC.DelRecord(l.ctx, &storageclient.DelRecordReq{
		Tp:  common.Channel_ToGroup,
		Mid: mid,
	}); err != nil {
		return err
	}

	action := &signal.SignalRevoke{
		Mid:      mid,
		Operator: operator,
		Self:     record.SenderId == operator,
	}
	err = l.svcCtx.SignalHub.RevokeGroupMessage(l.ctx, gid, action)
	return err
}
