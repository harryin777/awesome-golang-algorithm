package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeListNode(x []int) *ListNode {
	list := &ListNode{}
	head := list

	list.Val = x[0]
	for i := 1; i < len(x); i++ {
		list.Next = &ListNode{}
		list = list.Next
		list.Val = x[i]
	}
	return head
}
