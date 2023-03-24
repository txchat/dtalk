package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/app/services/group/internal/svc"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type DisbandGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDisbandGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DisbandGroupLogic {
	return &DisbandGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DisbandGroupLogic) DisbandGroup(in *group.DisbandGroupReq) (*group.DisbandGroupResp, error) {
	gid := in.GetGid()
	nowTS := util.TimeNowUnixMilli()

	members, err := l.svcCtx.Repo.GetUnLimitedMembers(gid)
	if err != nil {
		return nil, err
	}

	tx, err := l.svcCtx.Repo.NewTx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, _, err = l.svcCtx.Repo.UpdateGroupStatus(tx, &model.GroupInfo{
		GroupId:         gid,
		GroupStatus:     model.GroupStatusDisBand,
		GroupUpdateTime: nowTS,
	})
	if err != nil {
		return nil, err
	}

	for _, member := range members {
		_, _, err = l.svcCtx.Repo.UpdateGroupMemberRole(tx, &model.GroupMember{
			GroupId:               gid,
			GroupMemberId:         member.GroupMemberId,
			GroupMemberType:       model.GroupMemberTypeOther,
			GroupMemberUpdateTime: nowTS,
		})
		if err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	//signal and notice
	err = l.svcCtx.NoticeHub.GroupDeleted(l.ctx, gid, in.GetOperator())

	err = l.svcCtx.SignalHub.GroupDeleted(l.ctx, gid)

	err = l.svcCtx.UnRegisterGroup(l.ctx, gid)

	return &group.DisbandGroupResp{}, nil
}
