package msgfactory

import (
	"context"
	"time"

	"github.com/txchat/dtalk/pkg/util"
	groupApi "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/record/answer/model"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"
	"github.com/txchat/imparse/proto/common"
)

type Filters struct {
	groupRPCClient groupApi.GroupClient
}

func NewFilters(groupRPCClient groupApi.GroupClient) *Filters {
	return &Filters{
		groupRPCClient: groupRPCClient,
	}
}

func (fs *Filters) GetFilters() map[imparse.FrameType][]imparse.Filter {
	//filters
	return map[imparse.FrameType][]imparse.Filter{
		chat.GroupFrameType: {
			func(ctx context.Context, frame imparse.Frame) error {
				fm := frame.(*chat.GroupFrame)
				//判断群聊拦截
				if fm.GetMsgType() != common.MsgType_Notice {
					if ok, err := fs.checkInGroup(ctx, fm.GetFrom(), util.MustToInt64(fm.GetTarget())); !ok {
						if err != nil {
							return err
						}
						return model.ErrGroupMemberNotExists
					}
				}
				return nil
			},
		},
	}
}

func (fs *Filters) checkInGroup(ctx context.Context, uid string, gid int64) (isOk bool, err error) {
	var (
		req   groupApi.CheckInGroupRequest
		reply *groupApi.CheckInGroupReply
	)
	req.MemberId = uid
	req.GroupId = gid
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
	reply, err = fs.groupRPCClient.CheckInGroup(ctx, &req)
	if err != nil {
		return
	}

	return reply.IsOk, nil
}
