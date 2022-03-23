package service

import (
	"context"

	"github.com/txchat/dtalk/pkg/contextx"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
	"github.com/txchat/dtalk/service/group/model/types"
)

// GroupDisbandHttp 解散群
func (s *Service) GroupDisbandHttp(ctx context.Context, req *types.GroupDisbandRequest) (res *types.GroupDisbandResponse, err error) {
	groupId := req.Id
	personId := req.PersonId

	person, err := s.GetPersonByMemberIdAndGroupId(ctx, personId, groupId)
	if err != nil {
		return nil, err
	}

	// 只有群主可以解散群
	if err = person.IsOwner(); err != nil {
		return nil, err
	}

	err = s.GroupDisband(ctx, groupId, personId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) GroupDisband(ctx context.Context, groupId int64, opeId string) error {
	var err error
	log := s.GetLogWithTrace(ctx)

	// 执行解散群
	if err = s.ExecGroupDisband(groupId); err != nil {
		return err
	}

	go func() {
		// 发送给 alert
		if err = s.NoticeMsgDeleteGroup(contextx.ValueOnlyFrom(ctx), groupId, opeId); err != nil {
			log.Error().Err(err).Msg("GroupDisband alert")
		}
		// 发送给 pusher
		if err = s.PusherSignalDel(contextx.ValueOnlyFrom(ctx), groupId); err != nil {
			log.Error().Err(err).Msg("GroupDisband pusher")
		}
		// 发送给 logic
		if err = s.LogicNoticeDel(contextx.ValueOnlyFrom(ctx), groupId); err != nil {
			log.Error().Err(err).Msg("GroupDisband logic")
		}
	}()

	return nil
}

// ExecGroupDisband 执行解散群操作
func (s *Service) ExecGroupDisband(groupId int64) error {
	nowTime := s.getNowTime()
	groupMembers, err := s.GetMembersByGroupId(groupId)
	if err != nil {
		return err
	}

	tx, err := s.dao.NewTx()
	if err != nil {
		return err
	}
	defer tx.RollBack()

	groupInfo := &db.GroupInfo{
		GroupId:         groupId,
		GroupStatus:     biz.GroupStatusDisBand,
		GroupUpdateTime: nowTime,
	}
	if _, _, err = s.dao.UpdateGroupInfoStatusWithTx(tx, groupInfo); err != nil {
		return err
	}

	for _, member := range groupMembers {
		groupMemberInfo := &db.GroupMember{
			GroupId:               groupId,
			GroupMemberId:         member.GroupMemberId,
			GroupMemberUpdateTime: nowTime,
			GroupMemberType:       biz.GroupMemberTypeOther,
		}
		if _, _, err = s.dao.UpdateGroupMemberTypeWithTx(tx, groupMemberInfo); err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
