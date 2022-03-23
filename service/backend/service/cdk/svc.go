package cdk

import (
	chainTypes "github.com/33cn/chain33/types"
	"github.com/rs/zerolog"
	"github.com/txchat/dtalk/pkg/logger"
	"github.com/txchat/dtalk/pkg/tx"
	"github.com/txchat/dtalk/service/backend/config"
	"github.com/txchat/dtalk/service/backend/dao"
	"github.com/txchat/dtalk/service/backend/model/biz"
	idgen "github.com/txchat/dtalk/service/generator/api"
)

type ServiceContent struct {
	log            zerolog.Logger
	dao            *dao.Dao
	idGenRPCClient *idgen.Client
	// cdk订单处理 channal
	CdkOrderMessage chan *biz.CdkOrderMessage
	CdkMaxNumber    int64
	chain33Client   *tx.ChainClient
}

func NewServiceContent(env string, dao *dao.Dao, idgen *idgen.Client, cdkMaxNumber int64, chainCliConfig config.Chain33Client) *ServiceContent {
	s := &ServiceContent{
		log:             logger.New(env, "cdkOrder"),
		dao:             dao,
		idGenRPCClient:  idgen,
		CdkOrderMessage: make(chan *biz.CdkOrderMessage, 100),
		CdkMaxNumber:    cdkMaxNumber,
		chain33Client:   initChain33Client(chainCliConfig),
	}
	go s.ListenOrder()

	go s.CleanFrozenOrder()

	return s
}

func initChain33Client(chainCliConfig config.Chain33Client) *tx.ChainClient {
	title := chainCliConfig.Title
	cfg := tx.Config{
		Grpc: tx.Grpc{
			BlockChainAddr: chainCliConfig.BlockChainAddr,
		},
		Chain: tx.Chain{
			FeePrikey: "",
			FeeAddr:   "",
			Title:     title,
			BaseExec:  title + chainTypes.NoneX,
		},
		Encrypt: tx.Encrypt{
			Seed:     "",
			SignType: 1,
		},
	}

	return tx.MustNewChainClient(&cfg)
}
