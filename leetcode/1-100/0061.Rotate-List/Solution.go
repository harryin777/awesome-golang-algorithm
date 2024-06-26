package Solution

func Solution(head *ListNode, k int) *ListNode {
	if nil == head || 0 == k {
		return head
	}
	head, length := cycleList(head)
	if k >= length {
		k = k % length
	}
	for i := 0; i < length-k-1; i++ {
		head = head.Next
	}
	p := head.Next
	head.Next = nil
	return p
}

func cycleList(l *ListNode) (*ListNode, int) {
	head, length := l, 1
	for l.Next != nil {
		l = l.Next
		length++
	}
	l.Next = head
	return head, length
}

func rotateRight(head *ListNode, k int) *ListNode {
	pointer := head
	count := 1
	if head == nil || head.Next == nil {
		return head
	}
	for pointer != nil {
		if pointer.Next == nil {
			pointer.Next = head
			break
		}
		pointer = pointer.Next
		count++
	}
	rem := count - k%count
	for rem > 0 {
		pointer = pointer.Next
		rem--
	}
	ans := pointer.Next
	pointer.Next = nil
	return ans
}

func rotateRight2(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	length := 1
	i := head
	for i.Next != nil {
		i = i.Next
		length++
	}
	route := length - k%length
	if route == length {
		return head
	}
	i.Next = head

	for route > 0 {
		i = i.Next
		route--
	}
	res := i.Next
	i.Next = nil

	return res

}
