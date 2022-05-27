package leetcode

import (
	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/

// deleteMiddle You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.
// The middle node of a linked list of size n is the ⌊n / 2⌋th node from the start using 0-based indexing, where ⌊x⌋ denotes the largest integer less than or equal to x.
// For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2, respectively.
func deleteMiddle(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		// 空链表或者一个元素的链表，结果都是空链表
		return nil
	}

	fast, slow := head.Next, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}
