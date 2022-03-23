package dao

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/pkg/net/grpc"
	groupApi "github.com/txchat/dtalk/service/group/api"
	pusher "github.com/txchat/dtalk/service/record/pusher/api"
	"github.com/txchat/dtalk/service/record/store/config"
	"google.golang.org/grpc/resolver"
)

type Dao struct {
	appId          string
	log            zerolog.Logger
	conn           *mysql.MysqlConn
	redis          *redis.Pool
	groupRPCClient groupApi.GroupClient
	pusherCli      pusher.PusherClient
}

func New(c *config.Config) *Dao {
	d := &Dao{
		appId:          c.AppId,
		log:            zlog.Logger,
		conn:           newDB(c.MySQL),
		redis:          newRedis(c.Redis),
		groupRPCClient: newGroupClient(c),
		pusherCli:      newPushClient(c),
	}
	return d
}

func newDB(cfg *config.MySQL) *mysql.MysqlConn {
	c, err := mysql.NewMysqlConn(cfg.Host, fmt.Sprintf("%v", cfg.Port),
		cfg.User, cfg.Pwd, cfg.Db, "UTF8MB4")
	if err != nil {
		panic(err)
	}
	return c
}

func newRedis(c *config.Redis) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     c.Idle,
		MaxActive:   c.Active,
		IdleTimeout: time.Duration(c.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(c.Network, c.Addr,
				redis.DialConnectTimeout(time.Duration(c.DialTimeout)),
				redis.DialReadTimeout(time.Duration(c.ReadTimeout)),
				redis.DialWriteTimeout(time.Duration(c.WriteTimeout)),
				redis.DialPassword(c.Auth),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

func newPushClient(cfg *config.Config) pusher.PusherClient {
	rb := naming.NewResolver(cfg.PusherRPCClient.RegAddrs, cfg.PusherRPCClient.Schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", cfg.PusherRPCClient.Schema, cfg.PusherRPCClient.SrvName) // "schema://[authority]/service"
	fmt.Println("rpc client call addr:", addr)

	conn, err := grpc.NewGRPCConn(addr, time.Duration(cfg.PusherRPCClient.Dial))
	if err != nil {
		panic(err)
	}
	return pusher.NewPusherClient(conn)
}

func newGroupClient(cfg *config.Config) groupApi.GroupClient {
	rb := naming.NewResolver(cfg.GroupRPCClient.RegAddrs, cfg.GroupRPCClient.Schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", cfg.GroupRPCClient.Schema, cfg.GroupRPCClient.SrvName) // "schema://[authority]/service"
	fmt.Println("rpc client call addr:", addr)

	conn, err := grpc.NewGRPCConn(addr, time.Duration(cfg.GroupRPCClient.Dial))
	if err != nil {
		panic(err)
	}
	return groupApi.NewGroupClient(conn)
}

func (d *Dao) NewTx() (*mysql.MysqlTx, error) {
	return d.conn.NewTx()
}
