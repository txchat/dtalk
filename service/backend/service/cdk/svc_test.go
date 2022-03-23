package cdk

import (
	"github.com/txchat/dtalk/service/backend/config"
	"github.com/txchat/dtalk/service/backend/dao"
	idgen "github.com/txchat/dtalk/service/generator/api"
	"os"
	"testing"
	"time"
)

var (
	cfg            *config.Config
	srv            *ServiceContent
	Dao            *dao.Dao
	idGenRPCClient *idgen.Client
)

func TestMain(m *testing.M) {
	cfg = config.Default()
	Dao = dao.New(cfg)
	idGenRPCClient = idgen.New(cfg.IdGenRPCClient.RegAddrs, cfg.IdGenRPCClient.Schema, cfg.IdGenRPCClient.SrvName, time.Duration(cfg.IdGenRPCClient.Dial))
	srv = NewServiceContent(cfg.Env, Dao, idGenRPCClient)

	os.Exit(m.Run())
}
