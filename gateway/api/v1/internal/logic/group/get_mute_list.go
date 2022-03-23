package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) GetMuteList(req *types.GetMuteListReq) (*types.GetMuteListResp, error) {
	resp, err := l.svcCtx.GroupClient.GetMuteList(l.ctx, &pb.GetMuteListReq{
		GroupId:  req.Id,
		PersonId: l.getOpe(),
	})
	if err != nil {
		return nil, err
	}

	Members := NewTypesGroupMemberInfos(resp.Members)
	return &types.GetMuteListResp{Members: Members}, nil
}
