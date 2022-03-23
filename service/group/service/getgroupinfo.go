package service

import (
	"context"
	"github.com/txchat/dtalk/service/group/model/types"
)

// GetGroupInfoHttp 查询群资料
func (s *Service) GetGroupInfoHttp(ctx context.Context, req *types.GetGroupInfoRequest) (res *types.GetGroupInfoResponse, err error) {
	groupId := req.Id
	personId := req.PersonId

	group, err := s.GetGroupInfoByGroupId(ctx, groupId)
	if err != nil {
		return nil, err
	}

	personInfo, err := s.GetPersonByMemberIdAndGroupId(ctx, personId, group.GroupId)
	if err != nil {
		return nil, err
	}

	ownerInfo, err := s.GetMemberByMemberIdAndGroupId(ctx, group.GroupOwnerId, group.GroupId)
	if err != nil {
		return nil, err
	}

	memberInfos, err := s.GetGroupMembersByGroupIdWithLimit(group.GroupId, 0, req.DisPlayNum)
	if err != nil {
		return nil, err
	}

	groupType := group.ToTypes(ownerInfo.ToTypes(), personInfo.ToTypes())
	memberInfosType := make([]*types.GroupMember, 0)
	for _, member := range memberInfos {
		memberInfosType = append(memberInfosType, member.ToTypes())
	}

	res = &types.GetGroupInfoResponse{}
	res.GroupInfo = groupType
	res.Members = memberInfosType
	return res, nil
}
