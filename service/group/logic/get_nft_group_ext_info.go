package logic

import (
	"context"

	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type GetNFTGroupExtInfoLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewGetNFTGroupExtInfoLogic(ctx context.Context, svc *service.Service) *GetNFTGroupExtInfoLogic {
	return &GetNFTGroupExtInfoLogic{
		ctx: ctx,
		svc: svc,
	}
}

// GetNFTGroupExtInfo 查询NFT群拓展信息
func (l *GetNFTGroupExtInfoLogic) GetNFTGroupExtInfo(req *pb.GetNFTGroupExtInfoReq) (*pb.GetNFTGroupExtInfoResp, error) {
	condition, err := l.svc.GetNFTGroupExtInfoByGroupId(l.ctx, req.GetGroupId())
	if err != nil || condition == nil {
		return &pb.GetNFTGroupExtInfoResp{}, err
	}
	nfts := make([]*pb.Condition_NFT, len(condition.NFT))
	for i, nft := range condition.NFT {
		nfts[i] = &pb.Condition_NFT{
			Type: nft.Type,
			Name: nft.Name,
			Id:   nft.Id,
		}
	}
	return &pb.GetNFTGroupExtInfoResp{
		Condition: &pb.Condition{
			Type: condition.Type,
			Nft:  nfts,
		},
	}, nil
}
