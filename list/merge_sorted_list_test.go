package list

import (
	"fmt"
)

func ExampleMergeSortedList() {
	l1 := &Node{Value: 1, Next: &Node{2, &Node{3, nil}}}
	l2 := &Node{Value: 1, Next: &Node{2, &Node{3, nil}}}

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
