package logic

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/api/proto/chat"
	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/app/services/transfer/internal/svc"
	"github.com/txchat/dtalk/app/services/transfer/transfer"
	checker "github.com/txchat/dtalk/internal/recordutil/dtalk"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/zeromicro/go-zero/core/logx"
)

type PreSendMessageCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPreSendMessageCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreSendMessageCheckLogic {
	return &PreSendMessageCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PreSendMessageCheckLogic) PreSendMessageCheck(in *transfer.PreSendMessageCheckReq) (resp *transfer.PreSendMessageCheckResp, err error) {
	chatProto := in.GetMsg()
	switch chatProto.GetType() {
	case chat.Chat_message:
		var msg message.Message
		err = proto.Unmarshal(chatProto.GetBody(), &msg)
		if err != nil {
			resp = &transfer.PreSendMessageCheckResp{
				Result: SetResult(chat.SendMessageReply_InnerError, "", 0, false),
			}
			return
		}
		// 1. 检查消息协议格式是否正确
		if failedType := checker.CheckMessage(&msg); failedType != chat.SendMessageReply_IsOK {
			resp = &transfer.PreSendMessageCheckResp{
				Result: SetResult(failedType, "", 0, false),
			}
			return
		}
		// 2. 判断消息去重
		now := util.TimeNowUnixMilli()
		var mid string
		var state chat.SendMessageReply_FailedType
		mid, state, err = l.svcCtx.Repo.GetMidByCid(l.ctx, msg.GetCid())
		if err != nil {
			resp = &transfer.PreSendMessageCheckResp{
				Result: SetResult(chat.SendMessageReply_InnerError, "", now, false),
			}
			return
		}
		repeat := mid != ""
		defer func(repeat bool) {
			s := resp.GetResult().GetCode()
			if !repeat && s != chat.SendMessageReply_InnerError {
				err = l.svcCtx.Repo.AddIndexCidMid(l.ctx, msg.GetCid(), mid, s)
			}
		}(repeat)
		if !repeat {
			mid = genServerMid(msg.GetFrom())
			// 3. 检查是否可以被发出
			filter, ok := l.svcCtx.Filters[msg.GetChannelType()]
			if !ok || filter == nil {
				resp = &transfer.PreSendMessageCheckResp{
					Result: SetResult(chat.SendMessageReply_UnSupportedMessageType, mid, now, repeat),
				}
				return
			}
			if failedType := filter.Filter(&msg); failedType != chat.SendMessageReply_IsOK {
				resp = &transfer.PreSendMessageCheckResp{
					Result: SetResult(failedType, mid, now, repeat),
				}
				return
			}
			state = chat.SendMessageReply_IsOK
		}
		resp = &transfer.PreSendMessageCheckResp{
			Result: SetResult(state, mid, now, repeat),
		}
		return
	case chat.Chat_signal:
		return &transfer.PreSendMessageCheckResp{
			Result: SetResult(chat.SendMessageReply_UnSupportedMessageType, "", 0, false),
		}, nil
	default:
		return &transfer.PreSendMessageCheckResp{
			Result: SetResult(chat.SendMessageReply_UnSupportedMessageType, "", 0, false),
		}, nil
	}
}

func genServerMid(uid string) string {
	//todo 生成mid
	return fmt.Sprintf("%d-%s-%d", time.Now().UnixMilli(), uid, rand.Intn(100))
}

func SetResult(failedType chat.SendMessageReply_FailedType, mid string, datetime int64, repeat bool) *chat.SendMessageReply {
	return &chat.SendMessageReply{
		Code:     failedType,
		Mid:      mid,
		Datetime: datetime,
		Repeat:   repeat,
	}
}
