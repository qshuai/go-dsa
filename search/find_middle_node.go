package search

import (
	"github.com/qshuai/go-dsa/types"
)

// FindMiddleNode 给定一个链表的头结点，要求找到中间节点并返回
func FindMiddleNode[T any](node *types.ListNode[T]) *types.ListNode[T] {
	if node == nil {
		return nil
	}

	middle := node
	step := 0
	for node != nil {
		if step%2 != 0 {
			middle = middle.Next
		}

		node = node.Next
		step++
	}

	return middle
}
