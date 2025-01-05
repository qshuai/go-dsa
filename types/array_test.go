package types

import (
	"fmt"
)

func ExampleMoveArray() {
	// empty array
	var arr []int
	fmt.Println(arr)

	// nonempty array
	arr = []int{1, 2, 3, 4, 5, 6, 7}
	moveArray(arr, 3)
	fmt.Println(arr)

	// Output:
	// []
	// [5 6 7 1 2 3 4]
}
