package service

import (
	"context"

	"github.com/txchat/dtalk/service/group/model/biz"

	"github.com/txchat/dtalk/service/group/model/types"
)

// UpdateGroupNameSvc 更新群名称
func (s *Service) UpdateGroupNameSvc(ctx context.Context, req *types.UpdateGroupNameRequest) (res *types.UpdateGroupNameResponse, err error) {
	groupId := req.Id
	personId := req.PersonId
	groupName := req.Name
	groupPubName := req.PublicName

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

	if err := s.UpdateGroupName(ctx, group, groupName, groupPubName); err != nil {
		return nil, err
	}

	return res, nil
}

// updateGroupInfoName 更新群名称
func (s *Service) updateGroupInfoName(ctx context.Context, groupId int64, name, publicName string) error {
	err := s.dao.UpdateGroupInfoName(ctx, groupId, name, publicName)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateGroupName(ctx context.Context, group *biz.GroupInfo, name, pubName string) error {
	log := s.GetLogWithTrace(ctx)
	groupId := group.GroupId
	groupNameAlert := name
	groupName := name
	groupPubName := pubName

	if groupName == "" {
		groupName = group.GroupName
	}

	if groupPubName == "" {
		groupPubName = group.GroupPubName
	}

	if err := s.updateGroupInfoName(ctx, groupId, groupName, groupPubName); err != nil {
		return err
	}

	// 发送给 pusher
	if err := s.PusherSignalGroupName(ctx, groupId, groupNameAlert); err != nil {
		log.Error().Err(err).Msg("UpdateGroupNameSvc pusher")
	}

	// 发送给 alert
	if err := s.NoticeMsgUpdateGroupName(ctx, groupId, s.GetOpe(ctx), groupNameAlert); err != nil {
		log.Error().Err(err).Msg("UpdateGroupNameSvc alert")
	}

	return nil
}
