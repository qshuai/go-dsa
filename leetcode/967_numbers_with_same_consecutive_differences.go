package leetcode

// https://leetcode.com/problems/numbers-with-same-consecutive-differences/

// numsSameConsecDiff Return all non-negative integers of length n such that the absolute difference
// between every two consecutive digits is k.
// Note that every number in the answer must not have leading zeros. For example, 01 has one leading zero and is invalid.
// You may return the answer in any order.
//
// Constraints:
// 2 <= n <= 9
// 0 <= k <= 9
func numsSameConsecDiff(n int, k int) []int {
	collector := make([]int, 0)
	temp := make([]int, 0, n)
	pows := make([]int, n) // 避免重复运算
	pows[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		pows[i] = 10 * pows[i+1]
	}

	for i := 1; i < 10; i++ {
		temp = append(temp, i)
		dfsDigits(n, k, pows, &temp, &collector)
		temp = temp[:0]
	}

	return collector
}

func dfsDigits(n, k int, pows []int, temp, collector *[]int) {
	if len(*temp) == n {
		var item int
		for i := len(*temp) - 1; i >= 0; i-- {
			item += (*temp)[i] * pows[i]
		}
		*collector = append(*collector, item)
		return
	}

	// +k
	last := (*temp)[len(*temp)-1]
	if last+k < 10 {
		*temp = append(*temp, last+k)
		dfsDigits(n, k, pows, temp, collector)
		*temp = (*temp)[:len(*temp)-1]
	}

	// -k
	if last-k >= 0 && k != 0 {
		*temp = append(*temp, last-k)
		dfsDigits(n, k, pows, temp, collector)
		*temp = (*temp)[:len(*temp)-1]
	}
}
