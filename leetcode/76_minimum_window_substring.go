package leetcode

import "math"

// https://leetcode.com/problems/minimum-window-substring/
//
// Given two strings s and t of lengths m and n respectively, return the minimum window
// substring of s such that every character in t (including duplicates) is included in the window.
// If there is no such substring, return the empty string "".
//
// The testcases will be generated such that the answer is unique.
//
// Example 1:
// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
//
// Example 2:
// Input: s = "a", t = "a"
// Output: "a"
// Explanation: The entire string s is the minimum window.
//
// Example 3:
// Input: s = "a", t = "aa"
// Output: ""
// Explanation: Both 'a's from t must be included in the window.
// Since the largest window of s only has one 'a', return empty string.
//
// Constraints:
// m == s.length
// n == t.length
// 1 <= m, n <= 105
// s and t consist of uppercase and lowercase English letters.
//
// Follow up: Could you find an algorithm that runs in O(m + n) time?
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	counts := make([]int, 128)
	for _, char := range t {
		counts[char]++
	}

	start, length := -1, math.MaxInt
	for left, right, debt := 0, 0, len(t); right < len(s); right++ {
		if counts[s[right]] > 0 {
			debt--
		}
		counts[s[right]]--

		if debt == 0 {
			for counts[s[left]] < 0 {
				counts[s[left]]++
				left++
			}

			if right-left+1 < length {
				start = left
				length = right - left + 1
			}
		}
	}

	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}
