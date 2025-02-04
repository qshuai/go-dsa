package leetcode

// https://leetcode.com/problems/trapping-rain-water
//
// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water
// it can trap after raining.
//
// Example 1:
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case,
// 6 units of rain water (blue section) are being trapped.
//
// Example 2:
// Input: height = [4,2,0,3,2,5]
// Output: 9
//
// Constraints:
// n == height.length
// 1 <= n <= 2 * 10^4
// 0 <= height[i] <= 10^5
func trap(height []int) int {
	// i位置接的水量: max(0, min(max(0 ~ i-1), max(i+1 ~ n-1)) - height[i])
	// max(0 ~ i-1)代表height在0 ~ i-1范围内的最大值
	// max(i+1 ~ n-1)代表height在i+1 ~ n-1范围内的最大值
	lmax := make([]int, len(height))
	lmax[0] = height[0]
	for i := 1; i < len(height); i++ {
		lmax[i] = max(lmax[i-1], height[i])
	}

	rmax := make([]int, len(height))
	rmax[len(rmax)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rmax[i] = max(rmax[i+1], height[i])
	}

	var res int
	for i := 1; i < len(height)-1; i++ {
		res += max(0, min(lmax[i], rmax[i])-height[i])
	}

	return res
}

// trap2采用双指针方法求解，核心思想还是当求解i位置的接水量时，需要知道左侧和右侧的最大值
func trap2(height []int) int {
	lmax, rmax := height[0], height[len(height)-1]
	l, r := 1, len(height)-2

	var res int
	for l <= r {
		if lmax <= rmax {
			res += max(0, lmax-height[l])

			lmax = max(lmax, height[l])
			l++
		} else {
			res += max(0, rmax-height[r])

			rmax = max(rmax, height[r])
			r--
		}
	}

	return res
}
