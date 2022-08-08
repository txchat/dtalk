package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	xerror "github.com/txchat/dtalk/pkg/error"
	pb "github.com/txchat/dtalk/service/group/api"
	vip "github.com/txchat/dtalk/service/vip/api"
)

func (l *GroupLogic) CreateNFTGroup(req *types.CreateNFTGroupReq) (*types.CreateNFTGroupResp, error) {
	name := req.Name
	addr := l.getOpe()
	owner := &pb.GroupMemberInfo{
		Id:   addr,
		Name: "",
	}

	//检查是否有创群资格
	reply, err := l.svcCtx.VIPClient.GetVIP(l.ctx, &vip.GetVIPReq{
		Uid: addr,
	})
	if err != nil {
		return nil, err
	}
	if reply.GetVip() == nil || reply.GetVip().GetUid() != addr {
		return nil, xerror.NewError(xerror.PermissionDenied)
	}

	members := make([]*pb.GroupMemberInfo, 0, len(req.MemberIds))
	for _, memberId := range req.MemberIds {
		members = append(members, &pb.GroupMemberInfo{
			Id:   memberId,
			Name: "",
		})
	}

	resp1, err := l.svcCtx.GroupClient.CreateNFTGroup(l.ctx, &pb.CreateNFTGroupReq{
		Name:      name,
		GroupType: pb.GroupType_GROUP_TYPE_NFT,
		Owner:     owner,
		Members:   members,
	})
	if err != nil {
		return nil, err
	}

	groupId := resp1.GroupId

	resp2, err := l.svcCtx.GroupClient.GetPriGroupInfo(l.ctx, &pb.GetPriGroupInfoReq{
		GroupId:    groupId,
		PersonId:   addr,
		DisplayNum: int32(1 + len(members)),
	})
	if err != nil {
		return nil, err
	}

	Group := NewTypesGroupInfo(resp2.Group)
	Members := NewTypesGroupMemberInfos(resp2.Group.Members)

	return &types.CreateNFTGroupResp{
		GroupInfo: Group,
		Members:   Members,
	}, nil
}
