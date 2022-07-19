package logic

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/service"
)

type GetPubGroupInfoLogic struct {
	ctx context.Context
	svc *service.Service
}

func NewGetPubGroupInfoLogic(ctx context.Context, svc *service.Service) *GetPubGroupInfoLogic {
	return &GetPubGroupInfoLogic{
		ctx: ctx,
		svc: svc,
	}
}

// GetPubGroupInfo 查询群公开信息
func (l *GetPubGroupInfoLogic) GetPubGroupInfo(req *pb.GetPubGroupInfoReq) (*pb.GetPubGroupInfoResp, error) {
	group, err := l.svc.GetGroupInfoByGroupId(l.ctx, req.GroupId)
	if err != nil {
		return nil, err
	}

	owner, err := l.svc.GetMemberByMemberIdAndGroupId(l.ctx, group.GroupOwnerId, group.GroupId)
	if err != nil {
		return nil, err
	}

	groupInfo := NewRPCGroupInfo(group)
	groupInfo.Owner = NewRPCGroupMemberInfo(owner)

	person, err := l.svc.GetPersonByMemberIdAndGroupId(l.ctx, req.PersonId, req.GroupId)
	if err != nil {
		if xerror.NewError(xerror.GroupPersonNotExist).Error() != err.Error() {
			return nil, err
		}
	}

	if person == nil {
		groupInfo.AESKey = ""
	} else {
		groupInfo.Person = NewRPCGroupMemberInfo(person)
	}

	return &pb.GetPubGroupInfoResp{
		Group: groupInfo,
	}, nil
}
