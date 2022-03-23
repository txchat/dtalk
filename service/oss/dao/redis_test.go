package dao

import (
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
	"github.com/txchat/dtalk/pkg/oss"
)

func TestDao_SaveAssumeRole(t *testing.T) {
	type fields struct {
		log   zerolog.Logger
		db    *gorm.DB
		redis *redis.Pool
	}
	type args struct {
		appId   string
		ossType string
		cfg     *oss.Config
		data    *oss.AssumeRoleResp
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test add assumeRole",
			fields: fields{
				log:   testLog,
				redis: testRedis,
			},
			args: args{
				appId:   "test",
				ossType: "aliyun",
				cfg: &oss.Config{
					EndPoint:        "",
					AccessKeyId:     "",
					AccessKeySecret: "",
					Role:            "",
					Policy:          "",
					DurationSeconds: 20,
				},
				data: &oss.AssumeRoleResp{
					RequestId:       "123",
					Credentials:     oss.Credentials{},
					AssumedRoleUser: oss.AssumedRoleUser{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				log:   tt.fields.log,
				redis: tt.fields.redis,
			}
			if err := d.SaveAssumeRole(tt.args.appId, tt.args.ossType, tt.args.cfg, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SaveAssumeRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDao_GetAssumeRole(t *testing.T) {
	type fields struct {
		log   zerolog.Logger
		db    *gorm.DB
		redis *redis.Pool
	}
	type args struct {
		appId   string
		ossType string
		cfg     *oss.Config
		data    *oss.AssumeRoleResp
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test add assumeRole",
			fields: fields{
				log:   testLog,
				redis: testRedis,
			},
			args: args{
				appId:   "test",
				ossType: "aliyun",
				cfg: &oss.Config{
					EndPoint:        "",
					AccessKeyId:     "",
					AccessKeySecret: "",
					Role:            "",
					Policy:          "",
					DurationSeconds: 20,
				},
				data: &oss.AssumeRoleResp{
					RequestId:       "123",
					Credentials:     oss.Credentials{},
					AssumedRoleUser: oss.AssumedRoleUser{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				log:   tt.fields.log,
				redis: tt.fields.redis,
			}
			got, err := d.GetAssumeRole(tt.args.appId, tt.args.ossType, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAssumeRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
