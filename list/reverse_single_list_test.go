package list

import (
	"fmt"

	"github.com/qshuai/go-dsa/types"
)

func ExampleReverseSingleList() {
	list := types.ListNode{1, nil}
	list.Next = &types.ListNode{2, nil}
	list.Next.Next = &types.ListNode{3, nil}
	list.Next.Next.Next = &types.ListNode{4, nil}
	list.Next.Next.Next.Next = &types.ListNode{5, nil}

	reversed := ReverseSingleList(&list)
	for reversed != nil {
		fmt.Println(reversed.Value)

		reversed = reversed.Next
	}

	// Output:
	// 5
	// 4
	// 3
	// 2
	// 1
}
