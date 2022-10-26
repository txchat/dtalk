package logic

import (
	"context"
	"time"

	"github.com/txchat/dtalk/pkg/util"
	groupApi "github.com/txchat/dtalk/service/group/api"

	"github.com/txchat/dtalk/app/services/storage/internal/svc"
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
	var (
		req   groupApi.GetMemberIdsRequest
		reply *groupApi.GetMemberIdsReply
	)
	req.GroupId = util.MustToInt64(gid)
	ctx, cancel := context.WithDeadline(l.ctx, time.Now().Add(time.Second*3))
	defer cancel()
	reply, err := l.svcCtx.GroupRPC.GetMemberIds(ctx, &req)
	if err != nil {
		return nil, err
	}

	return reply.MemberIds, nil
}
