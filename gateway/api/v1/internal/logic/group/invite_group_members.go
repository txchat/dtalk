package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/slg"
	"github.com/txchat/dtalk/pkg/util"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) InviteGroupMembers(req *types.InviteGroupMembersReq) (*types.InviteGroupMembersResp, error) {
	groupId := req.Id
	invitedMembers := req.NewMemberIds
	//根据groupId查询基本群信息
	groupPubInfo, err := l.svcCtx.GroupClient.GetPubGroupInfo(l.ctx, &pb.GetPubGroupInfoReq{
		GroupId:  groupId,
		PersonId: l.getOpe(),
	})
	if err != nil {
		return nil, err
	}
	if groupPubInfo.GetGroup().GetType() == pb.GroupType_GROUP_TYPE_NFT {
		//如果是藏品群：1. 获取入群条件
		extInfo, err := l.svcCtx.GroupClient.GetNFTGroupExtInfo(l.ctx, &pb.GetNFTGroupExtInfoReq{
			GroupId: groupId,
		})
		if err != nil {
			return nil, err
		}
		//如果是藏品群：2. 查询用户是否有入群资格
		if conditionsRequest := extInfo.GetCondition(); conditionsRequest != nil {
			invitedMembers, err = l.nftInviteMembersFilter(conditionsRequest, req.NewMemberIds)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, xerror.NewError(xerror.PermissionDenied).SetExtMessage("group condition not find")
		}
	}

	_, err = l.svcCtx.GroupClient.InviteGroupMembers(l.ctx, &pb.InviteGroupMembersReq{
		GroupId:   groupId,
		InviterId: l.getOpe(),
		MemberIds: invitedMembers,
	})
	if err != nil {
		return nil, err
	}

	return &types.InviteGroupMembersResp{
		Id:    groupId,
		IdStr: util.ToString(groupId),
	}, nil
}

func (l *GroupLogic) nftInviteMembersFilter(condition *pb.Condition, members []string) ([]string, error) {
	//请求上链购接口判断 conditionsRequest.GetType() conditionsRequest.GetNft()
	ids := make([]string, len(condition.GetNft()))
	for i, nft := range condition.GetNft() {
		ids[i] = nft.GetId()
	}
	conditions := make([]*slg.UserCondition, len(members))
	for i, tarId := range members {
		item := &slg.UserCondition{
			UID:        tarId,
			HandleType: condition.GetType(),
			Conditions: ids,
		}
		conditions[i] = item
	}
	gps, err := l.svcCtx.SlgClient.LoadGroupPermission(conditions)
	if err != nil {
		return nil, err
	}
	filteredMembers := make([]string, 0)
	for _, memberId := range members {
		if gps.IsPermission(memberId) {
			filteredMembers = append(filteredMembers, memberId)
		}
	}
	return filteredMembers, nil
}
