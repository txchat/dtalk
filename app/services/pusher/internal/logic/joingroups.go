package logic

import (
	"context"
	"strconv"
	"time"

	"github.com/txchat/dtalk/app/services/group/groupclient"

	"github.com/txchat/dtalk/app/services/pusher/internal/svc"
	logic "github.com/txchat/im/api/logic/grpc"
	"github.com/zeromicro/go-zero/core/logx"
)

type JoinGroupsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupsLogic {
	return &JoinGroupsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinGroupsLogic) getAllJoinedGroups(uid string) (groups []int64, err error) {
	ctx, cancel := context.WithDeadline(l.ctx, time.Now().Add(time.Second*15))
	defer cancel()
	reply, err := l.svcCtx.GroupRPC.JoinedGroups(ctx, &groupclient.JoinedGroupsReq{
		Uid: uid,
	})
	if err != nil {
		return
	}

	return reply.GetGid(), nil
}

func (l *JoinGroupsLogic) JoinGroups(uid, key string) error {
	groups, err := l.getAllJoinedGroups(uid)
	if err != nil {
		return err
	}

	if len(groups) == 0 {
		return nil
	}

	var gids = make([]string, len(groups))
	for i, group := range groups {
		gids[i] = strconv.FormatInt(group, 10)
	}
	_, err = l.svcCtx.LogicRPC.JoinGroupsByKeys(l.ctx, &logic.GroupsKey{
		AppId: l.svcCtx.Config.AppID,
		Keys:  []string{key},
		Gid:   gids,
	})
	return err
}
