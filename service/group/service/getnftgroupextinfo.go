package service

import (
	"context"
	"errors"

	"github.com/txchat/dtalk/service/group/model"
	"github.com/txchat/dtalk/service/group/model/biz"
)

// GetNFTGroupExtInfoByGroupId 根据 GroupId 查询NFT群拓展信息
func (s *Service) GetNFTGroupExtInfoByGroupId(ctx context.Context, groupId int64) (*biz.NFTGroupCondition, error) {
	extInfo, err := s.dao.GetNFTGroupExtInfoByGroupId(groupId)
	if err != nil || extInfo == nil {
		if errors.Is(err, model.ErrRecordNotExist) {
			return nil, nil
		}
		return nil, err
	}

	conditions, err := s.dao.GetNFTGroupConditionsByGroupId(groupId)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotExist) {
			return nil, nil
		}
		return nil, err
	}

	nfts := make([]*biz.NFT, len(conditions))
	for i, condition := range conditions {
		nfts[i] = &biz.NFT{
			Type: condition.NFTType,
			Name: condition.NFTName,
			Id:   condition.NFTId,
		}
	}

	return &biz.NFTGroupCondition{
		Type: extInfo.ConditionType,
		NFT:  nfts,
	}, nil
}
