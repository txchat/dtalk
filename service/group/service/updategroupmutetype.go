package service

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/types"
)

// UpdateGroupMuteTypeSvc 更新群禁言设置
func (s *Service) UpdateGroupMuteTypeSvc(ctx context.Context, req *types.UpdateGroupMuteTypeRequest) (res *types.UpdateGroupMuteTypeResponse, err error) {
	groupId := req.Id
	personId := req.PersonId
	muteType := req.MuteType

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

	if err := s.UpdateGroupMuteType(ctx, group, muteType); err != nil {
		return nil, err
	}

	return res, nil
}

// CheckGroupMuteType 检查muteType是否合法
func (s *Service) CheckGroupMuteType(muteType int32) (int32, error) {
	switch muteType {
	case biz.GroupMuteTypeAny:
		return biz.GroupMuteTypeAny, nil
	case biz.GroupMuteTypeAdmin:
		return biz.GroupMuteTypeAdmin, nil
	default:
		return 0, xerror.NewError(xerror.ParamsError)
	}
}

// updateGroupMuteType 更新群禁言设置
func (s *Service) updateGroupMuteType(ctx context.Context, groupId int64, muteType int32) error {
	err := s.dao.UpdateGroupInfoMuteType(ctx, groupId, muteType)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateGroupMuteType(ctx context.Context, group *biz.GroupInfo, muteType int32) error {
	log := s.GetLogWithTrace(ctx)
	groupId := group.GroupId

	muteType, err := s.CheckGroupMuteType(muteType)
	if err != nil {
		return err
	}

	if err = s.updateGroupMuteType(ctx, groupId, muteType); err != nil {
		return err
	}

	// 发送给 pusher
	if err = s.PusherSignalMuteType(ctx, groupId, muteType); err != nil {
		log.Error().Err(err).Msg("UpdateGroupMuteType pusher")
	}

	// 发送给 alert
	if err = s.NoticeMsgUpdateGroupMuted(ctx, groupId, s.GetOpe(ctx), muteType); err != nil {
		log.Error().Err(err).Msg("UpdateGroupMuteType alert")
	}

	return nil
}
