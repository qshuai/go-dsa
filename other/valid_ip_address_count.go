package other

import "strconv"

// restoreIpAddresses givens a string s containing only digits, return the number of all possible
// valid IP addresses that can be formed by inserting dots into s. You are not allowed to reorder
// or remove any digits in s. You may return the valid IP addresses in any order.
//
// Constraints:
// 1 <= s.length <= 20
// s consists of digits only.
//
// Comment: 使用动态规划的思想，逐级递推，使得最后一个元素就是问题的解
func restoreIpAddresses(s string) int {
	const ipNums = 4
	if len(s) < ipNums {
		return 0
	}

	dp := make([][]int, ipNums+1)
	dp[0] = make([]int, len(s)+1)
	dp[0][0] = 1
	for i := 1; i < ipNums+1; i++ {
		dp[i] = make([]int, len(s)+1)
		for j := i; j <= len(s); j++ {
			if i == j {
				dp[i][j] = 1
				continue
			}

			x := 1
			for j >= x {
				ok, err := validIpSegment(s[j-x : j])
				if err != nil {
					return 0
				}
				if ok {
					dp[i][j] += dp[i-1][j-x]
				}
				x++
			}
		}
	}

	return dp[ipNums][len(s)]
}

func validIpSegment(seg string) (bool, error) {
	if len(seg) <= 0 || len(seg) > 3 {
		return false, nil
	}
	if len(seg) == 1 {
		return true, nil
	}

	if seg[0] == '0' {
		// 前导零是非法的字段
		return false, nil
	}

	const MaxSegVal = 255
	segVal, err := strconv.Atoi(seg)
	if err != nil {
		return false, err
	}

	return segVal <= MaxSegVal, nil
}
