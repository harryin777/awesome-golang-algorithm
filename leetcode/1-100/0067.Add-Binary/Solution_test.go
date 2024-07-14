package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []string
		expect string
	}{
		{"1 test 1", []string{"11", "1"}, "100"},
		{"2 test 2", []string{"1010", "1011"}, "10101"},
		{"3 test 3", []string{"101011", "101101"}, "1011000"},
		{"3 test 3", []string{"101111", "10"}, "110001"},
		{"3 test 3", []string{"0", "1"}, "1"},
		{"3 test 3", []string{"100", "110010"}, "110110"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := addBinary3(c.inputs[0], c.inputs[1])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
