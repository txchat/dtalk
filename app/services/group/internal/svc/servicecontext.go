package svc

import (
	"context"
	"errors"
	"fmt"

	"github.com/txchat/dtalk/app/services/generator/generatorclient"
	"github.com/txchat/dtalk/app/services/group/internal/config"
	"github.com/txchat/dtalk/app/services/group/internal/dao"
	"github.com/txchat/dtalk/internal/group"
	"github.com/txchat/dtalk/internal/notice"
	"github.com/txchat/dtalk/internal/signal"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/group/model"
	logic "github.com/txchat/im/api/logic/grpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	Repo         dao.GroupRepository
	GroupManager *group.GroupManager
	IDGenRPC     generatorclient.Generator
	SignalHub    signal.Signal
	NoticeHub    notice.Notice
	logicClient  logic.LogicClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, err := mysql.NewMysqlConn(c.MySQL.Host, fmt.Sprintf("%v", c.MySQL.Port),
		c.MySQL.User, c.MySQL.Pwd, c.MySQL.DB, "UTF8MB4")
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		Repo:   dao.NewGroupRepositoryMysql(conn),
		IDGenRPC: generatorclient.NewGenerator(zrpc.MustNewClient(c.IDGenRPC,
			zrpc.WithUnaryClientInterceptor(xerror.ErrClientInterceptor))),
	}
}

func (s *ServiceContext) RegisterGroupMembers(ctx context.Context, gid int64, members []string) error {
	reply, err := s.logicClient.JoinGroupsByMids(ctx, &logic.GroupsMid{
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
	reply, err := s.logicClient.DelGroups(ctx, &logic.DelGroupsReq{
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
	reply, err := s.logicClient.LeaveGroupsByMids(ctx, &logic.GroupsMid{
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
