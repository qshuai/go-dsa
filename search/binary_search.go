package search

// FindAny 二分查找算法
// 二分查找依赖数组这样的数据结构，并且要求数据有序。这就要求数据量不能太大，并且在业务使用中
// 不能频繁的插入和删除。
// 如果处理的数据量很小，完全没有必要用二分查找，顺序遍历就足够了。
func FindAny(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	// 优先和序列的极值进行比较
	if nums[0] > target || nums[len(nums)-1] < target {
		return -1
	}

	var mid int
	start, end := 0, len(nums)-1
	for start <= end {
		// Using the following code instead of [ mid = (start + end) / 2 ]
		// to prevent [start + end] integer overflow.
		mid = start + (end-start)>>1
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}

	}

	return -1
}

// FindAnyRecursive 二分查找的递归实现
func FindAnyRecursive(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	// 优先和序列的极值进行比较
	if nums[0] > target || nums[len(nums)-1] < target {
		return -1
	}

	return recursiveSearch(nums, 0, len(nums)-1, target)
}

// FirstIndex 查找第一个等于指定值的索引，如果没有找到返回-1
func FirstIndex(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	if nums[0] > target || nums[len(nums)-1] < target {
		return -1
	}

	var mid int
	start, end := 0, len(nums)-1
	for start <= end {
		mid = start + (end-start)>>1
		if nums[mid] >= target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		}
	}

	if start < len(nums) && nums[start] == target {
		return start
	}

	return -1
}

// LastIndex 查找最后一个等于指定值的索引，如果没有找到返回-1
func LastIndex(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	if nums[0] > target || nums[len(nums)-1] < target {
		return -1
	}

	var mid int
	start, end := 0, len(nums)/2
	for start <= end {
		mid = start + (end-start)>>1
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] <= target {
			start = mid + 1
		}
	}

	if end >= 0 && nums[end] == target {
		return end
	}

	return -1
}

// FirstGteIndex 查找第一个大于等于指定值的索引，如果没有找到返回-1
func FirstGteIndex(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	// 首先和序列的极大值进行比较
	if nums[len(nums)-1] < target {
		return -1
	}
	if nums[0] > target {
		return 0
	}

	var mid int
	start, end := 0, len(nums)-1
	for start <= end {
		mid = start + (end-start)>>1
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}

			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

// LastLteIndex 查找最后一个大于等于指定值的索引，如果没有找到返回-1
func LastLteIndex(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}
	if nums[0] > target {
		return -1
	}
	if nums[len(nums)-1] <= target {
		return len(nums) - 1
	}

	var mid int
	start, end := 0, len(nums)-1
	for start <= end {
		mid = start + (end-start)>>1
		if nums[mid] > target {
			end = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}

			start = mid + 1
		}
	}

	return -1
}

func recursiveSearch(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)>>1
	if nums[mid] < target {
		return recursiveSearch(nums, mid+1, end, target)
	} else if nums[mid] > target {
		return recursiveSearch(nums, start, mid-1, target)
	} else {
		return mid
	}
}
