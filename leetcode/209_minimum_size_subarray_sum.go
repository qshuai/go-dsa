package leetcode

import (
	"math"
)

// https://leetcode.com/problems/minimum-size-subarray-sum/

// Given an array of positive integers nums and a positive integer target, return the minimal
// length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is
// greater than or equal to target. If there is no such subarray, return 0 instead.
//
// Constraints:
// 1 <= target <= 10^9
// 1 <= nums.length <= 10^5
// 1 <= nums[i] <= 10^4
// Using slide window
func minSubArrayLen(target int, nums []int) int {
	var i, j, sum int
	res := math.MaxInt
	for ; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			length := j - i + 1
			if length < res {
				res = length
			}

			sum -= nums[i]
			i++
		}
	}

	if res == math.MaxInt {
		return 0
	}
	return res
}
