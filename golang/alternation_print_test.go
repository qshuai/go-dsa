package golang

import (
	"fmt"
	"sync"
)

func ExampleAlternationPrint() {
	letters := []string{"A", "B", "C", "D", "E", "F"}
	numbers := []int{1, 2, 3, 4, 5, 6}

	letterChan := make(chan string)
	numberChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// print letter
	go func() {
		defer wg.Done()

		for _, item := range numbers {
			fmt.Println(<-letterChan)
			numberChan <- item
		}

		close(numberChan)
	}()

	// print number
	go func() {
		defer wg.Done()

		for _, item := range letters {
			letterChan <- item
			fmt.Println(<-numberChan)
		}

		close(letterChan)
	}()

	wg.Wait()

	// Output:
	// A
	// 1
	// B
	// 2
	// C
	// 3
	// D
	// 4
	// E
	// 5
	// F
	// 6
}
