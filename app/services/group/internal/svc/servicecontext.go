package svc

import (
	"context"
	"errors"
	"fmt"

	"github.com/txchat/dtalk/app/services/answer/answerclient"
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
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/im/app/logic/logicclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	Repo         dao.GroupRepository
	GroupManager *group.GroupManager
	IDGenRPC     generatorclient.Generator
	SignalHub    signal.Signal
	NoticeHub    notice.Notice
	logicClient  logicclient.Logic
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, err := mysql.NewMysqlConn(c.MySQL.Host, fmt.Sprintf("%v", c.MySQL.Port),
		c.MySQL.User, c.MySQL.Pwd, c.MySQL.DB, "UTF8MB4")
	if err != nil {
		panic(err)
	}
	answerClient := answerclient.NewAnswer(zrpc.MustNewClient(c.AnswerRPC,
		zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock()))
	return &ServiceContext{
		Config:       c,
		Repo:         dao.NewGroupRepositoryMysql(conn),
		GroupManager: group.NewGroupManager(),
		IDGenRPC: generatorclient.NewGenerator(zrpc.MustNewClient(c.IDGenRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
		SignalHub: txchatSignalApi.NewSignalHub(answerClient),
		NoticeHub: txchatNoticeApi.NewNoticeHub(answerClient),
		logicClient: logicclient.NewLogic(zrpc.MustNewClient(c.LogicRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor), zrpc.WithNonBlock())),
	}
}

func (s *ServiceContext) RegisterGroupMembers(ctx context.Context, gid int64, members []string) error {
	reply, err := s.logicClient.JoinGroupsByMids(ctx, &logicclient.GroupsMid{
		AppId: s.Config.AppID,
		Gid:   []string{util.MustToString(gid)},
		Mids:  members,
	})
	if err != nil || reply.IsOk == false {
		if err.Error() == model.ErrPushMsgArg.Error() {
			return nil
		}
		if err == nil {
			err = errors.New(fmt.Sprintf("reply=%+v", reply))
		}
		return err
	}
	return nil
}

func (s *ServiceContext) UnRegisterGroup(ctx context.Context, gid int64) error {
	reply, err := s.logicClient.DelGroups(ctx, &logicclient.DelGroupsReq{
		AppId: s.Config.AppID,
		Gid:   []string{util.MustToString(gid)},
	})
	if err != nil || reply.IsOk == false {
		if err.Error() == model.ErrPushMsgArg.Error() {
			return nil
		}
		if err == nil {
			err = errors.New(fmt.Sprintf("reply=%+v", reply))
		}
		return err
	}
	return nil
}

func (s *ServiceContext) UnRegisterGroupMembers(ctx context.Context, gid int64, members []string) error {
	reply, err := s.logicClient.LeaveGroupsByMids(ctx, &logicclient.GroupsMid{
		AppId: s.Config.AppID,
		Gid:   []string{util.MustToString(gid)},
		Mids:  members,
	})
	if err != nil || reply.IsOk == false {
		if err.Error() == model.ErrPushMsgArg.Error() {
			return nil
		}
		if err == nil {
			err = errors.New(fmt.Sprintf("reply=%+v", reply))
		}
		return err
	}
	return nil
}
