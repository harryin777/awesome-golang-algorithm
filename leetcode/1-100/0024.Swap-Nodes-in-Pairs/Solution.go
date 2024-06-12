package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs_1(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}
func swapPairs_2(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}

func swapPairs2(head *ListNode) *ListNode {
	count := 1
	res := &ListNode{}
	tmp3 := res
	son1 := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp1 := son1
	son2 := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp2 := son2
	for head != nil {
		if count%2 == 1 {
			tmp1.Next = head
			tmp1 = tmp1.Next
		} else {
			tmp2.Next = head
			tmp2 = tmp2.Next
		}

		head = head.Next
		count++
	}
	tmp1.Next = nil
	tmp2.Next = nil
	for son1.Next != nil && son2.Next != nil {
		s2 := son2.Next
		tmp3.Next = s2
		tmp3 = tmp3.Next
		s1 := son1.Next
		tmp3.Next = s1
		tmp3 = tmp3.Next

	}
	if son1.Next != nil {
		tmp3.Next = son1.Next
	}
	if son2.Next != nil {
		tmp3.Next = son2.Next
	}

	return res.Next
}

func swapPairs3(head *ListNode) *ListNode {
	count := 1
	res := &ListNode{}
	tmp3 := res
	son1 := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp1 := son1
	son2 := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp2 := son2

	// Split the list into two parts
	for head != nil {
		if count%2 == 1 {
			tmp1.Next = head
			tmp1 = tmp1.Next
		} else {
			tmp2.Next = head
			tmp2 = tmp2.Next
		}
		head = head.Next
		count++
	}

	// Terminate the two split lists
	tmp1.Next = nil
	tmp2.Next = nil

	// Merge the two lists back in pairs
	son1 = son1.Next
	son2 = son2.Next
	tmp3 = res
	for son1 != nil && son2 != nil {
		tmp3.Next = son2
		son2 = son2.Next
		tmp3 = tmp3.Next
		tmp3.Next = son1
		son1 = son1.Next
		tmp3 = tmp3.Next
	}

	// Attach any remaining nodes
	if son1 != nil {
		tmp3.Next = son1
	}
	if son2 != nil {
		tmp3.Next = son2
	}

	return res.Next
}
