package Solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_coinChange10(t *testing.T) {
	type args struct {
		coins  []int
		amount int
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
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, coinChange1(tt.args.coins, tt.args.amount), "coinChange10(%v, %v)", tt.args.coins, tt.args.amount)
		})
	}
}
