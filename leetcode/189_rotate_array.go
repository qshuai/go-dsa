package leetcode

// https://leetcode.com/problems/rotate-array/

// rotate Given an array, rotate the array to the right by k steps, where k is non-negative.
func rotate(nums []int, k int) {
	k = k % len(nums)
	// 翻转整个数组
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}

	// 翻转前k个元素
	for i := 0; i < k/2; i++ {
		nums[i], nums[k-i-1] = nums[k-i-1], nums[i]
	}

	// 翻转后len(nums) - k 个元素
	for i := k; i < k+(len(nums)-k)/2; i++ {
		nums[i], nums[len(nums)-1-(i-k)] = nums[len(nums)-1-(i-k)], nums[i]
	}
}
