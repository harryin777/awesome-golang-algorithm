package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name      string
		flowerbed []int
		n         int
		expect    bool
	}{
		{"TestCase", []int{1, 0, 0, 0, 1}, 1, true},
		{"TestCase", []int{1, 0, 0, 0, 1}, 2, false},
		{"TestCase", []int{1, 0, 1, 0}, 0, true},
		{"TestCase", []int{0}, 1, true},
		{"TestCase", []int{1, 0, 1, 0, 1, 0, 1}, 1, false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.flowerbed, c.n)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.flowerbed, c.n)
			}
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				[]int{1, 0, 0, 0, 1},
				2},
			want: false,
		},
		{
			name: "t2",
			args: args{
				[]int{
					0, 0, 1,
					0, 0, 0,
					0, 1, 0,
					1, 0, 0,
					0, 1, 0,
					0, 1, 0,
					1, 0, 1,
					0, 0, 0,
					1, 0, 1,
					0, 1, 0,
					0, 1, 0,
					0, 0, 0,
					0, 1, 0,
					1, 0, 0,
					0, 1, 0,
					0, 1, 0,
					0, 0, 1,
					0, 0, 1,
					0, 0, 1,
					0, 0, 0,
					1, 0, 0,
					0, 0, 1,
					0, 0, 1,
					0, 0, 0,
					0, 1, 0,
					0, 0, 1,
					0, 1, 0,
					0, 0, 0,
					0, 0},
				17},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
