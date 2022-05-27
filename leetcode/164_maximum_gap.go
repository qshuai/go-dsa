package leetcode

// https://leetcode.com/problems/maximum-gap/

// maximumGap Given an integer array nums, return the maximum difference between two
// successive elements in its sorted form. If the array contains less than two elements, return 0.
// You must write an algorithm that runs in linear time and uses linear extra space.
func maximumGap(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	// 获取数据的最大最小值
	min, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	if max == min {
		return 0
	}
	gap := (max - min) / (len(nums) - 1)
	if gap == 0 {
		gap = 1
	}
	buckets := (max-min)/gap + 1
	maximum := make([]*int, buckets)
	minimum := make([]*int, buckets)
	for i := 0; i < len(nums); i++ {
		k := (nums[i] - min) / gap
		if maximum[k] == nil {
			maximum[k] = &nums[i]
		} else if nums[i] > *maximum[k] {
			maximum[k] = &nums[i]
		}

		if minimum[k] == nil {
			minimum[k] = &nums[i]
		} else if nums[i] < *minimum[k] {
			minimum[k] = &nums[i]
		}
	}

	var maxgap int
	var i, j int
	var diff int
	for i < buckets {
		j = i + 1
		for j < buckets && maximum[j] == nil && minimum[j] == nil {
			j++
		}
		if j >= buckets {
			break
		}

		diff = *minimum[j] - *maximum[i]
		if diff > maxgap {
			maxgap = diff
		}
		i = j
	}

	return maxgap
}
