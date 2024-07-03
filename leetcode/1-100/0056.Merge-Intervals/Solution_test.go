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
		{
			name: "t2",
			args: args{
				[][]int{
					{1, 3},
				},
			},
			want: [][]int{
				{1, 3},
			},
		},
		{
			name: "t3",
			args: args{
				[][]int{
					{1, 4},
					{0, 2},
					{3, 5},
				},
			},
			want: [][]int{
				{0, 5},
			},
		},
		{
			name: "t4",
			args: args{
				[][]int{
					{2, 3},
					{4, 5},
					{6, 7},
					{8, 9},
					{1, 10},
				},
			},
			want: [][]int{
				{1, 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge3(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge2() = %v, want %v", got, tt.want)
			}
		})
	}
}
