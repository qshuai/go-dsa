package leetcode

import "sort"

// https://leetcode.com/problems/merge-intervals
//
// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
// and return an array of the non-overlapping intervals that cover all the intervals in the input.
//
// Example 1:
// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
//
// Example 2:
// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.
//
// Constraints:
// 1 <= intervals.length <= 10^4
// intervals[i].length == 2
// 0 <= starti <= endi <= 10^4
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	var res [][]int
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= prev[1] {
			prev = []int{prev[0], max(prev[1], intervals[i][1])}
		} else {
			res = append(res, prev)
			prev = intervals[i]
		}
	}

	res = append(res, prev)
	return res
}
