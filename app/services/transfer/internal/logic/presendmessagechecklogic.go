package logic

import (
	"context"

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

func (l *PreSendMessageCheckLogic) PreSendMessageCheck(in *transfer.PreSendMessageCheckReq) (*transfer.PreSendMessageCheckResp, error) {
	chatProto := in.GetMsg()
	switch chatProto.GetType() {
	case chat.Chat_message:
		var msg *message.Message
		err := proto.Unmarshal(chatProto.GetBody(), msg)
		if err != nil {
			return &transfer.PreSendMessageCheckResp{
				Result: GetResultFailed(chat.SendMessageReply_InnerError),
			}, nil
		}
		// 1. 检查消息协议格式是否正确
		if failedType := checker.CheckMessage(msg); failedType != chat.SendMessageReply_IsOK {
			return &transfer.PreSendMessageCheckResp{
				Result: GetResultFailed(failedType),
			}, nil
		}
		// 2. 判断消息去重
		now := util.TimeNowUnixMilli()
		mid, err := l.svcCtx.Repo.GetMidByCid(l.ctx, msg.GetCid())
		if err != nil {
			return &transfer.PreSendMessageCheckResp{
				Result: GetResultFailed(chat.SendMessageReply_InnerError),
			}, nil
		}
		repeat := mid == ""
		if !repeat {
			mid = genServerMid(msg.GetFrom())
			// 3. 检查是否可以被发出
			filter, ok := l.svcCtx.Filters[msg.GetChannelType()]
			if !ok || filter == nil {
				return &transfer.PreSendMessageCheckResp{
					Result: GetResultFailed(chat.SendMessageReply_UnSupportedMessageType),
				}, nil
			}
			if failedType := filter.Filter(msg); failedType != chat.SendMessageReply_IsOK {
				return &transfer.PreSendMessageCheckResp{
					Result: GetResultFailed(failedType),
				}, nil
			}
		}
		return &transfer.PreSendMessageCheckResp{
			Result: GetResultSuccess(mid, now, repeat),
		}, nil
	case chat.Chat_signal:
		return &transfer.PreSendMessageCheckResp{
			Result: GetResultFailed(chat.SendMessageReply_UnSupportedMessageType),
		}, nil
	default:
		return &transfer.PreSendMessageCheckResp{
			Result: GetResultFailed(chat.SendMessageReply_UnSupportedMessageType),
		}, nil
	}
}

func genServerMid(uid string) string {
	//todo 生成mid
	return ""
}

func GetResultFailed(failedType chat.SendMessageReply_FailedType) *chat.SendMessageReply {
	return &chat.SendMessageReply{
		Code:     failedType,
		Mid:      "",
		Datetime: 0,
		Repeat:   false,
	}
}

func GetResultSuccess(mid string, datetime int64, repeat bool) *chat.SendMessageReply {
	return &chat.SendMessageReply{
		Code:     chat.SendMessageReply_IsOK,
		Mid:      mid,
		Datetime: datetime,
		Repeat:   repeat,
	}
}
