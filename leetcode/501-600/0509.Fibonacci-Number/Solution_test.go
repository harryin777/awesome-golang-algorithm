package Solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{2},
			want: 1,
		},
		{
			name: "t2",
			args: args{
				4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, fib(tt.args.n), "fib(%v)", tt.args.n)
		})
	}
}
