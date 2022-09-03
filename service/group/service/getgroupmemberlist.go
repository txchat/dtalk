package service

import (
	"context"

	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model/types"
)

// GetGroupMemberListSvc 查询群成员列表
func (s *Service) GetGroupMemberListSvc(ctx context.Context, req *types.GetGroupMemberListRequest) (res *types.GetGroupMemberListResponse, err error) {
	groupId := req.Id
	personId := req.PersonId

	_, err = s.GetGroupInfoByGroupId(ctx, groupId)
	if err != nil {
		return nil, err
	}

	_, err = s.GetPersonByMemberIdAndGroupId(ctx, personId, groupId)
	if err != nil {
		return nil, err
	}

	groupMembers, err := s.GetMembersByGroupId(groupId)
	if err != nil {
		return nil, err
	}

	groupMemberTypes := make([]*types.GroupMember, 0)
	for _, groupMember := range groupMembers {
		groupMemberTypes = append(groupMemberTypes, groupMember.ToTypes())
	}

	res = &types.GetGroupMemberListResponse{
		Id:      groupId,
		IdStr:   util.MustToString(groupId),
		Members: groupMemberTypes,
	}

	return res, nil

}
