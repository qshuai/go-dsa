package leetcode

import (
	"math"
)

// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/

// minOperations You are given an integer array nums and an integer x. In one operation, you can either remove
// the leftmost or the rightmost element from the array nums and subtract its value from x. Note that this
// modifies the array for future operations.
// Return the minimum number of operations to reduce x to exactly 0 if it is possible, otherwise, return -1.
//
// Constraints:
// 1 <= nums.length <= 10^5
// 1 <= nums[i] <= 10^4
// 1 <= x <= 10^9
func minOperations(nums []int, x int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	// 转换成中间连续子序列的和等于指定值
	x = sum - x
	if x == 0 {
		return len(nums)
	}

	max := math.MinInt
	var sumWin int
	for i, j := 0, 0; i < len(nums); i++ {
		sumWin += nums[i]

		if sumWin == x && max < i-j+1 {
			max = i - j + 1
		}

		for sumWin > x && j < i {
			sumWin -= nums[j]
			j++
			if sumWin == x && max < i-j+1 {
				max = i - j + 1
			}
		}
	}

	if max == math.MinInt {
		return -1
	}
	return len(nums) - max
}
