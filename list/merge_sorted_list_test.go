package list

import (
	"fmt"

	"github.com/qshuai/go-dsa/types"
)

func ExampleMergeSortedList() {
	l1 := types.NewListNode(1, 3)
	l2 := types.NewListNode(1, 3)

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
