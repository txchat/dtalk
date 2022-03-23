package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/txchat/dtalk/pkg/contextx"
	xerror "github.com/txchat/dtalk/pkg/error"
	xrand "github.com/txchat/dtalk/pkg/rand"
	"github.com/txchat/dtalk/service/group/model"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
	"github.com/txchat/dtalk/service/group/model/types"
	"github.com/txchat/dtalk/pkg/util"
)

//CreateGroupSvc 创建群, 返回群信息
func (s *Service) CreateGroupSvc(ctx context.Context, req *types.CreateGroupRequest) (res *types.CreateGroupResponse, err error) {
	groupId, err := s.getLogId(context.Background())
	if err != nil {
		return nil, err
	}

	// 判断群人数
	if err := s.CheckGroupMemberNum(util.ToInt32(len(req.Members))+1, biz.GroupMaximum); err != nil {
		return nil, err
	}

	// 创建群
	nowTime := s.getNowTime()

	groupMembers := make([]*db.GroupMember, 0, len(req.MemberIds)+1)
	groupMembers = append(groupMembers, &db.GroupMember{
		GroupId:               groupId,
		GroupMemberId:         req.Owner.MemberId,
		GroupMemberName:       req.Owner.MemberName,
		GroupMemberType:       req.Owner.MemberType,
		GroupMemberJoinTime:   nowTime,
		GroupMemberUpdateTime: nowTime,
	})
	for i := 0; i < len(req.MemberIds); i++ {
		member := req.Members[i]
		if member.MemberId == req.Owner.MemberId {
			continue
		}
		groupMembers = append(groupMembers, &db.GroupMember{
			GroupId:               groupId,
			GroupMemberId:         member.MemberId,
			GroupMemberName:       member.MemberName,
			GroupMemberType:       member.MemberType,
			GroupMemberJoinTime:   nowTime,
			GroupMemberUpdateTime: nowTime,
		})
	}

	groupMarkId, err := s.getRandomGroupMarkId()
	if err != nil {
		return nil, err
	}
	group := &db.GroupInfo{
		GroupId:         groupId,
		GroupMarkId:     groupMarkId,
		GroupName:       req.Name,
		GroupAvatar:     req.Avatar,
		GroupMemberNum:  util.ToInt32(len(groupMembers)),
		GroupMaximum:    biz.GroupMaximum,
		GroupIntroduce:  req.Introduce,
		GroupStatus:     biz.GroupStatusNormal,
		GroupOwnerId:    req.Owner.MemberId,
		GroupJoinType:   biz.GroupJoinTypeAny,
		GroupMuteType:   biz.GroupMuteTypeAny,
		GroupFriendType: biz.GroupFriendTypeAllow,
		GroupCreateTime: nowTime,
		GroupUpdateTime: nowTime,
		GroupAESKey:     xrand.NewAESKey256(),
		GroupPubName:    req.Name,
		GroupType:       biz.GroupTypeNormal,
	}

	err = s.createGroup(ctx, group, groupMembers)
	if err != nil {
		return nil, err
	}

	res = &types.CreateGroupResponse{
		GroupInfo: &types.GroupInfo{
			Id:         groupId,
			IdStr:      util.ToString(groupId),
			MarkId:     groupMarkId,
			Name:       req.Name,
			Avatar:     req.Avatar,
			Introduce:  req.Introduce,
			Owner:      &req.Owner,
			MemberNum:  util.ToInt32(1 + len(req.Members)),
			Maximum:    biz.GroupMaximum,
			Status:     biz.GroupStatusNormal,
			CreateTime: nowTime,
			JoinType:   biz.GroupJoinTypeAny,
			MuteType:   biz.GroupMuteTypeAny,
			FriendType: biz.GroupFriendTypeAllow,
			MuteNum:    0,
			PublicName: req.Name,
			AESKey:     group.GroupAESKey,
			GroupType:  biz.GroupTypeNormal,
		},
		Members: req.Members,
	}
	return res, nil
}

// getRandomGroupMarkId 得到不重复的 8 位数字群组 id
func (s *Service) getRandomGroupMarkId() (string, error) {
	for {
		groupMarkId := xrand.NewNumber(8)
		if _, err := s.dao.GetGroupInfoByGroupMarkId(groupMarkId); err != nil {
			if errors.Is(err, model.ErrRecordNotExist) {
				return groupMarkId, nil
			}
			return "", err
		}
	}
}

// getLogId 由 generator 服务生成唯一 id
func (s *Service) getLogId(ctx context.Context) (id int64, err error) {

	reply, err := s.idGenRPCClient.GetID()
	if err != nil {
		return 0, err
	}

	return reply, nil
}

// execGroupCreate 执行创建群操作
func (s *Service) execGroupCreate(groupInfo *db.GroupInfo, members []*db.GroupMember) error {
	tx, err := s.dao.NewTx()
	if err != nil {
		return err
	}
	defer tx.RollBack()
	if _, _, err = s.dao.InsertGroupInfo(tx, groupInfo); err != nil {
		return err
	}

	if err = s.InsertGroupMembers(tx, members); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *Service) createGroup(ctx context.Context, group *db.GroupInfo, groupMembers []*db.GroupMember) error {
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

	if err = s.execGroupCreate(group, groupMembers); err != nil {
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

func (s *Service) CreateGroup(ctx context.Context, group *biz.GroupInfo, owner *biz.GroupMember, members []*biz.GroupMember) (int64, error) {
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

	err = s.createGroup(ctx, groupPo, membersPo)
	if err != nil {
		return 0, err
	}

	return groupId, nil
}
