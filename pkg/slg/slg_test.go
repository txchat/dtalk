package slg

import (
	"reflect"
	"testing"
)

func TestClient_LoadGroupPermission(t *testing.T) {
	type fields struct {
		client BackendRPCClient
	}
	type args struct {
		ucs []*UserCondition
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
				client: NewHTTPClient("", ""),
			},
			args: args{
				ucs: []*UserCondition{
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
			c := &Client{
				client: tt.fields.client,
			}
			got, err := c.LoadGroupPermission(tt.args.ucs)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadGroupPermission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadGroupPermission() got = %v, want %v", got, tt.want)
			}
		})
	}
}
