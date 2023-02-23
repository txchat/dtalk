package checker

import (
	"context"

	"github.com/txchat/dtalk/api/proto/chat"
	"github.com/txchat/dtalk/api/proto/message"
	"github.com/txchat/dtalk/app/services/group/groupclient"
	"github.com/txchat/dtalk/pkg/util"
)

type Filter interface {
	Filter(msg *message.Message) chat.SendMessageReply_FailedType
}

type PrivateFilter struct {
}

func (pf *PrivateFilter) Filter(msg *message.Message) chat.SendMessageReply_FailedType {
	return chat.SendMessageReply_IsOK
}

type GroupFilter struct {
	groupRPCClient groupclient.Group
}

func (gf *GroupFilter) Filter(msg *message.Message) chat.SendMessageReply_FailedType {
	ok, err := gf.memberOfGroup(context.Background(), msg.GetFrom(), util.MustToInt64(msg.GetTarget()))
	if err != nil {
		return chat.SendMessageReply_InnerError
	}
	if !ok {
		return chat.SendMessageReply_InsufficientPermission
	}
	return chat.SendMessageReply_IsOK
}

func (gf *GroupFilter) memberOfGroup(ctx context.Context, uid string, gid int64) (bool, error) {
	reply, err := gf.groupRPCClient.CheckMemberInGroup(ctx, &groupclient.CheckMemberInGroupReq{
		Gid: gid,
		Mid: uid,
	})
	if err != nil {
		return false, err
	}
	return reply.GetOk(), nil
}
