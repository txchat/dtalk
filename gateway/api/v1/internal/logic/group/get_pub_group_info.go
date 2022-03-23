package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) GetPubGroupInfo(req *types.GetGroupPubInfoReq) (*types.GetGroupPubInfoResp, error) {
	resp, err := l.svcCtx.GroupClient.GetPubGroupInfo(l.ctx, &pb.GetPubGroupInfoReq{
		GroupId:  req.Id,
		PersonId: l.getOpe(),
	})
	if err != nil {
		return nil, err
	}

	res := &types.GetGroupPubInfoResp{}
	Group := NewTypesGroupInfo(resp.Group)
	res.GroupInfo = Group
	return res, nil
}
