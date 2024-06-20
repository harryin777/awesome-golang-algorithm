package Solution

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	nextHalfHead := slow.Next
	slow.Next = nil
	l1 := sortList(head)
	l2 := sortList(nextHalfHead)

	dummy := &ListNode{}
	cur := dummy
	for l1 != nil || l2 != nil {
		if l1 == nil || (l2 != nil && l1.Val >= l2.Val) {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}

		cur = cur.Next
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
func sortList2(head *ListNode) *ListNode {
	arr := make([]*ListNode, 0, 10)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	arr = mergeSort1(arr)

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

func mergeSort1(arr []*ListNode) []*ListNode {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) >> 1
	m1 := mergeSort1(arr[:mid])
	m2 := mergeSort1((arr[mid:]))
	return mergeSort2(m1, m2)
}

func mergeSort2(left, right []*ListNode) []*ListNode {
	i, j := 0, 0
	res := make([]*ListNode, 0, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i].Val <= right[j].Val {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)

	return res
}

func quickSort(arr []*ListNode, left, right int) []*ListNode {
	if left > right {
		return arr
	}

	i, j := left, right
	tmp := arr[i]
	for i < j {
		for i < j && tmp.Val <= arr[j].Val {
			j--
		}
		for i < j && tmp.Val >= arr[i].Val {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[left] = arr[i]
	arr[i] = tmp

	arr = quickSort(arr, i+1, right)
	arr = quickSort(arr, left, i-1)
	return arr
}
