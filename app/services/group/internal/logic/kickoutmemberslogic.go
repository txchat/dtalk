package logic

import (
	"context"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/app/services/group/internal/svc"
	"github.com/txchat/dtalk/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type KickOutMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKickOutMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KickOutMembersLogic {
	return &KickOutMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KickOutMembersLogic) KickOutMembers(in *group.KickOutMembersReq) (*group.KickOutMembersResp, error) {
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

	for _, mid := range in.GetMid() {
		_, _, err = l.svcCtx.Repo.UpdateGroupMemberRole(tx, &model.GroupMember{
			GroupId:               in.GetGid(),
			GroupMemberId:         mid,
			GroupMemberUpdateTime: nowTs,
			GroupMemberType:       model.GroupMemberTypeOther,
		})
	}

	if _, _, err = l.svcCtx.Repo.UpdateGroupMembersNumber(tx, &model.GroupInfo{
		GroupId:         in.GetGid(),
		GroupMemberNum:  gInfo.GroupMemberNum - int32(len(in.GetMid())),
		GroupUpdateTime: nowTs,
	}); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	err = l.svcCtx.NoticeHub.GroupKickOutMembers(l.ctx, in.GetGid(), in.GetOperator(), in.GetMid())
	err = l.svcCtx.SignalHub.GroupRemoveMembers(l.ctx, in.GetGid(), in.GetMid())
	err = l.svcCtx.UnRegisterGroupMembers(l.ctx, in.GetGid(), in.GetMid())
	return &group.KickOutMembersResp{
		Number: gInfo.GroupMemberNum - int32(len(in.GetMid())),
	}, nil
}
