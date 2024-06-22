package Solution

import (
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	type args struct {
		columnTitle string
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
				"A",
			},
			want: 1,
		},
		{
			name: "t2",
			args: args{
				"AB",
			},
			want: 28,
		},
		{
			name: "t3",
			args: args{
				"ZY",
			},
			want: 701,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.columnTitle); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
