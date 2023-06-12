package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/internal/model"
	xgroup "github.com/txchat/dtalk/internal/group"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInviteMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteMembersLogic {
	return &InviteMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InviteMembersLogic) InviteMembers(in *group.InviteMembersReq) (*group.InviteMembersResp, error) {
	timeTs := util.TimeNowUnixMilli()
	gInfo, err := l.svcCtx.Repo.GetGroupById(in.GetGid())
	if err != nil {
		return nil, err
	}

	g := xgroup.NewGroup(gInfo.GroupId, gInfo.GroupOwnerId, int(gInfo.GroupMaximum), int(gInfo.GroupMemberNum), gInfo.GroupCreateTime)

	for _, mid := range in.GetMid() {
		err = g.Invite(mid, "")
		if err != nil {
			return nil, err
		}
	}

	members := make([]*model.GroupMember, 0, g.MemberCount())
	for _, member := range g.Members() {
		members = append(members, &model.GroupMember{
			GroupId:               g.Id(),
			GroupMemberId:         member.Id(),
			GroupMemberName:       member.Nickname(),
			GroupMemberType:       int32(member.Role()),
			GroupMemberJoinTime:   g.CreateTime(),
			GroupMemberUpdateTime: g.CreateTime(),
		})
	}
	tx, err := l.svcCtx.Repo.NewTx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if _, _, err = l.svcCtx.Repo.InsertGroupMembers(tx, members, timeTs); err != nil {
		return nil, err
	}

	if _, _, err = l.svcCtx.Repo.UpdateGroupMembersNumber(tx, &model.GroupInfo{
		GroupId:         g.Id(),
		GroupMemberNum:  int32(g.MemberCount()),
		GroupUpdateTime: timeTs,
	}); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	//signal and notice
	membersId := xgroup.Members(g.Members()).ToArray()

	err = l.svcCtx.RegisterGroupMembers(l.ctx, g.Id(), membersId)

	err = l.svcCtx.SignalHub.GroupAddNewMembers(l.ctx, g.Id(), membersId)

	err = l.svcCtx.NoticeHub.GroupAddNewMembers(l.ctx, g.Id(), in.GetOperator(), membersId)

	return &group.InviteMembersResp{
		Number: int32(g.MemberCount()),
	}, nil
}
