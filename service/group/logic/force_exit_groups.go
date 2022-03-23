package logic

import (
	"context"
	"github.com/rs/zerolog"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type ForceExitGroupsLogic struct {
	ctx context.Context
	svc *service.Service
	log zerolog.Logger
}

func NewForceExitGroupsLogic(ctx context.Context, svc *service.Service) *ForceExitGroupsLogic {
	return &ForceExitGroupsLogic{
		ctx: ctx,
		svc: svc,
		log: svc.GetLog(),
	}
}

// ForceExitGroups .
// todo: 没有好的想法
// 相当于多次 DeleteMember
func (l *ForceExitGroupsLogic) ForceExitGroups(req *pb.ForceExitGroupsReq) (*pb.ForceExitGroupsResp, error) {
	_, err := FilteredMemberId(req.Member.Id)
	if err != nil {
		return nil, err
	}

	dml := NewForceDeleteMemberLogic(l.ctx, l.svc)
	for _, groupId := range req.GroupIds {
		_, err := dml.ForceDeleteMember(&pb.ForceDeleteMemberReq{
			MemberId: req.Member.Id,
			GroupId:  groupId,
		})
		if err != nil {
			l.log.Error().Err(err).Msg("ForceExitGroups")
		}
	}

	return &pb.ForceExitGroupsResp{}, nil
}
