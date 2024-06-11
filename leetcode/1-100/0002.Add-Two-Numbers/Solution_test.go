package Solution

import (
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 *ListNode
		input2 *ListNode
		expect *ListNode
	}{
		{
			"Test Case ",
			MarshalSliceToListNode([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
			MarshalSliceToListNode([]int{5, 6, 4}),
			MarshalSliceToListNode([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+strconv.Itoa(i), func(t *testing.T) {
			got := addTwoNumbers2(c.input1, c.input2)
			if !isEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.input1)
			}
		})
	}
}
