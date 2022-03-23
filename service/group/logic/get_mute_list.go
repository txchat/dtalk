package logic

import (
	"context"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type GetMuteListLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewGetMuteListLogic(ctx context.Context, svc *service.Service) *GetMuteListLogic {
	return &GetMuteListLogic{
		ctx: ctx,
		svc: svc,
	}
}

// GetMuteList 查询群禁言列表
func (l *GetMuteListLogic) GetMuteList(req *pb.GetMuteListReq) (*pb.GetMuteListResp, error) {
	groupId := req.GroupId
	personId := req.PersonId

	_, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, err
	}

	person, err := l.svc.GetPersonByMemberIdAndGroupId(l.ctx, personId, groupId)
	if err != nil {
		return nil, err
	}
	if err = person.IsAdmin(); err != nil {
		return nil, err
	}

	muteList, err := l.svc.GetGroupMembersMutedByGroupId(groupId)
	if err != nil {
		return nil, err
	}

	return &pb.GetMuteListResp{
		Members: NewRPCGroupMemberInfos(muteList),
	}, nil
}
