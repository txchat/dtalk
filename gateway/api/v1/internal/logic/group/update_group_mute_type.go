package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) UpdateGroupMuteType(req *types.UpdateGroupMuteTypeReq) (*types.UpdateGroupMuteTypeResp, error) {
	_, err := l.svcCtx.GroupClient.UpdateGroupMuteType(l.ctx, &pb.UpdateGroupMuteTypeReq{
		GroupId:       req.Id,
		PersonId:      l.getOpe(),
		GroupMuteType: pb.GroupMuteType(req.MuteType),
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateGroupMuteTypeResp{}, nil
}
