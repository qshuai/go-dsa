package leetcode

import "github.com/qshuai/go-dsa/types"

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
//
// Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
// Return the linked list sorted as well.
//
// Examples:
// Input: head = [1,1,2]
// Output: [1,2]
// Input: head = [1,1,2,3,3]
// Output: [1,2,3]
//
// Constraints:
// The number of nodes in the list is in the range [0, 300].
// -100 <= Node.val <= 100
// The list is guaranteed to be sorted in ascending order.
func deleteDuplicates[T comparable](head *types.ListNode[T]) *types.ListNode[T] {
	res := head

	for head != nil && head.Next != nil {
		if head.Value == head.Next.Value {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	return res
}
