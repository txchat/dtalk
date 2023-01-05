package logic

import (
	"context"
	"time"

	"github.com/txchat/dtalk/app/services/group/groupclient"

	"github.com/txchat/dtalk/app/services/storage/internal/svc"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMembersLogic {
	return &GetGroupMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupMembersLogic) GetGroupMembersLogic(gid string) ([]string, error) {
	ctx, cancel := context.WithDeadline(l.ctx, time.Now().Add(time.Second*3))
	defer cancel()
	reply, err := l.svcCtx.GroupRPC.GroupLimitedMembers(ctx, &groupclient.GroupLimitedMembersReq{
		Gid: util.MustToInt64(gid),
	})
	if err != nil {
		return nil, err
	}
	mid := make([]string, 0, len(reply.GetMembers()))
	for _, member := range reply.GetMembers() {
		mid = append(mid, member.GetUid())
	}
	return mid, nil
}
