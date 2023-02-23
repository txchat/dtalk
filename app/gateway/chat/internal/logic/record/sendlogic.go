package record

import (
	"context"
	"io/ioutil"
	"mime/multipart"

	"github.com/txchat/dtalk/app/services/transfer/transferclient"

	"github.com/txchat/dtalk/api/proto/chat"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/im/api/protocol"
	"github.com/txchat/im/app/logic/logicclient"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type SendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &SendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *SendLogic) Send(req *types.SendReq, fh *multipart.FileHeader) (resp *types.SendResp, err error) {
	// todo: add your logic here and delete this line
	uid := l.custom.UID
	f, err := fh.Open()
	if err != nil {
		l.Errorf("UploadFile fh.Open err, err: %v", err)
		return nil, err
	}
	defer f.Close()

	body, err := ioutil.ReadAll(f)
	if err != nil {
		err = xerror.ErrSendMsgFailed
		return
	}
	if len(body) == 0 {
		err = xerror.ErrSendMsgFailed
		return
	}

	chatProto := chat.Chat{
		Type: chat.Chat_message,
		Seq:  0,
		Body: body,
	}
	presendResp, err := l.svcCtx.TransferRPC.PreSendMessageCheck(l.ctx, &transferclient.PreSendMessageCheckReq{
		Msg: &chatProto,
	})
	if err != nil {
		return
	}

	switch x := presendResp.GetResult().GetResult().(type) {
	case *chat.SendMessageReply_Failed_:
		return
	case *chat.SendMessageReply_Success_:
		if x.Success.GetRepeat() {
			return
		}
	}

	l.svcCtx.LogicRPC.SendByUID(l.ctx, &logicclient.SendByUIDReq{
		AppId: "",
		Uid:   uid,
		Op:    protocol.Op_Message,
		Body:  chatProto,
	})
	return
}
