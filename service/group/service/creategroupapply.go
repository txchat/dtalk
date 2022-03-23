package service

import (
	"context"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
	"github.com/txchat/dtalk/service/group/model/types"
)

func (s *Service) CreateGroupApplySvc(ctx context.Context, req *types.CreateGroupApplyReq) (res *types.CreateGroupApplyResp, err error) {
	personId := req.PersonId
	groupId := util.ToInt64(req.Id)

	_, err = s.GetGroupInfoByGroupId(ctx, groupId)
	if err != nil {
		return nil, err
	}

	// 判断是否已经在群内
	if isExit, err := s.CheckInGroup(personId, groupId); err != nil {
		return nil, err
	} else if isExit == true {
		return nil, xerror.NewError(xerror.GroupMemberExist)
	}

	err = s.ExecCreateGroupApply(groupId, "", []string{personId}, req.ApplyNote)
	if err != nil {
		return nil, err
	}

	return &types.CreateGroupApplyResp{}, nil
}

func (s *Service) ExecCreateGroupApply(groupId int64, inviterId string, memberIds []string, applyNote string) error {
	groupApplys := make([]*db.GroupApply, 0)
	nowTime := s.getNowTime()
	applyId, err := s.getLogId(context.Background())
	if err != nil {
		return err
	}
	for _, memberId := range memberIds {
		groupApply := &db.GroupApply{
			Id:           applyId,
			GroupId:      groupId,
			InviterId:    inviterId,
			MemberId:     memberId,
			ApplyNote:    applyNote,
			OperatorId:   "",
			ApplyStatus:  biz.GroupApplyWait,
			RejectReason: "",
			CreateTime:   nowTime,
			UpdateTime:   nowTime,
		}
		groupApplys = append(groupApplys, groupApply)
	}

	tx, err := s.dao.NewTx()
	if err != nil {
		return err
	}
	defer tx.RollBack()

	err = s.dao.InsertGroupApplys(tx, groupApplys)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
