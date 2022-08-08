package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	"github.com/txchat/dtalk/pkg/util"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) InviteGroupMembers(req *types.InviteGroupMembersReq) (*types.InviteGroupMembersResp, error) {
	groupId := req.Id

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
		if conditions := extInfo.GetCondition(); conditions != nil {
			//TODO 请求上链购接口判断 conditions.GetType() conditions.GetNft()
			var allowed = true
			if !allowed {

			}
		}
	}

	_, err = l.svcCtx.GroupClient.InviteGroupMembers(l.ctx, &pb.InviteGroupMembersReq{
		GroupId:   groupId,
		InviterId: l.getOpe(),
		MemberIds: req.NewMemberIds,
	})
	if err != nil {
		return nil, err
	}

	return &types.InviteGroupMembersResp{
		Id:    groupId,
		IdStr: util.ToString(groupId),
	}, nil
}
