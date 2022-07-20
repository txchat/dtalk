package dao

import (
	"testing"
	"time"

	"github.com/inconshreveable/log15"
	"github.com/jinzhu/gorm"
	"github.com/txchat/dtalk/service/backup/model"
)

func TestDao_Query(t *testing.T) {
	type fields struct {
		log log15.Logger
		db  *gorm.DB
	}
	type args struct {
		tp    uint32
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.AddrBackup
		wantErr bool
	}{
		{
			name: "test query",
			fields: fields{
				log: testLog,
				db:  testConn,
			},
			args: args{
				tp:    model.Phone,
				query: "18217379846",
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
			got, err := d.Query(tt.args.tp, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got record:%v", got)
		})
	}
}

func TestDao_UpdateAddrBackup(t *testing.T) {
	type fields struct {
		log log15.Logger
		db  *gorm.DB
	}
	type args struct {
		tp uint32
		r  *model.AddrBackup
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test bind",
			fields: fields{
				log: testLog,
				db:  testConn,
			},
			args: args{
				tp: model.Phone,
				r: &model.AddrBackup{
					Address:    "test_address",
					Area:       "",
					Phone:      "test_phone",
					Mnemonic:   "test_mne1",
					UpdateTime: time.Now(),
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
			if err := d.UpdateAddrBackup(tt.args.tp, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAddrBackup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDao_UpdateAddrRelate(t *testing.T) {
	type fields struct {
		log log15.Logger
		db  *gorm.DB
	}
	type args struct {
		tp uint32
		r  *model.AddrRelate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test relate",
			fields: fields{
				log: testLog,
				db:  testConn,
			},
			args: args{
				tp: model.Phone,
				r: &model.AddrRelate{
					Address:    "test_address",
					Area:       "",
					Phone:      "test_phone2",
					Mnemonic:   "test_mne2",
					UpdateTime: time.Now(),
				},
			},
			wantErr: false,
		},
		{
			name: "test relate2",
			fields: fields{
				log: testLog,
				db:  testConn,
			},
			args: args{
				tp: model.Phone,
				r: &model.AddrRelate{
					Address:    "test_address",
					Area:       "",
					Phone:      "test_phone",
					Mnemonic:   "test_mne3",
					UpdateTime: time.Now(),
				},
			},
			wantErr: false,
		},
		{
			name: "test relate2",
			fields: fields{
				log: testLog,
				db:  testConn,
			},
			args: args{
				tp: model.Phone,
				r: &model.AddrRelate{
					Address:    "test_address2",
					Area:       "",
					Phone:      "test_phone2",
					Mnemonic:   "test_mne22",
					UpdateTime: time.Now(),
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
			if err := d.UpdateAddrRelate(tt.args.tp, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAddrRelate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
