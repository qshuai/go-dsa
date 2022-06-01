package leetcode

// https://leetcode.com/problems/fibonacci-number/

// fib 计算斐波那契数列
func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

// fibLoop fib loop implementation
func fibLoop(n int) int {
	if n <= 1 {
		return n
	}

	prev, cur := 0, 1
	for i := 2; i <= n; i++ {
		prev, cur = cur, prev+cur
	}

	return cur
}
