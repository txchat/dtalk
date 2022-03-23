package dao

import (
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog"
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/device/model"
	"gopkg.in/Shopify/sarama.v1"
	"testing"
	"time"
)

func TestDao_GetAllDevices(t *testing.T) {
	type fields struct {
		log        zerolog.Logger
		redis      *redis.Pool
		offPushPub sarama.SyncProducer
	}
	type args struct {
		uid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Device
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				log:        testLog,
				redis:      testRedis,
				offPushPub: nil,
			},
			args: args{
				uid: "test-uid",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				log:   tt.fields.log,
				redis: tt.fields.redis,
			}
			got, err := d.GetAllDevices(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllDevices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, info := range got {
				t.Log("got", info)
			}
		})
	}
}

func TestDao_AddDeviceInfo(t *testing.T) {
	type fields struct {
		log   zerolog.Logger
		redis *redis.Pool
	}
	type args struct {
		device *model.Device
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
				log:   testLog,
				redis: testRedis,
			},
			args: args{
				device: &model.Device{
					Uid:         "test-uid",
					ConnectId:   "test-conn",
					DeviceType:  int32(0),
					Username:    "test-name",
					DeviceToken: "test-deviceToken",
					IsEnabled:   false,
					AddTime:     uint64(util.TimeNowUnixNano() / int64(time.Millisecond)),
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
			if err := d.AddDeviceInfo(tt.args.device); (err != nil) != tt.wantErr {
				t.Errorf("AddDeviceInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDao_EnableDeviceInfo(t *testing.T) {
	type fields struct {
		log   zerolog.Logger
		redis *redis.Pool
	}
	type args struct {
		uid    string
		connId string
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
				log:   testLog,
				redis: testRedis,
			},
			args: args{
				uid:    "test-uid",
				connId: "test-conn",
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
			if err := d.EnableDevice(tt.args.uid, tt.args.connId); (err != nil) != tt.wantErr {
				t.Errorf("EnableDevice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDao_setExpire(t *testing.T) {
	type fields struct {
		log   zerolog.Logger
		redis *redis.Pool
	}
	type args struct {
		device *model.Device
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
				log:   testLog,
				redis: testRedis,
			},
			args: args{
				device: &model.Device{
					Uid:         "115Mkz6vxW61Phj3UM3MPft1mCLrEiXUrQ",
					ConnectId:   "f471d531-2f77-452c-b80b-47b232aeb772",
					DeviceType:  0,
					Username:    "dld",
					DeviceToken: "",
					IsEnabled:   false,
					AddTime:     1631764139040,
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
			if err := d.setExpire(tt.args.device); (err != nil) != tt.wantErr {
				t.Errorf("setExpire() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
