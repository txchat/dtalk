package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) GetPriGroupInfo(req *types.GetGroupInfoReq) (*types.GetGroupInfoResp, error) {
	if req.DisPlayNum == 0 {
		req.DisPlayNum = 10
	}

	resp, err := l.svcCtx.GroupClient.GetPriGroupInfo(l.ctx, &pb.GetPriGroupInfoReq{
		GroupId:    req.Id,
		PersonId:   l.getOpe(),
		DisplayNum: int32(req.DisPlayNum),
	})
	if err != nil {
		return nil, err
	}

	Group := NewTypesGroupInfo(resp.Group)
	Members := NewTypesGroupMemberInfos(resp.Group.Members)

	return &types.GetGroupInfoResp{
		GroupInfo: Group,
		Members:   Members,
	}, nil
}
