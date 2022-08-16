package logic

import (
	"context"

	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type GetNFTGroupsExtInfoLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewGetNFTGroupsExtInfoLogic(ctx context.Context, svc *service.Service) *GetNFTGroupsExtInfoLogic {
	return &GetNFTGroupsExtInfoLogic{
		ctx: ctx,
		svc: svc,
	}
}

// GetNFTGroupsExtInfo 获取所有藏品群拓展信息
func (l *GetNFTGroupsExtInfoLogic) GetNFTGroupsExtInfo(req *pb.GetNFTGroupsExtInfoReq) (*pb.GetNFTGroupsExtInfoResp, error) {
	groups, conditions, err := l.svc.GetNFTGroupsExtInfoByGroupId(l.ctx)
	if err != nil {
		return &pb.GetNFTGroupsExtInfoResp{}, err
	}
	pbConditions := make([]*pb.GetNFTGroupsExtInfoResp_ConditionUnionGroupId, 0)
	for i, condition := range conditions {
		nfts := make([]*pb.Condition_NFT, len(condition.NFT))
		for i, nft := range condition.NFT {
			nfts[i] = &pb.Condition_NFT{
				Type: nft.Type,
				Name: nft.Name,
				Id:   nft.Id,
			}
		}

		item := &pb.GetNFTGroupsExtInfoResp_ConditionUnionGroupId{
			GroupId: groups[i],
			Condition: &pb.Condition{
				Type: condition.Type,
				Nft:  nfts,
			},
		}
		pbConditions = append(pbConditions, item)
	}

	return &pb.GetNFTGroupsExtInfoResp{
		Conditions: pbConditions,
	}, nil
}
