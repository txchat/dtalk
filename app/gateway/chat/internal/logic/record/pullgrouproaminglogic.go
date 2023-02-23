package record

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/chat/internal/model"

	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/app/services/storage/storageclient"
	xhttp "github.com/txchat/dtalk/pkg/net/http"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type PullGroupRoamingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewPullGroupRoamingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullGroupRoamingLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &PullGroupRoamingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *PullGroupRoamingLogic) PullGroupRoaming(req *types.PullGroupRoamingReq) (resp *types.PullGroupRoamingResp, err error) {
	uid := l.custom.UID
	rpcResp, err := l.svcCtx.StorageRPC.GetChatSessionMsg(l.ctx, &storageclient.GetChatSessionMsgReq{
		Tp:     message.Channel_Group,
		Mid:    req.Cursor,
		From:   uid,
		Target: req.Session,
		Size:   req.Size,
	})
	if err != nil {
		return
	}
	resp = &types.PullGroupRoamingResp{Records: model.ToChatRecord(rpcResp.GetRecords())}
	return
}
