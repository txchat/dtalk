package service

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/types"
)

// UpdateGroupFriendTypeSvc 更新群内加好友设置
func (s *Service) UpdateGroupFriendTypeSvc(ctx context.Context, req *types.UpdateGroupFriendTypeRequest) (res *types.UpdateGroupFriendTypeResponse, err error) {
	groupId := req.Id
	personId := req.PersonId
	friendType := req.FriendType

	group, err := s.GetGroupInfoByGroupId(ctx, groupId)
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

	if err := s.UpdateGroupFriendType(ctx, group, friendType); err != nil {
		return nil, err
	}

	return res, nil
}

// CheckGroupFriendType 检查friendType是否合法
func (s *Service) CheckGroupFriendType(friendType int32) (int32, error) {
	switch friendType {
	case biz.GroupFriendTypeAllow:
		return biz.GroupFriendTypeAllow, nil
	case biz.GroupFriendTypeDeny:
		return biz.GroupFriendTypeDeny, nil
	default:
		return 0, xerror.NewError(xerror.ParamsError)
	}
}

// updateGroupFriendType 更新群内加好友设置
func (s *Service) updateGroupFriendType(ctx context.Context, groupId int64, friendType int32) error {
	err := s.dao.UpdateGroupInfoFriendType(ctx, groupId, friendType)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateGroupFriendType(ctx context.Context, group *biz.GroupInfo, friendType int32) error {
	log := s.GetLogWithTrace(ctx)
	groupId := group.GroupId

	friendType, err := s.CheckGroupFriendType(friendType)
	if err != nil {
		return err
	}

	if err = s.updateGroupFriendType(ctx, groupId, friendType); err != nil {
		return err
	}

	// 发送给 pusher
	if err = s.PusherSignalFriendType(ctx, groupId, friendType); err != nil {
		log.Error().Err(err).Msg("UpdateGroupFriendTypeSvc pusher")
	}

	return nil
}
