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
		expect int
	}{
		//{"TestCase", []int{0, 0}, 1},
		//{"TestCase", []int{3, 3, 3, 2, 5}, 3},
		//{"TestCase", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
		{"TestCase", []int{1, 7, 4, 9, 2, 5}, 6},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+strconv.Itoa(i), func(t *testing.T) {
			got := wiggleMaxLength2(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
			got = wiggleMaxLength1(c.inputs)
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
