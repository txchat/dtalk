package huaweiyun

import (
	"encoding/json"
	"testing"

	"github.com/txchat/dtalk/pkg/oss"
)

var (
	AccessKeyId     = "VS1SNCU6SM7NSRLFT0HF"
	AccessKeySecret = "pRI6OjPTrc0atS3Do4PFgSOyx7IOnaXXbDf5ZBd2"
	//endPointRegion := "cn-east-3"
	policy = `
{
	"Version": "1.1",
	"Statement": [
		{
			"Action": [
				"obs:object:*"
			],
			"Effect": "Allow"
		}
	]
}`
	RegionId        = "cn-east-3"
	DurationSeconds = 3600
	Bucket          = "chy-cdn"
	EndPoint        = "obs.cn-east-3.myhuaweicloud.com"
)

func TestHuaweiyun_AssumeRole(t *testing.T) {
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
			name: "huaweiyun_getTempSk_test",
			fields: fields{
				cfg: &oss.Config{
					RegionId:        RegionId,
					Bucket:          Bucket,
					EndPoint:        EndPoint,
					AccessKeyId:     AccessKeyId,
					AccessKeySecret: AccessKeySecret,
					Role:            "",
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
			huawei := New(tt.fields.cfg)
			got, err := huawei.AssumeRole()
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
