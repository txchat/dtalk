package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/pkg/util"
)

func (l *GroupLogic) ChangeOwner(req *types.ChangeOwnerReq) (*types.ChangeOwnerResp, error) {
	groupId := util.ToInt64(req.Id)
	personId := l.getOpe()
	memberId := req.MemberId

	_, err := l.svcCtx.GroupClient.ChangeOwner(l.ctx, &pb.ChangeOwnerReq{
		GroupId:  groupId,
		PersonId: personId,
		MemberId: memberId,
	})
	if err != nil {
		return nil, err
	}

	return &types.ChangeOwnerResp{}, nil
}
