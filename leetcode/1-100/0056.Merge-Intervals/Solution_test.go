package Solution

import (
	"reflect"
	"testing"
)

func Test_merge2(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				[][]int{
					{1, 4},
					{5, 6},
				},
			},
			want: [][]int{
				{1, 4},
				{5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge2(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge2() = %v, want %v", got, tt.want)
			}
		})
	}
}
