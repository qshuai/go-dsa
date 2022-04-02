package leetcode

import (
	"github.com/qshuai/go-dsa/types"
)

func swapPairs(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cursor := head
	head = cursor.Next
	var prev *types.ListNode
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
