package leetcode

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// lengthOfLongestSubstring Given a string s, find the length of the longest substring
// without repeating characters.
func lengthOfLongestSubstring(s string) int {
	chars := make([]byte, 128)

	left, right := 0, 0
	res := 0
	for right < len(s) {
		rs := s[right]
		chars[rs]++

		for ; chars[rs] > 1; left++ {
			chars[s[left]]--
		}

		b := right - left + 1
		if res < b {
			res = b
		}
		right++
	}

	return res
}
