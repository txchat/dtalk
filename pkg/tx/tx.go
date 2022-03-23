package tx

import (
	"context"
	"math/rand"
	"time"

	chainCommon "github.com/33cn/chain33/common"
	"github.com/33cn/chain33/common/address"
	"github.com/33cn/chain33/rpc/grpcclient"
	coinsTypes "github.com/33cn/chain33/system/dapp/coins/types"
	chainTypes "github.com/33cn/chain33/types"
	chainUtil "github.com/33cn/chain33/util"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type ChainClient struct {
	client chainTypes.Chain33Client // chain33提供的客户端
	cfg    *Config
	//FeePri crypto.PrivKey // 代扣私钥
}

// NewChainClient 新建上链客户端
func NewChainClient(c *Config) (*ChainClient, error) {
	if c == nil {
		return nil, errors.New("ChainClient: illegal ChainClient configure")
	}
	initConfig(c)
	conn, err := grpc.Dial(grpcclient.NewMultipleURL(c.Grpc.BlockChainAddr), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := chainTypes.NewChain33Client(conn)
	_, err = client.GetWalletStatus(context.Background(), &chainTypes.ReqNil{})
	if err != nil {
		return nil, errors.Errorf("GetWalletStatus fail Err:%v", err.Error())
	}
	c.Chain.BaseExec = c.Chain.Title + chainTypes.NoneX
	c.Chain.CoinExec = c.Chain.Title + c.Chain.CoinExec
	if err != nil {
		return nil, errors.Errorf("NewChainClient AesCbc Err:%v", err.Error())
	}
	//return &ChainClient{client: client, cfg: c, FeePri: chainUtil.HexToPrivkey(c.Chain.FeePrikey)}, nil
	return &ChainClient{client: client, cfg: c}, nil
}

// MustNewChainClient 新建上链客户端
func MustNewChainClient(c *Config) *ChainClient {
	f, err := NewChainClient(c)
	if err != nil {
		panic(err)
	}

	return f
}

func initConfig(c *Config) {
	const defaultExec = "testproofv2"
	if c.Chain.ReceiveCoinAmount == 0 {
		c.Chain.ReceiveCoinAmount = 100000000
	}
	//if c.ProofCh == nil {
	//	c.ProofCh = &ProofCh{100, 3, 20, 3}
	//}
	//if c.Organization == nil {
	//	c.Organization = &Organization{10, 3, 10}
	//}
}

func (c *ChainClient) newTxInfo(exec, prikey string, payload []byte, feerate int64) (*TxInfo, error) {
	//if prikey != c.cfg.Chain.FeePrikey {
	//	pridata, err := hex.DecodeString(prikey)
	//	if err != nil {
	//		return nil, err
	//	}
	//	pridata, err = c.cipher.Decrypt(pridata)
	//	if err != nil {
	//		return nil, err
	//	}
	//	prikey = string(pridata)
	//}

	return &TxInfo{
		Exec:    exec,
		Payload: payload,
		FeeRate: feerate,
		Prikey:  prikey,
	}, nil
}

// CreateTx 创建交易
func (c *ChainClient) createTx(t *TxInfo) *chainTypes.Transaction {
	tx := &chainTypes.Transaction{
		Execer:  []byte(t.Exec),
		Payload: t.Payload,
		Nonce:   rand.Int63(),
		To:      address.ExecAddress(t.Exec),
	}
	if t.FeeRate != noFeeRate {
		if t.FeeRate == 0 {
			t.FeeRate = minFeeRate
		}
		tx.SetRealFee(t.FeeRate)
	}

	return tx
}

// CreateTx 创建签名交易
func (c *ChainClient) createSignTx(t *TxInfo) *chainTypes.Transaction {
	tx := c.createTx(t)
	priv := chainUtil.HexToPrivkey(t.Prikey)
	tx.Sign(c.cfg.Encrypt.SignType, priv)
	return tx
}

// CreateBaseTx 创建代扣基础交易
func (c *ChainClient) createBaseTx(exec string) *chainTypes.Transaction {
	tx := &chainTypes.Transaction{
		Execer: []byte(exec),
		Nonce:  rand.Int63(),
		To:     address.ExecAddress(exec),
	}
	return tx
}

// CreateBaseTx 创建项目方扣费交易
func (c *ChainClient) createCoinTx() *chainTypes.Transaction {
	transfer := &chainTypes.AssetsTransfer{
		Cointoken: "",
		Amount:    c.cfg.Chain.ReceiveCoinAmount,
		Note:      nil,
		To:        c.cfg.Chain.ReceiveCoinAddr,
	}
	cbvalue := &coinsTypes.CoinsAction_Transfer{
		Transfer: transfer,
	}
	cb := &coinsTypes.CoinsAction{
		Value: cbvalue,
		Ty:    coinsTypes.CoinsActionTransfer,
	}
	payloadData, _ := proto.Marshal(cb)
	tx := &chainTypes.Transaction{
		Execer:  []byte(c.cfg.Chain.CoinExec),
		Nonce:   rand.Int63(),
		To:      address.ExecAddress(c.cfg.Chain.CoinExec),
		Payload: payloadData,
	}
	return tx
}

//// CreateTxGroup 创建交易组 代扣方式
//func (c *ChainClient) createTxGroup(t *TxInfo) (*chainTypes.Transaction, error) {
//	var txs []*chainTypes.Transaction
//	tx1 := c.createBaseTx(c.cfg.Chain.BaseExec)
//	tx2 := c.createTx(t)
//	txs = append(txs, tx1, tx2)
//	if c.cfg.Chain.Coin != "" {
//		tx3 := c.createCoinTx()
//		txs = append(txs, tx3)
//	}
//	txGroup, err := chainTypes.CreateTxGroup(txs, minFeeRate)
//	if err != nil {
//		return nil, err
//	}
//	txGroup.SignN(0, c.cfg.Encrypt.SignType, c.FeePri)
//	priv := chainUtil.HexToPrivkey(t.Prikey)
//	txGroup.SignN(1, c.cfg.Encrypt.SignType, priv)
//	if c.cfg.Chain.Coin != "" {
//		txGroup.SignN(2, c.cfg.Encrypt.SignType, c.FeePri)
//	}
//	return txGroup.Tx(), nil
//}

func (c *ChainClient) sendTxInfo(tx *chainTypes.Transaction) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	reply, err := c.client.SendTransaction(ctx, tx)
	if err != nil {
		// h, _ := c.client.GetLastHeader(context.Background(), nil)
		return "", err
	}
	if !reply.IsOk {
		return "", err
	}

	return chainCommon.ToHex(reply.Msg), nil
}
