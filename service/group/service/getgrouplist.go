package service

import (
	"context"

	"github.com/txchat/dtalk/service/group/model/types"
)

// TODO check

// GetGroupListSvc 查询群列表
func (s *Service) GetGroupListSvc(ctx context.Context, req *types.GetGroupListRequest) (res *types.GetGroupListResponse, err error) {
	personId := req.PersonId

	groupIds, err := s.GetGroupIdsByMemberId(personId)
	if err != nil {
		return nil, err
	}

	res = &types.GetGroupListResponse{
		Groups: make([]*types.GroupInfo, len(groupIds), len(groupIds)),
	}

	for i, id := range groupIds {
		group, err := s.GetGroupInfoByGroupId(ctx, id)
		if err != nil {
			return nil, err
		}
		ownerInfo, err := s.GetMemberByMemberIdAndGroupId(ctx, group.GroupOwnerId, group.GroupId)
		if err != nil {
			return nil, err
		}
		personInfo, err := s.GetPersonByMemberIdAndGroupId(ctx, personId, group.GroupId)
		if err != nil {
			return nil, err
		}

		groupType := group.ToTypes(ownerInfo.ToTypes(), personInfo.ToTypes())

		res.Groups[i] = groupType
	}

	return res, nil
}
