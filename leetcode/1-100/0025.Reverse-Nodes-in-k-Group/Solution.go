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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup3(head *ListNode, k int) *ListNode {
	arr := make([][]*ListNode, k)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]*ListNode, 0, 10)
	}
	tmp1 := head
	tmp3 := head
	total := 0
	for tmp3 != nil {
		tmp3 = tmp3.Next
		total++
	}

	count := 0
	reminder := total / k
	for count < reminder*k {
		arr[count%k] = append(arr[count%k], &ListNode{
			Val: tmp1.Val,
		})
		tmp1 = tmp1.Next
		count++
	}

	newRoot := &ListNode{
		Val: 0,
	}
	tmp2 := newRoot
	count = 0
	for count < reminder*k {
		tmp2.Next = arr[k-count%k-1][0]
		arr[k-count%k-1] = arr[k-count%k-1][1:]
		tmp2 = tmp2.Next
		count++
	}
	if tmp1 != nil {
		tmp2.Next = tmp1
	}

	return newRoot.Next
}
