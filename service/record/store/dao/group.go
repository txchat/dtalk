package dao

import (
	"context"
	"time"

	"github.com/txchat/dtalk/pkg/util"
	groupApi "github.com/txchat/dtalk/service/group/api"
)

func (d *Dao) AllGroupMembers(ctx context.Context, gid string) ([]string, error) {
	var (
		req   groupApi.GetMemberIdsRequest
		reply *groupApi.GetMemberIdsReply
	)
	req.GroupId = util.MustToInt64(gid)
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
	reply, err := d.groupRPCClient.GetMemberIds(ctx, &req)
	if err != nil {
		return nil, err
	}

	return reply.MemberIds, nil
}
