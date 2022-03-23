package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) UpdateGroupJoinType(req *types.UpdateGroupJoinTypeReq) (*types.UpdateGroupJoinTypeResp, error) {
	_, err := l.svcCtx.GroupClient.UpdateGroupJoinType(l.ctx, &pb.UpdateGroupJoinTypeReq{
		GroupId:       req.Id,
		PersonId:      l.getOpe(),
		GroupJoinType: pb.GroupJoinType(req.JoinType),
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateGroupJoinTypeResp{}, nil
}
