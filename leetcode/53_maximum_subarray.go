package leetcode

// https://leetcode.com/problems/maximum-subarray/

// maxSubArray Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.
func maxSubArray(nums []int) int {
	res, tmpMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tmpMax = max(nums[i], nums[i]+tmpMax)
		if tmpMax > res {
			res = tmpMax
		}
	}

	return res
}
