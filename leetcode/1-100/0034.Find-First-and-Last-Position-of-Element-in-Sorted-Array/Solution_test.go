package Solution

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		expect []int
	}{
		{"TestCacse 1", [][]int{{5, 7, 7, 8, 8, 10}, {8}}, []int{3, 4}},
		{"TestCacse 2", [][]int{{5, 7, 7, 8, 8, 10}, {6}}, []int{-1, -1}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := searchRange(c.inputs[0], c.inputs[1][0])
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
		inputs [][]int
		expect []int
	}{
		{"TestCacse 1", [][]int{{5, 7, 7, 8, 8, 10}, {8}}, []int{3, 4}},
		{"TestCacse 2", [][]int{{5, 7, 7, 8, 8, 10}, {6}}, []int{-1, -1}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := searchRange2(c.inputs[0], c.inputs[1][0])
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}

func Test1(t *testing.T) {
	//fmt.Println(searchRange3([]int{5, 7, 7, 8, 8}, 8))
	//fmt.Println(searchRange3([]int{5, 7, 7, 8, 8, 10}, 8))
	//fmt.Println(searchRange6([]int{5, 7, 7, 8, 8, 10}, 8))
	//fmt.Println(LeftBound([]int{5, 7, 7, 8, 8, 10}, 8))
	//fmt.Println(RightBound([]int{5, 7, 7, 8, 8, 10}, 8))
	//fmt.Println(searchRange4([]int{5, 7, 7, 8, 8}, 8))

	fmt.Println(LeftBound([]int{5, 7, 7, 8, 8, 8}, 10))
	fmt.Println(RightBound([]int{5, 7, 7, 8, 8, 8}, 10))

}

func Test2(t *testing.T) {
	fmt.Println(search([]int{2, 2}, 3))
}
