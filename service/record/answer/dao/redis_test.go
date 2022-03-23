package dao

import (
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog"
	"github.com/txchat/dtalk/service/record/answer/model"
)

func TestDao_AddRecordSeqIndex(t *testing.T) {
	type fields struct {
		appId string
		log   zerolog.Logger
		redis *redis.Pool
	}
	type args struct {
		uid string
		m   *model.MsgIndex
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
				appId: "dtalk",
				log:   testLog,
				redis: testRedis,
			},
			args: args{
				uid: "2",
				m: &model.MsgIndex{
					Mid:        "1",
					Seq:        "1",
					SenderId:   "2",
					CreateTime: 0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				appId: tt.fields.appId,
				log:   tt.fields.log,
				redis: tt.fields.redis,
			}
			if err := d.AddRecordSeqIndex(tt.args.uid, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("AddRecordSeqIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDao_GetRecordSeqIndex(t *testing.T) {
	type fields struct {
		appId string
		log   zerolog.Logger
		redis *redis.Pool
	}
	type args struct {
		uid string
		seq string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.MsgIndex
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				appId: "dtalk",
				log:   testLog,
				redis: testRedis,
			},
			args: args{
				uid: "2",
				seq: "1",
			},
			want:    nil,
			wantErr: false,
		}, {
			name: "",
			fields: fields{
				appId: "dtalk",
				log:   testLog,
				redis: testRedis,
			},
			args: args{
				uid: "2",
				seq: "2",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				appId: tt.fields.appId,
				log:   tt.fields.log,
				redis: tt.fields.redis,
			}
			got, err := d.GetRecordSeqIndex(tt.args.uid, tt.args.seq)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRecordSeqIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got:%v\n", got)
		})
	}
}
