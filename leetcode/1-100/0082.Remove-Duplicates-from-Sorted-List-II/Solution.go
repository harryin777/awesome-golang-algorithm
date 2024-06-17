package Solution

import "sort"

/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func Solution(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	p := new(ListNode)
	p.Next = head
	prev, current := p, p.Next
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			for current.Next != nil && current.Val == current.Next.Val {
				current = current.Next
			}
			current = current.Next
			prev.Next = current
		} else {
			prev = prev.Next
			current = current.Next
		}
	}
	return p.Next
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tmp1 := head
	res := &ListNode{
		Val:  0,
		Next: head,
	}
	tmp2 := res
	c := make(map[int]int)
	for tmp1 != nil {
		c[tmp1.Val]++
		tmp1 = tmp1.Next
	}
	arr := make([]int, 0, len(c))
	for key, count := range c {
		if count > 1 {
			continue
		}
		arr = append(arr, key)
	}
	sort.Ints(arr)
	if len(arr) == 0 {
		return nil
	}
	for _, val := range arr {
		tmp2.Next = &ListNode{
			Val: val,
		}
		tmp2 = tmp2.Next
	}

	return res.Next
}
