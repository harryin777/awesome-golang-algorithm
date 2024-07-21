package Solution

func Solution(head *ListNode) {
	if nil == head || nil == head.Next {
		return
	}
	mid := findMid(head)
	first, second := head, reverseList(mid.Next)
	mid.Next = nil

	newHead := new(ListNode)
	current := newHead
	for first != nil || second != nil {
		if first != nil {
			current.Next = &ListNode{first.Val, nil}
			current = current.Next
			first = first.Next
		}
		if second != nil {
			current.Next = &ListNode{second.Val, nil}
			current = current.Next
			second = second.Next
		}
	}
	head.Next = newHead.Next.Next
}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	newHead := head
	for head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = newHead
		newHead = next
	}
	return newHead
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	arr := make([]*ListNode, 0, 10)
	tmp2 := head
	for tmp2 != nil {
		arr = append(arr, tmp2)
		tmp2 = tmp2.Next
	}

	tmp := head
	l, r := 1, len(arr)-1
	for l <= r {
		tmp.Next = arr[r]
		r--
		tmp = tmp.Next
		if r < l {
			break
		}
		tmp = tmp.Next
		tmp.Next = arr[l]
		l++
		tmp = tmp.Next
	}
	tmp.Next = nil

}
