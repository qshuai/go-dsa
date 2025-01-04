package leetcode

import (
	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/swap-nodes-in-pairs/

// swapPairs Given a linked list, swap every two adjacent nodes and return its head. You must solve
// the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)
func swapPairs[T any](head *types.ListNode[T]) *types.ListNode[T] {
	if head == nil || head.Next == nil {
		return head
	}

	cursor := head
	head = cursor.Next
	var prev *types.ListNode[T]
	for cursor != nil && cursor.Next != nil {
		tmp := cursor.Next.Next
		cursor.Next.Next = cursor
		if prev == nil {
			prev = cursor
		} else {
			prev.Next = cursor.Next
			prev = cursor
		}

		cursor.Next = tmp
		cursor = tmp
	}

	return head
}
