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
		{"TestCase", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := maxSubArray(c.inputs)
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
		{"TestCase", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := maxSubArray2(c.inputs)
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

func TestMaxSubArray3(t *testing.T) {
	fmt.Println(maxSubArray3([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
