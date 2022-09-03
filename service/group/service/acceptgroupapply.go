package service

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
	"github.com/txchat/dtalk/service/group/model/types"
)

// AcceptGroupApplySvc 接受 申请加入团队 的审批
func (s *Service) AcceptGroupApplySvc(ctx context.Context, req *types.AcceptGroupApplyReq) (res *types.AcceptGroupApplyResp, err error) {
	//log := s.GetLogWithTrace(ctx)

	personId := req.PersonId
	groupId := util.MustToInt64(req.Id)
	applyId := util.MustToInt64(req.ApplyId)

	// 查询审批详情
	groupApply, err := s.getGroupApplyById(applyId)
	if err != nil {
		return nil, err
	}
	// 判断审批是否被处理过
	if err = groupApply.IsWait(); err != nil {
		return nil, err
	}

	// 查询操作人详情并判断是否有权限操作
	person, err := s.GetPersonByMemberIdAndGroupId(ctx, personId, groupId)
	if err != nil {
		return nil, err
	}
	if err = person.IsAdmin(); err != nil {
		return nil, err
	}

	// 查询群详情
	group, err := s.GetGroupInfoByGroupId(ctx, groupId)
	if err != nil {
		return nil, err
	}
	// 判断群人数上限
	if err = group.TryJoin(1); err != nil {
		return nil, err
	}

	// 处理审批
	groupApply.ApplyStatus = biz.GroupApplyAccept
	groupApply.OperatorId = personId
	err = s.updateGroupApply(groupApply)
	if err != nil {
		return nil, err
	}

	// 判断是否已经在群内
	if isExit, err := s.CheckInGroup(groupApply.MemberId, groupId); err != nil {
		return nil, err
	} else if isExit == true {
		err = xerror.NewError(xerror.GroupInviteMemberExist)
		return nil, err
	}

	// 加人
	nowTime := s.getNowTime()
	newMember := &db.GroupMember{
		GroupId:               groupId,
		GroupMemberId:         groupApply.MemberId,
		GroupMemberName:       "",
		GroupMemberType:       0,
		GroupMemberJoinTime:   nowTime,
		GroupMemberUpdateTime: nowTime,
	}

	opeId := groupApply.InviterId
	if opeId == "" {
		opeId = groupApply.OperatorId
	}

	if err = s.AddGroupMembers(ctx, group.GroupId, []*db.GroupMember{newMember}, opeId); err != nil {
		return nil, err
	}

	return &types.AcceptGroupApplyResp{}, nil
}

func (s *Service) updateGroupApply(biz *biz.GroupApplyBiz) error {
	err := s.dao.UpdateGroupApply(biz.OperatorId, biz.ApplyStatus, biz.RejectReason, biz.ApplyId)
	if err != nil {
		return err
	}

	return nil
}
