package service

import (
	"context"
	"github.com/txchat/dtalk/service/group/model/biz"

	"github.com/txchat/dtalk/service/group/model/types"
)

// UpdateGroupAvatarSvc 更新群头像
func (s *Service) UpdateGroupAvatarSvc(ctx context.Context, req *types.UpdateGroupAvatarRequest) (res *types.UpdateGroupAvatarResponse, err error) {
	groupId := req.Id
	personId := req.PersonId
	groupAvatar := req.Avatar

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

	if err := s.UpdateGroupAvatar(ctx, group, groupAvatar); err != nil {
		return nil, err
	}

	return res, nil
}

// updateGroupAvatar 更新db群头像
func (s *Service) updateGroupAvatar(ctx context.Context, groupId int64, avatar string) error {
	err := s.dao.UpdateGroupInfoAvatar(ctx, groupId, avatar)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateGroupAvatar(ctx context.Context, group *biz.GroupInfo, avatar string) error {
	log := s.GetLogWithTrace(ctx)
	groupId := group.GroupId

	if err := s.updateGroupAvatar(ctx, groupId, avatar); err != nil {
		return err
	}

	// 发送给 pusher
	if err := s.PusherSignalGroupAvatar(ctx, groupId, avatar); err != nil {
		log.Error().Err(err).Msg("UpdateGroupAvatar pusher")
	}

	return nil
}
