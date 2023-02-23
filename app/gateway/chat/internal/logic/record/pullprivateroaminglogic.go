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

type PullPrivateRoamingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewPullPrivateRoamingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullPrivateRoamingLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &PullPrivateRoamingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *PullPrivateRoamingLogic) PullPrivateRoaming(req *types.PullPrivateRoamingReq) (resp *types.PullPrivateRoamingResp, err error) {
	uid := l.custom.UID
	rpcResp, err := l.svcCtx.StorageRPC.GetChatSessionMsg(l.ctx, &storageclient.GetChatSessionMsgReq{
		Tp:     message.Channel_Private,
		Mid:    req.Cursor,
		From:   uid,
		Target: req.Session,
		Size:   req.Size,
	})
	if err != nil {
		return
	}

	resp = &types.PullPrivateRoamingResp{Records: model.ToChatRecord(rpcResp.GetRecords())}
	return
}
