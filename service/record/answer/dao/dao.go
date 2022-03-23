package dao

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/pkg/net/grpc"
	idgen "github.com/txchat/dtalk/service/generator/api"
	groupApi "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/record/answer/config"
	"github.com/txchat/dtalk/service/record/kafka/publisher"
	"google.golang.org/grpc/resolver"
	kafka "gopkg.in/Shopify/sarama.v1"
)

type Dao struct {
	appId string
	//log            zerolog.Logger
	redis          *redis.Pool
	mqPub          kafka.SyncProducer
	idGenRPCClient idgen.GeneratorClient
	groupRPCClient groupApi.GroupClient
}

func New(c *config.Config) *Dao {
	d := &Dao{
		appId: c.AppId,
		//log:            zlog.Logger,
		redis:          newRedis(c.Redis),
		mqPub:          publisher.NewKafkaPub(c.MQPub.Brokers),
		idGenRPCClient: newIdGenClient(c),
		groupRPCClient: newGroupClient(c),
	}
	return d
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

func newIdGenClient(cfg *config.Config) idgen.GeneratorClient {
	rb := naming.NewResolver(cfg.IdGenRPCClient.RegAddrs, cfg.IdGenRPCClient.Schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", cfg.IdGenRPCClient.Schema, cfg.IdGenRPCClient.SrvName) // "schema://[authority]/service"
	fmt.Println("rpc client call addr:", addr)

	conn, err := grpc.NewGRPCConn(addr, time.Duration(cfg.IdGenRPCClient.Dial))
	if err != nil {
		panic(err)
	}
	return idgen.NewGeneratorClient(conn)
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
