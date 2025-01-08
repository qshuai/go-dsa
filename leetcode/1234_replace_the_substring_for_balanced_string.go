package leetcode

// https://leetcode.com/problems/replace-the-substring-for-balanced-string/
//
// You are given a string s of length n containing only four kinds of characters: 'Q', 'W', 'E', and 'R'.
// A string is said to be balanced if each of its characters appears n / 4 times where n is the length of
// the string.
// Return the minimum length of the substring that can be replaced with any other string of the same length
// to make s balanced. If s is already balanced, return 0.
//
// Example 1:
// Input: s = "QWER"
// Output: 0
// Explanation: s is already balanced.
//
// Example 2:
// Input: s = "QQWE"
// Output: 1
// Explanation: We need to replace a 'Q' to 'R', so that "RQWE" (or "QRWE") is balanced.
//
// Example 3:
// Input: s = "QQQW"
// Output: 2
// Explanation: We can replace the first "QQ" to "ER".
//
// Constraints:
// n == s.length
// 4 <= n <= 10^5
// n is a multiple of 4.
// s contains only 'Q', 'W', 'E', and 'R'.
func balancedString(s string) int {
	target := len(s) / 4
	counts := make([]int, 4)
	for _, char := range s {
		counts[charIndex(char)]++
	}

	var debt int
	for idx, count := range counts {
		if count <= target {
			counts[idx] = 0
		} else if count > target {
			counts[idx] = count - target
			debt += count - target
		}
	}
	if debt == 0 {
		return 0
	}

	length := len(s)
	for left, right := 0, 0; right < len(s); right++ {
		if counts[charIndex(s[right])] > 0 {
			debt--
		}
		counts[charIndex(s[right])]--

		if debt == 0 {
			for counts[charIndex(s[left])] < 0 {
				counts[charIndex(s[left])]++
				left++
			}

			if right-left+1 < length {
				length = right - left + 1
			}
		}
	}

	return length
}

func charIndex[T ~uint8 | ~int32](b T) int {
	switch b {
	case 'Q':
		return 0
	case 'W':
		return 1
	case 'E':
		return 2
	default:
		return 3
	}
}
