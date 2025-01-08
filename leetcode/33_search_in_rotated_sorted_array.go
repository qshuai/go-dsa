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
	if len(nums) <= 0 {
		return -1
	}

	start, end := 0, len(nums)-1
	var mid int
	for start <= end {
		mid = start + (end-start)>>1
		if target == nums[mid] {
			return mid
		}

		if nums[mid] < nums[end] {
			// 如果中间元素小于尾部元素，说明后半部分是有序的，前半部分是循环有序数组
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
				continue
			}

			end = mid - 1
		} else {
			// 如果中间元素大于尾部元素，说明前半部分是有序的，后半部分是循环有序数组
			if target < nums[mid] && target >= nums[start] {
				end = mid - 1
				continue
			}

			start = mid + 1
		}
	}

	return -1
}

// SearchInRotatedArray2 使用穷举所有条件进行的编码
func SearchInRotatedArray2(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	var mid int
	start, end := 0, len(nums)-1
	for start <= end {
		mid = start + (end-start)>>1
		if target < nums[mid] {
			if target > nums[start] {
				if nums[mid] > nums[start] {
					end = mid - 1
				} else if nums[mid] < nums[start] {
					// impossible
				} else {
					start = mid + 1
				}
			} else if target < nums[start] {
				if nums[mid] >= nums[start] {
					start = mid + 1
				} else if nums[mid] < nums[start] {
					end = mid - 1
				}
			} else {
				return start
			}
		} else if target > nums[mid] {
			if target > nums[start] {
				if nums[mid] >= nums[start] {
					start = mid + 1
				} else if nums[mid] < nums[start] {
					end = mid - 1
				}
			} else if target < nums[start] {
				if nums[mid] > nums[start] {
					// impossible
				} else if nums[mid] <= nums[start] {
					start = mid + 1
				}
			} else {
				return start
			}
		} else {
			return mid
		}
	}

	return -1
}
