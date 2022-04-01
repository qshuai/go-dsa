package leetcode

import (
	"github.com/qshuai/go-dsa/types"
)

// swapNodes You are given the head of a linked list, and an integer k.
// Return the head of the linked list after swapping the values of the kth node
// from the beginning and the kth node from the end (the list is 1-indexed).
func swapNodes(head *types.ListNode, k int) *types.ListNode {
	slow, fast := head, head
	for i := 1; i < k; i++ {
		fast = fast.Next
	}

	prev := fast
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	prev.Value, slow.Value = slow.Value, prev.Value
	return head
}
