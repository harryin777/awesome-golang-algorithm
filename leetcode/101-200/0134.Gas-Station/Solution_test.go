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
		inputs struct {
			p1 []int
			p2 []int
		}
		expect int
	}{
		//{
		//	name: "t1",
		//	inputs: struct {
		//		p1 []int
		//		p2 []int
		//	}{
		//		p1: []int{1, 2, 3, 4, 5},
		//		p2: []int{3, 4, 5, 1, 2},
		//	},
		//	expect: 3,
		//},
		{
			name: "t2",
			inputs: struct {
				p1 []int
				p2 []int
			}{
				p1: []int{2, 3, 4},
				p2: []int{3, 4, 3},
			},
			expect: -1,
		},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := canCompleteCircuit(c.inputs.p1, c.inputs.p2)
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
