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
		inputs string
		expect []int
	}{
		{"TestCase1", "ababcbacadefegdehijhklij", []int{9, 7, 8}},
		//{"TestCase2", "eccbbbbdec", []int{10}},
		//{"TestCase3", "abc", []int{1, 1, 1}},
		//{"TestCase4", "aabbadec", []int{5, 1, 1, 1}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := partitionLabels(c.inputs)
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
