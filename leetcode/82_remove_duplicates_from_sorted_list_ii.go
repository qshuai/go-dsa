package leetcode

import (
	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/

// deleteDuplicates Given the head of a sorted linked list, delete all nodes that have duplicate numbers,
// leaving only distinct numbers from the original list. Return the linked list sorted as well.
//
// Constraints:
// The number of nodes in the list is in the range [0, 300].
// -100 <= TreeNode.val <= 100
// The list is guaranteed to be sorted in ascending order.
func deleteDuplicates(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	root := &types.ListNode{}
	tail := root
	var prev *types.ListNode
	cursor := head
	for cursor != nil {
		if prev != nil && cursor.Value.(int) == prev.Value.(int) {

		} else if cursor.Next != nil && cursor.Value.(int) == cursor.Next.Value.(int) {
			prev = cursor
		} else {
			prev = cursor

			tail.Next = &types.ListNode{Value: cursor.Value}
			tail = tail.Next
		}

		cursor = cursor.Next
	}

	return root.Next
}
