package Solution

import "awesome-golang-algorithm/utils"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fastNode := head
	slowNode := head
	for n > 0 {
		fastNode = fastNode.Next
		n--
	}

	if fastNode != nil {
		for fastNode.Next != nil {
			fastNode = fastNode.Next
			slowNode = slowNode.Next
		}
		slowNode.Next = slowNode.Next.Next

	} else {
		head = head.Next
	}
	return head
}

func removeNthFromEnd2(head *utils.ListNode, n int) *utils.ListNode {
	l1 := head
	l2 := head
	for i := 0; i < n; i++ {
		l1 = l1.Next
	}
	if l1 == nil {
		head = head.Next
		return head
	} else {
		for l1.Next != nil {
			l2 = l2.Next
			l1 = l1.Next
		}
		l2.Next = l2.Next.Next
	}

	return head
}
