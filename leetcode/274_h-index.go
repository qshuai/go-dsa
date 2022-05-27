package leetcode

// https://leetcode.com/problems/h-index/

// hIndex
//
// Constraints:
// n == citations.length
// 1 <= n <= 5000
// 0 <= citations[i] <= 1000
func hIndex(citations []int) int {
	buckets := make([]int, len(citations)+1)
	for _, refs := range citations {
		if refs >= len(citations) {
			buckets[len(citations)]++
			continue
		}

		buckets[refs]++
	}

	var articles int
	for i := len(buckets) - 1; i >= 0; i-- {
		articles += buckets[i]
		if articles >= i {
			return i
		}
	}

	return 0
}
