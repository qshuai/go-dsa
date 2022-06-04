package leetcode

import (
	"sort"
	"strconv"
)

// https://leetcode.com/problems/relative-ranks/

// findRelativeRanks returns an array answer of size n
// where answer[i] is the rank of the ith athlete.
//
// Constraints:
// n == score.length
// 1 <= n <= 104
// 0 <= score[i] <= 106
// All the values in score are unique.
func findRelativeRanks(score []int) []string {
	scoreWithIdx := make([][2]int, len(score))
	for i, item := range score {
		scoreWithIdx[i][0] = item
		scoreWithIdx[i][1] = i
	}

	sort.Slice(scoreWithIdx, func(i, j int) bool {
		return scoreWithIdx[i][0] > scoreWithIdx[j][0]
	})

	ans := make([]string, len(score))
	medal := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	for idx, item := range scoreWithIdx {
		if idx < len(medal) {
			ans[item[1]] = medal[idx]
			continue
		}

		ans[item[1]] = strconv.Itoa(idx + 1)
	}

	return ans
}
