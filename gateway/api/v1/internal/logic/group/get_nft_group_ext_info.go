package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) GetNFTGroupExtInfo(req *types.GetNFTGroupExtInfoReq) (*types.GetNFTGroupExtInfoResp, error) {
	resp, err := l.svcCtx.GroupClient.GetNFTGroupExtInfo(l.ctx, &pb.GetNFTGroupExtInfoReq{
		GroupId: req.Id,
	})
	if err != nil || resp == nil {
		return nil, err
	}
	condition := resp.GetCondition()
	if condition == nil {
		return nil, nil
	}
	nfts := make([]*types.NFT, len(condition.GetNft()))
	for i, nft := range condition.GetNft() {
		nfts[i] = &types.NFT{
			Type: nft.GetType(),
			Name: nft.GetName(),
			ID:   nft.GetId(),
		}
	}
	return &types.GetNFTGroupExtInfoResp{
		Condition: &types.Condition{
			Type: condition.GetType(),
			NFT:  nfts,
		},
	}, nil
}
