package leetcode

// https://leetcode.com/problems/sort-colors/
// tag:
//   - bucket sort

// sortColors Given an array nums with n objects colored red, white, or blue, sort them in-place so that
// objects of the same color are adjacent, with the colors in the order red, white, and blue.
// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
// You must solve this problem without using the library's sort function.
func sortColors(nums []int) {
	if len(nums) <= 0 {
		return
	}

	red := make([]int, 0)
	white := make([]int, 0)
	blue := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			red = append(red, 0)
		case 1:
			white = append(white, 1)
		case 2:
			blue = append(blue, 2)
		}
	}
	nums = nums[:0]
	nums = append(nums, red...)
	nums = append(nums, white...)
	nums = append(nums, blue...)
}
