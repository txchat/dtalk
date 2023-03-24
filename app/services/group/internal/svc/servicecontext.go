package svc

import (
	"context"
	"fmt"

	"github.com/txchat/dtalk/app/services/transfer/transferclient"

	"github.com/txchat/dtalk/app/services/generator/generatorclient"
	"github.com/txchat/dtalk/app/services/group/internal/config"
	"github.com/txchat/dtalk/app/services/group/internal/dao"
	"github.com/txchat/dtalk/app/services/group/internal/model"
	"github.com/txchat/dtalk/internal/group"
	"github.com/txchat/dtalk/internal/notice"
	txchatNoticeApi "github.com/txchat/dtalk/internal/notice/txchat"
	"github.com/txchat/dtalk/internal/signal"
	txchatSignalApi "github.com/txchat/dtalk/internal/signal/txchat"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/im/app/logic/logicclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	Repo         dao.GroupRepository
	GroupManager *group.Manager
	IDGenRPC     generatorclient.Generator
	SignalHub    signal.Signal
	NoticeHub    notice.Notice
	logicClient  logicclient.Logic
}

func NewServiceContext(c config.Config) *ServiceContext {
	transferClient := transferclient.NewTransfer(zrpc.MustNewClient(c.TransferRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock()))
	return &ServiceContext{
		Config:       c,
		Repo:         dao.NewGroupRepositoryMysql(c.MySQL),
		GroupManager: group.NewGroupManager(),
		IDGenRPC: generatorclient.NewGenerator(zrpc.MustNewClient(c.IDGenRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		SignalHub: txchatSignalApi.NewSignalHub(transferClient),
		NoticeHub: txchatNoticeApi.NewNoticeHub(transferClient),
		logicClient: logicclient.NewLogic(zrpc.MustNewClient(c.LogicRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
	}
}

func (s *ServiceContext) RegisterGroupMembers(ctx context.Context, gid int64, members []string) error {
	reply, err := s.logicClient.JoinGroupByUID(ctx, &logicclient.JoinGroupByUIDReq{
		AppId: s.Config.AppID,
		Gid:   []string{util.MustToString(gid)},
		Uid:   members,
	})
	if err != nil || !reply.IsOk {
		if err.Error() == model.ErrPushMsgArg.Error() {
			return nil
		}
		if err == nil {
			err = fmt.Errorf("reply=%+v", reply)
		}
		return err
	}
	return nil
}

func (s *ServiceContext) UnRegisterGroup(ctx context.Context, gid int64) error {
	reply, err := s.logicClient.DelGroup(ctx, &logicclient.DelGroupReq{
		AppId: s.Config.AppID,
		Gid:   []string{util.MustToString(gid)},
	})
	if err != nil || !reply.IsOk {
		if err.Error() == model.ErrPushMsgArg.Error() {
			return nil
		}
		if err == nil {
			err = fmt.Errorf("reply=%+v", reply)
		}
		return err
	}
	return nil
}

func (s *ServiceContext) UnRegisterGroupMembers(ctx context.Context, gid int64, members []string) error {
	reply, err := s.logicClient.LeaveGroupByUID(ctx, &logicclient.LeaveGroupByUIDReq{
		AppId: s.Config.AppID,
		Gid:   []string{util.MustToString(gid)},
		Uid:   members,
	})
	if err != nil || !reply.IsOk {
		if err.Error() == model.ErrPushMsgArg.Error() {
			return nil
		}
		if err == nil {
			err = fmt.Errorf("reply=%+v", reply)
		}
		return err
	}
	return nil
}
