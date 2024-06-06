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
		s1     string
		n1     int
		s2     string
		n2     int
		expect int
	}{
		//{"TestCase", "acb", 4, "ab", 2, 2},
		{"TestCase", "aaa", 3, "aa", 1, 4},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := getMaxRepetitions(c.s1, c.n1, c.s2, c.n2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with Name: %v",
					c.expect, got, c.name)
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
