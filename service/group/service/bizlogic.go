package service

import (
	"context"
	"fmt"
	"github.com/txchat/dtalk/pkg/util"
	"time"

	"github.com/pkg/errors"
	"github.com/txchat/dtalk/service/group/model"
	logic "github.com/txchat/im/api/logic/grpc"
)

func (s *Service) LogicNoticeJoin(ctx context.Context, groupId int64, groupMemberIds []string) error {
	gid := make([]string, 1, 1)
	gid[0] = util.ToString(groupId)

	groupMid := &logic.GroupsMid{
		AppId: s.cfg.AppId,
		Gid:   gid,
		Mids:  groupMemberIds,
	}

	reply, err := s.logicClient.JoinGroupsByMids(ctx, groupMid)
	if err != nil {
		if err.Error() == model.ErrPushMsgArg.Error() {
			return nil
		}
		return err
	} else if reply.IsOk == false {
		return errors.New(fmt.Sprintf("reply=%+v", reply))
	}

	return nil
}

func (s *Service) LogicNoticeLeave(ctx context.Context, groupId int64, groupMemberIds []string) error {
	// 保证离开前的通知收到
	time.Sleep(3 * time.Second)

	gid := make([]string, 1, 1)
	gid[0] = util.ToString(groupId)

	groupMid := &logic.GroupsMid{
		AppId: s.cfg.AppId,
		Gid:   gid,
		Mids:  groupMemberIds,
	}

	reply, err := s.logicClient.LeaveGroupsByMids(ctx, groupMid)
	if err != nil {
		if err.Error() == model.ErrPushMsgArg.Error() {
			return nil
		}
		return err
	} else if reply.IsOk == false {
		return errors.New(fmt.Sprintf("reply=%+v", reply))
	}

	return nil
}

func (s *Service) LogicNoticeDel(ctx context.Context, groupId int64) error {
	// 保证离开前的通知收到
	time.Sleep(3 * time.Second)

	gid := make([]string, 1, 1)
	gid[0] = util.ToString(groupId)

	delGroupsReq := &logic.DelGroupsReq{
		AppId: s.cfg.AppId,
		Gid:   gid,
	}

	reply, err := s.logicClient.DelGroups(ctx, delGroupsReq)
	if err != nil {
		return err
	} else if reply.IsOk == false {
		return errors.New(fmt.Sprintf("reply=%+v", reply))
	}

	return nil
}
