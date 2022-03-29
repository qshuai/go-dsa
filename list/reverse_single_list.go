package list

import (
	"github.com/qshuai/go-dsa/types"
)

func ReverseSingleList(list *types.ListNode) *types.ListNode {
	cur := list
	var prev *types.ListNode

	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
}
