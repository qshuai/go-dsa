package leetcode

// https://leetcode.com/problems/climbing-stairs/

// climbStairs You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Constraints:
// 1 <= n <= 45
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}

	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < len(dp); i++ {
		// 到达i台阶，可以选择一步走到这里，也可选择两步走到这里
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}
