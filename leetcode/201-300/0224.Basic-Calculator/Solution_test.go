package Solution

import (
	"testing"
)

func Test_calculate2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				"3+2*(1+1)",
			},
			want: 7,
		},
		{
			name: "t2",
			args: args{
				"2+(1)",
			},
			want: 3,
		},
		{
			name: "t2",
			args: args{
				"2-4-(8+2-6+(8+4-(1)+8-10))",
			},
			want: -15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate2(tt.args.s); got != tt.want {
				t.Errorf("calculate2() = %v, want %v", got, tt.want)
			}
		})
	}
}
