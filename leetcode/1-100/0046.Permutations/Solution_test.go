package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect [][]int
	}{
		{"TestCase", []int{1, 1, 2}, [][]int{
			{1, 1, 2},
			{1, 2, 1},
			{1, 1, 2},
			{1, 2, 1},
			{2, 1, 1},
			{2, 1, 1},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := permute(c.inputs)
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

func TestDFS(t *testing.T) {
	//fmt.Println(permute3([]int{1, 2, 3}))
}
