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

type GetGroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewGetGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupListLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &GetGroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *GetGroupListLogic) GetGroupList(req *types.GetGroupListReq) (resp *types.GetGroupListResp, err error) {
	uid := l.custom.UID

	joinedResp, err := l.svcCtx.GroupRPC.JoinedGroups(l.ctx, &groupclient.JoinedGroupsReq{
		Uid: uid,
	})
	if err != nil {
		return
	}
	groups := make([]*types.GroupInfo, 0, len(joinedResp.GetGid()))
	for _, gid := range joinedResp.GetGid() {
		gInfoResp, err := l.svcCtx.GroupRPC.GroupInfo(l.ctx, &groupclient.GroupInfoReq{
			Gid: gid,
		})
		if err != nil {
			return nil, err
		}
		groupInfo := gInfoResp.GetGroup()
		if groupInfo == nil {
			return nil, xerror.ErrGroupPersonNotExist
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

		groups = append(groups, &types.GroupInfo{
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
		})
	}

	resp = &types.GetGroupListResp{Groups: groups}
	return
}
