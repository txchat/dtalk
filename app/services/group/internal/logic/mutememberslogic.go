package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/app/services/group/internal/svc"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type MuteMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMuteMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MuteMembersLogic {
	return &MuteMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MuteMembersLogic) MuteMembers(in *group.MuteMembersReq) (*group.MuteMembersResp, error) {
	nowTS := util.TimeNowUnixMilli()
	gid := in.GetGid()

	members := make([]*model.GroupMember, 0, len(in.GetMid()))
	for _, mid := range in.GetMid() {
		members = append(members, &model.GroupMember{
			GroupId:       gid,
			GroupMemberId: mid,
			GroupMemberMute: model.GroupMemberMute{
				GroupMemberMuteTime:       in.GetDeadline(),
				GroupMemberMuteUpdateTime: nowTS,
			},
		})
	}

	tx, err := l.svcCtx.Repo.NewTx()
	if err != nil {
		return nil, err
	}
	defer tx.RollBack()
	_, _, err = l.svcCtx.Repo.UpdateGroupMembersMuteTime(tx, members)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	// signal and notice
	err = l.svcCtx.SignalHub.UpdateMembersMuteTime(l.ctx, gid, in.GetDeadline(), in.GetMid())
	err = l.svcCtx.NoticeHub.UpdateMembersMuteTime(l.ctx, in.GetOperator(), gid, in.GetDeadline(), in.GetMid())
	return &group.MuteMembersResp{}, nil
}
