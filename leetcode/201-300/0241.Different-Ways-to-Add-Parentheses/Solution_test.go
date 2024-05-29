package Solution

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect []int
	}{
		{"TestCase", "2-1-1", []int{0, 2}},
		{"TestCase", "2*3-4*5", []int{-34, -14, -10, -10, 10}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := diffWaysToCompute(c.inputs)
			if len(got) != len(c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v", c.expect, got, c.inputs)
			}
			m := make(map[int]int)
			for v := range got {
				m[v]++
			}
			for v := range c.expect {
				if _, ok := m[v]; !ok {
					t.Fatalf("expected: %v, but got: %v, with inputs: %v", c.expect, got, c.inputs)
				}
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

func Test_diffWaysToCompute2(t *testing.T) {
	fmt.Println(diffWaysToCompute2("2-1-1"))
}
