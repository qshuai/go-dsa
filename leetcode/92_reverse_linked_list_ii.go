package leetcode

import (
	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/reverse-linked-list-ii/

// reverseBetween Given the head of a singly linked list and two integers left and right where left <= right,
// reverse the nodes of the list from position left to position right, and return the reversed list.
//
// Constraints:
// The number of nodes in the list is n.
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
func reverseBetween(head *types.ListNode, left int, right int) *types.ListNode {
	if left == right {
		return head
	}

	mock := &types.ListNode{Next: head}
	prev := mock
	cursor := head
	var lPoint, lPointNext *types.ListNode
	i := 1
	for i != left {
		prev = cursor
		cursor = cursor.Next
		i++
	}

	lPoint = prev
	lPointNext = cursor

	i++
	prev = cursor
	cursor = cursor.Next
	for i != right {
		temp := cursor.Next
		cursor.Next = prev
		prev = cursor
		cursor = temp

		i++
	}

	rPoint := cursor.Next
	cursor.Next = prev
	lPoint.Next = cursor
	lPointNext.Next = rPoint

	return mock.Next
}
