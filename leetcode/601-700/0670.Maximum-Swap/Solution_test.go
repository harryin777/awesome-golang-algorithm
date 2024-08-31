package Solution

import (
	"testing"
)

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
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
				98368,
			},
			want: 98863,
		},
		{
			name: "t2",
			args: args{
				987654,
			},
			want: 987654,
		},
		{
			name: "t3",
			args: args{
				1993,
			},
			want: 9913,
		},
		{
			name: "t4",
			args: args{
				99901,
			},
			want: 99910,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
