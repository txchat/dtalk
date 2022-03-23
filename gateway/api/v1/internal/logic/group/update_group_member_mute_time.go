package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) UpdateGroupMemberMuteTime(req *types.UpdateGroupMemberMuteTimeReq) (*types.UpdateGroupMemberMuteTimeResp, error) {
	resp, err := l.svcCtx.GroupClient.UpdateGroupMemberMuteTime(l.ctx, &pb.UpdateGroupMemberMuteTimeReq{
		GroupId:   req.Id,
		PersonId:  l.getOpe(),
		MemberIds: req.MemberIds,
		MuteTime:  req.MuteTime,
	})
	if err != nil {
		return nil, err
	}

	Members := NewTypesGroupMemberInfos(resp.Members)

	return &types.UpdateGroupMemberMuteTimeResp{Members: Members}, nil
}
