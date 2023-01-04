package group

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/groupclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type GetGroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewGetGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupInfoLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &GetGroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *GetGroupInfoLogic) GetGroupInfo(req *types.GetGroupInfoReq) (resp *types.GetGroupInfoResp, err error) {
	uid := l.custom.UID
	gid, err := util.ToInt64(req.IdStr)
	if err != nil {
		gid = req.Id
	}
	groupInfoResp, err := l.svcCtx.GroupRPC.GroupInfo(l.ctx, &groupclient.GroupInfoReq{
		Gid: gid,
	})
	if err != nil {
		return
	}
	groupInfo := groupInfoResp.GetGroup()
	if groupInfo == nil {
		err = xerror.ErrGroupNotExist
		return
	}

	var owner, person *types.GroupMember
	ownerResp, err := l.svcCtx.GroupRPC.MemberInfo(l.ctx, &groupclient.MemberInfoReq{
		Gid: gid,
		Uid: groupInfo.GetOwnerId(),
	})
	if err == nil {
		m := ownerResp.GetMember()
		owner = &types.GroupMember{
			MemberId:       m.GetUid(),
			MemberName:     m.GetNickname(),
			MemberType:     int32(m.GetRole()),
			MemberMuteTime: m.GetMutedTime(),
		}
	}

	personResp, err := l.svcCtx.GroupRPC.MemberInfo(l.ctx, &groupclient.MemberInfoReq{
		Gid: gid,
		Uid: uid,
	})
	if err == nil {
		m := personResp.GetMember()
		person = &types.GroupMember{
			MemberId:       m.GetUid(),
			MemberName:     m.GetNickname(),
			MemberType:     int32(m.GetRole()),
			MemberMuteTime: m.GetMutedTime(),
		}
	}

	membersResp, err := l.svcCtx.GroupRPC.GroupLimitedMembers(l.ctx, &groupclient.GroupLimitedMembersReq{
		Gid: gid,
		Num: 10,
	})
	if err != nil {
		return
	}

	members := make([]*types.GroupMember, 0, len(membersResp.GetMembers()))
	for _, m := range membersResp.GetMembers() {
		members = append(members, &types.GroupMember{
			MemberId:       m.GetUid(),
			MemberName:     m.GetNickname(),
			MemberType:     int32(m.GetRole()),
			MemberMuteTime: m.GetMutedTime(),
		})
	}

	resp = &types.GetGroupInfoResp{
		GroupInfo: types.GroupInfo{
			Id:         groupInfo.GetId(),
			IdStr:      util.MustToString(groupInfo.GetId()),
			MarkId:     groupInfo.GetMarkId(),
			Name:       groupInfo.GetName(),
			PublicName: groupInfo.GetMaskName(),
			Avatar:     groupInfo.GetAvatar(),
			Introduce:  groupInfo.GetIntroduce(),
			Owner:      owner,
			Person:     person,
			MemberNum:  groupInfo.GetMemberCount(),
			Maximum:    groupInfo.GetMaxMembersLimit(),
			Status:     int32(groupInfo.GetStatus()),
			CreateTime: groupInfo.GetCreateTime(),
			JoinType:   int32(groupInfo.GetJoinType()),
			MuteType:   int32(groupInfo.GetMuteType()),
			FriendType: int32(groupInfo.GetFriendType()),
			MuteNum:    groupInfo.GetMutedNumbers(),
			AdminNum:   groupInfo.GetManagerNumbers(),
			AESKey:     groupInfo.GetAESKey(),
			GroupType:  int32(groupInfo.GetType()),
		},
		Members: members,
	}
	return
}
