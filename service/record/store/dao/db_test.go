package dao

import (
	"testing"
	"time"

	group "github.com/txchat/dtalk/service/group/api"
	pusher "github.com/txchat/dtalk/service/record/pusher/api"

	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog"
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/record/store/model"
)

func TestDao_AppendMsgContent(t *testing.T) {
	type fields struct {
		log   zerolog.Logger
		conn  *mysql.MysqlConn
		redis *redis.Pool
	}
	type args struct {
		tx *mysql.MysqlTx
		m  *model.MsgContent
	}
	tx, err := testConn.NewTx()
	if err != nil {
		t.Errorf("AppendMsgContent() error = %v", err)
		return
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		want1   int64
		wantErr bool
	}{
		{
			name: "test append msg content",
			fields: fields{
				log:   testLog,
				conn:  testConn,
				redis: testRedis,
			},
			args: args{
				tx: tx,
				m: &model.MsgContent{
					Mid:        "1",
					Seq:        "1",
					SenderId:   "18XDVerjKrLi8xGXEkU8YFuMvqBBYnHdU1",
					ReceiverId: "1AsPsahP7FvpR7F2de1LhSB4SU5ShqZ7eu",
					MsgType:    0,
					Content:    "测试消息1",
					CreateTime: uint64(util.TimeNowUnixNano() / int64(time.Millisecond)),
					Source:     "",
					Reference:  "",
				},
			},
			want:    0,
			want1:   0,
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
			_, _, err := d.AppendMsgContent(tt.args.tx, tt.args.m)
			if (err != nil) != tt.wantErr {
				tx.RollBack()
				t.Errorf("AppendMsgContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			err = tx.Commit()
			if err != nil {
				t.Errorf("AppendMsgContent Commit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDao_AppendMsgRelation(t *testing.T) {
	type fields struct {
		log   zerolog.Logger
		conn  *mysql.MysqlConn
		redis *redis.Pool
	}
	type args struct {
		tx *mysql.MysqlTx
		m  *model.MsgRelation
	}
	tx, err := testConn.NewTx()
	if err != nil {
		t.Errorf("AppendMsgContent() error = %v", err)
		return
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		want1   int64
		wantErr bool
	}{
		{
			name: "tset append msg relation",
			fields: fields{
				log:   testLog,
				conn:  testConn,
				redis: testRedis,
			},
			args: args{
				tx: tx,
				m: &model.MsgRelation{
					Mid:        "1",
					OwnerUid:   "18XDVerjKrLi8xGXEkU8YFuMvqBBYnHdU1",
					OtherUid:   "1AsPsahP7FvpR7F2de1LhSB4SU5ShqZ7eu",
					Type:       0,
					CreateTime: uint64(util.TimeNowUnixNano() / int64(time.Millisecond)),
				},
			},
			want:    0,
			want1:   0,
			wantErr: false,
		}, {
			name: "tset append msg relation2",
			fields: fields{
				log:   testLog,
				conn:  testConn,
				redis: testRedis,
			},
			args: args{
				tx: tx,
				m: &model.MsgRelation{
					Mid:        "1",
					OwnerUid:   "1AsPsahP7FvpR7F2de1LhSB4SU5ShqZ7eu",
					OtherUid:   "18XDVerjKrLi8xGXEkU8YFuMvqBBYnHdU1",
					Type:       1,
					CreateTime: uint64(util.TimeNowUnixNano() / int64(time.Millisecond)),
				},
			},
			want:    0,
			want1:   0,
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
			_, _, err := d.AppendMsgRelation(tt.args.tx, tt.args.m)
			if (err != nil) != tt.wantErr {
				tx.RollBack()
				t.Errorf("AppendMsgRelation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	err = tx.Commit()
	if err != nil {
		t.Errorf("AppendMsgContent Commit() error = %v", err)
		return
	}
}

func TestDao_IncMsgVersion(t *testing.T) {
	type fields struct {
		log   zerolog.Logger
		conn  *mysql.MysqlConn
		redis *redis.Pool
	}
	type args struct {
		uid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "test inc msg id",
			fields: fields{
				log:   testLog,
				conn:  testConn,
				redis: testRedis,
			},
			args: args{
				uid: "1AsPsahP7FvpR7F2de1LhSB4SU5ShqZ7eu",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "test inc msg id",
			fields: fields{
				log:   testLog,
				conn:  testConn,
				redis: testRedis,
			},
			args: args{
				uid: "1AsPsahP7FvpR7F2de1LhSB4SU5ShqZ7eu",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "test inc msg id",
			fields: fields{
				log:   testLog,
				conn:  testConn,
				redis: testRedis,
			},
			args: args{
				uid: "1AsPsahP7FvpR7F2de1LhSB4SU5ShqZ7eu",
			},
			want:    0,
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
			got, err := d.IncMsgVersion(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("IncMsgVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got: %v", got)
		})
	}
}

func TestDao_UserLastMsg(t *testing.T) {
	type fields struct {
		log   zerolog.Logger
		conn  *mysql.MysqlConn
		redis *redis.Pool
	}
	type args struct {
		uid string
		num int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.MsgContent
		wantErr bool
	}{
		{
			name: "test get user last msg",
			fields: fields{
				log:   testLog,
				conn:  testConn,
				redis: testRedis,
			},
			args: args{
				uid: "1AsPsahP7FvpR7F2de1LhSB4SU5ShqZ7eu",
				num: 10,
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
			got, err := d.UserLastMsg(tt.args.uid, tt.args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserLastMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, item := range got {
				t.Logf("got num %v : %v\n", i, item)
			}
		})
	}
}

func TestDao_UserMsgAfter(t *testing.T) {
	type fields struct {
		log   zerolog.Logger
		conn  *mysql.MysqlConn
		redis *redis.Pool
	}
	type args struct {
		uid      string
		startMid int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.MsgContent
		wantErr bool
	}{
		{
			name: "test get user last msg",
			fields: fields{
				log:   testLog,
				conn:  testConn,
				redis: testRedis,
			},
			args: args{
				uid:      "1FdnxKR4r952x2HQA2BTTpFH6tgHYYNs3M",
				startMid: 0,
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
			got, err := d.UserMsgAfter(tt.args.uid, tt.args.startMid)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserMsgAfter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, item := range got {
				t.Logf("got num %v : %v\n", i, item)
			}
		})
	}
}

func TestDao_UnReceiveMsg(t *testing.T) {
	type fields struct {
		log   zerolog.Logger
		conn  *mysql.MysqlConn
		redis *redis.Pool
	}
	type args struct {
		uid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.MsgContent
		wantErr bool
	}{
		{
			name: "测试获取未读消息",
			fields: fields{
				log:   testLog,
				conn:  testConn,
				redis: testRedis,
			},
			args: args{
				uid: "17HTAXJqn3Wp5wMpdXopCkr4AZZjESrj1m",
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
			got, err := d.UnReceiveMsg(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnReceiveMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, r := range got {
				t.Logf("%d got:%v", i, r)
			}
		})
	}
}

func TestDao_MarkMsgReceived(t *testing.T) {
	type fields struct {
		log   zerolog.Logger
		conn  *mysql.MysqlConn
		redis *redis.Pool
	}
	type args struct {
		uid string
		mid int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		want1   int64
		wantErr bool
	}{
		{
			name: "测试修改state",
			fields: fields{
				log:   testLog,
				conn:  testConn,
				redis: testRedis,
			},
			args: args{
				uid: "",
				mid: 0,
			},
			want:    0,
			want1:   0,
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
			got, got1, err := d.MarkMsgReceived(tt.args.uid, tt.args.mid)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarkMsgReceived() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MarkMsgReceived() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MarkMsgReceived() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDao_DelMsgContent(t *testing.T) {
	type fields struct {
		appId string
		log   zerolog.Logger
		conn  *mysql.MysqlConn
		redis *redis.Pool
	}
	type args struct {
		mid int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		want1   int64
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				log:   testLog,
				conn:  testConnRecord,
				redis: testRedis,
			},
			args: args{
				mid: 176393658935808000,
			},
			want:    0,
			want1:   0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				log:   testLog,
				conn:  testConnRecord,
				redis: testRedis,
			}
			_, _, err := d.DelMsgContent(tt.args.mid)
			if (err != nil) != tt.wantErr {
				t.Errorf("DelMsgContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDao_DelGroupMsgContent(t *testing.T) {
	type fields struct {
		appId          string
		log            zerolog.Logger
		conn           *mysql.MysqlConn
		redis          *redis.Pool
		groupRPCClient group.GroupClient
		pusherCli      pusher.PusherClient
	}
	type args struct {
		mid int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		want1   int64
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				log:   testLog,
				conn:  testConnRecord,
				redis: testRedis,
			},
			args: args{
				mid: 130687370197471232,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				log:   testLog,
				conn:  testConnRecord,
				redis: testRedis,
			}
			_, _, err := d.DelGroupMsgContent(tt.args.mid)
			if (err != nil) != tt.wantErr {
				t.Errorf("DelGroupMsgContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDao_GetSpecifyRecord(t *testing.T) {
	type fields struct {
		appId          string
		log            zerolog.Logger
		conn           *mysql.MysqlConn
		redis          *redis.Pool
		groupRPCClient group.GroupClient
		pusherCli      pusher.PusherClient
	}
	type args struct {
		mid int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.MsgContent
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				log:   testLog,
				conn:  testConnRecord,
				redis: testRedis,
			},
			args: args{
				mid: 112919335671959552,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				log:   testLog,
				conn:  testConnRecord,
				redis: testRedis,
			}
			got, err := d.GetSpecifyRecord(tt.args.mid)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(got)
		})
	}
}

func TestDao_GetSpecifyGroupRecord(t *testing.T) {
	type fields struct {
		appId          string
		log            zerolog.Logger
		conn           *mysql.MysqlConn
		redis          *redis.Pool
		groupRPCClient group.GroupClient
		pusherCli      pusher.PusherClient
	}
	type args struct {
		mid int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.MsgContent
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				log:   testLog,
				conn:  testConnRecord,
				redis: testRedis,
			},
			args: args{
				mid: 130687370197471232,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				log:   testLog,
				conn:  testConnRecord,
				redis: testRedis,
			}
			got, err := d.GetSpecifyGroupRecord(tt.args.mid)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(got)
		})
	}
}
