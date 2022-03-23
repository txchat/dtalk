package service

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/types"
)

// UpdateGroupJoinTypeSvc 更新加群设置
// todo 改造完后就弃用
func (s *Service) UpdateGroupJoinTypeSvc(ctx context.Context, req *types.UpdateGroupJoinTypeRequest) (res *types.UpdateGroupJoinTypeResponse, err error) {
	groupId := req.Id
	personId := req.PersonId
	joinType := req.JoinType

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

	if err := s.UpdateGroupJoinType(ctx, group, joinType); err != nil {
		return nil, err
	}

	return res, nil
}

// CheckGroupJoinType 检查joinType是否合法
func (s *Service) CheckGroupJoinType(joinType int32) (int32, error) {
	switch joinType {
	case biz.GroupJoinTypeAny:
		return biz.GroupJoinTypeAny, nil
	case biz.GroupJoinTypeAdmin:
		return biz.GroupJoinTypeAdmin, nil
	case biz.GroupJoinTypeApply:
		return biz.GroupJoinTypeApply, nil
	default:
		return 0, xerror.NewError(xerror.ParamsError)
	}
}

// updateGroupJoinType 更新加群设置
func (s *Service) updateGroupJoinType(ctx context.Context, groupId int64, joinType int32) error {
	err := s.dao.UpdateGroupInfoJoinType(ctx, groupId, joinType)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateGroupJoinType(ctx context.Context, group *biz.GroupInfo, joinType int32) error {
	log := s.GetLogWithTrace(ctx)
	groupId := group.GroupId

	joinType, err := s.CheckGroupJoinType(joinType)
	if err != nil {
		return err
	}

	if err = s.updateGroupJoinType(ctx, groupId, joinType); err != nil {
		return err
	}

	// 发送给 pusher
	if err = s.PusherSignalJoinType(ctx, groupId, joinType); err != nil {
		log.Error().Err(err).Msg("UpdateGroupJoinType pusher")
	}

	return nil
}
