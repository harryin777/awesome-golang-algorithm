package Solution

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCacse 1", []int{1, 2, 3, 1}, 4},
		{"TestCacse 1", []int{2, 7, 9, 3, 1}, 12},
		{"TestCacse 1", []int{5, 2, 6, 7, 3, 1}, 14},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := rob(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCacse 1", []int{1, 2, 3, 1}, 4},
		{"TestCacse 1", []int{2, 7, 9, 3, 1}, 12},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := rob2(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestSolution3(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCacse 1", []int{1, 2, 3, 1}, 4},
		{"TestCacse 1", []int{2, 7, 9, 3, 1}, 12},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := rob3(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestSolution4(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCacse", []int{}, 0},
		{"TestCacse", []int{3}, 3},
		{"TestCacse", []int{1, 2, 3, 1}, 4},
		{"TestCacse", []int{2, 7, 9, 3, 1}, 12},
		{"TestCacse", []int{5, 2, 6, 7, 3, 1}, 14},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := rob4(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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

func Test1(t *testing.T) {
	//fmt.Println(rob11([]int{2, 7, 9, 3, 1}))
	//fmt.Println(rob11([]int{1, 1}))
	fmt.Println(rob11([]int{1, 3, 1, 3, 100}))
	fmt.Println(rob12([]int{1, 3, 1, 3, 100}))
}
