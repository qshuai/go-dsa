package leetcode

// SearchInRotatedArray returns the index of target in an ascending array that is possibly rotated.
// -1 returned if not found.
//
// Constraints:
// 1 <= nums.length <= 5000
// -10^4 <= nums[i] <= 10^4
// All values of nums are unique.
// nums is an ascending array that is possibly rotated.
// -10^4 <= target <= 10^4
func SearchInRotatedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = left + (right-left)>>1
		if target == nums[mid] {
			return mid
		}

		if nums[mid] >= nums[left] { // 左侧是有序的情况
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右侧是有序的情况
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
