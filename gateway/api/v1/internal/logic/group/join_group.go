package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/slg"
	"github.com/txchat/dtalk/pkg/util"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) JoinGroup(req *types.JoinGroupReq) (*types.JoinGroupResp, error) {
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
			//请求上链购接口判断 conditions.GetType() conditions.GetNft()
			ids := make([]string, len(conditions.GetNft()))
			for i, nft := range conditions.GetNft() {
				ids[i] = nft.GetId()
			}
			gps, err := l.svcCtx.SlgClient.LoadGroupPermission([]*slg.UserCondition{
				{
					UID:        req.InviterId,
					HandleType: conditions.GetType(),
					Conditions: ids,
				},
			})
			if err != nil {
				return nil, err
			}
			if !gps.IsPermission(req.InviterId) {
				return nil, xerror.NewError(xerror.PermissionDenied)
			}
		} else {
			return nil, xerror.NewError(xerror.PermissionDenied).SetExtMessage("group condition not find")
		}
	}

	_, err = l.svcCtx.GroupClient.InviteGroupMembers(l.ctx, &pb.InviteGroupMembersReq{
		GroupId:   groupId,
		InviterId: req.InviterId,
		MemberIds: []string{l.getOpe()},
	})
	if err != nil {
		return nil, err
	}

	return &types.JoinGroupResp{
		Id:    groupId,
		IdStr: util.ToString(groupId),
	}, nil
}
