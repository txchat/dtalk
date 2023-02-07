package record

import (
	"context"
	"io/ioutil"
	"mime/multipart"

	"github.com/txchat/im/api/protocol"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"
	"github.com/txchat/dtalk/app/services/answer/answerclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
)

type PushLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewPushLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &PushLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *PushLogic) Push(req *types.PushReq, fh *multipart.FileHeader) (resp *types.PushResp, err error) {
	uid := l.custom.UID
	f, err := fh.Open()
	if err != nil {
		l.Errorf("UploadFile fh.Open err, err: %v", err)
		return nil, err
	}
	defer f.Close()

	body, err := ioutil.ReadAll(f)
	if err != nil {
		err = xerror.ErrExec
		return
	}
	if len(body) == 0 {
		err = xerror.ErrExec
		return
	}
	p := protocol.Proto{
		Ver:  1,
		Op:   int32(protocol.Op_SendMsg),
		Seq:  0,
		Ack:  0,
		Body: body,
	}
	data, err := proto.Marshal(&p)
	if err != nil {
		err = xerror.ErrExec
		return
	}

	pushResp, err := l.svcCtx.AnswerRPC.PushCommonMsg(l.ctx, &answerclient.PushCommonMsgReq{
		Key:  "",
		From: uid,
		Body: data,
	})
	if err != nil {
		err = xerror.ErrSendMsgFailed
		return
	}

	resp = &types.PushResp{
		Mid:      pushResp.GetMid(),
		Datetime: pushResp.GetTime(),
	}
	return
}
