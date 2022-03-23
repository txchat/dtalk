package service

import (
	"context"

	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/types"
	"github.com/txchat/dtalk/pkg/util"
)

func (s *Service) GetGroupApplysSvc(ctx context.Context, req *types.GetGroupApplysReq) (res *types.GetGroupApplysResp, err error) {
	//personId := req.PersonId
	groupId := util.ToInt64(req.Id)

	groupApplyBizs, err := s.getGroupApplys(groupId, req.Count, req.Offset)
	if err != nil {
		return nil, err
	}

	res = &types.GetGroupApplysResp{}
	res.GroupApplys = make([]*types.GroupApplyInfo, 0)
	for _, groupApplyBiz := range groupApplyBizs {
		res.GroupApplys = append(res.GroupApplys, groupApplyBiz.ToTypes())
	}

	return res, nil
}

func (s *Service) getGroupApplys(groupId int64, limit, offset int32) ([]*biz.GroupApplyBiz, error) {
	groupApplys, err := s.dao.GetGroupApplys(groupId, limit, offset)
	if err != nil {
		return nil, err
	}
	groupApplyBizs := make([]*biz.GroupApplyBiz, 0)
	for _, groupApply := range groupApplys {
		groupApplyBizs = append(groupApplyBizs, groupApply.ToBiz())
	}
	return groupApplyBizs, nil
}
