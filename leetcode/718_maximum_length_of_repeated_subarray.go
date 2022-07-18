package leetcode

// https://leetcode.com/problems/maximum-length-of-repeated-subarray/

// Given two integer arrays nums1 and nums2, return the maximum length of a subarray that appears in both arrays.
// Constraints:
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 100
//
// Hint: dp[i][j] 代表nums1[i:]和nums2[j:]最长公共前缀
func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	dp[len(nums1)] = make([]int, len(nums2)+1)

	var ans int
	for i := len(nums1) - 1; i >= 0; i-- {
		dp[i] = make([]int, len(nums2)+1)
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if dp[i][j] > ans {
					ans = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	return ans
}
