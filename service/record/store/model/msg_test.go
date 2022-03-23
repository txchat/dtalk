package model

import (
	xproto "github.com/txchat/imparse/proto"
	"testing"
)

func TestParseSource(t *testing.T) {
	type args struct {
		m *xproto.Common
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "",
			args: args{
				m: &xproto.Common{
					ChannelType: 0,
					Mid:         0,
					Seq:         "",
					From:        "",
					Target:      "",
					MsgType:     0,
					Msg:         nil,
					Datetime:    0,
					Source:      nil,
				},
			},
			want: nil,
		}, {
			name: "",
			args: args{
				m: &xproto.Common{
					ChannelType: 0,
					Mid:         0,
					Seq:         "",
					From:        "",
					Target:      "",
					MsgType:     0,
					Msg:         nil,
					Datetime:    0,
					Source:      &xproto.Source{},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(string(ParseSource(tt.args.m)))
		})
	}
}
