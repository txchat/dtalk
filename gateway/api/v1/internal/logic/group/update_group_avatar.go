package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) UpdateGroupAvatar(req *types.UpdateGroupAvatarReq) (*types.UpdateGroupAvatarResp, error) {
	_, err := l.svcCtx.GroupClient.UpdateGroupAvatar(l.ctx, &pb.UpdateGroupAvatarReq{
		GroupId:  req.Id,
		PersonId: l.getOpe(),
		Avatar:   req.Avatar,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateGroupAvatarResp{}, nil
}
