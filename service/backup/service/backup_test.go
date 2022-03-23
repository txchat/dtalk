package service

import (
	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/service/backup/config"
	"github.com/txchat/dtalk/service/backup/dao"
	"github.com/txchat/dtalk/service/backup/model"
	"testing"
)

func TestService_PhoneIsBound(t *testing.T) {
	type fields struct {
		log           log15.Logger
		cfg           *config.Config
		dao           *dao.Dao
		smsValidate   model.Validate
		emailValidate model.Validate
	}
	type args struct {
		area  string
		phone string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test phone is bound",
			fields: fields{
				log:           testLog,
				cfg:           testCfg,
				dao:           testDao,
				smsValidate:   testSmsValidate,
				emailValidate: testEmailValidate,
			},
			args: args{
				area:  "",
				phone: "18217379846",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				log:           tt.fields.log,
				cfg:           tt.fields.cfg,
				dao:           tt.fields.dao,
				smsValidate:   tt.fields.smsValidate,
				emailValidate: tt.fields.emailValidate,
			}
			got, err := s.PhoneIsBound(tt.args.area, tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("PhoneIsBound() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got exists:%v", got)
		})
	}
}

func TestService_PhoneRetrieve(t *testing.T) {
	type fields struct {
		log           log15.Logger
		cfg           *config.Config
		dao           *dao.Dao
		smsValidate   model.Validate
		emailValidate model.Validate
	}
	type args struct {
		area  string
		phone string
		code  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.AddrBackup
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				log:           testLog,
				cfg:           testCfg,
				dao:           testDao,
				smsValidate:   testSmsValidate,
				emailValidate: testEmailValidate,
			},
			args: args{
				area:  "",
				phone: "test_phone",
				code:  "111111",
			},
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				log:           testLog,
				cfg:           testCfg,
				dao:           testDao,
				smsValidate:   testSmsValidate,
				emailValidate: testEmailValidate,
			},
			args: args{
				area:  "",
				phone: "test_phone2",
				code:  "111111",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				log:           tt.fields.log,
				cfg:           tt.fields.cfg,
				dao:           tt.fields.dao,
				smsValidate:   tt.fields.smsValidate,
				emailValidate: tt.fields.emailValidate,
			}
			got, err := s.PhoneRetrieve(tt.args.area, tt.args.phone, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("PhoneRetrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
