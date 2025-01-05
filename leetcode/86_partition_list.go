package leetcode

import (
	"golang.org/x/exp/constraints"

	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/partition-list
//
// Given the head of a linked list and a value x, partition it such that all nodes less than x come
// before nodes greater than or equal to x.
// You should preserve the original relative order of the nodes in each of the two partitions.
//
// Examples:
// Input: head = [1,4,3,2,5,2], x = 3
// Output: [1,2,2,4,3,5]
// Input: head = [2,1], x = 2
// Output: [1,2]
//
// Constraints:
// The number of nodes in the list is in the range [0, 200].
// -100 <= Node.val <= 100
// -200 <= x <= 200
func partitionList[T constraints.Ordered](head *types.ListNode[T], x T) *types.ListNode[T] {
	head = &types.ListNode[T]{Next: head}

	prev, cursor := head, head.Next
	var tail *types.ListNode[T]
	for cursor != nil {
		if cursor.Value < x {
			if prev.Next == cursor {
				prev = cursor
				cursor = cursor.Next
			} else {
				tmp := prev.Next
				prev.Next = cursor

				tmp2 := cursor.Next
				cursor.Next = tmp
				tail.Next = tmp2

				prev = cursor
				cursor = tmp2
			}
		} else {
			tail = cursor
			cursor = cursor.Next
		}
	}

	return head.Next
}
