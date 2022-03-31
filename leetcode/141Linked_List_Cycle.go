package leetcode

import (
	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/linked-list-cycle/

// hasCycle 检查单链表是否存在环
// 使用快慢指针遍历链表，快指针每次走两步，慢指针每次走一步；
// 边界条件：1. 慢指针在快指针的前面 2. 慢指针和快指针指向了同一个元素
func hasCycle(head *types.ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		if fast.Next == slow {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
