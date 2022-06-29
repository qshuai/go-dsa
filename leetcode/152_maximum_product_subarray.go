package leetcode

import "github.com/qshuai/go-dsa/utils"

// https://leetcode.com/problems/maximum-product-subarray/

// maxProduct Given an integer array nums, find a contiguous non-empty subarray within
// the array that has the largest product, and return the product.
//
// Constraints:
// 1 <= nums.length <= 2 * 104
// -10 <= nums[i] <= 10
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
func maxProduct(nums []int) int {
	max := nums[0]
	tmpMax, tmpMin := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			tmpMax, tmpMin = tmpMin, tmpMax
		}

		tmpMax = utils.Max(nums[i], tmpMax*nums[i])
		tmpMin = utils.Min(nums[i], tmpMin*nums[i])

		max = utils.Max(max, tmpMax)
	}

	return max
}
