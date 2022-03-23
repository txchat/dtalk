package service

import (
	"context"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/types"
)

// GetMuteListSvc 查询群内禁言列表
func (s *Service) GetMuteListSvc(ctx context.Context, req *types.GetMuteListRequest) (res *types.GetMuteListResponse, err error) {
	groupId := req.Id
	personId := req.PersonId

	_, err = s.GetGroupInfoByGroupId(ctx, groupId)
	if err != nil {
		return nil, err
	}

	person, err := s.GetPersonByMemberIdAndGroupId(ctx, personId, groupId)
	if err != nil {
		return nil, err
	}
	if err = person.IsAdmin(); err != nil {
		return nil, err
	}

	muteList, err := s.GetGroupMembersMutedByGroupId(groupId)
	if err != nil {
		return nil, err
	}

	memberTypes := make([]*types.GroupMember, 0)
	for _, member := range muteList {
		memberTypes = append(memberTypes, member.ToTypes())
	}

	res = &types.GetMuteListResponse{Members: memberTypes}
	return res, nil
}

// GetGroupMembersMutedByGroupId 查询群内被禁言的群成员信息
func (s *Service) GetGroupMembersMutedByGroupId(groupId int64) ([]*biz.GroupMember, error) {
	muteList, err := s.dao.GetGroupMembersMutedByGroupId(groupId)
	if err != nil {
		return nil, err
	}
	return muteList, nil
}
