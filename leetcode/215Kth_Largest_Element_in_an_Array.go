package leetcode

// findKthLargest 从大到小，找到排名第K的元素(值相同的元素，排名加1，也就是说排名是连续的[可以查看测试用例])
func findKthLargest(nums []int, k int) int {
	if k > len(nums) || k <= 0 || len(nums) <= 0 {
		return -1
	}

	var offset int
	for {
		pivot := partition(nums)
		if pivot+1+offset == k {
			return nums[pivot]
		} else if pivot+1+offset > k {
			nums = nums[:pivot]
		} else {
			offset += pivot
			nums = nums[pivot:]
		}
	}
}

// 分区，大的元素放在前面（索引小的位置）
func partition(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	i, j := 0, len(nums)-1
	for {
		if i >= j {
			break
		}

		for i < j && nums[i] >= nums[j] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
		for i < j && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	return i
}
