package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberExitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberExitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberExitLogic {
	return &MemberExitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberExitLogic) MemberExit(in *group.MemberExitReq) (*group.MemberExitResp, error) {
	nowTs := util.TimeNowUnixMilli()

	gInfo, err := l.svcCtx.Repo.GetGroupById(in.GetGid())
	if err != nil {
		return nil, err
	}

	tx, err := l.svcCtx.Repo.NewTx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, _, err = l.svcCtx.Repo.UpdateGroupMemberRole(tx, &model.GroupMember{
		GroupId:               in.GetGid(),
		GroupMemberId:         in.GetOperator(),
		GroupMemberUpdateTime: nowTs,
		GroupMemberType:       model.GroupMemberTypeOther,
	})

	if _, _, err = l.svcCtx.Repo.UpdateGroupMembersNumber(tx, &model.GroupInfo{
		GroupId:         in.GetGid(),
		GroupMemberNum:  gInfo.GroupMemberNum - 1,
		GroupUpdateTime: nowTs,
	}); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	err = l.svcCtx.NoticeHub.GroupSignOut(l.ctx, in.GetGid(), in.GetOperator())
	err = l.svcCtx.SignalHub.GroupRemoveMembers(l.ctx, in.GetGid(), []string{in.GetOperator()})
	err = l.svcCtx.UnRegisterGroupMembers(l.ctx, in.GetGid(), []string{in.GetOperator()})
	return &group.MemberExitResp{
		Number: gInfo.GroupMemberNum - 1,
	}, nil
}
