package Solution

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	// 测试用例
	cases := []struct {
		name   string
		inputs [][]int
		expect int
	}{
		{
			"TestCases 1",
			[][]int{
				{1, 1, 0},
				{1, 1, 1},
				{0, 1, 1},
			},
			1,
		},
		{
			name: "t2",
			inputs: [][]int{
				{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1},
			},
			expect: 1,
		},
	}

	// 开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := findCircleNum3(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v", c.expect, got, c.inputs)
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

func TestFindCircleNum2(t *testing.T) {
	fmt.Println(findCircleNum2([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}
