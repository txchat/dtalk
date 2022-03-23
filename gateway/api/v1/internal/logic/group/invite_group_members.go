package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/pkg/util"
)

func (l *GroupLogic) InviteGroupMembers(req *types.InviteGroupMembersReq) (*types.InviteGroupMembersResp, error) {
	groupId := req.Id

	_, err := l.svcCtx.GroupClient.InviteGroupMembers(l.ctx, &pb.InviteGroupMembersReq{
		GroupId:   groupId,
		InviterId: l.getOpe(),
		MemberIds: req.NewMemberIds,
	})
	if err != nil {
		return nil, err
	}

	return &types.InviteGroupMembersResp{
		Id:    groupId,
		IdStr: util.ToString(groupId),
	}, nil
}
