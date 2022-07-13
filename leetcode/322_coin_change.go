package leetcode

import (
	"math"
)

// https://leetcode.com/problems/coin-change/

// coinChange You are given an integer array coins representing coins of different denominations
// and an integer amount representing a total amount of money.
// Return the fewest number of coins that you need to make up that amount. If that amount of money
// cannot be made up by any combination of the coins, return -1.
// You may assume that you have an infinite number of each kind of coin.
//
// Constraints:
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 231 - 1
// 0 <= amount <= 104
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		ans := math.MaxInt
		for j := 0; j < len(coins); j++ {
			if coins[j] > i {
				continue
			}
			if coins[j] == i {
				ans = 1
				break
			}
			if dp[i-coins[j]] == -1 {
				continue
			}

			nums := 1 + dp[i-coins[j]]
			if nums < ans {
				ans = nums
			}
		}

		if ans == math.MaxInt {
			dp[i] = -1
		} else {
			dp[i] = ans
		}
	}

	return dp[amount]
}
