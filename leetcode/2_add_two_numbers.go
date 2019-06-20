package leetcode

// Reference: https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret ListNode
	tmp := &ret

	var sum, carry, store int
	var num1, num2 int
	for {
		if l1 != nil && l2 == nil {
			num1 = l1.Val
			num2 = 0

			l1 = l1.Next

			tmp.Next = &ListNode{}
		} else if l1 == nil && l2 != nil {
			num1 = 0
			num2 = l2.Val

			l2 = l2.Next

			tmp.Next = &ListNode{}
		} else if l1 != nil && l2 != nil {
			num1 = l1.Val
			num2 = l2.Val

			l1 = l1.Next
			l2 = l2.Next

			tmp.Next = &ListNode{}
		} else {
			if carry != 0 {
				tmp.Next.Val = carry
			} else {
				tmp.Next = nil
			}

			break
		}

		sum = num1 + num2 + carry
		carry = sum / 10
		store = sum % 10

		tmp.Val = store
		if l1 != nil || l2 != nil {
			tmp = tmp.Next
		}
	}

	return &ret
}
