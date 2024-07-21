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
		inputs int
		expect []string
	}{
		{"1 test 1", 3, []string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		}},
		{
			name:   "t2",
			inputs: 1,
			expect: []string{"()"},
		},
		{
			name:   "t3",
			inputs: 2,
			expect: []string{"()()", "(())"},
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := generateParenthesis3(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}

func Test1(t *testing.T) {
	//fmt.Println(validCheck("()()(()))", 1))
	fmt.Println(generateParenthesis2(3))
}
