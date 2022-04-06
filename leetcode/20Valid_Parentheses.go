package leetcode

// https://leetcode.com/problems/valid-parentheses/

// isValid Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
//
// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
func isValid(s string) bool {
	arr := make([]byte, len(s))
	offset := -1
	for i := 0; i < len(s); i++ {
		if offset < 0 {
			offset++
			arr[offset] = s[i]
			continue
		}

		switch s[i] {
		case ')':
			if arr[offset] != '(' {
				return false
			}
			offset--
		case '}':
			if arr[offset] != '{' {
				return false
			}
			offset--
		case ']':
			if arr[offset] != '[' {
				return false
			}
			offset--
		default:
			offset++
			arr[offset] = s[i]
		}
	}

	return offset < 0
}
