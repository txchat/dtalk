package logic

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MembersInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMembersInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MembersInfoLogic {
	return &MembersInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MembersInfoLogic) MembersInfo(in *group.MembersInfoReq) (*group.MembersInfoResp, error) {
	gid := in.GetGid()
	members := make([]*group.MemberInfo, 0, len(in.GetUid()))
	for _, uid := range in.GetUid() {
		m, err := l.svcCtx.Repo.GetMemberById(gid, uid)
		if err != nil {
			if err == xerror.ErrGroupMemberNotExist {
				continue
			}
			return nil, err
		}
		members = append(members, &group.MemberInfo{
			Gid:        gid,
			Uid:        uid,
			Nickname:   m.GroupMemberName,
			Role:       group.RoleType(m.GroupMemberType),
			MutedTime:  m.GroupMemberMuteTime,
			JoinedTime: m.GroupMemberJoinTime,
		})
	}
	return &group.MembersInfoResp{
		Members: members,
	}, nil
}
