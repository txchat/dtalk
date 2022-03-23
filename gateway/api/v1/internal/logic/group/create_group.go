package logic

import (
	"github.com/txchat/dtalk/gateway/api/v1/internal/types"
	pb "github.com/txchat/dtalk/service/group/api"
)

func (l *GroupLogic) CreateGroup(req *types.CreateGroupReq) (*types.CreateGroupResp, error) {
	name := req.Name
	addr := l.getOpe()
	owner := &pb.GroupMemberInfo{
		Id:   addr,
		Name: "",
	}
	members := make([]*pb.GroupMemberInfo, 0, len(req.MemberIds))
	for _, memberId := range req.MemberIds {
		members = append(members, &pb.GroupMemberInfo{
			Id:   memberId,
			Name: "",
		})
	}

	resp1, err := l.svcCtx.GroupClient.CreateGroup(l.ctx, &pb.CreateGroupReq{
		Name:      name,
		GroupType: pb.GroupType_GROUP_TYPE_NORMAL,
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

	return &types.CreateGroupResp{
		GroupInfo: Group,
		Members:   Members,
	}, nil
}
