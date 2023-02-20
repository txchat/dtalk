package record

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type LastPrivateArrivalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	//custom *xhttp.Custom
}

func NewLastPrivateArrivalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LastPrivateArrivalLogic {
	//c, ok := xhttp.FromContext(ctx)
	//if !ok {
	//    c = &xhttp.Custom{}
	//}
	return &LastPrivateArrivalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		//custom: c,
	}
}

func (l *LastPrivateArrivalLogic) LastPrivateArrival(req *types.LastPrivateArrivalReq) (resp *types.LastPrivateArrivalResp, err error) {
	// todo: add your logic here and delete this line

	return
}
