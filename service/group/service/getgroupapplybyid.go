package service

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/types"
)

func (s *Service) GetGroupApplyByIdSvc(ctx context.Context, req *types.GetGroupApplyByIdReq) (res *types.GetGroupApplysResp, err error) {
	//personId := req.PersonId
	applyId := util.ToInt64(req.ApplyId)

	groupApplyBiz, err := s.getGroupApplyById(applyId)
	if err != nil {
		return nil, err
	}

	return &types.GetGroupApplysResp{
		GroupApplys: []*types.GroupApplyInfo{groupApplyBiz.ToTypes()},
	}, nil
}

func (s *Service) getGroupApplyById(id int64) (*biz.GroupApplyBiz, error) {
	groupApply, err := s.dao.GetGroupApplyById(id)
	if err != nil {
		return nil, err
	}

	if groupApply == nil {
		return nil, xerror.NewError(xerror.GroupApplyNotExist)
	}

	return groupApply.ToBiz(), nil
}
