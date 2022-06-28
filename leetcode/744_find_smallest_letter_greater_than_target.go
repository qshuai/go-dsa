package leetcode

// https://leetcode.com/problems/find-smallest-letter-greater-than-target/

// nextGreatestLetter Given a characters array letters that is sorted in non-decreasing order and
// a character target, return the smallest character in the array that is larger than target.
//
// Constraints:
// 2 <= letters.length <= 104
// letters[i] is a lowercase English letter.
// letters is sorted in non-decreasing order.
// letters contain at least two different characters.
// target is a lowercase English letter.
func nextGreatestLetter(letters []byte, target byte) byte {
	begin, end := 0, len(letters)-1
	var mid int
	for begin <= end {
		mid = begin + (end-begin)>>1
		if letters[mid] <= target {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}

	if begin >= len(letters) {
		// wrap around
		begin = 0
	}
	return letters[begin]
}
