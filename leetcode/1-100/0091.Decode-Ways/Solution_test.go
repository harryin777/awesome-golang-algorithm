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
		inputs string
		expect int
	}{
		{"TestCacse 1", "12", 2},
		{"TestCacse 1", "226", 3},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := numDecodings(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
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

func Test_numDecodings2(t *testing.T) {
	fmt.Println(numDecodings2("06"))
}
