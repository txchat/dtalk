package logic

import (
	"context"

	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/service"
)

type CreateNFTGroupLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewCreateNFTGroupLogic(ctx context.Context, svc *service.Service) *CreateNFTGroupLogic {
	return &CreateNFTGroupLogic{
		ctx: ctx,
		svc: svc,
	}
}

// CreateGroup 创建群聊
func (l *CreateNFTGroupLogic) CreateNFTGroup(req *pb.CreateNFTGroupReq) (*pb.CreateNFTGroupResp, error) {
	var err error
	req.Owner.Id, err = FilteredMemberId(req.Owner.Id)
	if err != nil {
		return nil, err
	}

	group := &biz.GroupInfo{
		GroupName: req.Name,
		GroupType: int32(req.GroupType),
	}

	var nftGroupCondition *biz.NFTGroupCondition
	if condition := req.GetCondition(); condition != nil {
		var nfts = make([]*biz.NFT, len(condition.GetNft()))
		for i, nft := range condition.GetNft() {
			nfts[i] = &biz.NFT{
				Type: nft.Type,
				Name: nft.Name,
				Id:   nft.Id,
			}
		}
		nftGroupCondition = &biz.NFTGroupCondition{
			Type: condition.GetType(),
			NFT:  nfts,
		}
	}

	owner := &biz.GroupMember{
		GroupMemberId:   req.Owner.Id,
		GroupMemberName: req.Owner.Name,
	}

	members := make([]*biz.GroupMember, 0, len(req.Members))
	for _, member := range req.Members {
		if member.Id == owner.GroupMemberId {
			continue
		}

		member.Id, err = FilteredMemberId(member.Id)
		if err != nil {
			return nil, err
		}

		members = append(members, &biz.GroupMember{
			GroupMemberId:   member.Id,
			GroupMemberName: member.Name,
		})
	}

	groupId, err := l.svc.CreateNFTGroup(l.ctx, nftGroupCondition, group, owner, members)
	if err != nil {
		return nil, err
	}

	return &pb.CreateNFTGroupResp{
		GroupId: groupId,
	}, nil
}
