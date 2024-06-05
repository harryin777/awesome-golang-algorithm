package Solution

func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head
	count := 0
	for curr != nil && count != k {
		curr = curr.Next
		count++
	}

	if count == k {
		curr = reverseKGroup(curr, k)
		for count > 0 {
			tmp := head.Next
			head.Next = curr
			curr = head
			head = tmp
			count--
		}
		head = curr
	}
	return head
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	newRoot := &ListNode{0, nil}
	tmp := newRoot
	vals := make([]int, 0, k)
	pace := 0
	for head != nil {
		pace++
		vals = append(vals, head.Val)
		if pace == k {
			tmp.Next = reverse(vals)
			tmp = tmp.Next
			for tmp.Next != nil {
				tmp = tmp.Next
			}
			vals = vals[:0]
			pace = 0
		}
		head = head.Next
	}
	if len(vals) != 0 {
		tmp.Next = assemble(vals)
	}

	return newRoot.Next
}

func assemble(vals []int) *ListNode {
	res := &ListNode{0, nil}
	tmp := res
	for i := 0; i < len(vals); i++ {
		if tmp.Next == nil && i != len(vals)-1 {
			tmp.Next = &ListNode{0, nil}
		}
		tmp.Val = vals[i]
		tmp = tmp.Next
	}
	return res
}

func reverse(vals []int) *ListNode {
	res := &ListNode{0, nil}
	tmp := res
	for i := len(vals) - 1; i >= 0; i-- {
		if tmp.Next == nil && i != 0 {
			tmp.Next = &ListNode{0, nil}
		}
		tmp.Val = vals[i]
		tmp = tmp.Next
	}
	return res
}
