package leetcode

import "github.com/qshuai/go-dsa/types"

// https://leetcode.com/problems/rotate-list/
//
// Given the head of a linked list, rotate the list to the right by k places.
//
// Example:
// Input: head = [1,2,3,4,5], k = 2
// Output: [4,5,1,2,3]
// Input: head = [0,1,2], k = 4
// Output: [2,0,1]
//
// Constraints:
// The number of nodes in the list is in the range [0, 500].
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 10^9
func rotateRight[T any](head *types.ListNode[T], k int) *types.ListNode[T] {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 计算链表长度（因为k有可能很大，而链表长度的整数倍滚动后还是链表本身，故此处可以进行优化）
	var length int
	var tail *types.ListNode[T]
	cur := head
	for cur != nil {
		length++

		tail = cur
		cur = cur.Next
	}

	k = k % length
	if k == 0 {
		return head
	}

	var prev *types.ListNode[T]
	cur = head
	for i := 0; i < length-k; i++ {
		prev = cur
		cur = cur.Next
	}

	prev.Next = nil
	tail.Next = head

	return cur
}
