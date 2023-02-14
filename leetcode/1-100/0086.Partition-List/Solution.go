package Solution

func Solution(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	before, after := &ListNode{}, &ListNode{}
	bh, ah := before, after
	for ; head != nil; head = head.Next {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
	}
	after.Next, before.Next = nil, ah.Next
	return bh.Next
}

func partition(head *ListNode, x int) *ListNode {
	lAarr := make([]int, 0, 10)
	rAarr := make([]int, 0, 10)
	for head != nil {
		if head.Val < x {
			lAarr = append(lAarr, head.Val)
		} else {
			rAarr = append(rAarr, head.Val)
		}
	}
	for i := len(rAarr) - 1; i >= 0; i-- {
		lAarr = append(lAarr, rAarr[i])
	}
	res := ListNode{
		Val:  0,
		Next: nil,
	}
	res.Next = &ListNode{
		Val:  lAarr[0],
		Next: nil,
	}
	tmp := res
	for i := len(lAarr); i >= 0; i++ {
		tmp.Next = &ListNode{
			Val:  lAarr[i],
			Next: nil,
		}
		tmp = tmp.Next
	}
	return head
}
