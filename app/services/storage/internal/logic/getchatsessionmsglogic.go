package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/storage/internal/model"

	"github.com/txchat/dtalk/api/proto/message"

	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/dtalk/app/services/storage/storage"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatSessionMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChatSessionMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatSessionMsgLogic {
	return &GetChatSessionMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetChatSessionMsgLogic) GetChatSessionMsg(in *storage.GetChatSessionMsgReq) (*storage.GetChatSessionMsgReply, error) {
	var records []*model.MsgContent
	var err error
	switch in.GetTp() {
	case message.Channel_Private:
		records, err = l.svcCtx.Repo.GetPrivateChatSessionMsg(in.GetFrom(), in.GetTarget(), in.GetMid(), in.GetSize())
		if err != nil {
			return nil, err
		}
	case message.Channel_Group:
		records, err = l.svcCtx.Repo.GetGroupChatSessionMsg(in.GetFrom(), in.GetTarget(), in.GetMid(), in.GetSize())
		if err != nil {
			return nil, err
		}
	default:
		return nil, model.ErrChannelType
	}
	return &storage.GetChatSessionMsgReply{
		Records: model.ChatMessagesRepoToRPC(records),
	}, nil
}
