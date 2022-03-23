package tx

import (
	"context"
	"time"

	chainCommon "github.com/33cn/chain33/common"
	chainTypes "github.com/33cn/chain33/types"
)

func (c *ChainClient) QueryHash(hash string) (*chainTypes.TransactionDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	data, err := chainCommon.FromHex(hash)
	if err != nil {
		return nil, err
	}
	reqHash := &chainTypes.ReqHash{
		Hash: data,
	}
	return c.client.QueryTransaction(ctx, reqHash)
}

func (c *ChainClient) GetRealTx(hash string) (*chainTypes.TransactionDetail, error) {
	td, err := c.QueryHash(hash)
	if err != nil {
		return td, err
	}
	if td.Tx == nil || td.Tx.Next == nil {
		return td, nil
	}
	return c.QueryHash(chainCommon.ToHex(td.Tx.Next))
}
