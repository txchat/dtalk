package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) UpdateGroupFriendType(req *types.UpdateGroupFriendTypeReq) (*types.UpdateGroupFriendTypeResp, error) {
	_, err := l.svcCtx.GroupClient.UpdateGroupFriendType(l.ctx, &pb.UpdateGroupFriendTypeReq{
		GroupId:         req.Id,
		PersonId:        l.getOpe(),
		GroupFriendType: pb.GroupFriendType(req.FriendType),
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateGroupFriendTypeResp{}, nil
}
