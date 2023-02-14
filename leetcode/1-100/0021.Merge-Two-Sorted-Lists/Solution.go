package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmp := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	} else {
		tmp.Next = l2
	}

	return head.Next
}

func mergeTwoLists1(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	n := &ListNode{}
	if l1.Val < l2.Val {
		n.Val = l1.Val
		n.Next = mergeTwoLists1(l1.Next, l2)
	} else {
		n.Val = l2.Val
		n.Next = mergeTwoLists1(l1, l2.Next)
	}
	return n
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1 == nil && list2 == nil {
		return nil
	}

	p := ListNode{
		Val:  -1,
		Next: nil,
	}

	tmp := &p

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			tmp.Next = list2
			list2 = list2.Next
		} else {
			tmp.Next = list1
			list1 = list1.Next
		}
		tmp = tmp.Next
	}

	if list1 == nil {
		tmp.Next = list2
	}
	if list2 == nil {
		tmp.Next = list1
	}
	return p.Next
}
