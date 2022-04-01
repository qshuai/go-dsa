package list

import (
	"fmt"

	"github.com/qshuai/go-dsa/types"
)

func ExampleReverseSingleList() {
	list := types.NewListNode(1, 5)

	reversed := ReverseSingleList(list)
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
