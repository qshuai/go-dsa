package leetcode

import "github.com/qshuai/go-dsa/types"

// https://leetcode.com/problems/reverse-linked-list/description/
//
// Given the head of a singly linked list, reverse the list, and return the reversed list.
//
// Examples:
// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
// Input: head = [1,2]
// Output: [2,1]
// Input: head = []
// Output: []
//
// Constraints:
// The number of nodes in the list is the range [0, 5000].
// -5000 <= Node.val <= 5000
func reverseList[T any](head *types.ListNode[T]) *types.ListNode[T] {
	if head == nil || head.Next == nil {
		return head
	}

	prev, cur := (*types.ListNode[T])(nil), head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
}
