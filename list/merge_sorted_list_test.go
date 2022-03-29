package list

import (
	"fmt"

	"github.com/qshuai/go-dsa/types"
)

func ExampleMergeSortedList() {
	l1 := &types.ListNode{Value: 1, Next: &types.ListNode{2, &types.ListNode{3, nil}}}
	l2 := &types.ListNode{Value: 1, Next: &types.ListNode{2, &types.ListNode{3, nil}}}

	list := mergeSortedList(l1, l2)
	for list != nil {
		fmt.Println(list.Value)
		list = list.Next
	}

	// Output:
	// 1
	// 1
	// 2
	// 2
	// 3
	// 3
}
