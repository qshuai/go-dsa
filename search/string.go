package search

import (
	"github.com/qshuai/go-dsa/utils"
)

// IndexOfUsingBF returns the index of the first instance of sub string
// from str, or -1 if not matched substring present from str.
// Brute Force algorithm
func IndexOfUsingBF(str, sub string) int {
	if len(str) < len(sub) {
		return -1
	}

	for i := 0; i <= len(str)-len(sub); i++ {
		matched := true
		for j := 0; j < len(sub); j++ {
			if str[i+j] != sub[j] {
				matched = false
				break
			}
		}

		if matched {
			return i
		}
	}

	return -1
}

// IndexOfUsingBM return the first index of matched sub string, or -1
// if not found
// Using Boyer-Moore algorithm
// Constraints:
// str and sub are consist of ascii character
func IndexOfUsingBM(str, sub string) int {
	if len(str) < len(sub) {
		return -1
	}

	bc := buildAsciiArray(sub)
	suffix, prefix := buildMatchedSegment(sub)
	for i := 0; i <= len(str)-len(sub); {
		j := len(sub) - 1
		for ; j >= 0; j-- {
			if sub[j] != str[i+j] {
				break
			}
		}

		if j < 0 {
			// 找到匹配的子串
			return i
		}

		x := j - bc[str[i+j]] + 1
		var y int
		if j < len(sub)-1 {
			y = moveByMatchedSegment(j, len(sub), suffix, prefix)
		}

		i += utils.Max(x, y)
	}

	return -1
}

func buildAsciiArray(m string) []int {
	if len(m) <= 0 {
		return nil
	}

	ret := make([]int, 256)
	for i := 0; i < len(m); i++ {
		ret[m[i]] = i + 1
	}

	return ret
}

func buildMatchedSegment(m string) ([]int, []bool) {
	suffix := make([]int, len(m))
	prefix := make([]bool, len(m))
	for i := 0; i < len(m)-1; i++ {
		var k int
		j := i
		for j >= 0 && m[j] == m[len(m)-1-k] {
			j--
			k++
			suffix[k] = j + 2
		}

		if j == -1 {
			prefix[k] = true
		}
	}

	return suffix, prefix
}

func moveByMatchedSegment(j, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j
	if suffix[k] != 0 {
		return j - suffix[k] + 2
	}

	for i := j + 2; i < m-1; i++ {
		if prefix[m-i] {
			return i
		}
	}

	return m
}
