package service

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/group/model/types"
)

// GetGroupPubInfoSvc 查询群资料
func (s *Service) GetGroupPubInfoSvc(ctx context.Context, req *types.GetGroupPubInfoRequest) (res *types.GetGroupPubInfoResponse, err error) {
	groupId := req.Id
	personId := req.PersonId

	group, err := s.GetGroupInfoByGroupId(ctx, groupId)
	if err != nil {
		return nil, err
	}

	personInfo, err := s.GetPersonByMemberIdAndGroupId(ctx, personId, group.GroupId)
	if err != nil {
		if xerror.NewError(xerror.GroupPersonNotExist).Error() != err.Error() {
			return nil, err
		}
	}

	ownerInfo, err := s.GetMemberByMemberIdAndGroupId(ctx, group.GroupOwnerId, group.GroupId)
	if err != nil {
		return nil, err
	}

	res = &types.GetGroupPubInfoResponse{}
	var personType *types.GroupMember
	if personInfo != nil {
		personType = personInfo.ToTypes()
	}
	res.GroupInfo = group.ToTypes(ownerInfo.ToTypes(), personType)
	if personInfo == nil {
		res.GroupInfo.AESKey = ""
	}

	return res, nil
}
