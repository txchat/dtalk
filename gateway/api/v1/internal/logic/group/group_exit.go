package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) GroupExit(req *types.GroupExitReq) (*types.GroupExitResp, error) {
	_, err := l.svcCtx.GroupClient.GroupExit(l.ctx, &pb.GroupExitReq{
		GroupId:  req.Id,
		PersonId: l.getOpe(),
	})
	if err != nil {
		return nil, err
	}

	return &types.GroupExitResp{}, nil
}
