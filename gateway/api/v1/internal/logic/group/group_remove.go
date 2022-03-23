package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) GroupRemove(req *types.GroupRemoveReq) (*types.GroupRemoveResp, error) {
	resp, err := l.svcCtx.GroupClient.GroupRemove(l.ctx, &pb.GroupRemoveReq{
		GroupId:   req.Id,
		PersonId:  l.getOpe(),
		MemberIds: req.MemberIds,
	})
	if err != nil {
		return nil, err
	}

	return &types.GroupRemoveResp{
		MemberNum: resp.MemberNum,
		MemberIds: resp.MemberIds,
	}, nil
}
