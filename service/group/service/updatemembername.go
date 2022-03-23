package service

import (
	"context"
	"github.com/txchat/dtalk/service/group/model/biz"

	"github.com/txchat/dtalk/service/group/model/types"
)

// UpdateMemberNameSvc 更新群成员昵称
func (s *Service) UpdateMemberNameSvc(ctx context.Context, req *types.UpdateGroupMemberNameRequest) (res *types.UpdateGroupMemberNameResponse, err error) {
	groupId := req.Id
	personId := req.PersonId
	memberName := req.MemberName

	group, err := s.GetGroupInfoByGroupId(ctx, groupId)
	if err != nil {
		return nil, err
	}

	person, err := s.GetPersonByMemberIdAndGroupId(ctx, personId, groupId)
	if err != nil {
		return nil, err
	}

	if err := s.UpdateMemberName(ctx, group, person, memberName); err != nil {
		return nil, err
	}

	return res, nil
}

//updateGroupMemberName 更新群成员昵称
func (s *Service) updateGroupMemberName(groupId int64, memberId, memberName string) error {
	err := s.dao.UpdateGroupMemberName(groupId, memberId, memberName)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateMemberName(ctx context.Context, group *biz.GroupInfo, person *biz.GroupMember, memberName string) error {
	groupId := group.GroupId
	personId := person.GroupMemberId

	if err := s.updateGroupMemberName(groupId, personId, memberName); err != nil {
		return err
	}

	return nil
}
