package leetcode

import (
	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

// removeNthFromEnd Given the head of a linked list, remove the nth node
// from the end of the list and return its head.
// Constraints:
//	The number of nodes in the list is sz.
//	1 <= sz <= 30
//	0 <= TreeNode.val <= 100
//	1 <= n <= sz
func removeNthFromEnd(head *types.ListNode, n int) *types.ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	for i := 1; i <= n; i++ {
		fast = fast.Next
	}

	// the target element is the head of linked list
	if fast == nil {
		return head.Next
	}

	prev := slow
	for fast != nil {
		fast = fast.Next

		prev = slow
		slow = slow.Next
	}

	prev.Next = slow.Next
	return head
}

func removeNthFromEnd2(head *types.ListNode, n int) *types.ListNode {
	if head == nil {
		return nil
	}

	var a [30]*types.ListNode
	var cursor int
	for head != nil {
		a[cursor] = head
		cursor++
		head = head.Next
	}

	offset := cursor - n
	if offset < 0 || offset > cursor {
		return head
	}

	if offset == 0 {
		head = a[1]
		return head
	}

	a[offset-1].Next = a[offset].Next
	return a[0]
}
