package dao

import (
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog"
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/record/store/model"
)

func TestDao_AddRecordCache(t *testing.T) {
	type fields struct {
		log   zerolog.Logger
		conn  *mysql.MysqlConn
		redis *redis.Pool
	}
	type args struct {
		uid string
		ver uint64
		m   *model.MsgCache
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test add record cache",
			fields: fields{
				log:   testLog,
				conn:  testConn,
				redis: testRedis,
			},
			args: args{
				uid: "test1",
				ver: 1,
				m: &model.MsgCache{
					Mid:        "1",
					Seq:        "1",
					SenderId:   "18XDVerjKrLi8xGXEkU8YFuMvqBBYnHdU1",
					ReceiverId: "1AsPsahP7FvpR7F2de1LhSB4SU5ShqZ7eu",
					MsgType:    0,
					Content:    "测试消息1",
					CreateTime: uint64(util.TimeNowUnixNano() / int64(time.Millisecond)),
					Prev:       0,
					Version:    1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				log:   tt.fields.log,
				conn:  tt.fields.conn,
				redis: tt.fields.redis,
			}
			if err := d.AddRecordCache(tt.args.uid, tt.args.ver, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("AddRecordCache() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDao_UserRecords(t *testing.T) {
	type fields struct {
		log   zerolog.Logger
		conn  *mysql.MysqlConn
		redis *redis.Pool
	}
	type args struct {
		uid string
		ver uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.MsgCache
		wantErr bool
	}{
		{
			name: "test get user records",
			fields: fields{
				log:   testLog,
				conn:  testConn,
				redis: testRedis,
			},
			args: args{
				uid: "test1",
				ver: 2,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				log:   tt.fields.log,
				conn:  tt.fields.conn,
				redis: tt.fields.redis,
			}
			got, err := d.UserRecords(tt.args.uid, tt.args.ver)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, item := range got {
				t.Logf("got %v: %v", i, item)
			}
		})
	}
}
