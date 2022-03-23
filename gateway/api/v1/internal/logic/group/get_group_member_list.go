package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/pkg/util"
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
