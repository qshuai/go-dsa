package golang

import "fmt"

func ExampleSequence() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, c := range numbers {
		ch := make(chan int)
		go func() {
			fmt.Println(<-ch)
		}()

		ch <- c
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}
