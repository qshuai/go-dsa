package leetcode

// https://leetcode.com/problems/edit-distance/
//
// Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
//
// You have the following three operations permitted on a word:
//
// Insert a character
// Delete a character
// Replace a character
//
// Constraints:
// 0 <= word1.length, word2.length <= 500
// word1 and word2 consist of lowercase English letters.
func minDistance(word1 string, word2 string) int {
	// dp[i][j]代表长度为i的word1，转变成长度为j的word2，所需要的最小编辑距离。
	// 那么本题求解的答案就为dp[len(word1)][len(word2)]的值。为了包含该数据，
	// 故将矩阵的行和列增加1。
	dp := make([][]int, len(word1)+1)
	for idx := range dp {
		dp[idx] = make([]int, len(word2)+1)
	}

	dp[0][0] = 0 // word1、word2长度都为0时，不需要编辑

	// 第一行：word1长度为0，word2长度递增，为了使word1能够编辑为word2，故需要不断进行插入操作
	for i := 1; i <= len(word2); i++ {
		dp[0][i] = dp[0][i-1] + 1
	}

	// 第一列：word2长度为0，word1长度递增，为了使word1能够编辑为word2，故需要不断进行删除操作
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			// 长度分别为i、j的单词word1、word2，其最后一个字符索引分别为i-1、j-1
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j],
					dp[i][j-1],
					dp[i-1][j-1],
				) + 1
			}
		}
	}

	return dp[len(word1)][len(word2)]
}
