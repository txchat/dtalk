package dao

import (
	"context"
	groupApi "github.com/txchat/dtalk/service/group/api"
	"time"
)

func (d *Dao) CheckInGroup(ctx context.Context, uid string, gid int64) (isOk bool, err error) {
	var (
		req   groupApi.CheckInGroupRequest
		reply *groupApi.CheckInGroupReply
	)
	req.MemberId = uid
	req.GroupId = gid
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
	reply, err = d.groupRPCClient.CheckInGroup(ctx, &req)
	if err != nil {
		return
	}

	return reply.IsOk, nil
}
