package logic

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/txchat/dtalk/app/services/generator/generatorclient"
	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/app/services/group/internal/svc"
	xgroup "github.com/txchat/dtalk/internal/group"
	xerror "github.com/txchat/dtalk/pkg/error"
	xrand "github.com/txchat/dtalk/pkg/rand"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGroupLogic) CreateGroup(in *group.CreateGroupReq) (*group.CreateGroupResp, error) {
	if len(in.GetMembers()) > l.svcCtx.Config.Group.MaxMembers {
		return nil, xerror.ErrGroupMemberLimit
	}

	genIDResp, err := l.svcCtx.IDGenRPC.GetID(l.ctx, &generatorclient.GetIDReq{})
	if err != nil {
		return nil, err
	}
	gid := genIDResp.GetId()
	markId, err := l.randomMarkId()
	if err != nil {
		return nil, err
	}

	g := l.svcCtx.GroupManager.CreateNewGroup(xgroup.TypeOfGroup(in.GetType()), gid, in.GetName(), markId, in.GetOwner(), l.svcCtx.Config.Group.MaxMembers)

	for _, member := range in.Members {
		mid, ok := checkMemberID(member.GetId())
		if !ok {
			return nil, xerror.ErrInvalidParams
		}
		member.Id = mid
		g.Invite(mid, member.GetName())
	}

	groupInfo := &model.GroupInfo{
		GroupId:         g.Id(),
		GroupMarkId:     g.MarkId(),
		GroupName:       g.Name(),
		GroupAvatar:     g.Avatar(),
		GroupMemberNum:  int32(g.MemberCount()),
		GroupMaximum:    int32(g.MaxMembers()),
		GroupIntroduce:  "",
		GroupStatus:     model.GroupStatusServing,
		GroupOwnerId:    g.Owner(),
		GroupJoinType:   int32(g.JoinPermission()),
		GroupMuteType:   int32(g.MutePermission()),
		GroupFriendType: int32(g.FriendshipPermission()),
		GroupAESKey:     g.AesKey(),
		GroupPubName:    g.Name(),
		GroupType:       int32(in.GetType()),
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
	defer tx.RollBack()

	if _, _, err = l.svcCtx.Repo.InsertGroupInfo(tx, groupInfo); err != nil {
		return nil, err
	}

	if _, _, err = l.svcCtx.Repo.InsertGroupMembers(tx, members, g.CreateTime()); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	//signal and notice
	membersId := xgroup.Members(g.Members()).ToArray()
	err = l.svcCtx.RegisterGroup(l.ctx, gid, membersId)

	err = l.svcCtx.SignalHub.GroupAddNewMembers(l.ctx, gid, membersId)

	err = l.svcCtx.NoticeHub.GroupAddNewMembers(l.ctx, gid, g.Owner(), membersId)

	return &group.CreateGroupResp{
		Id:         g.Id(),
		CreateTime: g.CreateTime(),
	}, nil
}

func checkMemberID(mid string) (string, bool) {
	mid = strings.TrimSpace(mid)
	if mid == "" || len(mid) > 40 {
		return "", false
	}
	return mid, true
}

func (l *CreateGroupLogic) randomMarkId() (string, error) {
	for i := 0; i < 10; i++ {
		markId := xrand.NewNumber(8)
		if _, err := l.svcCtx.Repo.GetGroupByMarkId(markId); err != nil {
			if errors.Is(err, xerror.ErrNotFound) {
				return markId, nil
			}
			return "", err
		}
	}
	return "", xerror.ErrExec
}
