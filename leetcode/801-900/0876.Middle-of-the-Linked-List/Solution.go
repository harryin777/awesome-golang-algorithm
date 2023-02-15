package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(head *ListNode) *ListNode {
	tort, hare := head, head
	for hare != nil && hare.Next != nil {
		tort = tort.Next
		hare = hare.Next.Next
	}
	return tort
}

func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
