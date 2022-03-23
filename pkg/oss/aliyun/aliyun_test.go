package aliyun

import (
	"encoding/json"
	"testing"

	"github.com/txchat/dtalk/pkg/oss"
)

var (
	AccessKeyId     = "LTAI5tHhExpCkBhMgc9A1Apz"
	AccessKeySecret = "vb6VaS7ir1aPuN9dFh4e3HkJyzL8iD"
	role            = `{
"roleSessionName": "otc-test",
"roleArn": "acs:ram::1483023416825300:role/otc-test"
}`
	policy          = ``
	RegionId        = "cn-hangzhou"
	DurationSeconds = 3600
	Bucket          = "chy-otc-chat"
	EndPoint        = "oss-cn-shanghai.aliyuncs.com"
)

func TestAliyun_AssumeRole(t *testing.T) {
	type fields struct {
		cfg *oss.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    *oss.AssumeRoleResp
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				cfg: &oss.Config{
					RegionId:        RegionId,
					Bucket:          Bucket,
					EndPoint:        EndPoint,
					AccessKeyId:     AccessKeyId,
					AccessKeySecret: AccessKeySecret,
					Role:            role,
					Policy:          policy,
					DurationSeconds: DurationSeconds,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ali := &Aliyun{
				cfg: tt.fields.cfg,
			}
			got, err := ali.AssumeRole()
			if (err != nil) != tt.wantErr {
				t.Errorf("AssumeRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			data, err := json.Marshal(got)
			if err != nil {
				t.Errorf("Marshal result error = %v", err)
				return
			}
			t.Log("got:", string(data))
		})
	}
}
