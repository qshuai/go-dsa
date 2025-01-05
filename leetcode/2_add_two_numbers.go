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
	res := types.NewListNode[T]()
	cursor := res

	var num1, num2, carry T
	for l1 != nil || l2 != nil {
		if l1 != nil {
			num1 = l1.Value
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = l2.Value
			l2 = l2.Next
		}

		addition := num1 + num2 + carry
		cursor.Next = types.NewListNodeWithValue[T](addition % 10)
		cursor = cursor.Next

		// next loop
		num1, num2, carry = 0, 0, 0
		if addition >= 10 {
			carry = 1
		}
	}

	if carry > 0 {
		cursor.Next = types.NewListNodeWithValue[T](1)
	}

	return res.Next
}
