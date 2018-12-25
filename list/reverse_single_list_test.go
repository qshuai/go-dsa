package list

import (
	"fmt"
)

func ExampleReverseSingleList() {
	list := Node{1, nil}
	list.Next = &Node{2, nil}
	list.Next.Next = &Node{3, nil}
	list.Next.Next.Next = &Node{4, nil}
	list.Next.Next.Next.Next = &Node{5, nil}

	reversed := reverseSingleList(&list)
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
