package slg

import (
	"reflect"
	"testing"
)

func TestHTTPClient_GroupPermissionVerification(t *testing.T) {
	type fields struct {
		url  string
		salt string
	}
	type args struct {
		conditions []*UserCondition
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    GroupPermission
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				url:  "",
				salt: "",
			},
			args: args{
				conditions: []*UserCondition{
					{
						UID:        "12quUKnXMaHfxYvUB9bePW3k4eSj6H4ADo",
						HandleType: 0,
						Conditions: []string{"2bb5d007c391489a9a55eeb90372344d"},
					},
				},
			},
			want: GroupPermission{
				"12quUKnXMaHfxYvUB9bePW3k4eSj6H4ADo": true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HTTPClient{
				host: tt.fields.url,
				salt: tt.fields.salt,
			}
			got, err := c.GroupPermissionVerification(tt.args.conditions)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupPermissionVerification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupPermissionVerification() got = %v, want %v", got, tt.want)
			}
		})
	}
}
