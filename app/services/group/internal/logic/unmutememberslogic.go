package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnMuteMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnMuteMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnMuteMembersLogic {
	return &UnMuteMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnMuteMembersLogic) UnMuteMembers(in *group.UnMuteMembersReq) (*group.UnMuteMembersResp, error) {
	nowTS := util.TimeNowUnixMilli()
	gid := in.GetGid()

	members := make([]*model.GroupMember, 0, len(in.GetMid()))
	for _, mid := range in.GetMid() {
		members = append(members, &model.GroupMember{
			GroupId:       gid,
			GroupMemberId: mid,
			GroupMemberMute: model.GroupMemberMute{
				GroupMemberMuteTime:       0,
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
	err = l.svcCtx.SignalHub.UpdateMembersMuteTime(l.ctx, gid, 0, in.GetMid())
	return &group.UnMuteMembersResp{}, nil
}
