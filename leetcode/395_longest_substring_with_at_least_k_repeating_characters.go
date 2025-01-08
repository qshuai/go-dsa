package leetcode

// https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/
//
// Given a string s and an integer k, return the length of the longest substring of s such that
// the frequency of each character in this substring is greater than or equal to k.
// if no such substring exists, return 0.
//
// Example 1:
// Input: s = "aaabb", k = 3
// Output: 3
// Explanation: The longest substring is "aaa", as 'a' is repeated 3 times.
//
// Example 2:
// Input: s = "ababbc", k = 2
// Output: 5
// Explanation: The longest substring is "ababb", as 'a' is repeated 2 times and 'b' is repeated 3 times.
//
// Constraints:
// 1 <= s.length <= 10^4
// s consists of only lowercase English letters.
// 1 <= k <= 10^5
func longestSubstring(s string, k int) int {
	if k > len(s) {
		return 0
	}

	var res int
	for requires := 1; requires <= min(26, len(s)); requires++ {
		counts := make([]int, 26)
		var chars, satisfy int
		for left, right := 0, 0; right < len(s); right++ {
			counts[s[right]-'a']++
			if counts[s[right]-'a'] == 1 {
				chars++
			}
			if counts[s[right]-'a'] == k {
				satisfy++
			}

			for chars > requires {
				if counts[s[left]-'a'] == 1 {
					chars--
				}
				if counts[s[left]-'a'] == k {
					satisfy--
				}

				counts[s[left]-'a']--
				left++
			}

			if satisfy == requires {
				res = max(res, right-left+1)
			}
		}
	}

	return res
}
