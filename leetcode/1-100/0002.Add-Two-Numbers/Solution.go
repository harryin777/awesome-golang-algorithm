package Solution

import (
	"math"
	"strconv"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{Val: 0, Next: nil}
	n1, n2, tmp := l1, l2, node
	sum := 0

	for n1 != nil || n2 != nil {
		sum /= 10
		if n1 != nil {
			sum += n1.Val
			n1 = n1.Next

		}
		if n2 != nil {
			sum += n2.Val
			n2 = n2.Next
		}
		tmp.Next = &ListNode{Val: sum % 10}
		tmp = tmp.Next
	}
	if sum/10 != 0 {
		tmp.Next = &ListNode{Val: 1}
	}
	return node.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	val1 := getNumVal(l1)
	val2 := getNumVal(l2)
	return getListNode(val1 + val2)
}

// 这个方法不行,会数据范围越界
func getNumVal(l *ListNode) int {
	val := 0
	count := 0
	for l != nil {
		val += l.Val * int(math.Pow(10, float64(count)))
		count++
		l = l.Next
	}

	return val
}

func getListNode(val int) *ListNode {
	valStr := strconv.Itoa(val)
	res := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := res
	for i := len(valStr) - 1; i >= 0; i-- {
		valInt, _ := strconv.ParseInt(string(valStr[i]), 0, 64)
		tmp.Next = &ListNode{
			Val:  int(valInt),
			Next: nil,
		}
		tmp = tmp.Next

	}

	return res.Next
}

func MarshalSliceToListNode(nums []int) *ListNode {
	head := &ListNode{}
	tmp := head
	for _, v := range nums {
		tmp.Next = &ListNode{Val: v, Next: nil}
		tmp = tmp.Next
	}
	return head.Next
}
