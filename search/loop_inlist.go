package search

import (
	"github.com/qshuai/go-dsa/types"
)

// ContainLoopInList 判断链表是否为循环链表
// 采用快慢指针，一个走的快些一个走的慢些 如果最终相遇了就说明是环.
// 就相当于在一个环形跑道里跑步，速度不一样的最终一定会相遇。
func ContainLoopInList[T comparable](node *types.ListNode[T]) bool {
	// empty list
	if node == nil {
		return false
	}

	slow := node
	fast := node.Next
	for fast != nil {
		// same pointer
		if slow.Value == fast.Value {
			return true
		}

		if fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
