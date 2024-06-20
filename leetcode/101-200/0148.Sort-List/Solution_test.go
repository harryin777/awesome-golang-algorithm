package Solution

import (
	"reflect"
	"testing"
)

// 压力测试
func Test_quickSort2(t *testing.T) {
	type args struct {
		arr   []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				[]int{4, 2, 1, 3},
				0, 3,
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort2(tt.args.arr, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort2() = %v, want %v", got, tt.want)
			}
		})
	}
}
