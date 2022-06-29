package leetcode

import "github.com/qshuai/go-dsa/utils"

// https://leetcode.com/problems/maximum-subarray/

// maxSubArray Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.
func maxSubArray(nums []int) int {
	max, tmpMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tmpMax = utils.Max(nums[i], nums[i]+tmpMax)
		if tmpMax > max {
			max = tmpMax
		}
	}

	return max
}
