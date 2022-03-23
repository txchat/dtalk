package cdk

import (
	"reflect"
	"testing"

	chainTypes "github.com/33cn/chain33/types"
	"github.com/rs/zerolog"
	"github.com/txchat/dtalk/pkg/tx"
	"github.com/txchat/dtalk/service/backend/model/biz"
)

func TestDealFailedCdkOrder(t *testing.T) {
	ids := []int64{1, 2}
	err := srv.dealFailedCdkOrder(ids)
	if err != nil {
		t.Log(err)
	} else {
		t.Log("success")
	}

}

func TestServiceContent_checkBlockTxResult(t *testing.T) {
	title := "user.p.testproofv2."
	var config = tx.Config{
		Grpc: tx.Grpc{
			BlockChainAddr: "172.16.101.87:8902",
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
	type fields struct {
		log             zerolog.Logger
		CdkOrderMessage chan *biz.CdkOrderMessage
		CdkMaxNumber    int64
		chain33Client   *tx.ChainClient
	}
	type args struct {
		hash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    biz.TxResult
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				CdkOrderMessage: nil,
				CdkMaxNumber:    0,
				chain33Client:   tx.MustNewChainClient(&config),
			},
			args: args{
				hash: "",
			},
			want: biz.TxResult{
				Success: true,
				To:      "19rJWrjA8qgczqHPWLfFaSxzsSYdsepoNg",
				Amount:  100000000,
				Symbol:  "ZJY3TUY",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServiceContent{
				log:             tt.fields.log,
				CdkOrderMessage: tt.fields.CdkOrderMessage,
				CdkMaxNumber:    tt.fields.CdkMaxNumber,
				chain33Client:   tt.fields.chain33Client,
			}
			got, err := s.checkBlockTxResult(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkBlockTxResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkBlockTxResult() got = %v, want %v", got, tt.want)
			}
		})
	}
}
