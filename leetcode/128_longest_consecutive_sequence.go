package leetcode

// https://leetcode.com/problems/longest-consecutive-sequence/solution/

// longestConsecutive Given an unsorted array of integers nums, return the length of the
// longest consecutive elements sequence.
// You must write an algorithm that runs in O(n) time.
//
// Constraints:
// 0 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9
func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = struct{}{}
	}

	var res int
	for num, _ := range m {
		if _, ok := m[num-1]; ok {
			continue
		}

		cnt := 1
		for tmp := num + 1; ; tmp++ {
			if _, ok := m[tmp]; !ok {
				break
			}

			cnt++
		}
		if cnt > res {
			res = cnt
		}
	}

	return res
}

// Solution-2 OOM
func longestConsecutive2(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	min, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	bucket := make([]bool, max-min+1)
	for i := 0; i < len(nums); i++ {
		bucket[nums[i]-min] = true
	}

	var ans, tmp int
	for i := 0; i < len(bucket); i++ {
		if bucket[i] {
			tmp++
		} else {
			if tmp > ans {
				ans = tmp
			}
			tmp = 0
		}
	}
	if tmp > ans {
		ans = tmp
	}

	return ans
}
