package service

import (
	"context"

	"github.com/txchat/dtalk/pkg/contextx"

	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
)

func (s *Service) ChangeOwner(ctx context.Context, group *biz.GroupInfo, owner, member *biz.GroupMember) error {
	groupId := group.GroupId
	ownerId := owner.GroupMemberId
	memberId := member.GroupMemberId
	var err error
	log := s.GetLogWithTrace(ctx)

	if err = s.ExecChangeOwner(groupId, ownerId, memberId); err != nil {
		return err
	}

	go func() {
		// 发送给 pusher
		if err = s.PusherSignalMemberType(contextx.ValueOnlyFrom(ctx), groupId, memberId, biz.GroupMemberTypeOwner); err != nil {
			log.Error().Err(err).Msg("ChangeOwnerHttp pusher")
		}

		// 发送给 pusher
		if err = s.PusherSignalMemberType(contextx.ValueOnlyFrom(ctx), groupId, ownerId, biz.GroupMemberTypeNormal); err != nil {
			log.Error().Err(err).Msg("ChangeOwnerHttp pusher")
		}

		// 发送给 alert
		if err = s.NoticeMsgUpdateGroupOwner(contextx.ValueOnlyFrom(ctx), groupId, ownerId, memberId); err != nil {
			log.Error().Err(err).Msg("ChangeOwnerHttp alert")
		}
	}()

	// 解除新群主禁言
	nowTime := s.getNowTime()
	groupMemberMute := make([]*db.GroupMemberMute, 1, 1)
	groupMemberMute[0] = &db.GroupMemberMute{
		GroupId:                   groupId,
		GroupMemberId:             memberId,
		GroupMemberMuteTime:       0,
		GroupMemberMuteUpdateTime: nowTime,
	}
	if err = s.execSetMemberMuteTimes(ctx, groupMemberMute); err != nil {
		return err
	}

	// 发送给 pusher
	if err = s.PusherSignalMemberMuteTime(contextx.ValueOnlyFrom(ctx), groupId, []string{memberId}, 0); err != nil {
		log.Error().Err(err).Msg("UpdateMembersMuteTimeSvc pusher")
	}

	return nil
}

// ExecChangeOwner 转让群主
func (s *Service) ExecChangeOwner(groupId int64, ownerId, memberId string) error {
	// todo
	// 如果并发换群主可能出现两个群主的情况
	nowTime := s.getNowTime()
	tx, err := s.dao.NewTx()
	if err != nil {
		return err
	}
	defer tx.RollBack()

	groupInfo := &db.GroupInfo{
		GroupId:         groupId,
		GroupOwnerId:    memberId,
		GroupUpdateTime: nowTime,
	}
	if _, _, err = s.dao.UpdateGroupInfoOwnerIdWithTx(tx, groupInfo); err != nil {
		return err
	}

	owner := &db.GroupMember{
		GroupId:               groupId,
		GroupMemberId:         ownerId,
		GroupMemberType:       biz.GroupMemberTypeNormal,
		GroupMemberUpdateTime: nowTime,
	}

	if _, _, err := s.dao.UpdateGroupMemberTypeWithTx(tx, owner); err != nil {
		return err
	}

	member := &db.GroupMember{
		GroupId:               groupId,
		GroupMemberId:         memberId,
		GroupMemberType:       biz.GroupMemberTypeOwner,
		GroupMemberUpdateTime: nowTime,
	}

	if _, _, err = s.dao.UpdateGroupMemberTypeWithTx(tx, member); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
