package logic

import (
	"context"

	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type GetGroupMemberInfoLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewGetGroupMemberInfoLogic(ctx context.Context, svc *service.Service) *GetGroupMemberInfoLogic {
	return &GetGroupMemberInfoLogic{
		ctx: ctx,
		svc: svc,
	}
}

// GetGroupMemberInfo 查询一个人的信息
func (l *GetGroupMemberInfoLogic) GetGroupMemberInfo(req *pb.GetGroupMemberInfoReq) (*pb.GetGroupMemberInfoResp, error) {
	groupId := req.GroupId
	memberId := req.MemberId
	personId := req.PersonId

	_, err := l.svc.GetGroupInfoByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, err
	}

	_, err = l.svc.GetPersonByMemberIdAndGroupId(l.ctx, personId, groupId)
	if err != nil {
		return nil, err
	}

	member, err := l.svc.GetMemberByMemberIdAndGroupId(l.ctx, memberId, groupId)
	if err != nil {
		return nil, err
	}

	return &pb.GetGroupMemberInfoResp{
		Member: NewRPCGroupMemberInfo(member),
	}, nil
}
