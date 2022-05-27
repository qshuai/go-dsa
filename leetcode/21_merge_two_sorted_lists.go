package leetcode

import (
	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/merge-two-sorted-lists/

// mergeSortedList You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists in a one sorted list. The list should be made by splicing together
// the nodes of the first two lists.
//
// Return the head of the merged linked list.
func mergeSortedList(list1 *types.ListNode, list2 *types.ListNode) *types.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// singly linked list with guard
	ret := &types.ListNode{}
	cursor := ret
	for list1 != nil && list2 != nil {
		if list1.Value.(int) <= list2.Value.(int) {
			cursor.Next = list1
			list1 = list1.Next
		} else {
			cursor.Next = list2
			list2 = list2.Next
		}

		cursor = cursor.Next
	}
	if list1 != nil {
		cursor.Next = list1
	}
	if list2 != nil {
		cursor.Next = list2
	}

	return ret.Next
}
