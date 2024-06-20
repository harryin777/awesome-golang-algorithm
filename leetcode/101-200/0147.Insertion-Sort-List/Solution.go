package Solution

func Solution(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	newHead := new(ListNode)
	newHead.Next = &ListNode{head.Val, nil}

	head = head.Next
	for head != nil {
		prev, current := newHead, newHead.Next
		for current != nil {
			if head.Val < current.Val {
				prev.Next = &ListNode{head.Val, current}
				break
			}
			prev, current = prev.Next, current.Next
		}
		if nil == current && head.Val >= prev.Val {
			prev.Next = &ListNode{head.Val, nil}
		}
		head = head.Next
	}
	return newHead.Next
}

func insertionSortList(head *ListNode) *ListNode {
	dummy := new(ListNode)
	for head != nil {
		next, p := head.Next, dummy

		for p.Next != nil && p.Next.Val <= head.Val {
			p = p.Next
		}

		head.Next = p.Next
		p.Next = head
		head = next
	}
	return dummy.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList2(head *ListNode) *ListNode {
	arr := make([]*ListNode, 0, 10)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	for i := 0; i < len(arr); i++ {
		min := arr[i]
		var j int
		for j = i; j > 0; j-- {
			if min.Val <= arr[j-1].Val {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		arr[j] = min
	}

	res := &ListNode{}
	tmp := res
	for index, val := range arr {
		tmp.Next = val
		tmp = tmp.Next
		if index == len(arr)-1 {
			tmp.Next = nil
		}
	}

	return res.Next

}
