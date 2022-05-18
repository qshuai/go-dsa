package sort

// https://leetcode.com/problems/sort-an-array/

// BucketSort 桶排序
// Constraints:
// 1 <= nums.length <= 5 * 10^4
// -5 * 104 <= nums[i] <= 5 * 10^4
func BucketSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	min, max := nums[0], nums[1]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	buckets := make([]int, max-min+1)
	for _, num := range nums {
		buckets[num-min]++
	}

	var offset int
	for idx, bucket := range buckets {
		for bucket > 0 {
			nums[offset] = min + idx
			offset++
			bucket--
		}
	}

	return nums
}
