package leetcode

// https://leetcode.com/problems/find-all-anagrams-in-a-string/

// findAnagrams Given two strings s and p, return an array of all the start indices of p's anagrams in s.
// You may return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
//
// Constraints:
// 1 <= s.length, p.length <= 3 * 10^4
// s and p consist of lowercase English letters.
func findanagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}

	const StartPoint = 'a'
	var sCount [26]int
	var pCount [26]int
	for i := 0; i < len(p); i++ {
		sCount[s[i]-StartPoint]++
		pCount[p[i]-StartPoint]++
	}

	var ans []int
	for i := 0; i <= len(s)-len(p); i++ {
		if pCount == sCount {
			ans = append(ans, i)
		}

		sCount[s[i]-StartPoint]--
		if i < len(s)-len(p) {
			sCount[s[i+len(p)]-StartPoint]++
		}
	}

	return ans
}
