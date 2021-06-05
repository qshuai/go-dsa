package leetcode

// https://leetcode.com/problems/majority-element/
// majorityElement Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority
// element always exists in the array.
// eg: [3,2,3] -> 3; [2,2,1,1,1,2,2] -> 2; [1,2,2,2,1] -> 2
func majorityElement(nums []int) int {
	var times, target int
	for _, num := range nums {
		if times == 0 {
			target = num
			times++
			continue
		}

		if target == num {
			times++
		} else {
			times--
		}
	}

	return target
}
