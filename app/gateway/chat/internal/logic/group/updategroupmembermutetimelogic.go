package group

import (
	"context"

	"github.com/txchat/dtalk/app/gateway/chat/internal/model"
	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/groupclient"
	xerror "github.com/txchat/dtalk/pkg/error"
	xhttp "github.com/txchat/dtalk/pkg/net/http"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/gateway/chat/internal/svc"
	"github.com/txchat/dtalk/app/gateway/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	//xhttp "github.com/txchat/dtalk/pkg/net/http"
)

type UpdateGroupMemberMuteTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	custom *xhttp.Custom
}

func NewUpdateGroupMemberMuteTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupMemberMuteTimeLogic {
	c, ok := xhttp.FromContext(ctx)
	if !ok {
		c = &xhttp.Custom{}
	}
	return &UpdateGroupMemberMuteTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		custom: c,
	}
}

func (l *UpdateGroupMemberMuteTimeLogic) UpdateGroupMemberMuteTime(req *types.UpdateGroupMemberMuteTimeReq) (resp *types.UpdateGroupMemberMuteTimeResp, err error) {
	uid := l.custom.UID
	gid, err := util.ToInt64(req.IdStr)
	if err != nil {
		gid = req.Id
	}
	nowTs := util.TimeNowUnixMilli()
	var deadline int64 = 0

	if req.MuteTime > 0 {
		deadline = nowTs + req.MuteTime
	}

	operatorResp, err := l.svcCtx.GroupRPC.MemberInfo(l.ctx, &groupclient.MemberInfoReq{
		Gid: gid,
		Uid: uid,
	})
	if err != nil {
		return nil, err
	}
	operator := operatorResp.GetMember()
	if operator.GetRole() < group.RoleType_Manager {
		return nil, xerror.ErrGroupHigherPermission
	}

	if req.MuteTime != model.MuteForever && req.MuteTime > (24*60*60*1000) {
		return nil, xerror.ErrInvalidParams
	}

	membersResp, err := l.svcCtx.GroupRPC.MembersInfo(l.ctx, &groupclient.MembersInfoReq{
		Gid: gid,
		Uid: req.MemberIds,
	})

	members := make([]string, 0)
	groupMemberReply := make([]*types.GroupMember, 0)
	for _, member := range membersResp.GetMembers() {
		if member.GetRole() > group.RoleType_NormalMember {
			continue
		}
		members = append(members, member.GetUid())
		groupMemberReply = append(groupMemberReply, &types.GroupMember{
			MemberId:       member.GetUid(),
			MemberName:     member.GetNickname(),
			MemberType:     int32(member.GetRole()),
			MemberMuteTime: deadline,
		})
	}

	if deadline > 0 {
		_, err = l.svcCtx.GroupRPC.MuteMembers(l.ctx, &groupclient.MuteMembersReq{
			Gid:      gid,
			Operator: uid,
			Mid:      members,
			Deadline: nowTs + req.MuteTime,
		})
		if err != nil {
			return nil, err
		}
	} else {
		_, err = l.svcCtx.GroupRPC.UnMuteMembers(l.ctx, &groupclient.UnMuteMembersReq{
			Gid:      gid,
			Operator: uid,
			Mid:      members,
		})
		if err != nil {
			return nil, err
		}
	}

	resp = &types.UpdateGroupMemberMuteTimeResp{
		Members: groupMemberReply,
	}
	return
}
