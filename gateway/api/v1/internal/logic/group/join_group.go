package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/pkg/util"
)

func (l *GroupLogic) JoinGroup(req *types.JoinGroupReq) (*types.JoinGroupResp, error) {
	groupId := req.Id

	_, err := l.svcCtx.GroupClient.InviteGroupMembers(l.ctx, &pb.InviteGroupMembersReq{
		GroupId:   groupId,
		InviterId: req.InviterId,
		MemberIds: []string{l.getOpe()},
	})
	if err != nil {
		return nil, err
	}

	return &types.JoinGroupResp{
		Id:    groupId,
		IdStr: util.ToString(groupId),
	}, nil
}
