package service

import (
	"context"
	"errors"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/group/model"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
)

// GetGroupInfoByGroupId 根据 GroupID 查询群信息
func (s *Service) GetGroupInfoByGroupId(ctx context.Context, groupId int64) (*biz.GroupInfo, error) {
	groupInfo, err := s.dao.GetGroupInfoByGroupId(ctx, groupId)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotExist) {
			return nil, xerror.NewError(xerror.GroupNotExist)
		}
		return nil, err
	}

	if err = groupInfo.IsNormal(); err != nil {
		return nil, err
	}

	return groupInfo, nil
}

// GetGroupInfoByGroupMarkId 根据 GroupMarkId 查询群信息
// 没用到
func (s *Service) GetGroupInfoByGroupMarkId(groupMarkId string) (*biz.GroupInfo, error) {
	groupInfo, err := s.dao.GetGroupInfoByGroupMarkId(groupMarkId)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotExist) {
			return nil, xerror.NewError(xerror.GroupNotExist)
		}
		return nil, err
	}

	return groupInfo, nil
}

// UpdateGroupInfoMemberNum 更新群成员数量
func (s *Service) UpdateGroupInfoMemberNum(ctx context.Context, groupId int64) (int32, error) {
	n, err := s.dao.UpdateGroupInfoMemberNum(ctx, groupId)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// GetGroupIdsByMemberId 查询用户所有加入的群 ID
func (s *Service) GetGroupIdsByMemberId(memberId string) ([]int64, error) {
	groupIds, err := s.dao.GetGroupIdsByMemberId(memberId)
	if err != nil {
		return nil, err
	}
	return groupIds, nil
}

// GetGroupsByGroupIds 查询用户所有加入的群
// todo: 有 bug
func (s *Service) GetGroupsByGroupIds(ctx context.Context, groupIds []int64) ([]*biz.GroupInfo, error) {
	groups, err := s.dao.GetGroupInfosByGroupIds(ctx, groupIds)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

// CheckGroupMemberNum 检查新增群员是否会超过群人数上限
func (s *Service) CheckGroupMemberNum(newMembers int32, groupMaximum int32) error {
	if newMembers > groupMaximum {
		return xerror.NewError(xerror.GroupMemberLimit)
	}
	return nil
}

// UpdateGroupType 更新群类型
// rpc 专用
func (s *Service) UpdateGroupType(groupId int64, groupType int32) error {
	// todo check groupType
	_, _, err := s.dao.MaintainGroupType(&db.GroupInfo{
		GroupId:   groupId,
		GroupType: groupType,
	})
	if err != nil {
		return err
	}

	return nil
}
