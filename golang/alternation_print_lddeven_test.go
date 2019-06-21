package golang

import (
	"fmt"
	"sync"
)

// 两个线程交替打印10以内的奇偶数
func ExamplePrintOddEvenAlternately() {
	var mtx sync.Mutex
	num := 1

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for {
			mtx.Lock()
			if num > 10 {
				wg.Done()
				mtx.Unlock()
				break
			}

			if num%2 != 0 {
				fmt.Println("goroutine 1:", num)
				num++
			}
			mtx.Unlock()
		}
	}()

	go func() {
		for {
			mtx.Lock()
			if num > 10 {
				wg.Done()
				mtx.Unlock()
				break
			}

			if num%2 == 0 {
				fmt.Println("goroutine 2:", num)
				num++
			}
			mtx.Unlock()
		}
	}()

	wg.Wait()

	// Output:
	// goroutine 1: 1
	// goroutine 2: 2
	// goroutine 1: 3
	// goroutine 2: 4
	// goroutine 1: 5
	// goroutine 2: 6
	// goroutine 1: 7
	// goroutine 2: 8
	// goroutine 1: 9
	// goroutine 2: 10
}
