package service

import (
	"context"

	"github.com/txchat/dtalk/service/group/model/biz"
)

// GetNFTGroupsExtInfoByGroupId 获取所有群拓展信息
func (s *Service) GetNFTGroupsExtInfoByGroupId(ctx context.Context) ([]int64, []*biz.NFTGroupCondition, error) {
	groupsID, err := s.dao.GetNFTGroupsGroupId()
	if err != nil {
		return nil, nil, err
	}
	var conditions = make([]*biz.NFTGroupCondition, 0)
	var groups = make([]int64, 0)
	for _, groupId := range groupsID {
		condition, err := s.GetNFTGroupExtInfoByGroupId(ctx, groupId)
		if err != nil {
			continue
		}
		conditions = append(conditions, condition)
		groups = append(groups, groupId)
	}

	return groups, conditions, nil
}
