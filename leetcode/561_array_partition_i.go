package leetcode

// https://leetcode.com/problems/array-partition-i/

// arrayPairSum
//
// Constraints:
// 1 <= n <= 104
// nums.length == 2 * n
// -104 <= nums[i] <= 104
func arrayPairSum(nums []int) int {
	min, max := nums[0], nums[0]
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

	var isLeft bool
	var sum int
	var num, quotient, modulo int
	for idx, item := range buckets {
		if item <= 0 {
			continue
		}

		num = idx + min
		if isLeft {
			isLeft = false
			item--

			if item <= 0 {
				continue
			}
		}

		quotient = item / 2
		modulo = item % 2

		sum += num * quotient
		if modulo > 0 {
			isLeft = true
			sum += num
		}
	}

	return sum
}
