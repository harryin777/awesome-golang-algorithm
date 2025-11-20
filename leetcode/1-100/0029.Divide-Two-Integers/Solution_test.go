package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"1 test 1", []int{10, -3}, -3},
		//{"1 test 2", []int{7, -3}, -2},
		//{"1 test 2", []int{-2147483648, 2}, -1073741824},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := divide2(c.inputs[0], c.inputs[1])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
			//ret2 := divide2(c.inputs[0], c.inputs[1])
			//if !reflect.DeepEqual(ret2, c.expect) {
			//	t.Fatalf("expected: %v, but got: %v, with inputs: %v",
			//		c.expect, ret, c.inputs)
			//}
		})
	}
}
