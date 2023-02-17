package msgfactory

import (
	"context"
	"time"

	"github.com/txchat/dtalk/api/proto/common"
	"github.com/txchat/dtalk/app/services/group/groupclient"
	"github.com/txchat/dtalk/internal/bizproto"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/imparse"
)

type Filters struct {
	groupRPCClient groupclient.Group
}

func NewFilters(groupRPCClient groupclient.Group) *Filters {
	return &Filters{
		groupRPCClient: groupRPCClient,
	}
}

func (fs *Filters) GetFilters() map[imparse.FrameType][]imparse.Filter {
	//filters
	return map[imparse.FrameType][]imparse.Filter{
		bizproto.GroupFrameType: {
			func(ctx context.Context, frame imparse.Frame) error {
				fm := frame.(*bizproto.GroupFrame)
				//判断群聊拦截
				if fm.GetMsgType() != common.MsgType_Notice {
					if ok, err := fs.checkInGroup(ctx, fm.GetFrom(), util.MustToInt64(fm.GetTarget())); !ok {
						if err != nil {
							return err
						}
						return xerror.ErrPermissionDenied
					}
				}
				return nil
			},
		},
	}
}

func (fs *Filters) checkInGroup(ctx context.Context, uid string, gid int64) (isOk bool, err error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
	reply, err := fs.groupRPCClient.CheckMemberInGroup(ctx, &groupclient.CheckMemberInGroupReq{
		Gid: gid,
		Mid: uid,
	})
	if err != nil {
		return
	}

	return reply.GetOk(), nil
}
