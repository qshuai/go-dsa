package list

import (
	"github.com/qshuai/go-dsa/types"
)

// mergeSortedList 将两个有序的链表合并，且合并之后的链表依然有序
// 注意该方法默认链表元素的类型为int
func mergeSortedList(l1 *types.ListNode, l2 *types.ListNode) *types.ListNode {
	if l1 == nil || l2 == nil {
		if l1 != nil {
			return l1
		}

		if l2 != nil {
			return l2
		}

		return nil
	}

	var ret *types.ListNode
	lastNode := ret
	var head *types.ListNode
	for l1 != nil && l2 != nil {
		if l1.Value.(int) < l2.Value.(int) {
			ret = l1
			l1 = l1.Next
		} else {
			ret = l2
			l2 = l2.Next
		}

		// the initial state
		if head == nil {
			head = ret
			lastNode = head
		} else {
			lastNode.Next = ret
			lastNode = ret
		}
	}

	if l1 != nil {
		ret = l1
	}

	if l2 != nil {
		ret = l2
	}

	lastNode.Next = ret

	return head
}
