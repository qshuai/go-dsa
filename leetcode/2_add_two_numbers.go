package leetcode

import (
	"golang.org/x/exp/constraints"

	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/add-two-numbers/

// addTwoNumbers You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit. Add
// the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
func addTwoNumbers[T constraints.Integer](l1 *types.ListNode[T], l2 *types.ListNode[T]) *types.ListNode[T] {
	var ret types.ListNode[T]
	tmp := &ret

	var sum, carry, store T
	var num1, num2 T
	for {
		if l1 != nil && l2 == nil {
			num1 = l1.Value
			num2 = 0

			l1 = l1.Next

			tmp.Next = &types.ListNode[T]{}
		} else if l1 == nil && l2 != nil {
			num1 = 0
			num2 = l2.Value

			l2 = l2.Next

			tmp.Next = &types.ListNode[T]{}
		} else if l1 != nil && l2 != nil {
			num1 = l1.Value
			num2 = l2.Value

			l1 = l1.Next
			l2 = l2.Next

			tmp.Next = &types.ListNode[T]{}
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
