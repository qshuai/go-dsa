package leetcode

// https://leetcode.com/problems/jump-game-ii/
//
// You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].
//
// Each element nums[i] represents the maximum length of a forward jump from index i. In other words,
// if you are at nums[i], you can jump to any nums[i + j] where:
// 0 <= j <= nums[i] and
// i + j < n
//
// Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can
// reach nums[n - 1].
//
// Example 1:
// Input: nums = [2,3,1,1,4]
// Output: 2
// Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
//
// Example 2:
// Input: nums = [2,3,0,1,4]
// Output: 2
//
// Constraints:
// 1 <= nums.length <= 10^4
// 0 <= nums[i] <= 1000
// It's guaranteed that you can reach nums[n - 1].
func jump(nums []int) int {
	var count, maxPos, farthest int
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])

		if i == maxPos {
			count++
			maxPos = farthest

			if maxPos == len(nums)-1 {
				return count
			}
		}
	}

	return count
}

func jump2(nums []int) int {
	minStep := len(nums)
	dfsJump(nums, 0, 0, &minStep)

	return minStep
}

func dfsJump(nums []int, pos, step int, minStep *int) {
	if pos > len(nums)-1 {
		return
	}
	if pos == len(nums)-1 {
		// 到达终点
		*minStep = min(*minStep, step)
		return
	}
	if step >= *minStep {
		return
	}

	for i := nums[pos]; i >= 1; i-- {
		dfsJump(nums, pos+i, step+1, minStep)
	}
}
