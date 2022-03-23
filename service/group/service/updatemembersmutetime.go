package service

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
	"github.com/txchat/dtalk/service/group/model/types"
)

// UpdateMembersMuteTimeSvc 设置群成员禁言时间
func (s *Service) UpdateMembersMuteTimeSvc(ctx context.Context, req *types.UpdateGroupMemberMuteTimeRequest) (res *types.UpdateGroupMemberMuteTimeResponse, err error) {
	groupId := req.Id
	memberIds := req.MemberIds
	personId := req.PersonId
	muteTime := req.MuteTime
	//nowTime := s.getNowTime()
	//log := s.GetLogWithTrace(ctx)
	var members []*biz.GroupMember

	//if muteTime != biz.MuteMaximum {
	//	if muteTime > (24 * 60 * 60 * 1000) {
	//		return nil, xerror.NewError(xerror.ParamsError)
	//	}
	//	if muteTime != 0 {
	//		muteTime += nowTime
	//	}
	//}

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

	// 过滤
	for _, memberId := range memberIds {
		member, err := s.GetMemberByMemberIdAndGroupId(ctx, memberId, groupId)
		if err != nil {
			return nil, err
		}
		if err := member.IsAdmin(); err == nil {
			return nil, xerror.NewError(xerror.GroupMutePermission)
		}

		members = append(members, member)
	}

	members, err = s.UpdateMembersMuteTime(ctx, group, members, muteTime)
	if err != nil {
		return nil, err
	}

	//groupMemberMute := make([]*db.GroupMemberMute, len(memberIds), len(memberIds))
	//for i, memberId := range memberIds {
	//	groupMemberMute[i] = &db.GroupMemberMute{
	//		GroupId:                   groupId,
	//		GroupMemberId:             memberId,
	//		GroupMemberMuteTime:       muteTime,
	//		GroupMemberMuteUpdateTime: nowTime,
	//	}
	//}
	//if err = s.execSetMemberMuteTimes(ctx, groupMemberMute); err != nil {
	//	return nil, err
	//}
	//
	//// 发送给 pusher
	//if err = s.PusherSignalMemberMuteTime(ctx, groupId, memberIds, muteTime); err != nil {
	//	log.Error().Err(err).Msg("UpdateMembersMuteTimeSvc pusher")
	//}
	//
	//// 发送给 alert
	//if muteTime != 0 {
	//	if err = s.NoticeMsgUpdateGroupMemberMutedTime(ctx, groupId, personId, memberIds); err != nil {
	//		log.Error().Err(err).Msg("UpdateMembersMuteTimeSvc alert")
	//	}
	//}

	res = &types.UpdateGroupMemberMuteTimeResponse{
		Members: make([]*types.GroupMember, 0, len(members)),
	}

	for _, member := range members {
		res.Members = append(res.Members, member.ToTypes())
	}

	return res, nil
}

// execSetMemberMuteTimes 执行设置群员禁言
func (s *Service) execSetMemberMuteTimes(ctx context.Context, memberMutes []*db.GroupMemberMute) error {
	tx, err := s.dao.NewTx()
	if err != nil {
		return err
	}
	defer tx.RollBack()
	if err = s.dao.UpdateGroupMemberMuteTimes(ctx, tx, memberMutes); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateMembersMuteTime(ctx context.Context, group *biz.GroupInfo, members []*biz.GroupMember, muteTime int64) ([]*biz.GroupMember, error) {
	nowTime := s.getNowTime()
	groupId := group.GroupId
	log := s.GetLogWithTrace(ctx)
	var memberIds []string

	if muteTime != biz.MuteMaximum {
		if muteTime > (24 * 60 * 60 * 1000) {
			return nil, xerror.NewError(xerror.ParamsError)
		}
		if muteTime != 0 {
			muteTime += nowTime
		}
	}

	groupMemberMutes := make([]*db.GroupMemberMute, 0, len(members))
	for _, member := range members {
		groupMemberMutes = append(groupMemberMutes, &db.GroupMemberMute{
			GroupId:             groupId,
			GroupMemberId:       member.GroupMemberId,
			GroupMemberMuteTime: muteTime,
		})
		memberIds = append(memberIds, member.GroupMemberId)
		member.GroupMemberMuteTime = muteTime
	}

	if err := s.execSetMemberMuteTimes(ctx, groupMemberMutes); err != nil {
		return nil, err
	}

	// 发送给 pusher
	if err := s.PusherSignalMemberMuteTime(ctx, groupId, memberIds, muteTime); err != nil {
		log.Error().Err(err).Msg("UpdateMembersMuteTimeSvc pusher")
	}

	// 发送给 alert
	if muteTime != 0 {
		if err := s.NoticeMsgUpdateGroupMemberMutedTime(ctx, groupId, s.GetOpe(ctx), memberIds); err != nil {
			log.Error().Err(err).Msg("UpdateMembersMuteTimeSvc alert")
		}
	}

	return members, nil
}
