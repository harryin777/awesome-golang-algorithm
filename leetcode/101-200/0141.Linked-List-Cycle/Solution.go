package Solution

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 设置时间
func hasCycle1(head *ListNode) bool {
	start := time.Now()

	for time.Since(start) < 1000 {
		if head == nil {
			return false
		}
		// fmt.Println(head)

		head = head.Next
	}
	return true
}

// HashMap
func hasCycle2(head *ListNode) bool {
	listMap := make(map[string]int)
	for head != nil {
		if listMap[fmt.Sprintf("%v", &head)] > 1 {
			return true
		}

		head = head.Next
	}
	return false
}

// 快慢指针
func hasCycle3(head *ListNode) bool {
	quick := head
	slow := head

	for quick.Next.Next != nil && slow.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
		if slow == quick {
			return true
		}
	}
	return false
}

func MakeListNode(x []int) *ListNode {
	list := &ListNode{}
	head := list

	list.Val = x[0]
	for i := 1; i < len(x); i++ {
		list.Next = &ListNode{}
		list = list.Next
		list.Val = x[i]
	}
	return head
}
