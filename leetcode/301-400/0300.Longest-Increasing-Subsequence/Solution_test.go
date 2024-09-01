package Solution

import (
	"testing"
)

func Test_lengthOfLIS6(t *testing.T) {
	type args struct {
		nums []int
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
				[]int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS1(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS6() = %v, want %v", got, tt.want)
			}
		})
	}
}
