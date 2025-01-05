package golang

import (
	"fmt"
	"sync"
)

func ExampleSequence() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(numbers))

	res := make(chan int, len(numbers))
	for _, c := range numbers {
		go func() {
			res <- <-ch
			wg.Done()
		}()

		ch <- c
	}
	close(ch)
	wg.Wait()

	close(res)
	for v := range res {
		fmt.Println(v)
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
