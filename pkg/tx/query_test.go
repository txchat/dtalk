package tx

import (
	"encoding/json"
	"testing"

	chainCommon "github.com/33cn/chain33/common"

	chainTypes "github.com/33cn/chain33/types"
)

var (
	BlockChainAddr = "172.16.101.87:8902"

	title    = "user.p.testproofv2."
	coinExec = "coins" // 平行链币消耗执行器
)

var config = Config{
	Grpc: Grpc{
		BlockChainAddr: BlockChainAddr,
	},
	Chain: Chain{
		FeePrikey: "",
		FeeAddr:   "",
		Title:     title,
		BaseExec:  title + chainTypes.NoneX,
		CoinExec:  title + coinExec,
	},
	Encrypt: Encrypt{
		Seed:     "",
		SignType: 1,
	},
}

func initClient() *ChainClient {
	cfg := config
	return MustNewChainClient(&cfg)
}

func TestChainClient_QueryHash(t *testing.T) {
	type fields struct {
		cli *ChainClient
	}
	type args struct {
		hash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *chainTypes.TransactionDetail
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				cli: initClient(),
			},
			args: args{
				hash: "0xe5bee7afdf8794d76479f2d1e0bb5675ce6b9d54e86e0fcefcfa147436ce4abd",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.cli
			got, err := c.QueryHash(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			b, err := json.Marshal(got)
			if err != nil {
				t.Errorf("Json Marshal() error = %v", err)
				return
			}
			t.Logf("got json: %v\n", string(b))

			t.Logf("next hash: %v\n", chainCommon.ToHex(got.Tx.Next))

			//tx2 := chainTypes.Transaction{}
			//data, err := base64.StdEncoding.DecodeString(string(got.Tx.Next))
			//if err != nil {
			//	t.Errorf("base64.StdEncoding.DecodeString() error = %v", err)
			//	return
			//}
			//err = json.Unmarshal(data, &tx2)
			//if err != nil {
			//	t.Errorf("json.Unmarshal() error = %v", err)
			//	return
			//}
			//t.Logf("got json: %v\n", tx2)
		})
	}
}

func TestChainClient_GetRealTx(t *testing.T) {
	type fields struct {
		cli *ChainClient
	}
	type args struct {
		hash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *chainTypes.TransactionDetail
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				cli: initClient(),
			},
			args: args{
				hash: "0xfef6aab8dd9e65310982eda196a77c1c80494745b9be21fbf509326e010421eb",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.cli
			got, err := c.GetRealTx(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRealTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			b, err := json.Marshal(got)
			if err != nil {
				t.Errorf("Json Marshal() error = %v", err)
				return
			}
			t.Logf("got json: %v\n", string(b))
		})
	}
}
