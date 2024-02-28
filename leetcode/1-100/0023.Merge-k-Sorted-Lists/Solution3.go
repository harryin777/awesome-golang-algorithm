package Solution

import "container/heap"

func mergeKLists3(lists []*ListNode) *ListNode {
	pq := make(ListNodeQueue, len(lists))
	for i := range lists {
		pq[i] = lists[i]
	}
	heap.Init(&pq)
	if pq.Len() < 1 {
		return nil
	}

	result := heap.Pop(&pq).(*ListNode)
	if result == nil {
		return result
	}
	if result.Next != nil {
		heap.Push(&pq, result.Next)
	}
	tail := result
	for pq.Len() > 0 {
		tail.Next = heap.Pop(&pq).(*ListNode)
		tail = tail.Next
		if tail == nil {
			break
		}
		if tail.Next != nil {
			heap.Push(&pq, tail.Next)
		}
	}

	return result
}

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{
		-1,
		nil,
	}
	tmp := head

	goThrough := true
	for goThrough {
		goThrough = false
		index := 0
		currentVal := int(^uint(0) >> 1)
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				goThrough = true
				if lists[i].Val < currentVal {
					currentVal = lists[i].Val
					index = i
				}
			}
		}
		if !goThrough {
			break
		}
		lists[index] = lists[index].Next
		tmp.Next = &ListNode{
			currentVal,
			nil,
		}
		tmp = tmp.Next
	}

	return head.Next
}
