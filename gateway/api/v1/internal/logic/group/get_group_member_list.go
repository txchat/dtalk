package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	"github.com/txchat/dtalk/pkg/util"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) GetGroupMemberList(req *types.GetGroupMemberListReq) (*types.GetGroupMemberListResp, error) {
	groupId := req.Id

	resp, err := l.svcCtx.GroupClient.GetGroupMemberList(l.ctx, &pb.GetGroupMemberListReq{
		GroupId:  groupId,
		PersonId: l.getOpe(),
	})
	if err != nil {
		return nil, err
	}

	Members := NewTypesGroupMemberInfos(resp.Members)

	return &types.GetGroupMemberListResp{
		Id:      groupId,
		IdStr:   util.ToString(groupId),
		Members: Members,
	}, nil
}
