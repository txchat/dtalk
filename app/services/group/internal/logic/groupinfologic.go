package logic

import (
	"context"

	"github.com/txchat/dtalk/pkg/util"

	"github.com/txchat/dtalk/app/services/group/group"
	"github.com/txchat/dtalk/app/services/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupInfoLogic {
	return &GroupInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupInfoLogic) GroupInfo(in *group.GroupInfoReq) (*group.GroupInfoResp, error) {
	gInfo, err := l.svcCtx.Repo.GetGroupById(in.GetGid())
	if err != nil {
		return nil, err
	}
	mutedNumbers, err := l.svcCtx.Repo.GetGroupMutedNumbers(in.GetGid(), util.TimeNowUnixMilli())
	if err != nil {
		return nil, err
	}

	managerNumbers, err := l.svcCtx.Repo.GetGroupManagerNumbers(in.GetGid())
	if err != nil {
		return nil, err
	}

	return &group.GroupInfoResp{
		Group: &group.GroupInfo{
			Id:              gInfo.GroupId,
			MarkId:          gInfo.GroupMarkId,
			Name:            gInfo.GroupName,
			Avatar:          gInfo.GroupAvatar,
			MemberCount:     gInfo.GroupMemberNum,
			MaxMembersLimit: gInfo.GroupMaximum,
			Introduce:       gInfo.GroupIntroduce,
			Status:          group.GroupStatus(gInfo.GroupStatus),
			OwnerId:         gInfo.GroupOwnerId,
			CreateTime:      gInfo.GroupCreateTime,
			UpdateTime:      gInfo.GroupUpdateTime,
			JoinType:        group.GroupJoinType(gInfo.GroupJoinType),
			MuteType:        group.GroupMuteType(gInfo.GroupMuteType),
			FriendType:      group.GroupFriendlyType(gInfo.GroupFriendType),
			MutedNumbers:    mutedNumbers,
			ManagerNumbers:  managerNumbers,
			AESKey:          gInfo.GroupAESKey,
			MaskName:        gInfo.GroupPubName,
			Type:            group.GroupType(gInfo.GroupType),
		},
	}, nil
}
