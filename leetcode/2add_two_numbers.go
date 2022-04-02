package leetcode

import (
	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/add-two-numbers/

// addTwoNumbers You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit. Add
// the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
func addTwoNumbers(l1 *types.ListNode, l2 *types.ListNode) *types.ListNode {
	var ret types.ListNode
	tmp := &ret

	var sum, carry, store int
	var num1, num2 int
	for {
		if l1 != nil && l2 == nil {
			num1 = l1.Value.(int)
			num2 = 0

			l1 = l1.Next

			tmp.Next = &types.ListNode{}
		} else if l1 == nil && l2 != nil {
			num1 = 0
			num2 = l2.Value.(int)

			l2 = l2.Next

			tmp.Next = &types.ListNode{}
		} else if l1 != nil && l2 != nil {
			num1 = l1.Value.(int)
			num2 = l2.Value.(int)

			l1 = l1.Next
			l2 = l2.Next

			tmp.Next = &types.ListNode{}
		} else {
			if carry != 0 {
				tmp.Next.Value = carry
			} else {
				tmp.Next = nil
			}

			break
		}

		sum = num1 + num2 + carry
		carry = sum / 10
		store = sum % 10

		tmp.Value = store
		if l1 != nil || l2 != nil {
			tmp = tmp.Next
		}
	}

	return &ret
}
