package leetcode

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/permutation-sequence/

// getPermutation Given n and k, return the kth permutation sequence.
//
// Constraints:
// 1 <= n <= 9
// 1 <= k <= n!
func getPermutation(n int, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	sb := strings.Builder{}
	sb.Grow(n)
	visited := make([]bool, n)
	var num int
	for i := 0; i < n && k > 0; i++ {
		// index代表第i高位的数字，是第几种可能（从小到大排序）
		index := (k-1)/factorial[n-1-i] + 1
		var step int
		for j := 0; j < n; j++ {
			if visited[j] {
				// 如果已经使用了，就跳过
				continue
			}

			step++
			if step == index {
				// 从前往后遍历，如果走到该index的步骤，即为正确的解
				num = j + 1
				break
			}
		}
		visited[num-1] = true

		sb.WriteString(fmt.Sprintf("%d", num))
		k = k - (index-1)*factorial[n-i-1]
	}

	return sb.String()
}
