package service

import (
	"context"
	"errors"

	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/types"
)

// GetGroupInfoByConditionSvc 通过搜索得到群列表
func (s *Service) GetGroupInfoByConditionSvc(ctx context.Context, req *types.GetGroupInfoByConditionReq) (res *types.GetGroupInfoByConditionResp, err error) {
	//personId := req.PersonId

	groups := make([]*biz.GroupInfo, 0)

	switch req.Tp {
	case 0:
		groupInfo, err := s.dao.GetGroupInfoByGroupMarkId(req.Query)
		if err != nil && !errors.Is(err, model.ErrRecordNotExist) {
			return nil, err
		}
		if groupInfo != nil {
			if err := groupInfo.IsNormal(); err != nil {
				return nil, err
			}
			groups = append(groups, groupInfo)
		}
	case 1:
		groupId := util.ToInt64(req.Query)
		groupInfo, err := s.dao.GetGroupInfoByGroupId(ctx, groupId)
		if err != nil && !errors.Is(err, model.ErrRecordNotExist) {
			return nil, err
		}
		if groupInfo != nil {
			if err := groupInfo.IsNormal(); err != nil {
				return nil, err
			}
			groups = append(groups, groupInfo)
		}
	}

	groupBizInfos := make([]*types.GroupInfo, 0)
	for _, group := range groups {
		ownerInfo, err := s.GetMemberByMemberIdAndGroupId(ctx, group.GroupOwnerId, group.GroupId)
		if err != nil {
			return nil, err
		}

		groupBizInfo := group.ToTypes(ownerInfo.ToTypes(), nil)

		groupBizInfos = append(groupBizInfos, groupBizInfo)
	}

	return &types.GetGroupInfoByConditionResp{
		Groups: groupBizInfos,
	}, nil
}
