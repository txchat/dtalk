package debug

import "testing"

func TestGetMockCode(t *testing.T) {
	type args struct {
		mode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 5 length",
			args: args{
				mode: "FzmRandom5",
			},
			want: "11111",
		}, {
			name: "test default length",
			args: args{
				mode: "FzmRandom",
			},
			want: "111111",
		}, {
			name: "test 0 length",
			args: args{
				mode: "FzmRandom0",
			},
			want: "111111",
		}, {
			name: "test 6 length",
			args: args{
				mode: "FzmRandom6",
			},
			want: "111111",
		}, {
			name: "test 4 length",
			args: args{
				mode: "FzmRandom4",
			},
			want: "1111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMockCode(tt.args.mode); got != tt.want {
				t.Errorf("GetMockCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
