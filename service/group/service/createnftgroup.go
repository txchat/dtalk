package service

import (
	"context"

	"github.com/txchat/dtalk/pkg/contextx"
	xerror "github.com/txchat/dtalk/pkg/error"
	xrand "github.com/txchat/dtalk/pkg/rand"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
)

// execGroupCreate 执行创建群操作
func (s *Service) execNFTGroupCreate(extInfo *db.NFTGroupInfoExt, conditions []*db.NFTGroupCondition, groupInfo *db.GroupInfo, members []*db.GroupMember) error {
	tx, err := s.dao.NewTx()
	if err != nil {
		return err
	}
	defer tx.RollBack()
	if _, _, err = s.dao.InsertGroupInfo(tx, groupInfo); err != nil {
		return err
	}

	if extInfo != nil {
		if _, _, err = s.dao.InsertNFTGroupInfoExt(tx, extInfo); err != nil {
			return err
		}
	}

	if len(conditions) > 0 {
		if _, _, err = s.dao.InsertNFTGroupConditions(tx, conditions); err != nil {
			return err
		}
	}
	// 藏品群拓展信息
	if _, _, err = s.dao.InsertGroupInfo(tx, groupInfo); err != nil {
		return err
	}

	if err = s.InsertGroupMembers(tx, members); err != nil {
		return err
	}

	if tx.Commit() != nil {
		return err
	}

	return nil
}

func (s *Service) createNFTGroup(ctx context.Context, extInfo *db.NFTGroupInfoExt, conditions []*db.NFTGroupCondition, group *db.GroupInfo, groupMembers []*db.GroupMember) error {
	log := s.GetLogWithTrace(ctx)
	var err error
	groupId := group.GroupId
	groupMemberIds := make([]string, 0, len(groupMembers))
	groupOwnerId := ""
	for _, groupMember := range groupMembers {
		if groupMember.GroupMemberType == biz.GroupMemberTypeOwner {
			groupOwnerId = groupMember.GroupMemberId
			continue
		}
		groupMemberIds = append(groupMemberIds, groupMember.GroupMemberId)
	}
	groupMemberIds = append(groupMemberIds, groupOwnerId)

	if err = s.execNFTGroupCreate(extInfo, conditions, group, groupMembers); err != nil {
		return err
	}

	go func() {
		// 发送给 logic
		if err = s.LogicNoticeJoin(contextx.ValueOnlyFrom(ctx), groupId, groupMemberIds); err != nil {
			log.Error().Err(err).Msg("createGroup logic")
		}

		// 发送给 pusher
		if err = s.PusherSignalJoin(contextx.ValueOnlyFrom(ctx), groupId, groupMemberIds); err != nil {
			log.Error().Err(err).Msg("createGroup pusher")
		}

		if err = s.NoticeMsgSignInGroup(contextx.ValueOnlyFrom(ctx), groupId, s.GetOpe(ctx), groupMemberIds[0:len(groupMemberIds)-1]); err != nil {
			log.Error().Err(err).Msg("createGroup alert")
		}
	}()

	return nil
}

func (s *Service) CreateNFTGroup(ctx context.Context, condition *biz.NFTGroupCondition, group *biz.GroupInfo, owner *biz.GroupMember, members []*biz.GroupMember) (int64, error) {
	groupId, err := s.getLogId(ctx)
	if err != nil {
		return 0, err
	}

	groupMarkId, err := s.getRandomGroupMarkId()
	if err != nil {
		return 0, err
	}

	groupMemberNum := int32(1 + len(members))
	if groupMemberNum > biz.GroupMaximum {
		return 0, xerror.NewError(xerror.GroupMemberLimit)
	}

	groupPo := &db.GroupInfo{
		GroupId:         groupId,
		GroupMarkId:     groupMarkId,
		GroupName:       group.GroupName,
		GroupAvatar:     "",
		GroupMemberNum:  int32(1 + len(members)),
		GroupMaximum:    biz.GroupMaximum,
		GroupIntroduce:  "",
		GroupStatus:     biz.GroupStatusNormal,
		GroupOwnerId:    owner.GroupMemberId,
		GroupJoinType:   biz.GroupJoinTypeAny,
		GroupMuteType:   biz.GroupMuteTypeAny,
		GroupFriendType: biz.GroupFriendTypeAllow,
		GroupAESKey:     xrand.NewAESKey256(),
		GroupPubName:    group.GroupName,
		GroupType:       group.GroupType,
	}

	if group.GroupType == biz.GroupTypeNormal {
		groupPo.GroupJoinType = biz.GroupJoinTypeAny
	} else {
		groupPo.GroupJoinType = biz.GroupJoinTypeAdmin
	}

	ownerPo := &db.GroupMember{
		GroupId:         groupId,
		GroupMemberId:   owner.GroupMemberId,
		GroupMemberName: owner.GroupMemberName,
		GroupMemberType: biz.GroupMemberTypeOwner,
	}

	membersPo := make([]*db.GroupMember, 0, len(members))
	membersPo = append(membersPo, ownerPo)
	for _, member := range members {
		membersPo = append(membersPo, &db.GroupMember{
			GroupId:         groupId,
			GroupMemberId:   member.GroupMemberId,
			GroupMemberName: member.GroupMemberName,
			GroupMemberType: biz.GroupMemberTypeNormal,
		})
	}

	var extInfoPo *db.NFTGroupInfoExt
	var conditionsPo []*db.NFTGroupCondition
	if condition != nil {
		extInfoPo = &db.NFTGroupInfoExt{
			GroupId:       groupId,
			ConditionType: condition.Type,
		}

		conditionsPo = make([]*db.NFTGroupCondition, len(condition.NFT))
		for i, nft := range condition.NFT {
			conditionsPo[i] = &db.NFTGroupCondition{
				GroupId: groupId,
				NFTType: nft.Type,
				NFTId:   nft.Id,
				NFTName: nft.Name,
			}
		}
	}

	err = s.createNFTGroup(ctx, extInfoPo, conditionsPo, groupPo, membersPo)
	if err != nil {
		return 0, err
	}

	return groupId, nil
}
