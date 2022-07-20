package dao

import (
	"testing"

	"github.com/inconshreveable/log15"
	"github.com/jinzhu/gorm"
	"github.com/txchat/dtalk/service/backup/model"
)

func TestDao_CreateAddressEnrolment(t *testing.T) {
	type fields struct {
		log log15.Logger
		db  *gorm.DB
	}
	type args struct {
		r *model.AddrMove
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				log: testLog,
				db:  testConn,
			},
			args: args{
				r: &model.AddrMove{
					BtyAddr: "testBty",
					BtcAddr: "testBtc",
					State:   0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				log: tt.fields.log,
				db:  tt.fields.db,
			}
			if err := d.CreateAddressEnrolment(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("CreateAddressEnrolment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDao_QueryAddressEnrolment(t *testing.T) {
	type fields struct {
		log log15.Logger
		db  *gorm.DB
	}
	type args struct {
		btyAddr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.AddrMove
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				log: testLog,
				db:  testConn,
			},
			args: args{
				btyAddr: "testBty",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				log: tt.fields.log,
				db:  tt.fields.db,
			}
			got, err := d.QueryAddressEnrolment(tt.args.btyAddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryAddressEnrolment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("got:", got)
		})
	}
}
