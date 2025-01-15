package leetcode

import (
	"golang.org/x/exp/constraints"

	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/merge-k-sorted-lists/description/
//
// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
// Merge all the linked-lists into one sorted linked-list and return it.
//
// Example 1:
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
//  1->4->5,
//  1->3->4,
//  2->6
// ]
// merging them into one sorted list:
// 1->1->2->3->4->4->5->6
//
// Example 2:
// Input: lists = []
// Output: []
//
// Example 3:
// Input: lists = [[]]
// Output: []
//
//
// Constraints:
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] is sorted in ascending order.
// The sum of lists[i].length will not exceed 10^4.

func mergeKLists[T constraints.Ordered](lists []*types.ListNode[T]) *types.ListNode[T] {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeSortedList(lists[0], lists[1])
	}

	a := mergeKLists(lists[:len(lists)/2])
	b := mergeKLists(lists[len(lists)/2:])
	return mergeSortedList(a, b)
}

func mergeSortedList[T constraints.Ordered](a, b *types.ListNode[T]) *types.ListNode[T] {
	res := &types.ListNode[T]{}
	head := res
	for a != nil && b != nil {
		if a.Value <= b.Value {
			head.Next = a
			a = a.Next
		} else {
			head.Next = b
			b = b.Next
		}

		head = head.Next
	}
	if a != nil {
		head.Next = a
	}
	if b != nil {
		head.Next = b
	}

	return res.Next
}
