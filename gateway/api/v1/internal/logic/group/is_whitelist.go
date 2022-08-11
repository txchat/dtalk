package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/vip/api"
)

func (l *GroupLogic) IsWhitelist() (*types.IsWhitelistResp, error) {
	resp, err := l.svcCtx.VIPClient.GetVIP(l.ctx, &pb.GetVIPReq{
		Uid: l.getOpe(),
	})
	if err != nil {
		return nil, err
	}
	isExist := resp.GetVip() != nil && resp.GetVip().GetUid() == l.getOpe()
	return &types.IsWhitelistResp{Exist: isExist}, nil
}
