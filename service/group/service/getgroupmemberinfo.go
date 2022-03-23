package service

import (
	"context"

	"github.com/txchat/dtalk/service/group/model/types"
)

// GetGroupMemberInfoSvc 查询群成员信息
func (s *Service) GetGroupMemberInfoSvc(ctx context.Context, req *types.GetGroupMemberInfoRequest) (res *types.GetGroupMemberInfoResponse, err error) {
	groupId := req.Id
	memberId := req.MemberId
	personId := req.PersonId

	_, err = s.GetGroupInfoByGroupId(ctx, groupId)
	if err != nil {
		return nil, err
	}

	_, err = s.GetPersonByMemberIdAndGroupId(ctx, personId, groupId)
	if err != nil {
		return nil, err
	}

	member, err := s.GetMemberByMemberIdAndGroupId(ctx, memberId, groupId)
	if err != nil {
		return nil, err
	}

	res = &types.GetGroupMemberInfoResponse{}
	res.GroupMember = member.ToTypes()
	return res, nil
}
