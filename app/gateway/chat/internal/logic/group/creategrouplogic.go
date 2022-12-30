package group

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"
	pb "github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/groupclient"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.CreateGroupReq) (resp *types.CreateGroupResp, err error) {
	owner := l.custom.UID

	//member check, default member name is empty
	members := make([]*pb.CreateGroupReq_MemberMinData, 0, len(req.MemberIds)+1)
	membersId := make([]string, 0, len(req.MemberIds)+1)
	members = append(members, &pb.CreateGroupReq_MemberMinData{
		Id: owner,
	})
	membersId = append(membersId, owner)
	for _, memberId := range req.MemberIds {
		members = append(members, &pb.CreateGroupReq_MemberMinData{
			Id: memberId,
		})
		membersId = append(membersId, memberId)
	}

	createResp, err := l.svcCtx.GroupRPC.CreateGroup(l.ctx, &groupclient.CreateGroupReq{
		Name:    req.Name,
		Type:    pb.GroupType_Normal,
		Owner:   owner,
		Members: members,
	})
	if err != nil {
		return nil, err
	}

	//get group info
	groupResp, err := l.svcCtx.GroupRPC.GroupInfo(l.ctx, &groupclient.GroupInfoReq{
		Gid: createResp.GetId(),
	})
	if err != nil {
		return nil, err
	}

	membersResp, err := l.svcCtx.GroupRPC.MemberInfo(l.ctx, &groupclient.MemberInfoReq{
		Gid: groupResp.GetGroup().GetId(),
		Uid: membersId,
	})
	if err != nil {
		return nil, err
	}

	var ownerInfo *types.GroupMember
	membersInfo := make([]*types.GroupMember, 0, len(membersResp.GetMembers()))
	for _, member := range membersResp.GetMembers() {
		m := &types.GroupMember{
			MemberId:       member.GetUid(),
			MemberName:     member.GetNickname(),
			MemberType:     int32(member.GetRole()),
			MemberMuteTime: member.GetMutedTime(),
		}
		if member.GetUid() == owner {
			ownerInfo = m
		}
		membersInfo = append(membersInfo, m)
	}

	resp = &types.CreateGroupResp{
		GroupInfo: types.GroupInfo{
			Id:         groupResp.GetGroup().GetId(),
			IdStr:      util.MustToString(groupResp.GetGroup().GetId()),
			MarkId:     groupResp.GetGroup().GetMarkId(),
			Name:       groupResp.GetGroup().GetName(),
			PublicName: groupResp.GetGroup().GetMaskName(),
			Avatar:     groupResp.GetGroup().GetAvatar(),
			Introduce:  groupResp.GetGroup().GetIntroduce(),
			Owner:      ownerInfo,
			Person:     ownerInfo,
			MemberNum:  groupResp.GetGroup().GetMemberCount(),
			Maximum:    groupResp.GetGroup().GetMaxMembersLimit(),
			Status:     int32(groupResp.GetGroup().GetStatus()),
			CreateTime: groupResp.GetGroup().GetCreateTime(),
			JoinType:   int32(groupResp.GetGroup().GetJoinType()),
			MuteType:   int32(groupResp.GetGroup().GetMuteType()),
			FriendType: int32(groupResp.GetGroup().GetFriendType()),
			MuteNum:    groupResp.GetGroup().GetMutedNumbers(),
			AdminNum:   groupResp.GetGroup().GetManagerNumbers(),
			AESKey:     groupResp.GetGroup().GetAESKey(),
			GroupType:  int32(groupResp.GetGroup().GetType()),
		},
		Members: membersInfo,
	}
	return
}
