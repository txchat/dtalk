package service

import (
	"context"

	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/types"
)

func (s *Service) RejectGroupApplySvc(ctx context.Context, req *types.RejectGroupApplyReq) (res *types.RejectGroupApplyResp, err error) {
	personId := req.PersonId
	groupId := util.ToInt64(req.Id)
	applyId := util.ToInt64(req.ApplyId)

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

	// 处理审批
	groupApply.ApplyStatus = biz.GroupApplyReject
	groupApply.OperatorId = personId
	err = s.updateGroupApply(groupApply)
	if err != nil {
		return nil, err
	}

	return &types.RejectGroupApplyResp{}, nil
}
