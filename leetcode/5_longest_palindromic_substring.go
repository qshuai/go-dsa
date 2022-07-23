package leetcode

// https://leetcode.com/problems/longest-palindromic-substring/

// longestPalindrome Given a string s, return the longest palindromic substring in s.
//
// Constraints:
// 1 <= s.length <= 1000
// s consist of only digits and English letters.
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	// dp[i][j]代表s[i:j+1]是否为回文字符串
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}

	var begin int
	maxLen := 1
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			// s[i] == s[j]: 表示头尾两个字符相同
			// j-i < 3: 表示如果头尾相邻或者头尾中间只有一个字符的情况下，肯定满足回文要求(前提是头尾相同)
			// dp[i+1][j-1]：代表去掉头尾，内部的字串是否为回文字符串
			dp[i][j] = s[i] == s[j] && (j-i < 3 || dp[i+1][j-1])

			if dp[i][j] && j-i+1 > maxLen {
				begin = i
				maxLen = j - i + 1
			}
		}
	}

	return s[begin : begin+maxLen]
}
