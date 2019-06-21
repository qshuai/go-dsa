package array

import (
	"fmt"
)

func ExampleMoveArray() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	moveArray(arr, 3)
	fmt.Println(arr)

	// Output:
	// [5 6 7 1 2 3 4]
}
