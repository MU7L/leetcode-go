package main

import (
	. "leetcode-go/ds"
)

//go:generate ../../testgen ./solution.go

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	// 1. 快慢指针找中点
	before := findBeforeMiddle(head)
	// 2. 删除节点
	before.Next = before.Next.Next
	return head
}

func findBeforeMiddle(head *ListNode) *ListNode {
	var beforeSlow *ListNode
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return beforeSlow
		}
		fast = fast.Next
		if fast.Next == nil {
			return slow
		}
		fast = fast.Next
		beforeSlow = slow
		slow = slow.Next
	}
	return nil
}
