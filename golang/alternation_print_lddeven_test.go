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

// java implement
// import java.util.concurrent.locks.Lock;
// import java.util.concurrent.locks.ReentrantLock;
//
// public class AlternativePrintOddEvenNumber {
//     private static volatile int num = 1;
//
//     private static Lock lock = new ReentrantLock();
//
//     static class OddNum implements Runnable {
//         public void run() {
//             while (true) {
//                 lock.lock();
//                 if (num > 10) {
//                     lock.unlock();
//                     break;
//                 }
//
//                 if (num % 2 != 0) {
//                     System.out.printf("%s: %d\n",Thread.currentThread().getName(), num++);
//                 }
//                 lock.unlock();
//             }
//         }
//     }
//
//     static class EvenNum implements Runnable {
//         public void run() {
//             while (true) {
//                 lock.lock();
//                 if (num > 10) {
//                     lock.unlock();
//                     break;
//                 }
//
//                 if (num % 2 == 0) {
//                     System.out.printf("%s: %d\n",Thread.currentThread().getName(), num++);
//                 }
//                 lock.unlock();
//             }
//         }
//     }
//
//     public static void main(String[] args) {
//         AlternativePrintOddEvenNumber r = new AlternativePrintOddEvenNumber();
//         Thread oddThread = new Thread(new OddNum());
//         oddThread.setName("thread-1");
//
//         Thread evenThread = new Thread(new EvenNum());
//         evenThread.setName("thread-2");
//
//         oddThread.start();
//         evenThread.start();
//     }
// }
